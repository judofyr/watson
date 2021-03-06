package vm

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/types"
)

func TestFeedInewPushesZero(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Inew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(0)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIincIncrementsTheValue(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Inew)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iinc)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(1)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIincFailsWhenTypeMismatch(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Nnew)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iinc)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIincFailsIfStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Iinc)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIshlShiftsTheTopBy1(t *testing.T) {
	var err error
	vm := NewVM()

	var before int64 = 123

	err = vm.pushInt(before)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Ishl)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(before * 2)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIshlFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Ishl)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedIshlFailsWhenTypeMismatch(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Ishl)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIaddAddsTwoIntegers(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(2)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(3)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIaddFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Iadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIaddFailsWhenStackIsInsufficiient(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIaddFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIaddFailsWhenArg2IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedInegNegatesTheTop(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Ineg)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(-1)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedInegFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Ineg)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedInegFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Ineg)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIshtShiftsArg2ToLeftByArg1WhenArg1IsPositive(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(0xabcd0)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(4)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Isht)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(0xabcd00)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIshtShiftsArg2ToRightByArg1WhenArg1IsNegative(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(0xabcd0)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(-4)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Isht)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(0xabcd)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIshtFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Isht)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIshtFailsWhenStackIsInsufficient(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Isht)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIshtFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(0xabcd0)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Isht)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIshtFailsWhenArg2IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(4)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Isht)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedItofConvertsArg1ToFloat(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(int64(math.Float64bits(1.234e-56)))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Itof)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewFloatValue(1.234e-56)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedItofFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Itof)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedItofFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Itof)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedItouConvertsArg1ToUint(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(-1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Itou)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewUintValue(0xffffffffffffffff)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedItouFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Itou)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedItouFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Itou)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedFinfPushesPositiveInf(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Finf)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewFloatValue(math.Inf(1))
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedFinfPushesNaN(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Fnan)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if !got.IsNaN() {
		t.Errorf("expected NaN but got %+v", got)
	}
}

func TestFeedFnegNegatesArg1(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushFloat(9.87456e78)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Fneg)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewFloatValue(-9.87456e78)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedFnegCanNegateInf(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushFloat(math.Inf(1))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Fneg)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewFloatValue(math.Inf(-1))
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedFnegFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Fneg)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedFnegFailsWhenArg1IsNotFloat(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Fneg)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedSnewPushesEmptyString(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Snew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewStringValue([]byte{})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedSaddAddsACharToString(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushString([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(0x21) // '!'
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewStringValue([]byte("hello!"))
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedSaddFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Sadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedSaddFailsWhenStackIsInsufficiient(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedSaddFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushString([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedSaddFailsWhenArg2IsNotString(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(0x21) // '!'
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedOnewPushesEmptyObject(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Onew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewObjectValue(map[string]*types.Value{})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedOaddAddsAKeyValuePairToAnObject(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushObject(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushString([]byte("year"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(2021)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Oadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
		"year":  types.NewIntValue(2021),
	})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedOaddAddsACopyOfAValue(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushObject(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushString([]byte("user"))
	if err != nil {
		t.Fatal(err)
	}
	addedVal := map[string]*types.Value{
		"name": types.NewStringValue([]byte("taro")),
		"age":  types.NewIntValue(20),
	}
	err = vm.pushObject(addedVal)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Oadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
		"user": types.NewObjectValue(map[string]*types.Value{
			"name": types.NewStringValue([]byte("taro")),
			"age":  types.NewIntValue(20),
		}),
	})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	got.Object["user"].Object["name"] = types.NewStringValue([]byte("jiro"))
	if diff := cmp.Diff(addedVal, got.Object["user"].Object); diff == "" {
		t.Errorf("the added value does not seem to be a clone of the value on the stack")
	}
}

func TestFeedOaddFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Oadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedOaddFailsWhenStackIsInsufficient1(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Oadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedOaddFailsWhenStackIsInsufficient2(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushString([]byte("hoge"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Oadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedOaddFailsWhenArg2IsNotString(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushObject(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(2021)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Oadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedOaddFailsWhenArg3IsNotString(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushString([]byte("year"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(2021)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Oadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedAnewPushesEmptyArray(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Anew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewArrayValue([]*types.Value{})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedAaddAppendsArg1ToArg2(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushArray([]*types.Value{types.NewIntValue(123)})
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushFloat(4.56)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Aadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewFloatValue(4.56),
	})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedAaddAppendsACopyOfAValue(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushArray([]*types.Value{types.NewStringValue([]byte("hello"))})
	if err != nil {
		t.Fatal(err)
	}
	addedVal := map[string]*types.Value{
		"name": types.NewStringValue([]byte("taro")),
		"age":  types.NewIntValue(20),
	}
	err = vm.pushObject(addedVal)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Aadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewArrayValue([]*types.Value{
		types.NewStringValue([]byte("hello")),
		types.NewObjectValue(map[string]*types.Value{
			"name": types.NewStringValue([]byte("taro")),
			"age":  types.NewIntValue(20),
		}),
	})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	got.Array[1].Object["name"] = types.NewStringValue([]byte("jiro"))
	if diff := cmp.Diff(addedVal, got.Array[1].Object); diff == "" {
		t.Errorf("the added value does not seem to be a clone of the value on the stack")
	}
}

func TestFeedAaddFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Aadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedAaddFailsWhenStackIsInsufficient(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Aadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedAaddFailsIfArg2IsNotArray(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Aadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedBnewPushesFalse(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Bnew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewBoolValue(false)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedBnegNegatesTheTop(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushBool(true)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Bneg)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewBoolValue(false)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedBnegFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Bneg)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedBnegFailsWhenArg1IsNotBool(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Bneg)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedNnewPushesNil(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Nnew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewNilValue()
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestGdupDuplicatesArg1(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(123)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Gdup)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 1 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewIntValue(123)
	clone, err := vm.pop()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	orig, err := vm.pop()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, orig); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestGdupPushesACopy(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushObject(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Gdup)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 1 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	clone, err := vm.pop()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	orig, err := vm.pop()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, orig); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Object["ebi"] = types.NewStringValue([]byte("shrimp"))
	if diff := cmp.Diff(clone, orig); diff == "" {
		t.Errorf("Gdup does not seem to copy arg1")
	}
}

func TestFeedGdupFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Gdup)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedGpopPopsOnce(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushString([]byte("first"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushString([]byte("second"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Gpop)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := types.NewStringValue([]byte("first"))
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedGpopFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Gpop)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestGswpSwapsArgs(t *testing.T) {
	var err error
	vm := NewVM()
	a := types.NewIntValue(1)
	b := types.NewStringValue([]byte("2"))

	err = vm.push(b)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.push(a)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Gswp)
	if err != nil {
		t.Fatal(err)
	}

	b1, err := vm.pop()
	if err != nil {
		t.Fatal(err)
	}
	a1, err := vm.pop()
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != -1 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", -1, vm.sp)
	}

	if diff := cmp.Diff(a, a1); diff != "" {
		t.Errorf("comparing arg1: mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(b, b1); diff != "" {
		t.Errorf("comparing arg2: mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedGswpFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Gswp)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedGswpFailsWhenStackIsInsufficient(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Gswp)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedMultiDoNothingWhenOpsIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.FeedMulti([]Op{})
	if err != nil {
		t.Fatal(err)
	}
	if vm.sp != -1 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", -1, vm.sp)
	}
}

func TestFeedMultiExecutesOpsSequentially(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.FeedMulti([]Op{Inew, Iinc, Iinc, Iinc})
	if err != nil {
		t.Fatal(err)
	}

	want := types.NewIntValue(3)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
