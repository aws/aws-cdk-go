package previewawsssmguiconnectmixins


// Determines where recordings of RDP connections are stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordingDestinationsProperty := &RecordingDestinationsProperty{
//   	S3Buckets: []interface{}{
//   		&S3BucketProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-recordingdestinations.html
//
type CfnPreferencesPropsMixin_RecordingDestinationsProperty struct {
	// The S3 bucket where RDP connection recordings are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-recordingdestinations.html#cfn-ssmguiconnect-preferences-recordingdestinations-s3buckets
	//
	S3Buckets interface{} `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
}

