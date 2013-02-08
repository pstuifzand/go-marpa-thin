package Func::TypeDB::HandleType;
use Moose;

has 'go_type' => (
    is  => 'ro',
    isa => 'Str',
);

sub convert {
    my $self = shift;
    my $fc = shift;
    my $type = $self->go_type;
    return <<"GO";
var _ret $type
_ret.handle = $fc
return _ret
GO
}

sub convert_to_c {
    my $self = shift;
    my $name = shift;
    return $name . '.handle';
}

sub actual_go_type {
    my $self = shift;
    return '*'.$self->go_type;
}

1;
