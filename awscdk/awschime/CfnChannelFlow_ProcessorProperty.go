package awschime


// Information about a processor in a channel flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorProperty := &ProcessorProperty{
//   	Configuration: &ProcessorConfigurationProperty{
//   		Lambda: &LambdaConfigurationProperty{
//   			InvocationType: jsii.String("invocationType"),
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	ExecutionOrder: jsii.Number(123),
//   	FallbackAction: jsii.String("fallbackAction"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processor.html
//
type CfnChannelFlow_ProcessorProperty struct {
	// A processor's metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processor.html#cfn-chime-channelflow-processor-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The sequence in which processors run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processor.html#cfn-chime-channelflow-processor-executionorder
	//
	ExecutionOrder *float64 `field:"required" json:"executionOrder" yaml:"executionOrder"`
	// Determines whether to continue or stop processing when communication with a processor fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processor.html#cfn-chime-channelflow-processor-fallbackaction
	//
	FallbackAction *string `field:"required" json:"fallbackAction" yaml:"fallbackAction"`
	// The name of the processor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processor.html#cfn-chime-channelflow-processor-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

