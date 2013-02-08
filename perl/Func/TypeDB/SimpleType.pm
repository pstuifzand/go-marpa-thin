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
