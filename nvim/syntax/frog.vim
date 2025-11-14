if exists("b:current_syntax")
  finish
endif

syn keyword frogKeyword FRG_Begin FRG_End If Else Begin End Repeat Until break continue
syn keyword frogType FRG_Int FRG_Real FRG_Strg
syn keyword frogStatement FRG_Print
syn keyword frogBoolean True False

syn match frogNumber "\v\<\d+\.?\d*\>"
syn region frogString start="\"" skip=/\\"/ end="\""
syn match frogComment "\v##.*$"
syn match frogOperator "\v(:=|==|!=|<=|>=|<|>|\+|-|\*|/|%)"
syn match frogDelimiter "[][(){}]"
syn match frogTerminator "#"

hi def link frogKeyword Keyword
hi def link frogType Type
hi def link frogStatement Statement
hi def link frogBoolean Boolean
hi def link frogNumber Number
hi def link frogString String
hi def link frogComment Comment
hi def link frogOperator Operator
hi def link frogDelimiter Delimiter
hi def link frogTerminator Delimiter

let b:current_syntax = "frog"

