package previewawsconnectmixins


// Configuration information of a Kinesis data stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisStreamConfigProperty := &KinesisStreamConfigProperty{
//   	StreamArn: jsii.String("streamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisstreamconfig.html
//
type CfnInstanceStorageConfigPropsMixin_KinesisStreamConfigProperty struct {
	// The Amazon Resource Name (ARN) of the data stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisstreamconfig.html#cfn-connect-instancestorageconfig-kinesisstreamconfig-streamarn
	//
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

