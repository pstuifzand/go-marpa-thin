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
