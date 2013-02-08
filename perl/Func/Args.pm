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
