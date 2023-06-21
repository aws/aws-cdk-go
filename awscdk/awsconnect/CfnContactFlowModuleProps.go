package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContactFlowModule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactFlowModuleProps := &CfnContactFlowModuleProps{
//   	Content: jsii.String("content"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	State: jsii.String("state"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnContactFlowModuleProps struct {
	// The content of the flow module.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the flow module.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the flow module.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state of the flow module.
	State *string `field:"optional" json:"state" yaml:"state"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

