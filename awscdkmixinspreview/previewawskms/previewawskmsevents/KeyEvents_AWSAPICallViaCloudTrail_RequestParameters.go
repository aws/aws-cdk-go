package previewawskmsevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	EncryptionContext: &EncryptionContext{
//   		AwsS3Arn: []*string{
//   			jsii.String("awsS3Arn"),
//   		},
//   	},
//   	KeyId: []*string{
//   		jsii.String("keyId"),
//   	},
//   	KeySpec: []*string{
//   		jsii.String("keySpec"),
//   	},
//   }
//
// Experimental.
type KeyEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// encryptionContext property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionContext is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionContext *KeyEvents_AWSAPICallViaCloudTrail_EncryptionContext `field:"optional" json:"encryptionContext" yaml:"encryptionContext"`
	// keyId property.
	//
	// Specify an array of string values to match this event if the actual value of keyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Key reference.
	//
	// Experimental.
	KeyId *[]*string `field:"optional" json:"keyId" yaml:"keyId"`
	// keySpec property.
	//
	// Specify an array of string values to match this event if the actual value of keySpec is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KeySpec *[]*string `field:"optional" json:"keySpec" yaml:"keySpec"`
}

