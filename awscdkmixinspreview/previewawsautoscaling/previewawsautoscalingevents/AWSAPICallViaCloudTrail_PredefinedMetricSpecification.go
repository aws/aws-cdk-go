package previewawsautoscalingevents


// Type definition for PredefinedMetricSpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predefinedMetricSpecification := &PredefinedMetricSpecification{
//   	PredefinedMetricType: []*string{
//   		jsii.String("predefinedMetricType"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_PredefinedMetricSpecification struct {
	// predefinedMetricType property.
	//
	// Specify an array of string values to match this event if the actual value of predefinedMetricType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PredefinedMetricType *[]*string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
}

