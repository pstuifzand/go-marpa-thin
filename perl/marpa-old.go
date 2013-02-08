/*
Copyright (C) 2013 Peter Stuifzand
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
#cgo LDFLAGS: -L. -lmarpa
#include <marpa.h>
#include <marpa_api.h>
*/
import "C"

type Config struct {
	config C.Marpa_Config
}
type Grammar struct {
	grammar C.Marpa_Grammar
}

type SymbolID struct {
	id C.Marpa_Symbol_ID
}

type RuleID struct {
	id C.Marpa_Rule_ID
}

func NewConfig() Config {
	var config Config
	C.marpa_c_init(&config.config)
	return config
}

func NewGrammar(config *Config) Grammar {
	var grammar Grammar
	grammar.grammar = C.marpa_g_new(&config.config)
	return grammar
}

func (g Grammar) StartSymbol() SymbolID {
	var symbol SymbolID
	symbol.id = C.marpa_g_start_symbol(g.grammar)
	return symbol
}

func (g Grammar) StartSymbolSet(sym SymbolID) SymbolID {
	var symbol SymbolID
	symbol.id = C.marpa_g_start_symbol_set(g.grammar, sym.id)
	return symbol
}

func (g Grammar) HighestSymbolID() int {
	return int(C.marpa_g_highest_symbol_id(g.grammar))
}

func (g Grammar) SymbolIsAccessible(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_accessible(g.grammar, sym_id.id))
}

func (g Grammar) SymbolIsNullable(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_nullable(g.grammar, sym_id.id))
}

func (g Grammar) SymbolIsNulling(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_nulling(g.grammar, sym_id.id))
}

func (g Grammar) SymbolIsProductive(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_productive(g.grammar, sym_id.id))
}

func (g Grammar) SymbolIsStart(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_start(g.grammar, sym_id.id))
}

func (g Grammar) SymbolIsTerminal(sym_id SymbolID) int {
	return int(C.marpa_g_symbol_is_terminal(g.grammar, sym_id.id))
}

func (g Grammar) SymbolIsTerminalSet(sym_id SymbolID, val int) int {
	return int(C.marpa_g_symbol_is_terminal_set(g.grammar, sym_id.id, C.int(val)))
}

func (g Grammar) SymbolNew() SymbolID {
	var symbol SymbolID
	symbol.id = C.marpa_g_symbol_new(g.grammar)
	return symbol
}

func (g Grammar) HighestRule() int {
	return int(C.marpa_g_highest_rule_id(g.grammar))
}

func (g Grammar) RuleIsAccessible(sym_id RuleID) int {
	return int(C.marpa_g_rule_is_accessible(g.grammar, sym_id.id))
}

func (g Grammar) RuleIsNullable(sym_id RuleID) int {
	return int(C.marpa_g_rule_is_nullable(g.grammar, sym_id.id))
}

func (g Grammar) RuleIsNulling(sym_id RuleID) int {
	return int(C.marpa_g_rule_is_nulling(g.grammar, sym_id.id))
}

func (g Grammar) RuleIsLoop(sym_id RuleID) int {
	return int(C.marpa_g_rule_is_loop(g.grammar, sym_id.id))
}

func (g Grammar) RuleIsProductive(sym_id RuleID) int {
	return int(C.marpa_g_rule_is_productive(g.grammar, sym_id.id))
}

func (g Grammar) RuleLength(sym_id RuleID) int {
	return int(C.marpa_g_rule_length(g.grammar, sym_id.id))
}

