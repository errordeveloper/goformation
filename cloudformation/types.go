package cloudformation

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	AccountID        = "AWS::AccountId"
	NotificationARNs = "AWS::NotificationARNs"
	NoValue          = "AWS::NoValue"
	Partition        = "AWS::Partition"
	Region           = "AWS::Region"
	StackID          = "AWS::StackId"
	StackName        = "AWS::StackName"

	Ref           = "Ref"
	FnBase64      = "Fn::Base64"
	FnCIDR        = "Fn::Cidr"
	FnAnd         = "Fn::And"
	FnEquals      = "Fn::Equals"
	FnIf          = "Fn::If"
	FnNot         = "Fn::Not"
	FnOr          = "Fn::Or"
	FnFindInMap   = "Fn::FindInMap"
	FnGetAtt      = "Fn::GetAtt"
	FnGetAZs      = "Fn::GetAZs"
	FnImportValue = "Fn::ImportValue"
	FnJoin        = "Fn::Join"
	FnSub         = "Fn::Sub"
	FnSelect      = "Fn::Select"
	FnSplit       = "Fn::Split"
)

var (
	RefAccountID        = MakeRef(AccountID)
	RefNotificationARNs = MakeRef(NotificationARNs)
	RefNoValue          = MakeRef(NoValue)
	RefPartition        = MakeRef(Partition)
	RefRegion           = MakeRef(Region)
	RefStackID          = MakeRef(StackID)
	RefStackName        = MakeRef(StackName)
)

type Value struct {
	value json.Marshaler
}

func (v *Value) UnmarshalJSON(b []byte) error {
	var obj interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	switch obj.(type) {
	case string:
		v.value = String(obj.(string))
		return nil
	case float64:
		v.value = Double(obj.(float64))
	case bool:
		v.value = Boolean(obj.(bool))
	default:
		return fmt.Errorf("cannot handle type %s", reflect.ValueOf(obj).Kind())
	}

	return nil
}

type String string

func NewString(v string) *Value { return &Value{String(v)} }

func (v String) MarshalJSON() ([]byte, error) {
	x := string(v)
	return json.Marshal(&x)
}

// func (v String) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, v) }

type StringSlice struct {
	s []string
}

func NewStringSlice(v ...string) *Value { return &Value{StringSlice{v}} }

func (v StringSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(&v.s)
}

// func (v StringSlice) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &v.s) }

type Long int64

func (v Long) MarshalJSON() ([]byte, error) {
	x := int64(v)
	return json.Marshal(&x)
}

// func (v Long) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &v) }

type Integer int

func NewInteger(v int) *Value { return &Value{Integer(v)} }

func (v Integer) MarshalJSON() ([]byte, error) {
	x := int(v)
	return json.Marshal(&x)
}

// func (v Integer) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &v) }

type Double float64

func NewDouble(v float64) *Value { return &Value{Double(v)} }

func (v Double) MarshalJSON() ([]byte, error) {
	x := float64(v)
	return json.Marshal(&x)
}

// func (v Double) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &v) }

type Boolean bool

var (
	trueValue  = NewBoolean(true)
	falseValue = NewBoolean(false)
)

func True() *Value  { return trueValue }
func False() *Value { return falseValue }

func NewBoolean(v bool) *Value { return &Value{Boolean(v)} }

func (v Boolean) MarshalJSON() ([]byte, error) {
	x := bool(v)
	return json.Marshal(&x)
}

// func (v Boolean) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &v) }

type Intrinsic struct {
	Value map[string]interface{}
}

func MakeIntrinsic(k string, v interface{}) Intrinsic {
	return Intrinsic{
		map[string]interface{}{
			k: v,
		},
	}
}

func MakeRef(r string) Intrinsic { return MakeIntrinsic(Ref, r) }

// TODO MakeFnBase64
// TODO MakeFnCIDR
// TODO MakeFnAnd
// TODO MakeFnEquals
// TODO MakeFnIf
// TODO MakeFnNot
// TODO MakeFnOr
// TODO MakeFnFindInMap

func MakeFnGetAtt(arg *Value) Intrinsic       { return MakeIntrinsic(FnGetAtt, arg) }
func MakeFnGetAttString(arg string) Intrinsic { return MakeFnGetAtt(NewString(arg)) }

// TODO MakeFnGetAZs

func MakeFnImportValue(arg *Value) Intrinsic       { return MakeIntrinsic(FnImportValue, arg) }
func MakeFnImportValueString(arg string) Intrinsic { return MakeFnImportValue(NewString(arg)) }

func MakeFnJoin(sep string, args []*Value) Intrinsic {
	return MakeIntrinsic(FnJoin,
		[]interface{}{
			sep,
			args,
		},
	)
}

func MakeFnSub(arg *Value) Intrinsic       { return MakeIntrinsic(FnSub, arg) }
func MakeFnSubString(arg string) Intrinsic { return MakeFnSub(NewString(arg)) }

// TODO MakeFnSelect

func MakeFnSplit(sep string, arg *Value) Intrinsic {
	return MakeIntrinsic(FnSplit,
		[]interface{}{
			sep,
			arg,
		},
	)
}

func MakeFnSplitString(sep string, arg string) Intrinsic {
	return MakeFnSplit(sep, NewString(arg))
}

func (v Intrinsic) MarshalJSON() ([]byte, error) { return json.Marshal(&v.Value) }

// func (v Intrinsic) UnmarshalJSON(b []byte) error { return json.Unmarshal(b, &v.Value) }
