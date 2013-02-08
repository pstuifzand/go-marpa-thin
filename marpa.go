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
package marpa

/*
#cgo LDFLAGS: libmarpa.a
#include <marpa.h>
#include <marpa_api.h>
*/
import "C"

type Config struct {
	handle C.Marpa_Config
}

type Grammar struct {
	handle C.Marpa_Grammar
}

type Recognizer struct {
	handle C.Marpa_Recognizer
}

type Bocage struct {
	handle C.Marpa_Bocage
}

type Tree struct {
	handle C.Marpa_Tree
}

type Value struct {
	handle C.Marpa_Value
}

type Order struct {
	handle C.Marpa_Order
}
type Event struct {
	handle C.Marpa_Event
}

type Earleme int
type ErrorCode int
type EarleySetID int
type SymbolID int
type RuleID int
type StepType int
type EventType int
type Rank int
