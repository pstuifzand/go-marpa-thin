go-marpa-thin
=============

Go module for libmarpa


FILE STRUCTURE
--------------

* marpa.go - types
* marpa-gen.go - generated but edited function from marpa\_api.h
* marpa-value.go - extra functions that are not ports of the library
* perl/ - directory contains the generator

LIBMARPA
--------

You need a recent version of libmarpa.a. I did the following to get one.

Download a recent version Marpa::R2 and unpack it. Build the Marpa::R2 module
and copy the libmarpa.a file to the directory where you unpacked everything.

BUILD
-----

Use `go build` to build the module.

LICENSE
-------
LGPL3+

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Lesser General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Lesser General Public License for more details.

    You should have received a copy of the GNU Lesser General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
