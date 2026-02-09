package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCommandPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCommandMixinProps := &CfnCommandMixinProps{
//   	CommandId: jsii.String("commandId"),
//   	CreatedAt: jsii.String("createdAt"),
//   	Deprecated: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	MandatoryParameters: []interface{}{
//   		&CommandParameterProperty{
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
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   			Value: &CommandParameterValueProperty{
//   				B: jsii.Boolean(false),
//   				Bin: jsii.String("bin"),
//   				D: jsii.Number(123),
//   				I: jsii.Number(123),
//   				L: jsii.String("l"),
//   				S: jsii.String("s"),
//   				Ul: jsii.String("ul"),
//   			},
//   			ValueConditions: []interface{}{
//   				&CommandParameterValueConditionProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					Operand: &CommandParameterValueComparisonOperandProperty{
//   						Number: jsii.String("number"),
//   						NumberRange: &CommandParameterValueNumberRangeProperty{
//   							Max: jsii.String("max"),
//   							Min: jsii.String("min"),
//   						},
//   						Numbers: []*string{
//   							jsii.String("numbers"),
//   						},
//   						String: jsii.String("string"),
//   						Strings: []*string{
//   							jsii.String("strings"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	Payload: &CommandPayloadProperty{
//   		Content: jsii.String("content"),
//   		ContentType: jsii.String("contentType"),
//   	},
//   	PayloadTemplate: jsii.String("payloadTemplate"),
//   	PendingDeletion: jsii.Boolean(false),
//   	Preprocessor: &CommandPreprocessorProperty{
//   		AwsJsonSubstitution: &AwsJsonSubstitutionCommandPreprocessorConfigProperty{
//   			OutputFormat: jsii.String("outputFormat"),
//   		},
//   	},
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
type CfnCommandMixinProps struct {
	// The unique identifier of the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-commandid
	//
	CommandId *string `field:"optional" json:"commandId" yaml:"commandId"`
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
	// The payload template associated with the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-payloadtemplate
	//
	PayloadTemplate *string `field:"optional" json:"payloadTemplate" yaml:"payloadTemplate"`
	// Indicates whether the command is pending deletion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-pendingdeletion
	//
	PendingDeletion interface{} `field:"optional" json:"pendingDeletion" yaml:"pendingDeletion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-preprocessor
	//
	Preprocessor interface{} `field:"optional" json:"preprocessor" yaml:"preprocessor"`
	// The customer role associated with the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags to be associated with the command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-command.html#cfn-iot-command-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

