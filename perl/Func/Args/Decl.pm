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
package Func::Args::Decl;
use Moose;

has 'type' => (
    is  => 'ro',
    isa => 'Str',
);
has 'name' => (
    is  => 'ro',
    isa => 'Str',
);

my $typedb = Func::TypeDB->new;

sub to_go_value {
    my $self = shift;
    my $info = $typedb->get_type($self->type);
    return $info->convert_to_c($self->name);
}

sub go_type {
    my $self = shift;
    my $info = $typedb->get_type($self->type);
    return $info->actual_go_type;
}

1;
