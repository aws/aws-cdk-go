package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

type ISubnet interface {
	awscdk.IResource
	// Associate a Network ACL with this subnet.
	AssociateNetworkAcl(id *string, acl INetworkAcl)
	// The Availability Zone the subnet is located in.
	AvailabilityZone() *string
	// Dependable that can be depended upon to force internet connectivity established on the VPC.
	InternetConnectivityEstablished() constructs.IDependable
	// The IPv4 CIDR block for this subnet.
	Ipv4CidrBlock() *string
	// The route table for this subnet.
	RouteTable() IRouteTable
	// The subnetId for this particular subnet.
	SubnetId() *string
}

// The jsii proxy for ISubnet
type jsiiProxy_ISubnet struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ISubnet) AssociateNetworkAcl(id *string, acl INetworkAcl) {
	if err := i.validateAssociateNetworkAclParameters(id, acl); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateNetworkAcl",
		[]interface{}{id, acl},
	)
}

func (j *jsiiProxy_ISubnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnet) InternetConnectivityEstablished() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"internetConnectivityEstablished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnet) Ipv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnet) RouteTable() IRouteTable {
	var returns IRouteTable
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnet) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

