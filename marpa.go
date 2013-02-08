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
