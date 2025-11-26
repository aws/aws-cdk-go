package previewawskinesisanalyticsv2mixins


// When you configure a SQL-based Kinesis Data Analytics application's output, identifies a Kinesis data stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisStreamsOutputProperty := &KinesisStreamsOutputProperty{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-kinesisstreamsoutput.html
//
type CfnApplicationOutputPropsMixin_KinesisStreamsOutputProperty struct {
	// The ARN of the destination Kinesis data stream to write to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-kinesisstreamsoutput.html#cfn-kinesisanalyticsv2-applicationoutput-kinesisstreamsoutput-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

