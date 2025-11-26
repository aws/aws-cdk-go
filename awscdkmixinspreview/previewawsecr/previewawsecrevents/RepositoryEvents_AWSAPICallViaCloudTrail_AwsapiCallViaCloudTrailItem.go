package previewawsecrevents


// Type definition for AWSAPICallViaCloudTrailItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsapiCallViaCloudTrailItem := &AwsapiCallViaCloudTrailItem{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// ARN property.
	//
	// Specify an array of string values to match this event if the actual value of ARN is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
}

