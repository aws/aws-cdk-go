package mixinsawsvpclattice


// Identifies the resource configuration in one of the following ways:.
//
// - *Amazon Resource Name (ARN)* - Supported resource-types that are provisioned by AWS services, such as RDS databases, can be identified by their ARN.
// - *Domain name* - Any domain name that is publicly resolvable.
// - *IP address* - For IPv4 and IPv6, only IP addresses in the VPC are supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceConfigurationDefinitionProperty := &ResourceConfigurationDefinitionProperty{
//   	ArnResource: jsii.String("arnResource"),
//   	DnsResource: &DnsResourceProperty{
//   		DomainName: jsii.String("domainName"),
//   		IpAddressType: jsii.String("ipAddressType"),
//   	},
//   	IpResource: jsii.String("ipResource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.html
//
type CfnResourceConfigurationPropsMixin_ResourceConfigurationDefinitionProperty struct {
	// The Amazon Resource Name (ARN) of the resource configuration.
	//
	// For the ARN syntax and format, see [ARN format](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html#arns-syntax) in the *AWS Identity and Access Management user guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationdefinition-arnresource
	//
	ArnResource *string `field:"optional" json:"arnResource" yaml:"arnResource"`
	// The DNS name of the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationdefinition-dnsresource
	//
	DnsResource interface{} `field:"optional" json:"dnsResource" yaml:"dnsResource"`
	// The IP address of the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-resourceconfigurationdefinition.html#cfn-vpclattice-resourceconfiguration-resourceconfigurationdefinition-ipresource
	//
	IpResource *string `field:"optional" json:"ipResource" yaml:"ipResource"`
}

