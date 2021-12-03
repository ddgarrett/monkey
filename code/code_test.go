package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		for i, b := range tt.expected {

			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, go=%d", i, b, instruction[i])
			}
		}
	}
}
