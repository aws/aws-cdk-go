package awscdkgluealpha


// Represents a type of a column in a table schema.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &TableProps{
//   	Database: myDatabase,
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []*column{
//   		&column{
//   			Name: jsii.String("year"),
//   			Type: glue.Schema_SMALL_INT(),
//   		},
//   		&column{
//   			Name: jsii.String("month"),
//   			Type: glue.Schema_SMALL_INT(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
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

