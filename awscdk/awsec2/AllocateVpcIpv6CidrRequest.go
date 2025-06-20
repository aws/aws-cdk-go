package awsec2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Request for allocation of the VPC IPv6 CIDR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//
//   allocateVpcIpv6CidrRequest := &AllocateVpcIpv6CidrRequest{
//   	Scope: construct,
//   	VpcId: jsii.String("vpcId"),
//   }
//
type AllocateVpcIpv6CidrRequest struct {
	// The VPC construct to attach to.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
	// The id of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

