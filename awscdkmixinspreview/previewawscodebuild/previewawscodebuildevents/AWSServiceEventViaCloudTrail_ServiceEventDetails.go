package previewawscodebuildevents


// Type definition for ServiceEventDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceEventDetails := &ServiceEventDetails{
//   	Response: []*string{
//   		jsii.String("response"),
//   	},
//   }
//
// Experimental.
type AWSServiceEventViaCloudTrail_ServiceEventDetails struct {
	// response property.
	//
	// Specify an array of string values to match this event if the actual value of response is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Response *[]*string `field:"optional" json:"response" yaml:"response"`
}

