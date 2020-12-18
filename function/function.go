package function

import "github.com/5anthosh/eval/evaluator/datatype"

//Function struct
type Function struct {
	Name         string
	Arity        int
	FunctionImpl func(argument []interface{}) (interface{}, error)
	ParamsType   []uint
}

//CheckNumberOfArgs in the function
func (f *Function) CheckNumberOfArgs(arguments []interface{}) bool {
	if f.Arity == -1 {
		return true
	}
	return f.Arity == -1 || f.Arity == len(arguments)
}

//CheckTypeOfArgs in the function
func (f *Function) CheckTypeOfArgs(arguments []interface{}) bool {
	if f.Arity == -1 {
		dtype := f.ParamsType[0]
		for _, arg := range arguments {
			if !datatype.Checkdatatype(arg, dtype) {
				return false
			}
		}
	}

	paramsLen := len(f.ParamsType)
	argsLen := len(arguments)
	for i := 0; i < paramsLen && i < argsLen; i++ {
		if !datatype.Checkdatatype(arguments[i], f.ParamsType[i]) {
			return false
		}
	}
	return true
}
