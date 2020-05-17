package get

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestGet_getArgs(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{name: "tes1", args: args{input: "/login?ph=1"}, want: map[string]string{"ph": "1"}},
		{name: "tes2", args: args{input: "/danmu/v3/?id=15666&max=1000"}, want: map[string]string{"id": "15666", "max": "1000"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Get{}
			if got := g.getArgs(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanStruct(t *testing.T) {
	type Test1 struct {
		Name string
		ID   int
		Sex  float64 `json:"s"`
	}
	type args struct {
		target interface{}
		args   map[string]string
	}
	tests := []struct {
		name string
		args args
		want *Test1
	}{
		{name: "test1", args: args{target: &Test1{}, args: map[string]string{"name": "1", "id": "19", "s": "1.3"}}, want: &Test1{"1", 19, 1.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if scanStruct(tt.args.target, tt.args.args); !reflect.DeepEqual(tt.args.target, tt.want) {
				t.Errorf("getArgs() = %v, want %v", tt.args.target, tt.want)
			}
		})
	}
}

// 测试错误的json是不是会影响get到的数据
func Test_Chan(a *testing.T) {
	type Test1 struct {
		Name string
		ID   int
		Sex  float64 `json:"s"`
	}
	var t = &Test1{}
	var args = map[string]string{"name": "1", "id": "19", "s": "1.3"}
	scanStruct(t, args)
	json.Unmarshal([]byte("nil"), t)
	fmt.Println(t)
}