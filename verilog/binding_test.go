package verilog_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/verilog"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte(
		`module mod ();
		always_comb foo = bar.baz[7:0];
		endmodule`), verilog.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (module_declaration (module_header (module_keyword) (simple_identifier)) (module_nonansi_header (list_of_ports)) (module_or_generate_item (always_construct (always_keyword) (statement (statement_item (blocking_assignment (operator_assignment (variable_lvalue (simple_identifier)) (assignment_operator) (expression (primary (simple_identifier) (select1 (member_identifier (simple_identifier)) (constant_range (constant_expression (constant_primary (primary_literal (integral_number (decimal_number (unsigned_number)))))) (constant_expression (constant_primary (primary_literal (integral_number (decimal_number (unsigned_number))))))))))))))))))",
		n.String(),
	)
}
