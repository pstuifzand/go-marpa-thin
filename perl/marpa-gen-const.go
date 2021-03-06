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

const (
	ERROR_COUNT                      = 92
	ERR_NONE                         = 0
	ERR_AHFA_IX_NEGATIVE             = 1
	ERR_AHFA_IX_OOB                  = 2
	ERR_ANDID_NEGATIVE               = 3
	ERR_ANDID_NOT_IN_OR              = 4
	ERR_ANDIX_NEGATIVE               = 5
	ERR_BAD_SEPARATOR                = 6
	ERR_BOCAGE_ITERATION_EXHAUSTED   = 7
	ERR_COUNTED_NULLABLE             = 8
	ERR_DEVELOPMENT                  = 9
	ERR_DUPLICATE_AND_NODE           = 10
	ERR_DUPLICATE_RULE               = 11
	ERR_DUPLICATE_TOKEN              = 12
	ERR_EIM_COUNT                    = 13
	ERR_EIM_ID_INVALID               = 14
	ERR_EVENT_IX_NEGATIVE            = 15
	ERR_EVENT_IX_OOB                 = 16
	ERR_GRAMMAR_HAS_CYCLE            = 17
	ERR_INACCESSIBLE_TOKEN           = 18
	ERR_INTERNAL                     = 19
	ERR_INVALID_AHFA_ID              = 20
	ERR_INVALID_AIMID                = 21
	ERR_INVALID_BOOLEAN              = 22
	ERR_INVALID_IRLID                = 23
	ERR_INVALID_ISYID                = 24
	ERR_INVALID_LOCATION             = 25
	ERR_INVALID_RULE_ID              = 26
	ERR_INVALID_START_SYMBOL         = 27
	ERR_INVALID_SYMBOL_ID            = 28
	ERR_I_AM_NOT_OK                  = 29
	ERR_MAJOR_VERSION_MISMATCH       = 30
	ERR_MICRO_VERSION_MISMATCH       = 31
	ERR_MINOR_VERSION_MISMATCH       = 32
	ERR_NOOKID_NEGATIVE              = 33
	ERR_NOT_PRECOMPUTED              = 34
	ERR_NOT_TRACING_COMPLETION_LINKS = 35
	ERR_NOT_TRACING_LEO_LINKS        = 36
	ERR_NOT_TRACING_TOKEN_LINKS      = 37
	ERR_NO_AND_NODES                 = 38
	ERR_NO_EARLEY_SET_AT_LOCATION    = 39
	ERR_NO_OR_NODES                  = 40
	ERR_NO_PARSE                     = 41
	ERR_NO_RULES                     = 42
	ERR_NO_START_SYMBOL              = 43
	ERR_NO_TOKEN_EXPECTED_HERE       = 44
	ERR_NO_TRACE_EIM                 = 45
	ERR_NO_TRACE_ES                  = 46
	ERR_NO_TRACE_PIM                 = 47
	ERR_NO_TRACE_SRCL                = 48
	ERR_NULLING_TERMINAL             = 49
	ERR_ORDER_FROZEN                 = 50
	ERR_ORID_NEGATIVE                = 51
	ERR_OR_ALREADY_ORDERED           = 52
	ERR_PARSE_EXHAUSTED              = 53
	ERR_PARSE_TOO_LONG               = 54
	ERR_PIM_IS_NOT_LIM               = 55
	ERR_POINTER_ARG_NULL             = 56
	ERR_PRECOMPUTED                  = 57
	ERR_PROGRESS_REPORT_EXHAUSTED    = 58
	ERR_PROGRESS_REPORT_NOT_STARTED  = 59
	ERR_RECCE_NOT_ACCEPTING_INPUT    = 60
	ERR_RECCE_NOT_STARTED            = 61
	ERR_RECCE_STARTED                = 62
	ERR_RHS_IX_NEGATIVE              = 63
	ERR_RHS_IX_OOB                   = 64
	ERR_RHS_TOO_LONG                 = 65
	ERR_SEQUENCE_LHS_NOT_UNIQUE      = 66
	ERR_SOURCE_TYPE_IS_AMBIGUOUS     = 67
	ERR_SOURCE_TYPE_IS_COMPLETION    = 68
	ERR_SOURCE_TYPE_IS_LEO           = 69
	ERR_SOURCE_TYPE_IS_NONE          = 70
	ERR_SOURCE_TYPE_IS_TOKEN         = 71
	ERR_SOURCE_TYPE_IS_UNKNOWN       = 72
	ERR_START_NOT_LHS                = 73
	ERR_SYMBOL_VALUED_CONFLICT       = 74
	ERR_TERMINAL_IS_LOCKED           = 75
	ERR_TOKEN_IS_NOT_TERMINAL        = 76
	ERR_TOKEN_LENGTH_LE_ZERO         = 77
	ERR_TOKEN_TOO_LONG               = 78
	ERR_TREE_EXHAUSTED               = 79
	ERR_TREE_PAUSED                  = 80
	ERR_UNEXPECTED_TOKEN_ID          = 81
	ERR_UNPRODUCTIVE_START           = 82
	ERR_VALUATOR_INACTIVE            = 83
	ERR_VALUED_IS_LOCKED             = 84
	ERR_RANK_TOO_LOW                 = 85
	ERR_RANK_TOO_HIGH                = 86
	ERR_SYMBOL_IS_NULLING            = 87
	ERR_SYMBOL_IS_UNUSED             = 88
	ERR_NO_SUCH_RULE_ID              = 89
	ERR_NO_SUCH_SYMBOL_ID            = 90
	ERR_BEFORE_FIRST_TREE            = 91
	EVENT_COUNT                      = 7
	EVENT_NONE                       = 0
	EVENT_COUNTED_NULLABLE           = 1
	EVENT_EARLEY_ITEM_THRESHOLD      = 2
	EVENT_EXHAUSTED                  = 3
	EVENT_LOOP_RULES                 = 4
	EVENT_NULLING_TERMINAL           = 5
	EVENT_SYMBOL_EXPECTED            = 6
	STEP_COUNT                       = 8
	STEP_INTERNAL1                   = 0
	STEP_RULE                        = 1
	STEP_TOKEN                       = 2
	STEP_NULLING_SYMBOL              = 3
	STEP_TRACE                       = 4
	STEP_INACTIVE                    = 5
	STEP_INTERNAL2                   = 6
	STEP_INITIAL                     = 7
)
