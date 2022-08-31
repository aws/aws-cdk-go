package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
//   	tableName: jsii.String("my_table"),
//   	columns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			type: glue.schema_STRING(),
//   		},
//   	},
//   	partitionKeys: []*column{
//   		&column{
//   			name: jsii.String("year"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   		&column{
//   			name: jsii.String("month"),
//   			type: glue.*schema_SMALL_INT(),
//   		},
//   	},
//   	dataFormat: glue.dataFormat_JSON(),
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/data-types.html
//
// Experimental.
type Schema interface {
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

// Experimental.
func NewSchema() Schema {
	_init_.Initialize()

	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"monocdk.aws_glue.Schema",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSchema_Override(s Schema) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_glue.Schema",
		nil, // no parameters
		s,
	)
}

// Creates an array of some other type.
// Experimental.
func Schema_Array(itemType *Type) *Type {
	_init_.Initialize()

	if err := validateSchema_ArrayParameters(itemType); err != nil {
		panic(err)
	}
	var returns *Type

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Schema",
		"array",
		[]interface{}{itemType},
		&returns,
	)

	return returns
}

// Fixed length character data, with a specified length between 1 and 255.
// Experimental.
func Schema_Char(length *float64) *Type {
	_init_.Initialize()

	if err := validateSchema_CharParameters(length); err != nil {
		panic(err)
	}
	var returns *Type

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Schema",
		"char",
		[]interface{}{length},
		&returns,
	)

	return returns
}

// Creates a decimal type.
//
// TODO: Bounds.
// Experimental.
func Schema_Decimal(precision *float64, scale *float64) *Type {
	_init_.Initialize()

	if err := validateSchema_DecimalParameters(precision); err != nil {
		panic(err)
	}
	var returns *Type

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Schema",
		"decimal",
		[]interface{}{precision, scale},
		&returns,
	)

	return returns
}

// Creates a map of some primitive key type to some value type.
// Experimental.
func Schema_Map(keyType *Type, valueType *Type) *Type {
	_init_.Initialize()

	if err := validateSchema_MapParameters(keyType, valueType); err != nil {
		panic(err)
	}
	var returns *Type

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Schema",
		"map",
		[]interface{}{keyType, valueType},
		&returns,
	)

	return returns
}

// Creates a nested structure containing individually named and typed columns.
// Experimental.
func Schema_Struct(columns *[]*Column) *Type {
	_init_.Initialize()

	if err := validateSchema_StructParameters(columns); err != nil {
		panic(err)
	}
	var returns *Type

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Schema",
		"struct",
		[]interface{}{columns},
		&returns,
	)

	return returns
}

// Variable length character data, with a specified length between 1 and 65535.
// Experimental.
func Schema_Varchar(length *float64) *Type {
	_init_.Initialize()

	if err := validateSchema_VarcharParameters(length); err != nil {
		panic(err)
	}
	var returns *Type

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.Schema",
		"varchar",
		[]interface{}{length},
		&returns,
	)

	return returns
}

func Schema_BIG_INT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"BIG_INT",
		&returns,
	)
	return returns
}

func Schema_BINARY() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"BINARY",
		&returns,
	)
	return returns
}

func Schema_BOOLEAN() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"BOOLEAN",
		&returns,
	)
	return returns
}

func Schema_DATE() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"DATE",
		&returns,
	)
	return returns
}

func Schema_DOUBLE() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"DOUBLE",
		&returns,
	)
	return returns
}

func Schema_FLOAT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"FLOAT",
		&returns,
	)
	return returns
}

func Schema_INTEGER() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"INTEGER",
		&returns,
	)
	return returns
}

func Schema_SMALL_INT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"SMALL_INT",
		&returns,
	)
	return returns
}

func Schema_STRING() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"STRING",
		&returns,
	)
	return returns
}

func Schema_TIMESTAMP() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"TIMESTAMP",
		&returns,
	)
	return returns
}

func Schema_TINY_INT() *Type {
	_init_.Initialize()
	var returns *Type
	_jsii_.StaticGet(
		"monocdk.aws_glue.Schema",
		"TINY_INT",
		&returns,
	)
	return returns
}

