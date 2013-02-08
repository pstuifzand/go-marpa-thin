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
