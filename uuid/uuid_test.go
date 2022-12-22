package uuid

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid, err := NewUUID()
	fmt.Printf("err:%v uuid:%v \n", err, uuid)

	uuidStr := uuid.String()
	fmt.Printf("uuidStr:%v\n", uuidStr)

	uuidCopy := uuid.Copy()
	fmt.Printf("uuidCopy:%v\n", uuidCopy)

	uuidRaw := uuid.Raw()
	fmt.Printf("uuidRaw:%v\n", uuidRaw)

}

func TestUUIDFromString(t *testing.T) {
	uid, err := UUIDFromString("blah")
	fmt.Println("uid:", uid) //uid: 00000000-0000-0000-0000-000000000000
	fmt.Println("err:", err) //err: invalid UUID: "blah"
	validUUID := "9f484882-2f18-4fd2-967d-db9663db7bea"
	uuid, err := UUIDFromString(validUUID)
	fmt.Printf("uuid:%v err:%v \n", uuid, err)

}

func TestIsValidUUIDString(t *testing.T) {

	tests := []struct {
		args string
		want bool
	}{
		{
			UUID{}.String(),
			true,
		},
		{
			"",
			false,
		},
		{
			"blah",
			false,
		},
		{
			"blah-9f484882-2f18-4fd2-967d-db9663db7bea",
			false,
		},
		{
			"9f484882-2f18-4fd2-967d-db9663db7bea-blah",
			false,
		},
		{
			"9f484882-2f18-4fd2-967d-db9663db7bea",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			got := IsValidUUIDString(tt.args)
			fmt.Println("got:", got)
			if got != tt.want {
				t.Errorf("IsValidUUIDString() = %v, want %v", got, tt.want)
			}
		})
	}
}
