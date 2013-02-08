package Func::TypeDB;
use strict;
use Moose;

my %info = (
    'int'                   => ['SimpleType', 'int' ],
    'int*'                  => ['SimpleType', '*int' ],
    'unsigned int'          => ['SimpleType', 'uint' ],
    'Marpa_Config*'         => ['HandleType', '*Config' ],
    'Marpa_Error_Code'      => ['SimpleType', 'ErrorCode' ],
    'Marpa_Symbol_ID'       => ['SimpleType', 'SymbolID' ],
    'Marpa_Grammar'         => ['HandleType', 'Grammar' ],
    'Marpa_Rule_ID'         => ['SimpleType', 'RuleID' ],
    'Marpa_Rank'            => ['SimpleType', 'Rank' ],
    'Marpa_Recognizer'      => ['HandleType', 'Recognizer' ],
    'Marpa_Earleme'         => ['SimpleType', 'Earleme' ],
    'Marpa_Earley_Set_ID'   => ['SimpleType', 'EarleySetID' ],
    'Marpa_Earley_Set_ID*'  => ['SimpleType', '*EarleySetID' ],
    'Marpa_Symbol_ID*'      => ['SimpleType', '[]SymbolID' ],
    'Marpa_Bocage'          => ['HandleType', 'Bocage' ],
    'Marpa_Order'           => ['HandleType', 'Order' ],
    'Marpa_Tree'            => ['HandleType', 'Tree' ],
    'Marpa_Value'           => ['HandleType', 'Value' ],
    'Marpa_Step_Type'       => ['SimpleType', 'StepType' ],
    'Marpa_Event*'          => ['HandleType', '*Event' ],
    'Marpa_Event_Type'      => ['SimpleType', 'EventType' ],
);

sub get_type {
    my ($self, $type) = @_;

    if ($type !~ m/^[\w\s]+\*?$/) {
        die "Wrong type: $type";
    }
    if (!exists $info{$type}) {
        die "Unknown type: $type";
    }

    my $class_type = $info{$type};

    my $class = 'Func::TypeDB::'.$class_type->[0];
    eval "require $class;1;";
    die $@ if $@;
    return $class->new(c_type => $type, go_type => $class_type->[1]);
}

1;
