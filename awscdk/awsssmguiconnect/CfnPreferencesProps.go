package awsssmguiconnect


// Properties for defining a `CfnPreferences`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPreferencesProps := &CfnPreferencesProps{
//   	ConnectionRecordingPreferences: &ConnectionRecordingPreferencesProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		RecordingDestinations: &RecordingDestinationsProperty{
//   			S3Buckets: []interface{}{
//   				&S3BucketProperty{
//   					BucketName: jsii.String("bucketName"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmguiconnect-preferences.html
//
type CfnPreferencesProps struct {
	// The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region.
	//
	// This includes details such as which S3 bucket recordings are stored in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmguiconnect-preferences.html#cfn-ssmguiconnect-preferences-connectionrecordingpreferences
	//
	ConnectionRecordingPreferences interface{} `field:"optional" json:"connectionRecordingPreferences" yaml:"connectionRecordingPreferences"`
}

