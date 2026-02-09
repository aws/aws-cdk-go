package awscdkgluealpha


// Properties for DATE partition projection configuration.
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
//   			"min": jsii.String("2020-01-01"),
//   			"max": jsii.String("2023-12-31"),
//   			"format": jsii.String("yyyy-MM-dd"),
//   			"interval": jsii.Number(1),
//   			 // optional, defaults to 1
//   			"intervalUnit": glue.DateIntervalUnit_DAYS,
//   		}),
//   	},
//   })
//
// Experimental.
type DatePartitionProjectionConfigurationProps struct {
	// Date format for partition values.
	//
	// Uses Java SimpleDateFormat patterns.
	// See: https://docs.oracle.com/javase/8/docs/api/java/text/SimpleDateFormat.html
	//
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// End date for the partition range (inclusive).
	//
	// Can be either:
	// - Fixed date in the format specified by `format` property
	// - Relative date using NOW syntax
	//
	// Same format constraints as `min`.
	// Experimental.
	Max *string `field:"required" json:"max" yaml:"max"`
	// Start date for the partition range (inclusive).
	//
	// Can be either:
	// - Fixed date in the format specified by `format` property
	//   (e.g., '2020-01-01' for format 'yyyy-MM-dd')
	// - Relative date using NOW syntax
	//   (e.g., 'NOW', 'NOW-3YEARS', 'NOW+1MONTH')
	// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection-supported-types.html#partition-projection-date-type
	//
	// Experimental.
	Min *string `field:"required" json:"min" yaml:"min"`
	// Interval between partition values.
	//
	// When the provided dates are at single-day or single-month precision,
	// the interval is optional and defaults to 1 day or 1 month, respectively.
	// Otherwise, interval is required.
	// Default: - 1 for single-day or single-month precision, otherwise required.
	//
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Unit for the interval.
	//
	// When the provided dates are at single-day or single-month precision,
	// the intervalUnit is optional and defaults to 1 day or 1 month, respectively.
	// Otherwise, the intervalUnit is required.
	// Default: - DAYS for single-day precision, MONTHS for single-month precision, otherwise required.
	//
	// Experimental.
	IntervalUnit DateIntervalUnit `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

