package awscdkgluealpha


// A required-together interval step for DATE partition projection.
//
// Bundling `interval` and `intervalUnit` into one value makes a partial step
// (one without the other) unrepresentable.
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
//   			// `step` bundles interval + unit (supply both or neither). Optional at day
//   			// precision or coarser; required when the format is sub-day (e.g. hours).
//   			"step": &DateProjectionStep{
//   				"interval": jsii.Number(1),
//   				"intervalUnit": glue.DateIntervalUnit_DAYS,
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type DateProjectionStep struct {
	// Interval between partition values.
	// Experimental.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Unit for the interval.
	// Experimental.
	IntervalUnit DateIntervalUnit `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

