package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRouteTable.
// Experimental.
type ILocalGatewayRouteTableRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocalGatewayRouteTable resource.
	// Experimental.
	LocalGatewayRouteTableRef() *LocalGatewayRouteTableReference
}

// The jsii proxy for ILocalGatewayRouteTableRef
type jsiiProxy_ILocalGatewayRouteTableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILocalGatewayRouteTableRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableRef) LocalGatewayRouteTableRef() *LocalGatewayRouteTableReference {
	var returns *LocalGatewayRouteTableReference
	_jsii_.Get(
		j,
		"localGatewayRouteTableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteTableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

