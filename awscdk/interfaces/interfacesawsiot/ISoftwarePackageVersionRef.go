package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SoftwarePackageVersion.
// Experimental.
type ISoftwarePackageVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SoftwarePackageVersion resource.
	// Experimental.
	SoftwarePackageVersionRef() *SoftwarePackageVersionReference
}

// The jsii proxy for ISoftwarePackageVersionRef
type jsiiProxy_ISoftwarePackageVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISoftwarePackageVersionRef) SoftwarePackageVersionRef() *SoftwarePackageVersionReference {
	var returns *SoftwarePackageVersionReference
	_jsii_.Get(
		j,
		"softwarePackageVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoftwarePackageVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoftwarePackageVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

