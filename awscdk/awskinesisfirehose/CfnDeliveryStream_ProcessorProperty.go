package awskinesisfirehose


// The `Processor` property specifies a data processor for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorProperty := &ProcessorProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Parameters: []interface{}{
//   		&ProcessorParameterProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html
//
type CfnDeliveryStream_ProcessorProperty struct {
	// The type of processor.
	//
	// Valid values: `Lambda` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html#cfn-kinesisfirehose-deliverystream-processor-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The processor parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html#cfn-kinesisfirehose-deliverystream-processor-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

