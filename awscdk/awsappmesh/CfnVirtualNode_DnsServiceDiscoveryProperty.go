package awsappmesh


// An object that represents the DNS service discovery information for your virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsServiceDiscoveryProperty := &DnsServiceDiscoveryProperty{
//   	Hostname: jsii.String("hostname"),
//
//   	// the properties below are optional
//   	IpPreference: jsii.String("ipPreference"),
//   	ResponseType: jsii.String("responseType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-dnsservicediscovery.html
//
type CfnVirtualNode_DnsServiceDiscoveryProperty struct {
	// Specifies the DNS service discovery hostname for the virtual node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-dnsservicediscovery.html#cfn-appmesh-virtualnode-dnsservicediscovery-hostname
	//
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The preferred IP version that this virtual node uses.
	//
	// Setting the IP preference on the virtual node only overrides the IP preference set for the mesh on this specific node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-dnsservicediscovery.html#cfn-appmesh-virtualnode-dnsservicediscovery-ippreference
	//
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
	// Specifies the DNS response type for the virtual node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-dnsservicediscovery.html#cfn-appmesh-virtualnode-dnsservicediscovery-responsetype
	//
	ResponseType *string `field:"optional" json:"responseType" yaml:"responseType"`
}

