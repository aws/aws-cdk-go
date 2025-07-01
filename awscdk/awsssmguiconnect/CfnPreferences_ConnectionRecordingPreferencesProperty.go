package awsssmguiconnect


// The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region .
//
// This includes details such as which S3 bucket recordings are stored in.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionRecordingPreferencesProperty := &ConnectionRecordingPreferencesProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	RecordingDestinations: &RecordingDestinationsProperty{
//   		S3Buckets: []interface{}{
//   			&S3BucketProperty{
//   				BucketName: jsii.String("bucketName"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-connectionrecordingpreferences.html
//
type CfnPreferences_ConnectionRecordingPreferencesProperty struct {
	// The ARN of a AWS KMS key that is used to encrypt data while it is being processed by the service.
	//
	// This key must exist in the same AWS Region as the node you start an RDP connection to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-connectionrecordingpreferences.html#cfn-ssmguiconnect-preferences-connectionrecordingpreferences-kmskeyarn
	//
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Determines where recordings of RDP connections are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-connectionrecordingpreferences.html#cfn-ssmguiconnect-preferences-connectionrecordingpreferences-recordingdestinations
	//
	RecordingDestinations interface{} `field:"required" json:"recordingDestinations" yaml:"recordingDestinations"`
}

