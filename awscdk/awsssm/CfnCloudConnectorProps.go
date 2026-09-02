package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCloudConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudConnectorProps := &CfnCloudConnectorProps{
//   	ConfigConnectorArn: jsii.String("configConnectorArn"),
//   	Configuration: &CloudConnectorConfigurationProperty{
//   		AzureConfiguration: &AzureConfigurationProperty{
//   			ApplicationId: jsii.String("applicationId"),
//   			TenantId: jsii.String("tenantId"),
//
//   			// the properties below are optional
//   			ApplicationDisplayName: jsii.String("applicationDisplayName"),
//   			Targets: &ConfigurationTargetsProperty{
//   				Subscriptions: []interface{}{
//   					&AzureSubscriptionProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						DisplayName: jsii.String("displayName"),
//   					},
//   				},
//   			},
//   			TenantDisplayName: jsii.String("tenantDisplayName"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
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
type CfnCloudConnectorProps struct {
	// The ARN of the AWS Config connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-configconnectorarn
	//
	ConfigConnectorArn *string `field:"required" json:"configConnectorArn" yaml:"configConnectorArn"`
	// The configuration for the cloud connector.
	//
	// Currently supports Azure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The display name of the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The IAM role ARN used by the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to apply to the cloud connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-cloudconnector.html#cfn-ssm-cloudconnector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