/*
Marpa_Error_Code marpa_check_version (unsigned int required_major, unsigned int required_minor, unsigned int required_micro );
int marpa_c_init ( Marpa_Config* config);
Marpa_Error_Code marpa_c_error ( Marpa_Config* config, const char** p_error_string );
Marpa_Grammar marpa_g_new ( Marpa_Config* configuration );
Marpa_Grammar marpa_g_ref (Marpa_Grammar g);
void marpa_g_unref (Marpa_Grammar g);
Marpa_Symbol_ID marpa_g_start_symbol (Marpa_Grammar g);
Marpa_Symbol_ID marpa_g_start_symbol_set ( Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_highest_symbol_id (Marpa_Grammar g);
int marpa_g_symbol_is_accessible (Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_symbol_is_nullable ( Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_symbol_is_nulling (Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_symbol_is_productive (Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_symbol_is_start ( Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_symbol_is_terminal ( Marpa_Grammar g, Marpa_Symbol_ID sym_id);
int marpa_g_symbol_is_terminal_set ( Marpa_Grammar g, Marpa_Symbol_ID sym_id, int value);
int marpa_g_symbol_is_valued ( Marpa_Grammar g, Marpa_Symbol_ID symbol_id);
int marpa_g_symbol_is_valued_set ( Marpa_Grammar g, Marpa_Symbol_ID symbol_id, int value);
Marpa_Symbol_ID marpa_g_symbol_new (Marpa_Grammar g);
int marpa_g_highest_rule_id (Marpa_Grammar g);
int marpa_g_rule_is_accessible (Marpa_Grammar g, Marpa_Rule_ID rule_id);
int marpa_g_rule_is_nullable ( Marpa_Grammar g, Marpa_Rule_ID ruleid);
int marpa_g_rule_is_nulling (Marpa_Grammar g, Marpa_Rule_ID ruleid);
int marpa_g_rule_is_loop (Marpa_Grammar g, Marpa_Rule_ID rule_id);
int marpa_g_rule_is_productive (Marpa_Grammar g, Marpa_Rule_ID rule_id);
int marpa_g_rule_length ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
Marpa_Symbol_ID marpa_g_rule_lhs ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
Marpa_Rule_ID marpa_g_rule_new (Marpa_Grammar g, Marpa_Symbol_ID lhs_id, Marpa_Symbol_ID *rhs_ids, int length);
Marpa_Symbol_ID marpa_g_rule_rhs ( Marpa_Grammar g, Marpa_Rule_ID rule_id, int ix);
int marpa_g_rule_is_proper_separation ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
int marpa_g_sequence_min ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
Marpa_Rule_ID marpa_g_sequence_new (Marpa_Grammar g, Marpa_Symbol_ID lhs_id, Marpa_Symbol_ID rhs_id, Marpa_Symbol_ID separator_id, int min, int flags );
int marpa_g_sequence_separator ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
int marpa_g_symbol_is_counted (Marpa_Grammar g, Marpa_Symbol_ID sym_id);
Marpa_Rank marpa_g_rule_rank ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
Marpa_Rank marpa_g_rule_rank_set ( Marpa_Grammar g, Marpa_Rule_ID rule_id, Marpa_Rank rank);
int marpa_g_rule_null_high ( Marpa_Grammar g, Marpa_Rule_ID rule_id);
int marpa_g_rule_null_high_set ( Marpa_Grammar g, Marpa_Rule_ID rule_id, int flag);
int marpa_g_precompute (Marpa_Grammar g);
int marpa_g_is_precomputed (Marpa_Grammar g);
int marpa_g_has_cycle (Marpa_Grammar g);
Marpa_Recognizer marpa_r_new ( Marpa_Grammar g );
Marpa_Recognizer marpa_r_ref (Marpa_Recognizer r);
void marpa_r_unref (Marpa_Recognizer r);
int marpa_r_start_input (Marpa_Recognizer r);
int marpa_r_alternative (Marpa_Recognizer r, Marpa_Symbol_ID token_id, int value, int length);
Marpa_Earleme marpa_r_earleme_complete (Marpa_Recognizer r);
unsigned int marpa_r_current_earleme (Marpa_Recognizer r);
Marpa_Earleme marpa_r_earleme ( Marpa_Recognizer r, Marpa_Earley_Set_ID set_id);
int marpa_r_earley_set_value ( Marpa_Recognizer r, Marpa_Earley_Set_ID earley_set);
unsigned int marpa_r_furthest_earleme (Marpa_Recognizer r);
Marpa_Earley_Set_ID marpa_r_latest_earley_set (Marpa_Recognizer r);
int marpa_r_latest_earley_set_value_set ( Marpa_Recognizer r, int value);
int marpa_r_earley_item_warning_threshold (Marpa_Recognizer r);
int marpa_r_earley_item_warning_threshold_set (Marpa_Recognizer r, int threshold);
int marpa_r_expected_symbol_event_set ( Marpa_Recognizer r, Marpa_Symbol_ID symbol_id, int value);
int marpa_r_terminals_expected ( Marpa_Recognizer r, Marpa_Symbol_ID* buffer);
int marpa_r_is_exhausted (Marpa_Recognizer r);
int marpa_r_progress_report_start ( Marpa_Recognizer r, Marpa_Earley_Set_ID set_id);
int marpa_r_progress_report_finish ( Marpa_Recognizer r );
Marpa_Rule_ID marpa_r_progress_item ( Marpa_Recognizer r, int* position, Marpa_Earley_Set_ID* origin );
Marpa_Bocage marpa_b_new (Marpa_Recognizer r, Marpa_Earley_Set_ID earley_set_ID);
Marpa_Bocage marpa_b_ref (Marpa_Bocage b);
void marpa_b_unref (Marpa_Bocage b);
Marpa_Order marpa_o_new ( Marpa_Bocage b);
Marpa_Order marpa_o_ref ( Marpa_Order o);
void marpa_o_unref ( Marpa_Order o);
int marpa_o_high_rank_only_set ( Marpa_Order o, int flag);
int marpa_o_high_rank_only ( Marpa_Order o);
int marpa_o_rank ( Marpa_Order o );
Marpa_Tree marpa_t_new (Marpa_Order o);
Marpa_Tree marpa_t_ref (Marpa_Tree t);
void marpa_t_unref (Marpa_Tree t);
int marpa_t_next ( Marpa_Tree t);
int marpa_t_parse_count ( Marpa_Tree t);
Marpa_Value marpa_v_new ( Marpa_Tree t );
Marpa_Value marpa_v_ref (Marpa_Value v);
void marpa_v_unref ( Marpa_Value v);
int marpa_v_symbol_is_valued ( Marpa_Value v, Marpa_Symbol_ID sym_id );
int marpa_v_symbol_is_valued_set ( Marpa_Value v, Marpa_Symbol_ID sym_id, int value );
int marpa_v_rule_is_valued ( Marpa_Value v, Marpa_Rule_ID rule_id );
int marpa_v_rule_is_valued_set ( Marpa_Value v, Marpa_Rule_ID rule_id, int value );
Marpa_Step_Type marpa_v_step ( Marpa_Value v);
Marpa_Event_Type marpa_g_event (Marpa_Grammar g, Marpa_Event* event, int ix);
int marpa_g_event_count ( Marpa_Grammar g );
Marpa_Error_Code marpa_g_error ( Marpa_Grammar g, const char** p_error_string);
Marpa_Error_Code marpa_g_error_clear ( Marpa_Grammar g );
Marpa_Rank marpa_g_default_rank ( Marpa_Grammar g);
Marpa_Rank marpa_g_default_rank_set ( Marpa_Grammar g, Marpa_Rank rank);
Marpa_Rank marpa_g_symbol_rank ( Marpa_Grammar g, Marpa_Symbol_ID sym_id);
Marpa_Rank marpa_g_symbol_rank_set ( Marpa_Grammar g, Marpa_Symbol_ID sym_id, Marpa_Rank rank);
*/
