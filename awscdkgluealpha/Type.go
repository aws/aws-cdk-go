package awscdkgluealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of a column in a table schema.
//
// Instances are opaque: obtain one from a `Schema` factory (for example
// `Schema.STRING`, `Schema.decimal(...)`, `Schema.array(...)`) or, for a type the
// `Schema` factories don't model, from `Schema.custom(...)`.
//
// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("data"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []Column{
//   		&Column{
//   			Name: jsii.String("date"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"date": glue.PartitionProjectionConfiguration_date(&DatePartitionProjectionConfigurationProps{
//   			"min": jsii.String("NOW-3YEARS"),
//   			"max": jsii.String("NOW"),
//   			"format": jsii.String("yyyy-MM-dd"),
//   		}),
//   	},
//   })
//
// Experimental.
type Type interface {
	// Glue InputString for this type.
	// Experimental.
	InputString() *string
	// Indicates whether this type is a primitive data type.
	// Experimental.
	IsPrimitive() *bool
}

// The jsii proxy struct for Type
type jsiiProxy_Type struct {
	_ byte // padding
}

func (j *jsiiProxy_Type) InputString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Type) IsPrimitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isPrimitive",
		&returns,
	)
	return returns
}


