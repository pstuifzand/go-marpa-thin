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
