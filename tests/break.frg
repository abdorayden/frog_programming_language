FRG_Begin
    FRG_Int i #
    i := 0 #
    Repeat
        i := i + 1 #
        FRG_Print i, " " #
        If [i == 3]
        Begin
            Break #
        End
    Until [i == 10]
    FRG_Print "done" #
FRG_End
