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
        FRG_Print "if"#
        FRG_Print "if"#
    End
    Else
    Begin
        FRG_Print "else"#
        FRG_Print "else"#
        FRG_Print "else"#
    End

    FRG_Print "out"#


FRG_End
