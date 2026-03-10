package previewawsiotanalyticsevents


// Type definition for Timeseries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeseries := &Timeseries{
//   	DimensionAttributes: []TimeseriesItem{
//   		&TimeseriesItem{
//   			Attribute: []*string{
//   				jsii.String("attribute"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   	MeasureAttributes: []TimeseriesItem{
//   		&TimeseriesItem{
//   			Attribute: []*string{
//   				jsii.String("attribute"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   	TimestampAttribute: []*string{
//   		jsii.String("timestampAttribute"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Timeseries struct {
	// dimensionAttributes property.
	//
	// Specify an array of string values to match this event if the actual value of dimensionAttributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DimensionAttributes *[]*AWSAPICallViaCloudTrail_TimeseriesItem `field:"optional" json:"dimensionAttributes" yaml:"dimensionAttributes"`
	// measureAttributes property.
	//
	// Specify an array of string values to match this event if the actual value of measureAttributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MeasureAttributes *[]*AWSAPICallViaCloudTrail_TimeseriesItem `field:"optional" json:"measureAttributes" yaml:"measureAttributes"`
	// timestampAttribute property.
	//
	// Specify an array of string values to match this event if the actual value of timestampAttribute is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TimestampAttribute *[]*string `field:"optional" json:"timestampAttribute" yaml:"timestampAttribute"`
}

