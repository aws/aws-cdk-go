package awscdkgluealpha


// Properties for INTEGER partition projection configuration.
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
//   			Name: jsii.String("year"),
//   			Type: glue.Schema_INTEGER(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"year": glue.PartitionProjectionConfiguration_integer(&IntegerPartitionProjectionConfigurationProps{
//   			"min": jsii.Number(2020),
//   			"max": jsii.Number(2023),
//   			"interval": jsii.Number(1),
//   			 // optional, defaults to 1
//   			"digits": jsii.Number(4),
//   		}),
//   	},
//   })
//
// Experimental.
type IntegerPartitionProjectionConfigurationProps struct {
	// Maximum value for the integer partition range (inclusive).
	// Experimental.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Minimum value for the integer partition range (inclusive).
	// Experimental.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Number of digits to pad the partition value with leading zeros.
	//
	// With digits: 4, partition values: 0001, 0002, ..., 0100
	// Default: - no static number of digits and no leading zeroes.
	//
	// Experimental.
	Digits *float64 `field:"optional" json:"digits" yaml:"digits"`
	// Interval between partition values.
	// Default: 1.
	//
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

