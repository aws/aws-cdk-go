package interfacesawscodeartifact

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodeartifact/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackageGroup.
// Experimental.
type IPackageGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PackageGroup resource.
	// Experimental.
	PackageGroupRef() *PackageGroupReference
}

// The jsii proxy for IPackageGroupRef
type jsiiProxy_IPackageGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPackageGroupRef) PackageGroupRef() *PackageGroupReference {
	var returns *PackageGroupReference
	_jsii_.Get(
		j,
		"packageGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackageGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPackageGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

