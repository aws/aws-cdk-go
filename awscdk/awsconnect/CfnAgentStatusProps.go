package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAgentStatus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentStatusProps := &CfnAgentStatusProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DisplayOrder: jsii.Number(123),
//   	ResetOrderNumber: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html
//
type CfnAgentStatusProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the agent status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The state of the agent status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// The description of the agent status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display order of the agent status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-displayorder
	//
	DisplayOrder *float64 `field:"optional" json:"displayOrder" yaml:"displayOrder"`
	// A number indicating the reset order of the agent status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-resetordernumber
	//
	ResetOrderNumber interface{} `field:"optional" json:"resetOrderNumber" yaml:"resetOrderNumber"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of agent status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-agentstatus.html#cfn-connect-agentstatus-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

