## This is a test program for the FROG language
FRG_Begin
    FRG_Int i, j, x1, x2 #
    FRG_Real x3 #
    FRG_Strg name #

    name := "rayden" #
    ##FRG_Print "Hello, ", name, "!" #

    i := 30 + 5 #
    ## i := 5 #

    If [ i >= 20 ]
    Begin
        x1 := 10 #
    End
    Else
    Begin
        x1 := 30 #
        x3 := x1 * 4 #
    End

    FRG_Print x1 , "\n", x3 #

    FRG_Print "Hello, ", name, "!" #

    If [False]
    Begin
        FRG_Print "Hello inside false condition"#
    End


    ## instruction de boucle Repeat
    ## Repeat
    ##     i := i - 5 #
    ##     FRG_Print i #
    ## Until [ i <= 15 ]

FRG_End
