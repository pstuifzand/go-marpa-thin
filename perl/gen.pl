#!/usr/bin/env perl
package main;
use strict;
use Data::Dumper;
use Parse::CDecl;
use Func::TypeDB;

push @ARGV, '/home/peter/work/Marpa-R2-2.044000/libmarpa_dist/marpa_api.h';

open my $fh, '>', 'marpa-gen.go' or die 'Can\'t open marpa-gen.go';

print {$fh} qq{package marpa\n};
print {$fh} qq{/*
#cgo LDFLAGS: -L. -lmarpa
#include <marpa.h>
#include <marpa_api.h>
*/
};
print {$fh} qq{import "C"\n};
#print {$fh} qq{import (\n"unsafe"\n)\n};
print {$fh} qq{\n};

my $p = Parse::CDecl->new();

my $comment_start = 0;
while (<ARGV>) {
    if (!$comment_start && m{/\*}) {
        $comment_start = 1;
        next;
    }
    elsif ($comment_start && m{\*/}) {
        $comment_start = 0;
        next;
    }
    elsif ($comment_start) {
        next;
    }

    my $decl = $p->parse($_);
    next if !$decl;

    next if ref($decl) eq 'ARRAY' && $decl->[0] eq '#define';
    next if $decl->name->name =~ m/^_/;
    next if $decl->name->name =~ m/_ref$|_unref$/;
    next if $decl->name->name eq 'marpa_c_error';
    next if $decl->name->name eq 'marpa_g_error';

    #print {$fh} "/*\n".Dumper($decl)."\n*/\n";

    print {$fh} $decl->go_decl;
    print {$fh} "\n";
}

close $fh;
