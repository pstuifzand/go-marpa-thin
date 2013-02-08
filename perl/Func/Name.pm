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
package Func::Name;
use Moose;

has 'name' => (
    is =>'ro',
    isa => 'Str',
);

sub c_name {
    my $self = shift;
    return 'C.'.$self->name;
}

my %marpa_type = (
    'c' => 'config',
    'g' => 'grammar',
    'r' => 'recognizer',
    'o' => 'order',
    'b' => 'bocage',
    't' => 'tree',
    'v' => 'value',
);

sub is_marpa_class_type {
    my $self = shift;
    return $self->name =~ m/^marpa_[cgrobtv]_/;
}

sub go_type {
    my $self = shift;
    my $name = $self->name;
    if ($self->is_marpa_class_type) {
        $name =~ s/^marpa_([cgrobtv])_.+$/$marpa_type{$1}/;
    }
    return _to_camel($name);
}

sub go_name {
    my $self = shift;
    my $arg = shift;
    $arg //= 1;

    my $name = $self->name;
    
    if ($self->is_marpa_class_type) {
        if ($arg) {
            $name =~ s/^marpa_([cgrobtv])_/$marpa_type{$1}_/;
        }
        else {
            $name =~ s/^marpa_([cgrobtv])_//;
        }
    }
    if ($name =~ m/^(\w+)_new$/) {
        $name = 'new_'.$1;
    }
    return _to_camel($name);
}

sub _to_camel {
    my $v = shift;
    $v =~ s/_(\w)/\U$1/g;
    return ucfirst $v;
}

1;
