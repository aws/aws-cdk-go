package previewawskmsevents


// Type definition for EncryptionContext.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionContext := &EncryptionContext{
//   	AwsS3Arn: []*string{
//   		jsii.String("awsS3Arn"),
//   	},
//   }
//
// Experimental.
type KeyEvents_AWSAPICallViaCloudTrail_EncryptionContext struct {
	// aws:s3:arn property.
	//
	// Specify an array of string values to match this event if the actual value of aws:s3:arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsS3Arn *[]*string `field:"optional" json:"awsS3Arn" yaml:"awsS3Arn"`
}

