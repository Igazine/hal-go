package ext

import (
	"math"
	"github.com/Igazine/hank-go"
)

const safeIntMax = 9007199254740991.0

func checkSafeInt(n float64, taskName string) (int64, *hank.Value) {
	if math.Abs(n) > safeIntMax || math.IsInf(n, 0) || math.IsNaN(n) {
		return 0, &hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.BitwiseOutOfBounds, Args: []hank.Value{{Type: hank.TypeNumber, Number: n}, {Type: hank.TypeString, String: taskName}}}}
	}
	return int64(n), nil
}

func fromSafeInt(n int64, taskName string) hank.Value {
	f := float64(n)
	if math.Abs(f) > safeIntMax {
		return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.BitwiseOutOfBounds, Args: []hank.Value{{Type: hank.TypeNumber, Number: f}, {Type: hank.TypeString, String: taskName}}}}
	}
	return hank.Value{Type: hank.TypeNumber, Number: f}
}

type PlatformExtension struct{}

func (e *PlatformExtension) Name() string {
	return "PlatformExtension"
}

func (e *PlatformExtension) GetTasks() map[string]hank.NativeFunc {
	tasks := make(map[string]hank.NativeFunc)

	tasks["bin_and"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		var a, b float64
		if len(args) < 2 { return hank.Value{Type: hank.TypeVoid} }
		if args[0].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_and"}}}}
		}
		if args[1].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_and"}}}}
		}
		a = args[0].Number
		b = args[1].Number

		ia, err := checkSafeInt(a, "bin_and")
		if err != nil { return *err }
		ib, err := checkSafeInt(b, "bin_and")
		if err != nil { return *err }
		return fromSafeInt(ia & ib, "bin_and")
	}

	tasks["bin_or"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		var a, b float64
		if len(args) < 2 { return hank.Value{Type: hank.TypeVoid} }
		if args[0].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_or"}}}}
		}
		if args[1].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_or"}}}}
		}
		a = args[0].Number
		b = args[1].Number

		ia, err := checkSafeInt(a, "bin_or")
		if err != nil { return *err }
		ib, err := checkSafeInt(b, "bin_or")
		if err != nil { return *err }
		return fromSafeInt(ia | ib, "bin_or")
	}

	tasks["bin_xor"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		var a, b float64
		if len(args) < 2 { return hank.Value{Type: hank.TypeVoid} }
		if args[0].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_xor"}}}}
		}
		if args[1].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_xor"}}}}
		}
		a = args[0].Number
		b = args[1].Number

		ia, err := checkSafeInt(a, "bin_xor")
		if err != nil { return *err }
		ib, err := checkSafeInt(b, "bin_xor")
		if err != nil { return *err }
		return fromSafeInt(ia ^ ib, "bin_xor")
	}

	tasks["bin_not"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) == 0 || args[0].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_not"}}}}
		}
		a := args[0].Number
		ia, err := checkSafeInt(a, "bin_not")
		if err != nil { return *err }
		return fromSafeInt(^ia, "bin_not")
	}

	tasks["bin_shiftL"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) < 2 { return hank.Value{Type: hank.TypeVoid} }
		if args[0].Type != hank.TypeNumber || args[1].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_shiftL"}}}}
		}
		a := args[0].Number
		b := args[1].Number

		ia, err := checkSafeInt(a, "bin_shiftL")
		if err != nil { return *err }
		return fromSafeInt(ia << uint(b), "bin_shiftL")
	}

	tasks["bin_shiftR"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) < 2 { return hank.Value{Type: hank.TypeVoid} }
		if args[0].Type != hank.TypeNumber || args[1].Type != hank.TypeNumber {
			return hank.Value{Type: hank.TypeError, Error: &hank.ErrorValue{Code: hank.TypeMismatch, Args: []hank.Value{{Type: hank.TypeString, String: "Number"}, {Type: hank.TypeString, String: "Any"}, {Type: hank.TypeString, String: "bin_shiftR"}}}}
		}
		a := args[0].Number
		b := args[1].Number

		ia, err := checkSafeInt(a, "bin_shiftR")
		if err != nil { return *err }
		return fromSafeInt(ia >> uint(b), "bin_shiftR")
	}

	return tasks
}
