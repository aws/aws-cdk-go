package previewawspipesmixins


// Represents the Amazon Data Firehose logging configuration settings for the pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firehoseLogDestinationProperty := &FirehoseLogDestinationProperty{
//   	DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-firehoselogdestination.html
//
type CfnPipePropsMixin_FirehoseLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Firehose delivery stream to which EventBridge delivers the pipe log records.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-firehoselogdestination.html#cfn-pipes-pipe-firehoselogdestination-deliverystreamarn
	//
	DeliveryStreamArn *string `field:"optional" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
}

