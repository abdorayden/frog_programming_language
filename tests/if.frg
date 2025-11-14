FRG_Begin

    FRG_Int true , false#
    true := 1==1#
    false := 1!=1#

    FRG_Int x , y #

    x := 5  #
    y := 5  #

    If [x == y]
    Begin
        FRG_Print "if"#
        FRG_Print "\n" #
        FRG_Print "if"#
        FRG_Print "\n" #
        FRG_Print "if"#
        FRG_Print "\n" #
    End
    Else
    Begin
        FRG_Print "else"#
        FRG_Print "\n" #
        FRG_Print "else"#
        FRG_Print "\n" #
        FRG_Print "else"#
        FRG_Print "\n" #
    End

    FRG_Print "out"#
    FRG_Print "\n" #


FRG_End
