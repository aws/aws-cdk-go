package previewawsiotanalyticsevents


// Type definition for Settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   settings := &Settings{
//   	Timeseries: &Timeseries{
//   		DimensionAttributes: []TimeseriesItem{
//   			&TimeseriesItem{
//   				Attribute: []*string{
//   					jsii.String("attribute"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   		MeasureAttributes: []TimeseriesItem{
//   			&TimeseriesItem{
//   				Attribute: []*string{
//   					jsii.String("attribute"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   		TimestampAttribute: []*string{
//   			jsii.String("timestampAttribute"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Settings struct {
	// timeseries property.
	//
	// Specify an array of string values to match this event if the actual value of timeseries is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Timeseries *AWSAPICallViaCloudTrail_Timeseries `field:"optional" json:"timeseries" yaml:"timeseries"`
}

