package awsec2


// The configuration that links an Amazon VPC IPAM scope to an external authority system.
//
// It specifies the type of external system and the external resource identifier that identifies your account or instance in that system.
//
// In IPAM, an external authority is a third-party IP address management system that provides CIDR blocks when you provision address space for top-level IPAM pools. This allows you to use your existing IP management system to control which address ranges are allocated to AWS while using Amazon VPC IPAM to manage subnets within those ranges.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamScopeExternalAuthorityConfigurationProperty := &IpamScopeExternalAuthorityConfigurationProperty{
//   	ExternalResourceIdentifier: jsii.String("externalResourceIdentifier"),
//   	IpamScopeExternalAuthorityType: jsii.String("ipamScopeExternalAuthorityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamscope-ipamscopeexternalauthorityconfiguration.html
//
type CfnIPAMScope_IpamScopeExternalAuthorityConfigurationProperty struct {
	// The identifier for the external resource managing this scope.
	//
	// For Infoblox integrations, this is the Infoblox resource identifier in the format `<version>.identity.account.<entity_realm>.<entity_id>` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamscope-ipamscopeexternalauthorityconfiguration.html#cfn-ec2-ipamscope-ipamscopeexternalauthorityconfiguration-externalresourceidentifier
	//
	ExternalResourceIdentifier *string `field:"required" json:"externalResourceIdentifier" yaml:"externalResourceIdentifier"`
	// The type of external authority managing this scope.
	//
	// Currently supports `Infoblox` for integration with Infoblox Universal DDI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamscope-ipamscopeexternalauthorityconfiguration.html#cfn-ec2-ipamscope-ipamscopeexternalauthorityconfiguration-ipamscopeexternalauthoritytype
	//
	IpamScopeExternalAuthorityType *string `field:"required" json:"ipamScopeExternalAuthorityType" yaml:"ipamScopeExternalAuthorityType"`
}

