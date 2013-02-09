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

func (value *Value) TokenID() SymbolID {
	return SymbolID(value.handle.t_token_id)
}
func (value *Value) Symbol() SymbolID {
	return SymbolID(value.handle.t_token_id)
}

func (value *Value) TokenValue() int {
	return int(value.handle.t_token_value)
}

func (value *Value) Result() int {
	return int(value.handle.t_arg_0)
}

func (value *Value) Rule() RuleID {
	return RuleID(value.handle.t_rule_id)
}

func (value *Value) Arg0() int {
	return int(value.handle.t_arg_0)
}
func (value *Value) ArgN() int {
	return int(value.handle.t_arg_n)
}
