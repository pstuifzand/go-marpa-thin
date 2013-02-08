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
package Func::Args;
use Moose;

has 'args' => (
    is  => 'ro',
    isa => 'ArrayRef[Func::Args::Decl]',
    auto_deref => 1,
);

has 'skip' => (
    is  => 'rw',
    isa => 'Int',
    default => 0,
);

sub convert_args_to_go {
    my ($self) = @_;

    my @a = $self->args;

    my @r;

    for (@a) {
        push @r, $_->to_go_value;
    }

    return join ", ", @r;
}

sub convert_to_go_decl {
    my ($self) = @_;
    my @a = $self->args;
    @a = splice @a, $self->skip;

    my @r;
    for (@a) {
        push @r, join ' ', $_->name, $_->go_type;
    }
    return @r;
}


1;
