package previewawscodebuildevents


// Type definition for Webhook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webhook := &Webhook{
//   	BranchFilter: []*string{
//   		jsii.String("branchFilter"),
//   	},
//   	LastModifiedSecret: []*string{
//   		jsii.String("lastModifiedSecret"),
//   	},
//   	PayloadUrl: []*string{
//   		jsii.String("payloadUrl"),
//   	},
//   	Url: []*string{
//   		jsii.String("url"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Webhook struct {
	// branchFilter property.
	//
	// Specify an array of string values to match this event if the actual value of branchFilter is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BranchFilter *[]*string `field:"optional" json:"branchFilter" yaml:"branchFilter"`
	// lastModifiedSecret property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedSecret is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedSecret *[]*string `field:"optional" json:"lastModifiedSecret" yaml:"lastModifiedSecret"`
	// payloadUrl property.
	//
	// Specify an array of string values to match this event if the actual value of payloadUrl is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PayloadUrl *[]*string `field:"optional" json:"payloadUrl" yaml:"payloadUrl"`
	// url property.
	//
	// Specify an array of string values to match this event if the actual value of url is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Url *[]*string `field:"optional" json:"url" yaml:"url"`
}

