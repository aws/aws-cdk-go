package previewawstranscribeevents


// Type definition for AnalyticsJobDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analyticsJobDetails := &AnalyticsJobDetails{
//   	Skipped: []CallAnalyticsSkippedFeature{
//   		&CallAnalyticsSkippedFeature{
//   			Feature: []*string{
//   				jsii.String("feature"),
//   			},
//   			Message: []*string{
//   				jsii.String("message"),
//   			},
//   			ReasonCode: []*string{
//   				jsii.String("reasonCode"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type CallAnalyticsJobStateChange_AnalyticsJobDetails struct {
	// Skipped property.
	//
	// Specify an array of string values to match this event if the actual value of Skipped is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Skipped *[]*CallAnalyticsJobStateChange_CallAnalyticsSkippedFeature `field:"optional" json:"skipped" yaml:"skipped"`
}

