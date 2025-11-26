package previewawsvoiceidevents


// Type definition for EnrollmentAudioProgress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enrollmentAudioProgress := &EnrollmentAudioProgress{
//   	AudioAggregationEndedAt: []*string{
//   		jsii.String("audioAggregationEndedAt"),
//   	},
//   	AudioAggregationStartedAt: []*string{
//   		jsii.String("audioAggregationStartedAt"),
//   	},
//   	AudioAggregationStatus: []*string{
//   		jsii.String("audioAggregationStatus"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdStartSessionAction_EnrollmentAudioProgress struct {
	// audioAggregationEndedAt property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationEndedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationEndedAt *[]*string `field:"optional" json:"audioAggregationEndedAt" yaml:"audioAggregationEndedAt"`
	// audioAggregationStartedAt property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationStartedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationStartedAt *[]*string `field:"optional" json:"audioAggregationStartedAt" yaml:"audioAggregationStartedAt"`
	// audioAggregationStatus property.
	//
	// Specify an array of string values to match this event if the actual value of audioAggregationStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AudioAggregationStatus *[]*string `field:"optional" json:"audioAggregationStatus" yaml:"audioAggregationStatus"`
}

