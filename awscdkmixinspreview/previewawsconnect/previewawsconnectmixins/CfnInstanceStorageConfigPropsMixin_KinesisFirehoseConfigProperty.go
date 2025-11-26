package previewawsconnectmixins


// Configuration information of a Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisFirehoseConfigProperty := &KinesisFirehoseConfigProperty{
//   	FirehoseArn: jsii.String("firehoseArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisfirehoseconfig.html
//
type CfnInstanceStorageConfigPropsMixin_KinesisFirehoseConfigProperty struct {
	// The Amazon Resource Name (ARN) of the delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instancestorageconfig-kinesisfirehoseconfig.html#cfn-connect-instancestorageconfig-kinesisfirehoseconfig-firehosearn
	//
	FirehoseArn *string `field:"optional" json:"firehoseArn" yaml:"firehoseArn"`
}

