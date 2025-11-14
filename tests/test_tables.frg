FRG_Begin
    FRG_Int[] array #
    FRG_Int len , nameslen#
    ## array := {0,0,0,0,0,0}#
    len := 6#
    array := [len]# ## initialized with zero

    FRG_Int i#
    i := 0#

    Repeat
        array[i] := i#
        FRG_Print array[i]#
        FRG_Print "\n"#
        i := i + 1#
    Until [ i >= len]

    FRG_Print "\n"#
    FRG_Print "\n"#

    i := 0#
    len := 15#
    array := [len]# ## initialized with zero
    Repeat
        array[i] := i#
        FRG_Print array[i]#
        FRG_Print "\n"#
        i := i + 1#
    Until [ i >= len]

    FRG_Strg[] names#
    nameslen := 2#
    i := 0#
    names := {
            "abdo",
            "maroua"
    }#
    Repeat
        FRG_Print names[i]#
        FRG_Print "\n"#
        i := i + 1#
    Until [ i >= nameslen]

FRG_End
