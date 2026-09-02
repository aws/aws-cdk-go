package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCloudConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCloudConnectorMixinProps := &CfnCloudConnectorMixinProps{
//   	ConfigConnectorArn: jsii.String("configConnectorArn"),
//   	Configuration: &CloudConnectorConfigurationProperty{
//   		AzureConfiguration: &AzureConfigurationProperty{
//   			ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   			ApplicationId: jsii.String("applicationId"),
//   			Targets: &ConfigurationTargetsProperty{
//   				Subscriptions: []interface{}{
//   					&AzureSubscriptionProperty{
//   						DisplayName: jsii.String("displayName"),
//   						Id: jsii.String("id"),
//   					},
//   				},
//   			},
//   			TenantDisplayName: jsii.String("tenantDisplayName"),
//   			TenantId: jsii.String("tenantId"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html
//
type CfnCloudConnectorMixinProps struct {
	// The ARN of the AWS Config connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-configconnectorarn
	//
	ConfigConnectorArn *string `field:"optional" json:"configConnectorArn" yaml:"configConnectorArn"`
	// The configuration for the cloud connector.
	//
	// Currently supports Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The description of the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The IAM role ARN used by the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags to apply to the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

