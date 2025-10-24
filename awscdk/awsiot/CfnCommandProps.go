package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCommand`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCommandProps := &CfnCommandProps{
//   	CommandId: jsii.String("commandId"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	Deprecated: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	MandatoryParameters: []interface{}{
//   		&CommandParameterProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DefaultValue: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   			Description: jsii.String("description"),
//   			Value: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	Payload: &CommandPayloadProperty{
//   		Content: jsii.String("content"),
//   		ContentType: jsii.String("contentType"),
//   	},
//   	PendingDeletion: jsii.Boolean(false),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html
//
type CfnCommandProps struct {
	// The unique identifier of the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-commandid
	//
	CommandId *string `field:"required" json:"commandId" yaml:"commandId"`
	// The timestamp, when the command was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Indicates whether the command has been deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-deprecated
	//
	Deprecated interface{} `field:"optional" json:"deprecated" yaml:"deprecated"`
	// The description of the command parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The timestamp, when the command was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-lastupdatedat
	//
	LastUpdatedAt *string `field:"optional" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-mandatoryparameters
	//
	MandatoryParameters interface{} `field:"optional" json:"mandatoryParameters" yaml:"mandatoryParameters"`
	// The namespace to which the command belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-payload
	//
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Indicates whether the command is pending deletion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-pendingdeletion
	//
	PendingDeletion interface{} `field:"optional" json:"pendingDeletion" yaml:"pendingDeletion"`
	// The customer role associated with the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags to be associated with the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

