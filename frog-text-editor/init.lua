local api = require("rmp.rmp")

local Terminal = api.Terminal
local VirtualTerminal = api.VirtualTerminal
local Frame = api.Frame
local Text = api.Text
local Code = api.Code
local CodeSyntax = api.CodeSyntax
local StatusBar = api.StatusBar
local Options = api.Options
local Scroller = api.Scroller

-- Vim-like Text Editor Class
local VimEditor = {}
VimEditor.__index = VimEditor

function VimEditor.new()
    local self = setmetatable({}, VimEditor)

    self.lines = { "" }
    self.cursor_line = 1
    self.cursor_col = 1
    self.scroll_offset_y = 0
    self.scroll_offset_x = 0
    self.file_path = nil
    self.modified = false
    self.view_height = 20
    self.view_width = 80
    self.current_syntax = "LUA"

    -- Vim modes
    self.mode = "NORMAL"
    self.syntax_mode = false
    self.syntax_options = { "LUA", "C", "PYTHON" }
    self.syntax_scroller = nil

    return self
end

function VimEditor:initSyntaxScroller()
    local opts = Options.new(self.syntax_options)
    self.syntax_scroller = Scroller.new(3, opts)
    self.syntax_scroller:setOptionsObj(opts)
end

function VimEditor:getSize()
    local h, w = Terminal:getSize()
    self.view_height = h - 5
    self.view_width = w - 8
    return h, w
end

function VimEditor:insertChar(char)
    if self.mode == "INSERT" then
        local line = self.lines[self.cursor_line]
        self.lines[self.cursor_line] = line:sub(1, self.cursor_col - 1) .. char .. line:sub(self.cursor_col)
        self.cursor_col = self.cursor_col + 1
        self.modified = true
    end
end

function VimEditor:deletChar()
    if self.mode == "INSERT" and self.cursor_col > 1 then
        local line = self.lines[self.cursor_line]
        self.lines[self.cursor_line] = line:sub(1, self.cursor_col - 2) .. line:sub(self.cursor_col)
        self.cursor_col = self.cursor_col - 1
        self.modified = true
    end
end

function VimEditor:deleteCharForward()
    if self.mode == "INSERT" then
        local line = self.lines[self.cursor_line]
        if self.cursor_col <= #line then
            self.lines[self.cursor_line] = line:sub(1, self.cursor_col - 1) .. line:sub(self.cursor_col + 1)
            self.modified = true
        end
    end
end

function VimEditor:newLine()
    if self.mode == "INSERT" then
        local line = self.lines[self.cursor_line]
        local before = line:sub(1, self.cursor_col - 1)
        local after = line:sub(self.cursor_col)

        self.lines[self.cursor_line] = before
        table.insert(self.lines, self.cursor_line + 1, after)

        self.cursor_line = self.cursor_line + 1
        self.cursor_col = 1
        self.modified = true
    end
end

