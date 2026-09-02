package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannelFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelFlowProps := &CfnChannelFlowProps{
//   	AppInstanceArn: jsii.String("appInstanceArn"),
//   	Name: jsii.String("name"),
//   	Processors: []interface{}{
//   		&ProcessorProperty{
//   			Configuration: &ProcessorConfigurationProperty{
//   				Lambda: &LambdaConfigurationProperty{
//   					InvocationType: jsii.String("invocationType"),
//   					ResourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   			ExecutionOrder: jsii.Number(123),
//   			FallbackAction: jsii.String("fallbackAction"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channelflow.html
//
type CfnChannelFlowProps struct {
	// The ARN of the app instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channelflow.html#cfn-chime-channelflow-appinstancearn
	//
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
	// The name of the channel flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channelflow.html#cfn-chime-channelflow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Information about the processor Lambda functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channelflow.html#cfn-chime-channelflow-processors
	//
	Processors interface{} `field:"required" json:"processors" yaml:"processors"`
	// The tags for the channel flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-channelflow.html#cfn-chime-channelflow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

