package previewawsdlmevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	PolicyId: []*string{
//   		jsii.String("policyId"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// PolicyId property.
	//
	// Specify an array of string values to match this event if the actual value of PolicyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyId *[]*string `field:"optional" json:"policyId" yaml:"policyId"`
}

