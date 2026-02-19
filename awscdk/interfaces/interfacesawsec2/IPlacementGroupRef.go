package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlacementGroup.
// Experimental.
type IPlacementGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PlacementGroup resource.
	// Experimental.
	PlacementGroupRef() *PlacementGroupReference
}

// The jsii proxy for IPlacementGroupRef
type jsiiProxy_IPlacementGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPlacementGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPlacementGroupRef) PlacementGroupRef() *PlacementGroupReference {
	var returns *PlacementGroupReference
	_jsii_.Get(
		j,
		"placementGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

