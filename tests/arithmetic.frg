FRG_Begin
    FRG_Int a, b #
    a := 10 #
    b := 5 #

    FRG_Print "a + b" #
    FRG_Print a + b #
    FRG_Print "a - b" #
    FRG_Print a - b #
    FRG_Print "a * b" #
    FRG_Print a * b #

    If [b == 0]
    Begin
        FRG_Print "u can't" #
    End 
    Else
    Begin
        FRG_Print "a / b" #
        FRG_Print a / b #
    End

    FRG_Int i #
    i := 1#
    Repeat
        FRG_Print i #
        i := i + 1 #
    Until [ i > 5 ]

FRG_End