function VimEditor:moveCursorUp()
    if self.cursor_line > 1 then
        self.cursor_line = self.cursor_line - 1
        local line = self.lines[self.cursor_line]
        self.cursor_col = math.min(self.cursor_col, #line + 1)
    end
end

function VimEditor:moveCursorDown()
    if self.cursor_line < #self.lines then
        self.cursor_line = self.cursor_line + 1
        local line = self.lines[self.cursor_line]
        self.cursor_col = math.min(self.cursor_col, #line + 1)
    end
end

function VimEditor:moveCursorLeft()
    if self.cursor_col > 1 then
        self.cursor_col = self.cursor_col - 1
    end
end

function VimEditor:moveCursorRight()
    local line = self.lines[self.cursor_line]
    if self.cursor_col < #line + 1 then
        self.cursor_col = self.cursor_col + 1
    end
end

function VimEditor:moveHome()
    self.cursor_col = 1
end

function VimEditor:moveEnd()
    self.cursor_col = #self.lines[self.cursor_line] + 1
end

function VimEditor:moveBeginning()
    self.cursor_line = 1
    self.cursor_col = 1
end

function VimEditor:moveEndFile()
    self.cursor_line = #self.lines
    self.cursor_col = 1
end

function VimEditor:deleteLine()
    if self.mode == "NORMAL" and #self.lines > 1 then
        table.remove(self.lines, self.cursor_line)
        self.cursor_line = math.min(self.cursor_line, #self.lines)
        self.modified = true
    end
end

function VimEditor:insertNewLineAbove()
    if self.mode == "NORMAL" then
        table.insert(self.lines, self.cursor_line, "")
        self.cursor_col = 1
        self.modified = true
        self.mode = "INSERT"
    end
end

function VimEditor:insertNewLineBelow()
    if self.mode == "NORMAL" then
        table.insert(self.lines, self.cursor_line + 1, "")
        self.cursor_line = self.cursor_line + 1
        self.cursor_col = 1
        self.modified = true
        self.mode = "INSERT"
    end
end

function VimEditor:updateScroll()
    if self.cursor_line < self.scroll_offset_y + 1 then
        self.scroll_offset_y = self.cursor_line - 1
    elseif self.cursor_line > self.scroll_offset_y + self.view_height then
        self.scroll_offset_y = self.cursor_line - self.view_height
    end

    if self.cursor_col < self.scroll_offset_x + 1 then
        self.scroll_offset_x = self.cursor_col - 1
    elseif self.cursor_col > self.scroll_offset_x + self.view_width - 5 then
        self.scroll_offset_x = self.cursor_col - self.view_width + 5
    end

    self.scroll_offset_x = math.max(0, self.scroll_offset_x)
    self.scroll_offset_y = math.max(0, self.scroll_offset_y)
end

function VimEditor:getContent()
    return table.concat(self.lines, "\n")
end

function VimEditor:setContent(content)
    self.lines = {}
    for line in content:gmatch("[^\n]*") do
        table.insert(self.lines, line)
    end
    if #self.lines == 0 then
        self.lines = { "" }
    end
    self.cursor_line = 1
    self.cursor_col = 1
    self.modified = false
end

function VimEditor:setSyntax(syntax_type)
    if syntax_type == "LUA" or syntax_type == "C" or syntax_type == "PYTHON" then
        self.current_syntax = syntax_type
        return true
    end
    return false
end

function VimEditor:renderWithSyntax(vterm)
    local line_num_width = #tostring(#self.lines) + 1
    local start_line = self.scroll_offset_y + 1

    for display_idx = 1, self.view_height do
        local line_idx = start_line + display_idx - 1

        if line_idx <= #self.lines then
            local line = self.lines[line_idx]
            local is_current_line = (line_idx == self.cursor_line)

            -- Line number
            local line_num_str = string.format("%" .. line_num_width .. "d", line_idx)
            local line_num_color = is_current_line and api.FGColors.Brights.Yellow or api.FGColors.NoBrights.Cyan
            vterm:writeText(2, display_idx + 1, line_num_str, line_num_color, api.BGColors.NoBrights.Black)
            vterm:writeText(line_num_width + 3, display_idx + 1, api.BoxDrawing.LightBorder[2],
                api.FGColors.NoBrights.White)

            -- Use Code class for syntax highlighting
            local code_obj = Code.new(line, self.current_syntax, line_num_width + 5, display_idx + 1)
            local code_vterm = code_obj:render()
            vterm:merge(code_vterm, 0, 0)

            -- Padding
            local line_display_len = #line - self.scroll_offset_x
            if line_display_len < self.view_width then
                vterm:writeText(line_num_width + 5 + line_display_len, display_idx + 1,
                    string.rep(" ", self.view_width - line_display_len))
            end
        else
            vterm:writeText(line_num_width + 3, display_idx + 1, "~", api.FGColors.NoBrights.White)
        end
    end

    -- Draw cursor based on mode
    local cursor_screen_y = self.cursor_line - self.scroll_offset_y + 1
    local cursor_screen_x = line_num_width + 5 + (self.cursor_col - self.scroll_offset_x - 1)

    if cursor_screen_y > 1 and cursor_screen_y <= self.view_height + 1 then
        if self.mode == "INSERT" then
            vterm:writeText(cursor_screen_x, cursor_screen_y, "|", api.FGColors.Brights.Green)
        else
            vterm:writeText(cursor_screen_x, cursor_screen_y, api.Bar_100_per, api.FGColors.Brights.White)
        end
    end
end

function VimEditor:renderSyntaxMenu(vterm)
    local visible, start_idx, focus_idx = self.syntax_scroller:getVisible()
    local menu_y = 3

    vterm:writeText(5, menu_y, "Select Syntax (Use j/k, Enter to select):", api.FGColors.Brights.Cyan)

    for i, item in ipairs(visible) do
        local y_pos = menu_y + i
        local fg = item.fg or api.FGColors.Brights.White
        local bg = api.BGColors.NoBrights.Black

        if i == focus_idx then
            fg = api.FGColors.Brights.Yellow
            bg = api.BGColors.NoBrights.Black
            vterm:writeText(5, y_pos, "> " .. item.text, fg, bg)
        else
            vterm:writeText(5, y_pos, "  " .. item.text, fg, bg)
        end
    end
end

function VimEditor:handleKey(key)
    if self.syntax_mode then
        if key == api.KEY_J then
            self.syntax_scroller:nextLine()
        elseif key == api.KEY_K then
            self.syntax_scroller:prevLine()
        elseif key == api.KEY_ENTER then
            local idx = self.syntax_scroller:getCursorPosition()
            local syntax = self.syntax_options[idx]
            if syntax then
                self:setSyntax(syntax)
            end
            self.syntax_mode = false
            self.mode = "NORMAL"
        elseif key == api.KEY_ESCAPE then
            self.syntax_mode = false
            self.mode = "NORMAL"
        end
        return
    end

    if self.mode == "INSERT" then
        if key >= api.KEY_A and key <= api.KEY_Z then
            local chars = { "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
                "t", "u", "v", "w", "x", "y", "z" }
            self:insertChar(chars[key - api.KEY_A + 1])
        elseif key >= api.KEY_SHIFT_A and key <= api.KEY_SHIFT_Z then
            local chars = { "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
                "T", "U", "V", "W", "X", "Y", "Z" }
            self:insertChar(chars[key - api.KEY_SHIFT_A + 1])
        elseif key >= api.KEY_0 and key <= api.KEY_9 then
            self:insertChar(tostring(key - api.KEY_0))
        elseif key == api.KEY_SPACE then
            self:insertChar(" ")
        elseif key == api.KEY_BACKSPACE then
            self:deletChar()
        elseif key == api.KEY_DELETE then
            self:deleteCharForward()
        elseif key == api.KEY_ENTER then
            self:newLine()
        elseif key == api.KEY_UP then
            self:moveCursorUp()
        elseif key == api.KEY_DOWN then
            self:moveCursorDown()
        elseif key == api.KEY_LEFT then
            self:moveCursorLeft()
        elseif key == api.KEY_RIGHT then
            self:moveCursorRight()
        elseif key == api.KEY_HOME then
            self:moveHome()
        elseif key == api.KEY_END then
            self:moveEnd()
        elseif key == api.KEY_ESCAPE then
            self.mode = "NORMAL"
            self.cursor_col = math.max(1, self.cursor_col - 1)
        elseif key == api.KEY_DOT then
            self:insertChar(".")
        elseif key == api.KEY_COMMA then
            self:insertChar(",")
        elseif key == api.KEY_SEMICOL then
            self:insertChar(";")
        elseif key == api.KEY_COLON then
            self:insertChar(":")
        elseif key == api.KEY_MINUS then
            self:insertChar("-")
        elseif key == api.KEY_PLUS then
            self:insertChar("+")
        elseif key == api.KEY_STAR then
            self:insertChar("*")
        elseif key == api.KEY_SLASH then
            self:insertChar("/")
        elseif key == api.KEY_OPEN_BRAKET then
            self:insertChar("[")
        elseif key == api.KEY_CLOSED_BRAKET then
            self:insertChar("]")
        elseif key == api.KEY_OPCURB then
            self:insertChar("{")
        elseif key == api.KEY_CLCURB then
            self:insertChar("}")
        elseif key == api.KEY_DBL_QUOTE then
            self:insertChar("\"")
        elseif key == api.KEY_SINGLE_QOUTE then
            self:insertChar("'")
        end
    elseif self.mode == "NORMAL" then
        if key == api.KEY_I then
            self.mode = "INSERT"
        elseif key == api.KEY_A then
            self.mode = "INSERT"
            self:moveCursorRight()
        elseif key == api.KEY_O then
            self:insertNewLineBelow()
        elseif key == api.KEY_SHIFT_O then
            self:insertNewLineAbove()
        elseif key == api.KEY_H then
            self:moveCursorLeft()
        elseif key == api.KEY_J then
            self:moveCursorDown()
        elseif key == api.KEY_K then
            self:moveCursorUp()
        elseif key == api.KEY_L then
            self:moveCursorRight()
        elseif key == api.KEY_G then
            self:moveBeginning()
        elseif key == api.KEY_SHIFT_G then
            self:moveEndFile()
        elseif key == api.KEY_D then
            self:deleteLine()
        elseif key == api.KEY_COLON then
            self.syntax_mode = true
            self:initSyntaxScroller()
        end
    end
end

-- Main Application
local window = Frame.new()
window:initMainFrame()
window:setFps(60)

local editor = VimEditor.new()
editor:getSize()

-- Create Status Bar
local statusbar = StatusBar.new(editor.view_height + 3, editor.view_width + 6)
statusbar:addComponent("VIM-Like RMP Text Editor with Syntax Highlighting", api.FGColors.Brights.Cyan,
    api.BGColors.NoBrights.Black, api.TextStyle.Bold, "left")
statusbar:addComponent("hjkl: Move | i/a: Insert | o/O: NewLine | dd: Delete | :: Syntax", api.FGColors.Brights.White,
    api.BGColors.NoBrights.Black, nil, "right")

-- Initial content with Lua code
local sample_code = [[-- VIM-like Editor with Code Syntax Highlighting
local function fibonacci(n)
  if n <= 1 then
    return n
  end
  return fibonacci(n - 1) + fibonacci(n - 2)
end

for i = 1, 10 do
  print(fibonacci(i))
end
-- Press ESC for NORMAL mode, 'i' to INSERT
-- Use hjkl to navigate, dd to delete line, ':' for syntax]]

editor:setContent(sample_code)

local key = api.Terminal:handleKey()

while key ~= api.KEY_CTRL_Q do
    local h, w = Terminal:getSize()

    if h ~= editor.view_height + 5 or w ~= editor.view_width + 8 then
        editor:getSize()
    end

    if key ~= api.NONE and key ~= nil then
        editor:handleKey(key)
    end

    editor:updateScroll()

    -- Render frame
    local vterm = VirtualTerminal.new()

    -- Title bar with mode indicator
    local mode_display = "[ " .. editor.mode .. " ]"
    local syntax_display = "[" .. editor.current_syntax .. "]"
    local title = " VIM Editor " .. syntax_display .. " "
    local title_padding = math.floor((editor.view_width - #title) / 2)
    vterm:writeText(1, 1,
        string.rep(api.BoxDrawing.LightBorder[1], title_padding) ..
        title .. string.rep(api.BoxDrawing.LightBorder[1], editor.view_width - title_padding - #title),
        api.FGColors.Brights.Magenta, api.BGColors.NoBrights.Black)

    if editor.syntax_mode then
        -- Show syntax menu
        editor:renderSyntaxMenu(vterm)
    else
        -- Editor content with syntax highlighting
        editor:renderWithSyntax(vterm)
    end

    -- Bottom separator
    local sep_y = editor.view_height + 2
    vterm:writeText(1, sep_y, string.rep(api.BoxDrawing.LightBorder[1], editor.view_width + 6),
        api.FGColors.NoBrights.White)

    -- Status info with mode color
    local mode_color = api.FGColors.Brights.White
    if editor.mode == "INSERT" then
        mode_color = api.FGColors.Brights.Green
    end

    local status = string.format("Ln %d, Col %d | %s | %s",
        editor.cursor_line,
        editor.cursor_col,
        editor.modified and "[+]" or "[-]",
        mode_display)
    vterm:writeText(2, sep_y, status, mode_color, api.BGColors.NoBrights.Black)

    -- Status bar
    vterm:merge(statusbar:render(), 0, sep_y + 1)

    window:add(vterm)
    window:run(key, nil, nil, nil)

    key = api.Terminal:handleKey()
end

window:cleanupMainFrame()
