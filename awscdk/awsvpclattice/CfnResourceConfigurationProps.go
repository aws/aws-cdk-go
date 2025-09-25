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
//   	Name: jsii.String("name"),
//   	ResourceConfigurationType: jsii.String("resourceConfigurationType"),
//
//   	// the properties below are optional
//   	AllowAssociationToSharableServiceNetwork: jsii.Boolean(false),
//   	PortRanges: []*string{
//   		jsii.String("portRanges"),
//   	},
//   	ProtocolType: jsii.String("protocolType"),
//   	ResourceConfigurationAuthType: jsii.String("resourceConfigurationAuthType"),
//   	ResourceConfigurationDefinition: &ResourceConfigurationDefinitionProperty{
//   		ArnResource: jsii.String("arnResource"),
//   		DnsResource: &DnsResourceProperty{
//   			DomainName: jsii.String("domainName"),
//   			IpAddressType: jsii.String("ipAddressType"),
//   		},
//   		IpResource: jsii.String("ipResource"),
//   	},
//   	ResourceConfigurationGroupId: jsii.String("resourceConfigurationGroupId"),
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
	// The name of the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of resource configuration. A resource configuration can be one of the following types:.
	//
	// - *SINGLE* - A single resource.
	// - *GROUP* - A group of resources. You must create a group resource configuration before you create a child resource configuration.
	// - *CHILD* - A single resource that is part of a group resource configuration.
	// - *ARN* - An AWS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationtype
	//
	ResourceConfigurationType *string `field:"required" json:"resourceConfigurationType" yaml:"resourceConfigurationType"`
	// Specifies whether the resource configuration can be associated with a sharable service network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-allowassociationtosharableservicenetwork
	//
	AllowAssociationToSharableServiceNetwork interface{} `field:"optional" json:"allowAssociationToSharableServiceNetwork" yaml:"allowAssociationToSharableServiceNetwork"`
	// (SINGLE, GROUP, CHILD) The TCP port ranges that a consumer can use to access a resource configuration (for example: 1-65535).
	//
	// You can separate port ranges using commas (for example: 1,2,22-30).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-portranges
	//
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// (SINGLE, GROUP) The protocol accepted by the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-protocoltype
	//
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// The auth type for the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationauthtype
	//
	ResourceConfigurationAuthType *string `field:"optional" json:"resourceConfigurationAuthType" yaml:"resourceConfigurationAuthType"`
	// Identifies the resource configuration in one of the following ways:.
	//
	// - *Amazon Resource Name (ARN)* - Supported resource-types that are provisioned by AWS services, such as RDS databases, can be identified by their ARN.
	// - *Domain name* - Any domain name that is publicly resolvable.
	// - *IP address* - For IPv4 and IPv6, only IP addresses in the VPC are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationdefinition
	//
	ResourceConfigurationDefinition interface{} `field:"optional" json:"resourceConfigurationDefinition" yaml:"resourceConfigurationDefinition"`
	// The ID of the group resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationgroupid
	//
	ResourceConfigurationGroupId *string `field:"optional" json:"resourceConfigurationGroupId" yaml:"resourceConfigurationGroupId"`
	// The ID of the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-resourcegatewayid
	//
	ResourceGatewayId *string `field:"optional" json:"resourceGatewayId" yaml:"resourceGatewayId"`
	// The tags for the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html#cfn-vpclattice-resourceconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

