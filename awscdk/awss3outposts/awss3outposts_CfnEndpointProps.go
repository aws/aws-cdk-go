package awss3outposts


// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &cfnEndpointProps{
//   	outpostId: jsii.String("outpostId"),
//   	securityGroupId: jsii.String("securityGroupId"),
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	accessType: jsii.String("accessType"),
//   	customerOwnedIpv4Pool: jsii.String("customerOwnedIpv4Pool"),
//   }
//
type CfnEndpointProps struct {
	// The ID of the Outpost.
	OutpostId *string `field:"required" json:"outpostId" yaml:"outpostId"`
	// The ID of the security group to use with the endpoint.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The ID of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The container for the type of connectivity used to access the Amazon S3 on Outposts endpoint.
	//
	// To use the Amazon VPC , choose `Private` . To use the endpoint with an on-premises network, choose `CustomerOwnedIp` . If you choose `CustomerOwnedIp` , you must also provide the customer-owned IP address pool (CoIP pool).
	//
	// > `Private` is the default access type value.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// The ID of the customer-owned IPv4 address pool (CoIP pool) for the endpoint.
	//
	// IP addresses are allocated from this pool for the endpoint.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
}

