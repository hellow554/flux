// DO NOT EDIT: This file is autogenerated via the builtin command.

package influxdb

import (
	flux "github.com/influxdata/flux"
	ast "github.com/influxdata/flux/ast"
)

func init() {
	flux.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 16,
					Line:   5,
				},
				File:   "",
				Source: "package influxdb\n\nbuiltin from\nbuiltin to\nbuiltin buckets",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   3,
					},
					File:   "",
					Source: "builtin from",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   3,
						},
						File:   "",
						Source: "from",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Name: "from",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 11,
						Line:   4,
					},
					File:   "",
					Source: "builtin to",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 11,
							Line:   4,
						},
						File:   "",
						Source: "to",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "to",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   5,
					},
					File:   "",
					Source: "builtin buckets",
					Start: ast.Position{
						Column: 1,
						Line:   5,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 16,
							Line:   5,
						},
						File:   "",
						Source: "buckets",
						Start: ast.Position{
							Column: 9,
							Line:   5,
						},
					},
				},
				Name: "buckets",
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "influxdb.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   1,
					},
					File:   "",
					Source: "package influxdb",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   1,
						},
						File:   "",
						Source: "influxdb",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "influxdb",
			},
		},
	}},
	Package: "influxdb",
	Path:    "influxdata/influxdb",
}
