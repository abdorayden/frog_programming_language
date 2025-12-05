## Copyright (C) by abdenour souane
##
## frog standerd library
##

FRG_Begin
    ## define standerd functions

    ## Enum
    ## examples:
    ##  FRG_Int SunDay , Monday#
    ##  SunDay := enum(0)#
    ##  Monday := enum(N_START)#

    FRG_Int N_START#
    ## non start , that's mean continue inc for each call
    N_START := -9999#
    FRG_Fn iota(FRG_Int start) : FRG_Int 
    Begin
        FRG_Int ___local_counter#
        ___local_counter := 0#
        FRG_Int ret#
        If [start != N_START] 
        Begin
             ___local_counter := start#
        End
        ret := ___local_counter#
        ___local_counter := ___local_counter + 1#
        iota := ret#
    End

    
    ## math functions
    ## stolen from math stdc
    FRG_Real E , PI,LOG2E,LOG10E,LN2,LN10,PI_2,PI_4,M_1_PI,M_2_PI,M_2_SQRTPI,SQRT2,SQRT1_2#

    E := 2.718281828459045#
    PI:= 3.141592653589793#

    ## /* log_2 e */
    LOG2E :=	1.4426950408889634074 # 	
    ## /* log_10 e */
    LOG10E :=	0.43429448190325182765 # 	
    ## /* log_e 2 */
    LN2 :=		0.69314718055994530942 # 	
    ## /* log_e 10 */
    LN10 :=		2.30258509299404568402 # 	
    ## /* pi/2 */
    PI_2 :=		1.57079632679489661923 # 	
    ## /* pi/4 */
    PI_4 :=		0.78539816339744830962 # 	
    ## /* 1/pi */
    M_1_PI :=		0.31830988618379067154 # 	
    ## /* 2/pi */
    M_2_PI :=		0.63661977236758134308 # 	
    ## /* 2/sqrt(pi) */
    M_2_SQRTPI :=	1.12837916709551257390 # 	
    ## /* sqrt(2) */
    SQRT2 :=	1.41421356237309504880 # 	
    ## /* 1/sqrt(2) */
    
    SQRT1_2 :=	0.70710678118654752440 # 	

    FRG_Fn add(FRG_Int x , FRG_Int y) : FRG_Int
    Begin
        add := x + y#
    End
    FRG_Fn sub(FRG_Int x , FRG_Int y) : FRG_Int
    Begin
        sub := x - y#
    End
    FRG_Fn mul(FRG_Int x , FRG_Int y) : FRG_Int
    Begin
        mul := x * y#
    End
    FRG_Fn div(FRG_Int x , FRG_Int y) : FRG_Int
    Begin
        If [y == 0]
        Begin
            FRG_Print "[ERROR] : y is equal to zero"#
            add := 0#
        End
        Else
        Begin
            add := x / y#
        End
    End

    FRG_Fn pow(FRG_Int x , FRG_Int of_what) : FRG_Int
    Begin
        If [of_what == 0]
        Begin
            pow := 1#
        End
        Else
        Begin
            FRG_Int i,res#
            i := 0#
            res := 1#
            Repeat
                res := x * res#
                i := i + 1#
            Until [i == of_what]
            pow := res#
        End
    End

    ## factorial
    FRG_Fn factorial(FRG_Int n) : FRG_Int
    Begin
        FRG_Int forRet#
        forRet := 1#
        If [ n == 0 ]
        Begin
            factorial := forRet#
        End
        Else
        Begin
            If [n == 1]
            Begin
                factorial := forRet#
            End
            Else
            Begin
                Repeat
                    forRet := forRet * n#
                    n := n - 1#
                Until [n == 1]
                factorial := forRet#
            End
        End
    End

    FRG_Fn pgcd(FRG_Int a , FRG_Int b) : FRG_Int 
    Begin
        If [b == 0]
        Begin
            pgcd := a#
        End
        Else
        Begin
            If [b > a]
            Begin
                FRG_Int temp#
                temp := a#
                a := b#
                b := temp#
            End
            FRG_Int c#
            Repeat
               c := a%b#
               a := b#
               b := c#
            Until [b == 0]

            pgcd := a#
        End
    End

    FRG_Fn ppcm(FRG_Int a , FRG_Int b) : FRG_Int
    Begin
        FRG_Int prod#
        prod := a*b#
        ppcm := prod / pgcd(a,b)#
    End

    ## sqrt
    ## resource : https://en.wikipedia.org/wiki/Square_root_algorithms,
    ##           https://github.com/MichaelDipperstein/sqrt
    FRG_Fn sqrt(FRG_Real tha_number) : FRG_Real
    Begin

        FRG_Real SQRT_TOLERANCE#
        SQRT_TOLERANCE := 0.0001#

        FRG_Real guess, min, max, delta#
        If [tha_number < 0.0]
        Begin
            FRG_Print "errno : EDOM"#
            sqrt := -1.0#
        End
        Else
        Begin
            ##     /* come up with initail guess and bounds */
            If [tha_number < 1.0]
            Begin
                guess := tha_number * 2.0#
                max := 1.0#
            End
            Else
            Begin
                guess := tha_number / 2.0#
                max := tha_number#
            End
            min := 0.0#
            delta := guess * guess#
            delta := delta - tha_number#

            Repeat
                delta := guess * guess#
                delta := delta - tha_number#
                If [delta > SQRT_TOLERANCE]
                Begin
                    ## /* guess is too high bisect min and guess */
                    max := guess#
                End
                Else
                Begin
                    If [delta < -SQRT_TOLERANCE]
                    Begin
                       ## /* guess is too low bisect max and guess */
                       min := guess#
                    End
                    Else
                    Begin
                        ##/* our guess is good enough */
                        Break#
                    End
                End
                ##/* bisect new bound to get new guess */
                guess := min + max#
                guess := guess / 2.0#
            Until [False]
        End
        sqrt := guess#
    End

    ## dynamic arrays functions
    FRG_Int INT , FLOAT , STRINGS#
    INT := iota(N_START)#
    FLOAT := iota(N_START)#
    STRINGS := iota(N_START)#

    FRG_Fn alloc_ints(FRG_Int size) : FRG_Int
    Begin
        alloc_ints := [size]#
    End

    FRG_Fn alloc_floats(FRG_Int size) : FRG_Real
    Begin
        alloc_floats := [size]#
    End

    FRG_Fn alloc_strings(FRG_Int size) : FRG_Strg
    Begin
        alloc_strings := [size]#
    End

    ## strings operations

FRG_End
