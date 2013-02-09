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
#include <marpa.h>
#include <marpa_api.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

func MarpaCheckVersion(required_major uint, required_minor uint, required_micro uint) ErrorCode {
	return ErrorCode(C.marpa_check_version(C.uint(required_major), C.uint(required_minor), C.uint(required_micro)))
}

func ConfigInit(config *Config) int {
	return int(C.marpa_c_init(&config.handle))
}

func NewGrammar(configuration *Config) *Grammar {
	var _ret Grammar
	_ret.handle = C.marpa_g_new(&configuration.handle)
	if _ret.handle == (C.Marpa_Grammar)(unsafe.Pointer(uintptr(0))) {
		panic("ERROR: marpa_g_new")
	}
	return &_ret

}

func (g *Grammar) StartSymbol() SymbolID {
	return SymbolID(C.marpa_g_start_symbol(g.handle))
}

func (g *Grammar) StartSymbolSet(sym_id SymbolID) SymbolID {
	return SymbolID(C.marpa_g_start_symbol_set(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) HighestSymbolId() int {
	return int(C.marpa_g_highest_symbol_id(g.handle))
}

func (g *Grammar) SymbolIsAccessible(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_accessible(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolIsNullable(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_nullable(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolIsNulling(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_nulling(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolIsProductive(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_productive(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolIsStart(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_start(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolIsTerminal(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_terminal(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolIsTerminalSet(sym_id SymbolID, value int) int {
	return int(C.marpa_g_symbol_is_terminal_set(g.handle, C.Marpa_Symbol_ID(sym_id), C.int(value)))
}

func (g *Grammar) SymbolIsValued(symbol_id SymbolID) int {
	return int(C.marpa_g_symbol_is_valued(g.handle, C.Marpa_Symbol_ID(symbol_id)))
}

func (g *Grammar) SymbolIsValuedSet(symbol_id SymbolID, value int) int {
	return int(C.marpa_g_symbol_is_valued_set(g.handle, C.Marpa_Symbol_ID(symbol_id), C.int(value)))
}

func (g *Grammar) NewSymbol() SymbolID {
	return SymbolID(C.marpa_g_symbol_new(g.handle))
}

func (g *Grammar) HighestRuleId() int {
	return int(C.marpa_g_highest_rule_id(g.handle))
}

func (g *Grammar) RuleIsAccessible(rule_id RuleID) int {
	return int(C.marpa_g_rule_is_accessible(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) RuleIsNullable(ruleid RuleID) int {
	return int(C.marpa_g_rule_is_nullable(g.handle, C.Marpa_Rule_ID(ruleid)))
}

func (g *Grammar) RuleIsNulling(ruleid RuleID) int {
	return int(C.marpa_g_rule_is_nulling(g.handle, C.Marpa_Rule_ID(ruleid)))
}

func (g *Grammar) RuleIsLoop(rule_id RuleID) int {
	return int(C.marpa_g_rule_is_loop(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) RuleIsProductive(rule_id RuleID) int {
	return int(C.marpa_g_rule_is_productive(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) RuleLength(rule_id RuleID) int {
	return int(C.marpa_g_rule_length(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) RuleLhs(rule_id RuleID) SymbolID {
	return SymbolID(C.marpa_g_rule_lhs(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) NewRule(lhs_id SymbolID, rhs_ids []SymbolID) RuleID {
	i_rhs := make([]C.Marpa_Symbol_ID, len(rhs_ids))
	for i, id := range rhs_ids {
		i_rhs[i] = C.Marpa_Symbol_ID(id)
	}
	return RuleID(C.marpa_g_rule_new(g.handle, C.Marpa_Symbol_ID(lhs_id), &i_rhs[0], C.int(len(rhs_ids))))
}

func (g *Grammar) RuleRhs(rule_id RuleID, ix int) SymbolID {
	return SymbolID(C.marpa_g_rule_rhs(g.handle, C.Marpa_Rule_ID(rule_id), C.int(ix)))
}

func (g *Grammar) RuleIsProperSeparation(rule_id RuleID) int {
	return int(C.marpa_g_rule_is_proper_separation(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) SequenceMin(rule_id RuleID) int {
	return int(C.marpa_g_sequence_min(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) NewSequence(lhs_id SymbolID, rhs_id SymbolID, separator_id SymbolID, min int, flags int) RuleID {
	return RuleID(C.marpa_g_sequence_new(g.handle, C.Marpa_Symbol_ID(lhs_id), C.Marpa_Symbol_ID(rhs_id), C.Marpa_Symbol_ID(separator_id), C.int(min), C.int(flags)))
}

func (g *Grammar) SequenceSeparator(rule_id RuleID) int {
	return int(C.marpa_g_sequence_separator(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) SymbolIsCounted(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_counted(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) RuleRank(rule_id RuleID) Rank {
	return Rank(C.marpa_g_rule_rank(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) RuleRankSet(rule_id RuleID, rank Rank) Rank {
	return Rank(C.marpa_g_rule_rank_set(g.handle, C.Marpa_Rule_ID(rule_id), C.Marpa_Rank(rank)))
}

func (g *Grammar) RuleNullHigh(rule_id RuleID) int {
	return int(C.marpa_g_rule_null_high(g.handle, C.Marpa_Rule_ID(rule_id)))
}

func (g *Grammar) RuleNullHighSet(rule_id RuleID, flag int) int {
	return int(C.marpa_g_rule_null_high_set(g.handle, C.Marpa_Rule_ID(rule_id), C.int(flag)))
}

func (g *Grammar) Precompute() int {
	return int(C.marpa_g_precompute(g.handle))
}

func (g *Grammar) IsPrecomputed() int {
	return int(C.marpa_g_is_precomputed(g.handle))
}

func (g *Grammar) HasCycle() int {
	return int(C.marpa_g_has_cycle(g.handle))
}

func NewRecognizer(g *Grammar) *Recognizer {
	var _ret Recognizer
	_ret.handle = C.marpa_r_new(g.handle)
	if _ret.handle == (C.Marpa_Recognizer)(unsafe.Pointer(uintptr(0))) {
		panic("ERROR marpa_r_new")
	}
	return &_ret
}

func (r *Recognizer) StartInput() int {
	return int(C.marpa_r_start_input(r.handle))
}

func (r *Recognizer) Alternative(token_id SymbolID, value int, length int) int {
	return int(C.marpa_r_alternative(r.handle, C.Marpa_Symbol_ID(token_id), C.int(value), C.int(length)))
}

func (r *Recognizer) EarlemeComplete() Earleme {
	return Earleme(C.marpa_r_earleme_complete(r.handle))
}

func (r *Recognizer) CurrentEarleme() uint {
	return uint(C.marpa_r_current_earleme(r.handle))
}

func (r *Recognizer) Earleme(set_id EarleySetID) Earleme {
	return Earleme(C.marpa_r_earleme(r.handle, C.Marpa_Earley_Set_ID(set_id)))
}

func (r *Recognizer) EarleySetValue(earley_set EarleySetID) int {
	return int(C.marpa_r_earley_set_value(r.handle, C.Marpa_Earley_Set_ID(earley_set)))
}

func (r *Recognizer) FurthestEarleme() uint {
	return uint(C.marpa_r_furthest_earleme(r.handle))
}

func (r *Recognizer) LatestEarleySet() EarleySetID {
	return EarleySetID(C.marpa_r_latest_earley_set(r.handle))
}

func (r *Recognizer) LatestEarleySetValueSet(value int) int {
	return int(C.marpa_r_latest_earley_set_value_set(r.handle, C.int(value)))
}

func (r *Recognizer) EarleyItemWarningThreshold() int {
	return int(C.marpa_r_earley_item_warning_threshold(r.handle))
}

func (r *Recognizer) EarleyItemWarningThresholdSet(threshold int) int {
	return int(C.marpa_r_earley_item_warning_threshold_set(r.handle, C.int(threshold)))
}

func (r *Recognizer) ExpectedSymbolEventSet(symbol_id SymbolID, value int) int {
	return int(C.marpa_r_expected_symbol_event_set(r.handle, C.Marpa_Symbol_ID(symbol_id), C.int(value)))
}

func (r *Recognizer) TerminalsExpected(buffer *SymbolID) int {
	var sid C.Marpa_Symbol_ID
	ret := int(C.marpa_r_terminals_expected(r.handle, &sid))
	*buffer = SymbolID(sid)
	return ret
}

func (r *Recognizer) IsExhausted() int {
	return int(C.marpa_r_is_exhausted(r.handle))
}

func (r *Recognizer) ProgressReportStart(set_id EarleySetID) int {
	return int(C.marpa_r_progress_report_start(r.handle, C.Marpa_Earley_Set_ID(set_id)))
}

func (r *Recognizer) ProgressReportFinish() int {
	return int(C.marpa_r_progress_report_finish(r.handle))
}

func (r *Recognizer) ProgressItem(position *int, origin *EarleySetID) RuleID {
	var esid C.Marpa_Earley_Set_ID
	var pos C.int
	ret := RuleID(C.marpa_r_progress_item(r.handle, &pos, &esid))
	*origin = EarleySetID(esid)
	*position = int(pos)
	return ret
}

func NewBocage(r *Recognizer, earley_set_ID EarleySetID) (*Bocage, error) {
	var _ret Bocage
	_ret.handle = C.marpa_b_new(r.handle, C.Marpa_Earley_Set_ID(earley_set_ID))
	if _ret.handle == (C.Marpa_Bocage)(unsafe.Pointer(uintptr(0))) {
		return nil, errors.New("marpa_b_new failed")
	}
	return &_ret, nil
}

func NewOrder(b *Bocage) *Order {
	var _ret Order
	_ret.handle = C.marpa_o_new(b.handle)
	if _ret.handle == (C.Marpa_Order)(unsafe.Pointer(uintptr(0))) {
		panic("ERROR: marpa_o_new")
	}
	return &_ret
}

func (o *Order) HighRankOnlySet(flag int) int {
	return int(C.marpa_o_high_rank_only_set(o.handle, C.int(flag)))
}

func (o *Order) HighRankOnly() int {
	return int(C.marpa_o_high_rank_only(o.handle))
}

func (o *Order) Rank() int {
	return int(C.marpa_o_rank(o.handle))
}

func NewTree(o *Order) *Tree {
	var _ret Tree
	_ret.handle = C.marpa_t_new(o.handle)
	if _ret.handle == (C.Marpa_Tree)(unsafe.Pointer(uintptr(0))) {
		panic("ERROR: marpa_t_new")
	}
	return &_ret
}

func (t *Tree) Next() int {
	return int(C.marpa_t_next(t.handle))
}

func (t *Tree) ParseCount() int {
	return int(C.marpa_t_parse_count(t.handle))
}

func NewValue(t *Tree) *Value {
	var _ret Value
	_ret.handle = C.marpa_v_new(t.handle)
	return &_ret
}

func (v *Value) SymbolIsValued(sym_id SymbolID) int {
	return int(C.marpa_v_symbol_is_valued(v.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (v *Value) SymbolIsValuedSet(sym_id SymbolID, value int) int {
	return int(C.marpa_v_symbol_is_valued_set(v.handle, C.Marpa_Symbol_ID(sym_id), C.int(value)))
}

/*
func (v *Value) RuleIsValued(rule_id RuleID) int {
return int(C.marpa_v_rule_is_valued(v.handle, C.Marpa_Rule_ID(rule_id)))
}
*/

func (v *Value) RuleIsValuedSet(rule_id RuleID, value int) int {
	return int(C.marpa_v_rule_is_valued_set(v.handle, C.Marpa_Rule_ID(rule_id), C.int(value)))
}

func (v *Value) Step() StepType {
	return StepType(C.marpa_v_step(v.handle))
}

func (g *Grammar) Event(event *Event, ix int) EventType {
	var evt C.Marpa_Event
	ev_type := EventType(C.marpa_g_event(g.handle, &evt, C.int(ix)))
	event.handle = evt
	return ev_type
}

func (g *Grammar) EventCount() int {
	return int(C.marpa_g_event_count(g.handle))
}

func (g *Grammar) Error() ErrorCode {
	return ErrorCode(C.marpa_g_error(g.handle, (**C.char)(nil)))
}

func (g *Grammar) ErrorClear() ErrorCode {
	return ErrorCode(C.marpa_g_error_clear(g.handle))
}

func (g *Grammar) DefaultRank() Rank {
	return Rank(C.marpa_g_default_rank(g.handle))
}

func (g *Grammar) DefaultRankSet(rank Rank) Rank {
	return Rank(C.marpa_g_default_rank_set(g.handle, C.Marpa_Rank(rank)))
}

func (g *Grammar) SymbolRank(sym_id SymbolID) Rank {
	return Rank(C.marpa_g_symbol_rank(g.handle, C.Marpa_Symbol_ID(sym_id)))
}

func (g *Grammar) SymbolRankSet(sym_id SymbolID, rank Rank) Rank {
	return Rank(C.marpa_g_symbol_rank_set(g.handle, C.Marpa_Symbol_ID(sym_id), C.Marpa_Rank(rank)))
}
