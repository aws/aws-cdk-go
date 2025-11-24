package mixinsawsivschat


// The FirehoseDestinationConfiguration property type specifies a Kinesis Firehose location where chat logs will be stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firehoseDestinationConfigurationProperty := &FirehoseDestinationConfigurationProperty{
//   	DeliveryStreamName: jsii.String("deliveryStreamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-firehosedestinationconfiguration.html
//
type CfnLoggingConfigurationPropsMixin_FirehoseDestinationConfigurationProperty struct {
	// Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-loggingconfiguration-firehosedestinationconfiguration.html#cfn-ivschat-loggingconfiguration-firehosedestinationconfiguration-deliverystreamname
	//
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
}

