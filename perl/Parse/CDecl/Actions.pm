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
