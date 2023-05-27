grammar AdvancedSearch;

/* This will be the entry point of our parser. */
entry
    : expr EOF
    ;

expr
    : singleExpr                    # SimpleExpr
    | lhs=singleExpr ',' rhs=expr   # CommaExpr
    ;


singleExpr
    : '!' singleExpr                                            # NegateExpr
    | '(' singleExpr ')'                                        # NestedExpr
    | lhs=singleExpr '&' rhs=singleExpr                         # AndExpr
    | lhs=singleExpr '|' rhs=singleExpr                         # OrExpr
    | op=(Atk | Def | Sta | Level) min=Number '-' max=Number    # RangeValueExpr
    | op=(Atk | Def | Sta | Level) val=Number                   # SingleValueExpr
    | min=Number '-' max=Number                                 # IvRangeValueExpr
    | val=Number                                                # IvValueExpr
    ;


Atk : 'A' | 'a'
    ;

Def : 'D' | 'd'
    ;

Sta : 'S' | 's'
    ;

Level
    : 'L' | 'l'
    ;


/* A number: can be an integer value */
Number
    :    ('0'..'9')+
    ;

WS
    : [ \t\n\r\f]+ -> skip
    ;

