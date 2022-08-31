package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
)

// Experimental.
type ISubnet interface {
	awscdk.IResource
	// Associate a Network ACL with this subnet.
	// Experimental.
	AssociateNetworkAcl(id *string, acl INetworkAcl)
	// The Availability Zone the subnet is located in.
	// Experimental.
	AvailabilityZone() *string
	// Dependable that can be depended upon to force internet connectivity established on the VPC.
	// Experimental.
	InternetConnectivityEstablished() awscdk.IDependable
	// The IPv4 CIDR block for this subnet.
	// Experimental.
	Ipv4CidrBlock() *string
	// The route table for this subnet.
	// Experimental.
	RouteTable() IRouteTable
	// The subnetId for this particular subnet.
	// Experimental.
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

func (j *jsiiProxy_ISubnet) InternetConnectivityEstablished() awscdk.IDependable {
	var returns awscdk.IDependable
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

