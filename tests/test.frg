## This is a test program for the FROG language
FRG_Begin
    FRG_Int i, j, x1, x2 #
    FRG_Real x3 #
    FRG_Strg name #

    name := "FROG" #
    i := 30 #

    If [ i <= 20 ]
        x1 := 10 #
    Else
    Begin
        x1 := 30 #
        x3 := x1 * 4 #
        FRG_Print x1, x3 #
    End

    FRG_Print "Hello, ", name, "!" #

    ## instruction de boucle Repeat
    Repeat
        i := i - 5 #
    Until [ i <= 15 ]
FRG_End
