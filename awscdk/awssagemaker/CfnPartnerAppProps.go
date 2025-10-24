package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPartnerApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPartnerAppProps := &CfnPartnerAppProps{
//   	AuthType: jsii.String("authType"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	Tier: jsii.String("tier"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ApplicationConfig: &PartnerAppConfigProperty{
//   		AdminUsers: []*string{
//   			jsii.String("adminUsers"),
//   		},
//   		Arguments: map[string]*string{
//   			"argumentsKey": jsii.String("arguments"),
//   		},
//   	},
//   	ClientToken: jsii.String("clientToken"),
//   	EnableIamSessionBasedIdentity: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MaintenanceConfig: &PartnerAppMaintenanceConfigProperty{
//   		MaintenanceWindowStart: jsii.String("maintenanceWindowStart"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html
//
type CfnPartnerAppProps struct {
	// Defines the authentication type used for the Partner AI App.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-authtype
	//
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The Amazon Resource Name (ARN) of the IAM role of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the Partner AI App.
	//
	// This name must be unique within your account and region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the tier or level of the Partner AI App.
	//
	// The tier size impacts the speed and capabilities of the application. For more information, see [Set up Partner AI Apps](https://docs.aws.amazon.com/sagemaker/latest/dg/partner-app-onboard.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-tier
	//
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Specifies the type of Partner AI App being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configuration settings for the Partner AI App.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-applicationconfig
	//
	ApplicationConfig interface{} `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// The client token for the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-clienttoken
	//
	// Deprecated: this property has been deprecated.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Enables IAM Session based Identity for PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-enableiamsessionbasedidentity
	//
	EnableIamSessionBasedIdentity interface{} `field:"optional" json:"enableIamSessionBasedIdentity" yaml:"enableIamSessionBasedIdentity"`
	// The AWS KMS customer managed key used to encrypt the data associated with the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A collection of settings that specify the maintenance schedule for the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-maintenanceconfig
	//
	MaintenanceConfig interface{} `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// A list of tags to apply to the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-partnerapp.html#cfn-sagemaker-partnerapp-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

