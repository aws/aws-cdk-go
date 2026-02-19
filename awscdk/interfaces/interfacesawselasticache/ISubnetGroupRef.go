package interfacesawselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetGroup.
// Experimental.
type ISubnetGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SubnetGroup resource.
	// Experimental.
	SubnetGroupRef() *SubnetGroupReference
}

// The jsii proxy for ISubnetGroupRef
type jsiiProxy_ISubnetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISubnetGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISubnetGroupRef) SubnetGroupRef() *SubnetGroupReference {
	var returns *SubnetGroupReference
	_jsii_.Get(
		j,
		"subnetGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

