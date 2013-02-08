# Generates go code from .h include file
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

package Func;
use strict;
use Moose;
use IO::String;
use Func::TypeDB;

has 'return_value' => (
    is  => 'ro',
    isa => 'Func::ReturnVal',
);

has 'name' => (
    is  => 'ro',
    isa => 'Func::Name',
);

has 'args' => (
    is  => 'ro',
    isa => 'Func::Args',
);

sub go_decl {
    my $self = shift;
    my $fh = IO::String->new;

    print {$fh} "func ";

    my $arg_0 = ($self->args->args)[0];


    my $name;
    if ('*'.$self->name->go_type eq $arg_0->go_type) {
        print {$fh} '(' . $arg_0->name . ' ' . $arg_0->go_type .') ';
        $self->args->skip(1);
        $name = $self->name->go_name(0);
    }
    else {
        $name = $self->name->go_name(1);
    }

    print {$fh} $name;

    print {$fh} "(";
    print {$fh} join ", ", $self->args->convert_to_go_decl;
    print {$fh} ") ";

    my $typedb = Func::TypeDB->new;
    my $type_info = $typedb->get_type($self->return_value->type);

    print {$fh} $type_info->actual_go_type;

    my $func_call = $self->name->c_name."(".$self->args->convert_args_to_go.")";

    print {$fh} " {\n";
    print {$fh} $type_info->convert($func_call);
    print {$fh} "\n}\n";
    return ${$fh->string_ref};
}

1;
