package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnProvisioningTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProvisioningTemplateProps := &cfnProvisioningTemplateProps{
//   	provisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	templateBody: jsii.String("templateBody"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	enabled: jsii.Boolean(false),
//   	preProvisioningHook: &provisioningHookProperty{
//   		payloadVersion: jsii.String("payloadVersion"),
//   		targetArn: jsii.String("targetArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	templateName: jsii.String("templateName"),
//   	templateType: jsii.String("templateType"),
//   }
//
type CfnProvisioningTemplateProps struct {
	// The role ARN for the role associated with the fleet provisioning template.
	//
	// This IoT role grants permission to provision a device.
	ProvisioningRoleArn *string `field:"required" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// The JSON formatted contents of the fleet provisioning template version.
	TemplateBody *string `field:"required" json:"templateBody" yaml:"templateBody"`
	// The description of the fleet provisioning template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// True to enable the fleet provisioning template, otherwise false.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Creates a pre-provisioning hook template.
	PreProvisioningHook interface{} `field:"optional" json:"preProvisioningHook" yaml:"preProvisioningHook"`
	// Metadata that can be used to manage the fleet provisioning template.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the fleet provisioning template.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// `AWS::IoT::ProvisioningTemplate.TemplateType`.
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}

