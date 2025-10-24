package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProvisioningTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProvisioningTemplateProps := &CfnProvisioningTemplateProps{
//   	ProvisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	TemplateBody: jsii.String("templateBody"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Enabled: jsii.Boolean(false),
//   	PreProvisioningHook: &ProvisioningHookProperty{
//   		PayloadVersion: jsii.String("payloadVersion"),
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateName: jsii.String("templateName"),
//   	TemplateType: jsii.String("templateType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html
//
type CfnProvisioningTemplateProps struct {
	// The role ARN for the role associated with the fleet provisioning template.
	//
	// This IoT role grants permission to provision a device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-provisioningrolearn
	//
	ProvisioningRoleArn *string `field:"required" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// The JSON formatted contents of the fleet provisioning template version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-templatebody
	//
	TemplateBody *string `field:"required" json:"templateBody" yaml:"templateBody"`
	// The description of the fleet provisioning template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// True to enable the fleet provisioning template, otherwise false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Creates a pre-provisioning hook template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-preprovisioninghook
	//
	PreProvisioningHook interface{} `field:"optional" json:"preProvisioningHook" yaml:"preProvisioningHook"`
	// Metadata that can be used to manage the fleet provisioning template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the fleet provisioning template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The type of the provisioning template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-provisioningtemplate.html#cfn-iot-provisioningtemplate-templatetype
	//
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}

