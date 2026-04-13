package previewawselasticloadbalancingevents


// Type definition for ResponseElementsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElementsItem := &ResponseElementsItem{
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElementsItem struct {
	// instanceId property.
	//
	// Specify an array of string values to match this event if the actual value of instanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

