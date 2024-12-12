package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnResourceConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceConfigurationProps := &CfnResourceConfigurationProps{
//   	AllowAssociationToSharableServiceNetwork: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	PortRanges: []*string{
//   		jsii.String("portRanges"),
//   	},
//   	ProtocolType: jsii.String("protocolType"),
//   	ResourceConfigurationAuthType: jsii.String("resourceConfigurationAuthType"),
//   	ResourceConfigurationDefinition: &ResourceConfigurationDefinitionProperty{
//   		IpResource: jsii.String("ipResource"),
//   	},
//   	ResourceConfigurationGroupId: jsii.String("resourceConfigurationGroupId"),
//   	ResourceConfigurationType: jsii.String("resourceConfigurationType"),
//   	ResourceGatewayId: jsii.String("resourceGatewayId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html
//
type CfnResourceConfigurationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-allowassociationtosharableservicenetwork
	//
	AllowAssociationToSharableServiceNetwork interface{} `field:"optional" json:"allowAssociationToSharableServiceNetwork" yaml:"allowAssociationToSharableServiceNetwork"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-portranges
	//
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-protocoltype
	//
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationauthtype
	//
	ResourceConfigurationAuthType *string `field:"optional" json:"resourceConfigurationAuthType" yaml:"resourceConfigurationAuthType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationdefinition
	//
	ResourceConfigurationDefinition interface{} `field:"optional" json:"resourceConfigurationDefinition" yaml:"resourceConfigurationDefinition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationgroupid
	//
	ResourceConfigurationGroupId *string `field:"optional" json:"resourceConfigurationGroupId" yaml:"resourceConfigurationGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationtype
	//
	ResourceConfigurationType *string `field:"optional" json:"resourceConfigurationType" yaml:"resourceConfigurationType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourcegatewayid
	//
	ResourceGatewayId *string `field:"optional" json:"resourceGatewayId" yaml:"resourceGatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

