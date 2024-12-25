package registry

type Keyable interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | string
}

type AbstractRegistry[Key Keyable, Value any] struct {
	data map[Key]Value
}

func (registry *AbstractRegistry[Key, Value]) Save(key Key, value Value) *AbstractRegistry[Key, Value] {
	registry.data[key] = value
	return registry
}

func (registry *AbstractRegistry[Key, Value]) Extract(key Key) Value {
	return registry.data[key]
}

func (registry *AbstractRegistry[Key, Value]) Bootstrap() *AbstractRegistry[Key, Value] {
	return registry
}
