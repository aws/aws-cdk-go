// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Represents a type of a column in a table schema.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &tableProps{
//   	database: myDatabase,
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
// Experimental.
type Type struct {
	// Glue InputString for this type.
	// Experimental.
	InputString *string `field:"required" json:"inputString" yaml:"inputString"`
	// Indicates whether this type is a primitive data type.
	// Experimental.
	IsPrimitive *bool `field:"required" json:"isPrimitive" yaml:"isPrimitive"`
}

