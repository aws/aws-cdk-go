package interfacesawsmediapackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackagingGroup.
// Experimental.
type IPackagingGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PackagingGroup resource.
	// Experimental.
	PackagingGroupRef() *PackagingGroupReference
}

// The jsii proxy for IPackagingGroupRef
type jsiiProxy_IPackagingGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPackagingGroupRef) PackagingGroupRef() *PackagingGroupReference {
	var returns *PackagingGroupReference
	_jsii_.Get(
		j,
		"packagingGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackagingGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackagingGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

