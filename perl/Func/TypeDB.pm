# Copyright (C) 2013 Peter Stuifzand
#
# This program is free software: you can redistribute it and/or modify it under
# the terms of the GNU Lesser General Public License as published by the Free
# Software Foundation, either version 3 of the License, or (at your option) any
# later version.
#
# This program is distributed in the hope that it will be useful, but WITHOUT
# ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
# FOR A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for more
# details.
#
# You should have received a copy of the GNU Lesser General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
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
