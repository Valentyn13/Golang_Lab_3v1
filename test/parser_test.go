package test

import (
	"image/color"
	"strings"
	"testing"

	"github.com/Valentyn13/Golang_Lab_3v1/painter"
	"github.com/Valentyn13/Golang_Lab_3v1/painter/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserStructure(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{ // test cases, for the checking loop =)
		{
			name:    "backgr rectangle",
			command: "bgrect 0 0 100 100",
			op:      &painter.BgRectangle{X1: 0, Y1: 0, X2: 100, Y2: 100},
		},
		{
			name:    "figure",
			command: "figure 250 250",
			op:      &painter.Figure{X: 250, Y: 250, C: color.RGBA{R: 255, G: 255, B: 0, A: 1}},
		},
		{
			name:    "move to cordinates",
			command: "move 200 200",
			op:      &painter.Move{X: 200, Y: 200},
		},
		{
			name:    "update",
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			name:    "invalid command",
			command: "invalidcommand",
			op:      nil,
		},
	}

	for _, tc := range tests { //created checking loop
		t.Run(tc.name, func(t *testing.T) {
			parser := &lang.Parser{}
			ops, err := parser.Parse(strings.NewReader(tc.command))
			if tc.op == nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err) // checking errors
				assert.IsType(t, tc.op, ops[1])
				assert.Equal(t, tc.op, ops[1])
			}
		})
	}
}

func TestParserFuncions(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{
		{
			name:    "white fill",
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:    "green fill",
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			name:    "reset screen",
			command: "reset",
			op:      painter.OperationFunc(painter.ClearScreen),
		},
	}
	// created Parser object
	parser := &lang.Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))

			require.NoError(t, err)
			require.Len(t, ops, 1)
			assert.IsType(t, tc.op, ops[0])

		})
	}
}
