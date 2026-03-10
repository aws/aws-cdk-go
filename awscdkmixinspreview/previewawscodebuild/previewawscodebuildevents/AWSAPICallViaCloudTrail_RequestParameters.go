package previewawscodebuildevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	AwsActId: []*string{
//   		jsii.String("awsActId"),
//   	},
//   	InvokedBy: []*string{
//   		jsii.String("invokedBy"),
//   	},
//   	PayerId: []*string{
//   		jsii.String("payerId"),
//   	},
//   	UserArn: []*string{
//   		jsii.String("userArn"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// awsActId property.
	//
	// Specify an array of string values to match this event if the actual value of awsActId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsActId *[]*string `field:"optional" json:"awsActId" yaml:"awsActId"`
	// invokedBy property.
	//
	// Specify an array of string values to match this event if the actual value of invokedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InvokedBy *[]*string `field:"optional" json:"invokedBy" yaml:"invokedBy"`
	// payerId property.
	//
	// Specify an array of string values to match this event if the actual value of payerId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PayerId *[]*string `field:"optional" json:"payerId" yaml:"payerId"`
	// userArn property.
	//
	// Specify an array of string values to match this event if the actual value of userArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserArn *[]*string `field:"optional" json:"userArn" yaml:"userArn"`
}

