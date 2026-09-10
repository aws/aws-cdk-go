package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// VPC network placement for a Glue `Connection`.
//
// A Glue connection targets a single subnet. Choose the placement with one of
// the mutually-exclusive factories — an explicit subnet, or a VPC to select one
// from — so a subnet paired with a VPC, or a subnet selection without a VPC,
// cannot be expressed.
//
// Example:
//   var securityGroup SecurityGroup
//   var vpc Vpc
//
//   glue.NewConnection(this, jsii.String("MyConnection"), &ConnectionProps{
//   	Type: glue.ConnectionType_NETWORK(),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	// vpcSubnets is optional - defaults to private subnets
//   	Network: glue.ConnectionNetwork_Vpc(vpc, &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	}),
//   })
//
// Experimental.
type ConnectionNetwork interface {
}

// The jsii proxy struct for ConnectionNetwork
type jsiiProxy_ConnectionNetwork struct {
	_ byte // padding
}

// Pin the connection to a specific subnet.
// Experimental.
func ConnectionNetwork_Subnet(subnet awsec2.ISubnet) ConnectionNetwork {
	_init_.Initialize()

	if err := validateConnectionNetwork_SubnetParameters(subnet); err != nil {
		panic(err)
	}
	var returns ConnectionNetwork

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.ConnectionNetwork",
		"subnet",
		[]interface{}{subnet},
		&returns,
	)

	return returns
}

// Select the connection's subnet from a VPC.
//
// Since a Glue connection targets
// a single subnet, the first subnet of the selection is used.
// Default: vpcSubnets - private subnets.
//
// Experimental.
func ConnectionNetwork_Vpc(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) ConnectionNetwork {
	_init_.Initialize()

	if err := validateConnectionNetwork_VpcParameters(vpc, vpcSubnets); err != nil {
		panic(err)
	}
	var returns ConnectionNetwork

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.ConnectionNetwork",
		"vpc",
		[]interface{}{vpc, vpcSubnets},
		&returns,
	)

	return returns
}

