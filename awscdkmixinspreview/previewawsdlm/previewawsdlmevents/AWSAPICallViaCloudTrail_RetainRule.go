package previewawsdlmevents


// Type definition for RetainRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retainRule := &RetainRule{
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RetainRule struct {
	// Count property.
	//
	// Specify an array of string values to match this event if the actual value of Count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
}

