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
