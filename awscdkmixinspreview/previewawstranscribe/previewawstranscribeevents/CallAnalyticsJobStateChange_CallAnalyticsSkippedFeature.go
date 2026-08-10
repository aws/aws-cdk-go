package previewawstranscribeevents


// Type definition for CallAnalyticsSkippedFeature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   callAnalyticsSkippedFeature := &CallAnalyticsSkippedFeature{
//   	Feature: []*string{
//   		jsii.String("feature"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	ReasonCode: []*string{
//   		jsii.String("reasonCode"),
//   	},
//   }
//
// Experimental.
type CallAnalyticsJobStateChange_CallAnalyticsSkippedFeature struct {
	// feature property.
	//
	// Specify an array of string values to match this event if the actual value of feature is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Feature *[]*string `field:"optional" json:"feature" yaml:"feature"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// reasonCode property.
	//
	// Specify an array of string values to match this event if the actual value of reasonCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReasonCode *[]*string `field:"optional" json:"reasonCode" yaml:"reasonCode"`
}

