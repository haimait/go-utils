package slice

import (
	"reflect"
	"testing"
)

func TestGetUniqPolices(t *testing.T) {
	type args struct {
		polices [][]string
	}
	roleKey := "admin"
	tests := []struct {
		name            string
		args            args
		wantUniqPolices [][]string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{polices: [][]string{
				{roleKey, "/api/v1/hello", "GET"},
				{roleKey, "/api/v1/add", "POST"},
				{roleKey, "/api/v1/delete", "DELETE"},
				{roleKey, "/api/v1/delete", "DELETE"},
			}},
			wantUniqPolices: [][]string{
				{roleKey, "/api/v1/hello", "GET"},
				{roleKey, "/api/v1/add", "POST"},
				{roleKey, "/api/v1/delete", "DELETE"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUniqPolices := GetUniqPolices(tt.args.polices); !reflect.DeepEqual(gotUniqPolices, tt.wantUniqPolices) {
				t.Errorf("GetUniqPolices() = %v, want %v", gotUniqPolices, tt.wantUniqPolices)
			}
		})
	}
}

func Test_arrayInGroups(t *testing.T) {
	type args struct {
		arr []int
		num int64
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "testarrayInGroups",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				num: 2,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
				{5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayInGroups(tt.args.arr, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayInGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayTwoGroupsOf(t *testing.T) {
	type args struct {
		arr [][]int
		num int64
	}
	tests := []struct {
		name string
		args args
		want [][][]int
	}{
		// TODO: Add test cases.
		{
			name: "testArrayTwoGroups02",
			args: args{
				arr: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
				num: 2,
			},
			want: [][][]int{
				{{1, 2, 3}, {4, 5, 6}},
				{{7, 8, 9}, {10, 11, 12}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayTwoGroups(tt.args.arr, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayTwoStringGroupsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStuGroups(t *testing.T) {
	type args struct {
		stu []student
		num int64
	}
	tests := []struct {
		name string
		args args
		want [][]student
	}{
		// TODO: Add test cases.
		{
			name: "TestArrayStuGroups01",
			args: args{
				stu: []student{
					{name: "小王子", age: 18},
					{name: "娜扎", age: 23},
					{name: "大王八", age: 9000},
					{name: "大王八1", age: 9000},
					{name: "大王八2", age: 9000},
					{name: "大王八3", age: 9000},
					{name: "大王八4", age: 9000},
				},
				num: 3,
			},
			want: [][]student{
				{{name: "小王子", age: 18}, {name: "娜扎", age: 23}, {name: "大王八", age: 9000}},
				{{name: "大王八1", age: 9000}, {name: "大王八2", age: 9000}, {name: "大王八3", age: 9000}},
				{{name: "大王八4", age: 9000}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayStuGroups(tt.args.stu, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStuGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceSortAsc(t *testing.T) {
	type args struct {
		list []Person
	}
	tests := []struct {
		name string
		args args
		want []Person
	}{
		// TODO: Add test cases.
		{
			name: "TestSliceSortAsc",
			args: args{
				list: []Person{
					{Name: "娜扎", Age: 23},
					{Name: "小王子", Age: 18},
					{Name: "大王八1", Age: 9001},
					{Name: "大王八", Age: 9000},
					{Name: "大王八2", Age: 9008},
					{Name: "大王八3", Age: 9007},
					{Name: "大王八4", Age: 9006},
				},
			},
			want: []Person{
				{Name: "小王子", Age: 18},
				{Name: "娜扎", Age: 23},
				{Name: "大王八", Age: 9000},
				{Name: "大王八1", Age: 9001},
				{Name: "大王八4", Age: 9006},
				{Name: "大王八3", Age: 9007},
				{Name: "大王八2", Age: 9008},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSortAsc(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceSortAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}
