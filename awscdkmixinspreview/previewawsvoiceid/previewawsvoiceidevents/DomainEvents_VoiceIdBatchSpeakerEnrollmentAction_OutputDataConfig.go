package previewawsvoiceidevents


// Type definition for OutputDataConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputDataConfig := &OutputDataConfig{
//   	KmsKeyId: []*string{
//   		jsii.String("kmsKeyId"),
//   	},
//   	S3Uri: []*string{
//   		jsii.String("s3Uri"),
//   	},
//   }
//
// Experimental.
type DomainEvents_VoiceIdBatchSpeakerEnrollmentAction_OutputDataConfig struct {
	// kmsKeyId property.
	//
	// Specify an array of string values to match this event if the actual value of kmsKeyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KmsKeyId *[]*string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// s3Uri property.
	//
	// Specify an array of string values to match this event if the actual value of s3Uri is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Uri *[]*string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

