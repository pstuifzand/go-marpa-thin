package marpa

func (value *Value) TokenID() SymbolID {
	return SymbolID(value.handle.t_token_id)
}

func (value *Value) TokenValue() int {
	return int(value.handle.t_token_value)
}
