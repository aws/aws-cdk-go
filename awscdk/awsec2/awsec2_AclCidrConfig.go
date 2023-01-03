package awsec2


// Acl Configuration for CIDR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclCidrConfig := &aclCidrConfig{
//   	cidrBlock: jsii.String("cidrBlock"),
//   	ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   }
//
// Experimental.
type AclCidrConfig struct {
	// Ipv4 CIDR.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Ipv6 CIDR.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
}

