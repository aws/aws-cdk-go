package interfacesawsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TestGridProject.
// Experimental.
type ITestGridProjectRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TestGridProject resource.
	// Experimental.
	TestGridProjectRef() *TestGridProjectReference
}

// The jsii proxy for ITestGridProjectRef
type jsiiProxy_ITestGridProjectRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITestGridProjectRef) TestGridProjectRef() *TestGridProjectReference {
	var returns *TestGridProjectReference
	_jsii_.Get(
		j,
		"testGridProjectRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITestGridProjectRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITestGridProjectRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

