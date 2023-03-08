package awsec2


// Describes the ICMP type and code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icmpProperty := &icmpProperty{
//   	code: jsii.Number(123),
//   	type: jsii.Number(123),
//   }
//
type CfnNetworkAclEntry_IcmpProperty struct {
	// The Internet Control Message Protocol (ICMP) code.
	//
	// You can use -1 to specify all ICMP codes for the given ICMP type. Requirement is conditional: Required if you specify 1 (ICMP) for the protocol parameter.
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// The Internet Control Message Protocol (ICMP) type.
	//
	// You can use -1 to specify all ICMP types. Conditional requirement: Required if you specify 1 (ICMP) for the `CreateNetworkAclEntry` protocol parameter.
	Type *float64 `field:"optional" json:"type" yaml:"type"`
}

