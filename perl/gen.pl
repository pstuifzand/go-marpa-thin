# gen.pl - Generates go code from .h include file
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

#!/usr/bin/env perl
package main;
use strict;
use Data::Dumper;
use Parse::CDecl;
use Func::TypeDB;

push @ARGV, '/home/peter/work/Marpa-R2-2.044000/libmarpa_dist/marpa_api.h';

open my $fh, '>', 'marpa-gen.go' or die 'Can\'t open marpa-gen.go';
open my $const_fh, '>', 'marpa-gen-const.go' or die 'Can\'t open marpa-gen.go';

my $license = <<"LICENSE";
/* Copyright (C) 2013 Peter Stuifzand

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Lesser General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for more
details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

LICENSE

print {$fh} $license;
print {$const_fh} $license;

print {$fh} qq{package marpa\n};
print {$const_fh} qq{package marpa\n};
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

print {$const_fh} "const (\n";

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

    #next if ref($decl) eq 'ARRAY' && $decl->[0] eq '#define';
    if (ref($decl) eq 'ARRAY' && $decl->[0] eq '#define') {
        my $name = $decl->[1];
        $name =~ s/^MARPA_//;
        print {$const_fh} "\t".$name." = ".$decl->[2]."\n";
        next;
    }


    next if $decl->name->name =~ m/^_/;
    next if $decl->name->name =~ m/_ref$|_unref$/;
    next if $decl->name->name eq 'marpa_c_error';
    next if $decl->name->name eq 'marpa_g_error';

    #print {$fh} "/*\n".Dumper($decl)."\n*/\n";

    print {$fh} $decl->go_decl;
    print {$fh} "\n";
}

print {$const_fh} ")\n";
close $const_fh;
close $fh;
