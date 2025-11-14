FRG_Begin
    FRG_Int i #
    i := 0 #
    Repeat
        i := i + 1 #
        If [i % 2 == 0]
        Begin
            Continue #
        End
        FRG_Print i, " " #
    Until [i == 5]
    FRG_Print "done" #
FRG_End
