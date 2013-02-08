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
package Parse::CDecl;
use Marpa::R2;
use Parse::CDecl::Actions;

sub new {
    my ($class) = @_;

    my $grammar = Marpa::R2::Scanless::G->new({
        action_object => 'Parse::CDecl::Actions',
        default_action => 'do_first_arg',
        source => \<<'SOURCE',

:start     ::= decl

decl       ::= type id ('(') type_vars (')' ';')  action => do_func_node
             | '#define' id number                action => do_define_node

type_vars  ::= type_var*                        action => do_func_args separator => comma

type_var   ::= type id                          action => do_func_args_decl

type       ::= type_id+                         action => do_join_s

type_id    ::= id                               action => do_join
             | id '*'                           action => do_join
             | id '*' '*'                       action => do_join

comma        ~ ','
id           ~ [a-zA-Z_] id_rest
id_rest      ~ [\w]*

number       ~ [\d]+

:discard     ~ ws
ws           ~ [\s]+

SOURCE
    });

    my $self = {
        grammar => $grammar,
    };
    return bless $self, $class;
}

sub parse {
    my ($self, $input) = @_;
    my $re = Marpa::R2::Scanless::R->new({grammar => $self->{grammar}});
    $re->read(\$input);
    my $val = $re->value;
    return $$val;
}

1;
