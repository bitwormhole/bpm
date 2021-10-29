package convert

import "github.com/bitwormhole/starter/collection"

// Adapter 是 Entity 和 Properties 之间的适配器
type Adapter interface {
	ForString(p *string, name string)
	ForBool(p *bool, name string)
	ForByte(p *byte, name string)
	ForRune(p *rune, name string)

	ForInt(p *int, name string)
	ForInt8(p *int8, name string)
	ForInt16(p *int16, name string)
	ForInt32(p *int32, name string)
	ForInt64(p *int64, name string)

	ForUint(p *uint, name string)
	ForUint8(p *uint8, name string)
	ForUint16(p *uint16, name string)
	ForUint32(p *uint32, name string)
	ForUint64(p *uint64, name string)

	ForFloat32(p *float32, name string)
	ForFloat64(p *float64, name string)
}

// AdapterBuilder 用来创建 Adapter
type AdapterBuilder interface {
	GetterFor(props collection.Properties) AdapterBuilder
	SetterFor(props collection.Properties) AdapterBuilder
	Type(value string) AdapterBuilder
	ID(value string) AdapterBuilder
	Create() Adapter
}

////////////////////////////////////////////////////////////////////////////////

// NewAdapterBuilder 创建一个新的 AdapterBuilder
func NewAdapterBuilder() AdapterBuilder {
	return &adapterBuilderImpl{}
}

////////////////////////////////////////////////////////////////////////////////

type adapterBuilderImpl struct {
	id       string
	typeName string
	isSetter bool
	isGetter bool
	props    collection.Properties
}

func (inst *adapterBuilderImpl) _Impl() AdapterBuilder {
	return inst
}

func (inst *adapterBuilderImpl) GetterFor(props collection.Properties) AdapterBuilder {
	inst.props = props
	inst.isGetter = true
	inst.isSetter = false
	return inst
}

func (inst *adapterBuilderImpl) SetterFor(props collection.Properties) AdapterBuilder {
	inst.props = props
	inst.isGetter = false
	inst.isSetter = true
	return inst
}

func (inst *adapterBuilderImpl) Type(value string) AdapterBuilder {
	inst.typeName = value
	return inst
}

func (inst *adapterBuilderImpl) ID(value string) AdapterBuilder {
	inst.id = value
	return inst
}

func (inst *adapterBuilderImpl) Create() Adapter {

	keyPrefix := inst.typeName + "."
	id := inst.id
	if len(id) > 0 {
		keyPrefix = keyPrefix + id + "."
	}

	if inst.isSetter {
		setter := &setterAdapter{}
		setter.keyPrefix = keyPrefix
		setter.props = inst.props
		setter.setter = inst.props.Setter()
		return setter
	} else if inst.isGetter {
		getter := &getterAdapter{}
		getter.keyPrefix = keyPrefix
		getter.props = inst.props
		getter.getter = inst.props.Getter()
		return getter
	}
	panic("not Getter or Setter")
}

////////////////////////////////////////////////////////////////////////////////

type getterAdapter struct {
	props     collection.Properties
	getter    collection.PropertyGetter
	keyPrefix string
}

func (inst *getterAdapter) _Impl() Adapter {
	return inst
}

func (inst *getterAdapter) ForString(p *string, name string) {
	*p = inst.getter.GetString(inst.keyPrefix+name, "")
}

func (inst *getterAdapter) ForBool(p *bool, name string) {
	*p = inst.getter.GetBool(inst.keyPrefix+name, false)

}

func (inst *getterAdapter) ForByte(p *byte, name string) {
	*p = inst.getter.GetUint8(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForRune(p *rune, name string) {
	*p = inst.getter.GetInt32(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForInt(p *int, name string) {
	*p = inst.getter.GetInt(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForInt8(p *int8, name string) {
	*p = inst.getter.GetInt8(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForInt16(p *int16, name string) {
	*p = inst.getter.GetInt16(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForInt32(p *int32, name string) {
	*p = inst.getter.GetInt32(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForInt64(p *int64, name string) {
	*p = inst.getter.GetInt64(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForUint(p *uint, name string) {
	*p = inst.getter.GetUint(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForUint8(p *uint8, name string) {
	*p = inst.getter.GetUint8(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForUint16(p *uint16, name string) {
	*p = inst.getter.GetUint16(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForUint32(p *uint32, name string) {
	*p = inst.getter.GetUint32(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForUint64(p *uint64, name string) {
	*p = inst.getter.GetUint64(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForFloat32(p *float32, name string) {
	*p = inst.getter.GetFloat32(inst.keyPrefix+name, 0)
}

func (inst *getterAdapter) ForFloat64(p *float64, name string) {
	*p = inst.getter.GetFloat64(inst.keyPrefix+name, 0)
}

////////////////////////////////////////////////////////////////////////////////

type setterAdapter struct {
	props     collection.Properties
	setter    collection.PropertySetter
	keyPrefix string
}

func (inst *setterAdapter) _Impl() Adapter {
	return inst
}

func (inst *setterAdapter) ForString(p *string, name string) {
	inst.setter.SetString(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForBool(p *bool, name string) {
	inst.setter.SetBool(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForByte(p *byte, name string) {
	inst.setter.SetUint8(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForRune(p *rune, name string) {
	inst.setter.SetInt32(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForInt(p *int, name string) {
	inst.setter.SetInt(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForInt8(p *int8, name string) {
	inst.setter.SetInt8(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForInt16(p *int16, name string) {
	inst.setter.SetInt16(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForInt32(p *int32, name string) {
	inst.setter.SetInt32(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForInt64(p *int64, name string) {
	inst.setter.SetInt64(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForUint(p *uint, name string) {
	inst.setter.SetUint(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForUint8(p *uint8, name string) {
	inst.setter.SetUint8(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForUint16(p *uint16, name string) {
	inst.setter.SetUint16(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForUint32(p *uint32, name string) {
	inst.setter.SetUint32(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForUint64(p *uint64, name string) {
	inst.setter.SetUint64(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForFloat32(p *float32, name string) {
	inst.setter.SetFloat32(inst.keyPrefix+name, *p)
}

func (inst *setterAdapter) ForFloat64(p *float64, name string) {
	inst.setter.SetFloat64(inst.keyPrefix+name, *p)
}

////////////////////////////////////////////////////////////////////////////////
