package awscdkgluealpha


// Properties for ENUM partition projection configuration.
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
//   			Name: jsii.String("region"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"region": glue.PartitionProjectionConfiguration_enum(&EnumPartitionProjectionConfigurationProps{
//   			"values": []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("us-west-2"),
//   				jsii.String("eu-west-1"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type EnumPartitionProjectionConfigurationProps struct {
	// Explicit list of partition values.
	//
	// Example:
	//   []*string{
	//   	"us-east-1",
	//   	"us-west-2",
	//   	"eu-west-1",
	//   }
	//
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

