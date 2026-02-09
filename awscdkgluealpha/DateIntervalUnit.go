package awscdkgluealpha


// Date interval unit for partition projection.
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
// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection-supported-types.html#partition-projection-date-type
//
// Experimental.
type DateIntervalUnit string

const (
	// Year interval.
	// Experimental.
	DateIntervalUnit_YEARS DateIntervalUnit = "YEARS"
	// Month interval.
	// Experimental.
	DateIntervalUnit_MONTHS DateIntervalUnit = "MONTHS"
	// Week interval.
	// Experimental.
	DateIntervalUnit_WEEKS DateIntervalUnit = "WEEKS"
	// Day interval (default).
	// Experimental.
	DateIntervalUnit_DAYS DateIntervalUnit = "DAYS"
	// Hour interval.
	// Experimental.
	DateIntervalUnit_HOURS DateIntervalUnit = "HOURS"
	// Minute interval.
	// Experimental.
	DateIntervalUnit_MINUTES DateIntervalUnit = "MINUTES"
	// Second interval.
	// Experimental.
	DateIntervalUnit_SECONDS DateIntervalUnit = "SECONDS"
)

