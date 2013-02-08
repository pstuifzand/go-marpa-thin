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
package Parse::CDecl::Actions;
use Func;
use Func::ReturnVal;
use Func::Name;
use Func::Args;
use Func::Args::Decl;

sub new {
    my ($class) = @_;
    return bless {}, $class;
}

sub do_first_arg {
    shift;return $_[0];
}

sub do_list {
    shift;return \@_;
}

sub do_join {
    shift;return join '', @_;
}

sub do_join_s {
    shift;return join ' ', @_;
}

sub do_func_node {
    shift;
    return Func->new({
        return_value => Func::ReturnVal->new({type => $_[0]}),
        name         => Func::Name->new({name=>$_[1]}),
        args         => $_[2],
    });
}

sub do_func_args_decl {
    shift; return Func::Args::Decl->new({type => $_[0], name => $_[1] });
}

sub do_func_args {
    shift; return Func::Args->new({ args => \@_ });
}


sub do_define_node {
    shift; return \@_;
}

1;
