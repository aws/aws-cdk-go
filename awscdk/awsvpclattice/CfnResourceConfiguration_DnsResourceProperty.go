package awsvpclattice


// The domain name of the resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsResourceProperty := &DnsResourceProperty{
//   	DomainName: jsii.String("domainName"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-dnsresource.html
//
type CfnResourceConfiguration_DnsResourceProperty struct {
	// The domain name of the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-dnsresource.html#cfn-vpclattice-resourceconfiguration-dnsresource-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The IP address type for the resource configuration.
	//
	// Dualstack is not currently supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-resourceconfiguration-dnsresource.html#cfn-vpclattice-resourceconfiguration-dnsresource-ipaddresstype
	//
	IpAddressType *string `field:"required" json:"ipAddressType" yaml:"ipAddressType"`
}

