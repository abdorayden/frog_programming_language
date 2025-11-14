FRG_Begin
    FRG_Fn printName(FRG_Strg name) : FRG_Int
    Begin
        FRG_Print "hello " , name #
        FRG_Print "\n"#

        ## ignore return value
        foo := 0#
    End

    FRG_Fn add(FRG_Int a , FRG_Int b) : FRG_Int
    Begin
        add := a + b#
    End

    printName("rayden")#
    FRG_Int res#
    res := add(10,1)#
    FRG_Print res #

FRG_End


