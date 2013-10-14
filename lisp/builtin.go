package lisp

import "fmt"
import "reflect"

type Builtin struct{}

var builtin = Builtin{}

var builtin_commands = map[string]string{
	"+":       "Add",
	"-":       "Sub",
	"*":       "Mul",
	">":       "Gt",
	"<":       "Lt",
	">=":      "Gte",
	"<=":      "Lte",
	"display": "Display",
}

func isBuiltin(cons Cons) bool {
	s := cons.car.String()
	if _, ok := builtin_commands[s]; ok {
		return true
	}
	return false
}

func runBuiltin(cons Cons) (val Value, err error) {
	cmd := builtin_commands[cons.car.String()]
	vars := cons.cdr.Cons().Sexp()
	values := []reflect.Value{}
	for _, i := range vars {
		if value, err := evalValue(i); err != nil {
			return Nil, err
		} else {
			values = append(values, reflect.ValueOf(value))
		}
	}
	result := reflect.ValueOf(&builtin).MethodByName(cmd).Call(values)
	val = result[0].Interface().(Value)
	err, _ = result[1].Interface().(error)
	return
}

func (Builtin) Display(vars ...Value) (Value, error) {
	if len(vars) == 1 {
		fmt.Println(vars[0])
	} else {
		return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	return Nil, nil
}

func (Builtin) Add(vars ...Value) (Value, error) {
	var sum float64
	for _, v := range vars {
		if v.typ == numberValue {
			sum += v.Number()
		} else {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Sub(vars ...Value) (Value, error) {
	if vars[0].typ != numberValue {
		return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.typ == numberValue {
			sum -= v.Number()
		} else {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Mul(vars ...Value) (Value, error) {
	if vars[0].typ != numberValue {
		return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
	}
	sum := vars[0].Number()
	for _, v := range vars[1:] {
		if v.typ == numberValue {
			sum *= v.Number()
		} else {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		}
	}
	return Value{numberValue, sum}, nil
}

func (Builtin) Gt(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() > v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Lt(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() < v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Gte(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() >= v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}

func (Builtin) Lte(vars ...Value) (Value, error) {
	for i := 1; i < len(vars); i++ {
		v1 := vars[i-1]
		v2 := vars[i]
		if v1.typ != numberValue || v2.typ != numberValue {
			return Nil, fmt.Errorf("Badly formatted arguments: %v", vars)
		} else if !(v1.Number() <= v2.Number()) {
			return False, nil
		}
	}
	return True, nil
}
