package compiler

import (
	"monkey/code"
	"testing"
)

func TestGlobalLetStatementsWithoutLet(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
            one = 1;
            two = 2;
            `,
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpSetGlobal, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpSetGlobal, 1),
			},
		},
		{
			input: `
            one = 1;
            one;
            `,
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpSetGlobal, 0),
				code.Make(code.OpGetGlobal, 0),
				code.Make(code.OpPop),
			},
		},
		{
			input: `
            one = 1;
            two = one;
            two;
            `,
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpSetGlobal, 0),
				code.Make(code.OpGetGlobal, 0),
				code.Make(code.OpSetGlobal, 1),
				code.Make(code.OpGetGlobal, 1),
				code.Make(code.OpPop),
			},
		},
	}

	runCompilerTests(t, tests)
}
func TestGlobalLetStatements(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
            let one = 1;
            let two = 2;
            `,
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpSetGlobal, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpSetGlobal, 1),
			},
		},
		{
			input: `
            let one = 1;
            one;
            `,
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpSetGlobal, 0),
				code.Make(code.OpGetGlobal, 0),
				code.Make(code.OpPop),
			},
		},
		{
			input: `
            let one = 1;
            let two = one;
            two;
            `,
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpSetGlobal, 0),
				code.Make(code.OpGetGlobal, 0),
				code.Make(code.OpSetGlobal, 1),
				code.Make(code.OpGetGlobal, 1),
				code.Make(code.OpPop),
			},
		},
	}

	runCompilerTests(t, tests)
}
