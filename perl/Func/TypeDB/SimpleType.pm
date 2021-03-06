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
package Func::TypeDB::SimpleType;
use Moose;

has 'go_type' => (
    is  => 'ro',
    isa => 'Str',
);

has 'c_type' => (
    is  => 'ro',
    isa => 'Str',
);

sub convert {
    my $self = shift;
    my $fc = shift;
    return 'return ' . $self->go_type . '('.$fc.')';
}

sub convert_to_c {
    my $self = shift;
    my $name = shift;
    if ($self->go_type eq 'uint') {
        return 'C.'.$self->go_type . '('.$name.')';
    }
    else {
        return 'C.'.$self->c_type . '('.$name.')';
    }
}

sub actual_go_type {
    my $self = shift;
    return $self->go_type;
}

1;
