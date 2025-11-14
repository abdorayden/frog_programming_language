FRG_Begin
    FRG_Int a, b #
    a := 10 #
    b := 5 #

    FRG_Print "a + b\n" #
    FRG_Print a + b #
    FRG_Print "\na - b\n" #
    FRG_Print a - b #
    FRG_Print "\na * b\n" #
    FRG_Print a * b #
    FRG_Print "\n" #

    If [b == 0]
    Begin
        FRG_Print "u can't" #
        FRG_Print "\n" #
    End 
    Else
    Begin
        FRG_Print "a / b" #
        FRG_Print "\n" #
        FRG_Print a / b #
        FRG_Print "\n" #
    End

    FRG_Int i #
    i := 1#
    Repeat
        FRG_Print i #
        FRG_Print "\n" #
        i := i + 1 #
    Until [ i > 5 ]

FRG_End
