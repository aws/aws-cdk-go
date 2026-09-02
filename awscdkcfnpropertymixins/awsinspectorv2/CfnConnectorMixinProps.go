package awsinspectorv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConnectorMixinProps := &CfnConnectorMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Provider: jsii.String("provider"),
//   	ProviderConfiguration: &ProviderConfigurationProperty{
//   		Azure: &AzureProviderConfigurationProperty{
//   			AutoInstallVmScanner: jsii.Boolean(false),
//   			AwsConfigConnectorArn: jsii.String("awsConfigConnectorArn"),
//   			AzureRegions: []*string{
//   				jsii.String("azureRegions"),
//   			},
//   			ScopeConfiguration: &AzureScopeConfigurationMapProperty{
//   				ContainerImageScanning: &ScopeConfigurationProperty{
//   					ScopeType: jsii.String("scopeType"),
//   					ScopeValues: []*string{
//   						jsii.String("scopeValues"),
//   					},
//   					State: jsii.String("state"),
//   					StateReason: jsii.String("stateReason"),
//   				},
//   				ServerlessScanning: &ScopeConfigurationProperty{
//   					ScopeType: jsii.String("scopeType"),
//   					ScopeValues: []*string{
//   						jsii.String("scopeValues"),
//   					},
//   					State: jsii.String("state"),
//   					StateReason: jsii.String("stateReason"),
//   				},
//   				VmScanning: &ScopeConfigurationProperty{
//   					ScopeType: jsii.String("scopeType"),
//   					ScopeValues: []*string{
//   						jsii.String("scopeValues"),
//   					},
//   					State: jsii.String("state"),
//   					StateReason: jsii.String("stateReason"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-connector.html
//
type CfnConnectorMixinProps struct {
	// Optional description of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-connector.html#cfn-inspectorv2-connector-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name for the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-connector.html#cfn-inspectorv2-connector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-connector.html#cfn-inspectorv2-connector-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-connector.html#cfn-inspectorv2-connector-providerconfiguration
	//
	ProviderConfiguration interface{} `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
	// Tags to apply to the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-connector.html#cfn-inspectorv2-connector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

