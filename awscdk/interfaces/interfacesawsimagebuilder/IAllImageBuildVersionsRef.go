package interfacesawsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AllImageBuildVersions.
// Experimental.
type IAllImageBuildVersionsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AllImageBuildVersions resource.
	// Experimental.
	AllImageBuildVersionsRef() *AllImageBuildVersionsReference
}

// The jsii proxy for IAllImageBuildVersionsRef
type jsiiProxy_IAllImageBuildVersionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAllImageBuildVersionsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAllImageBuildVersionsRef) AllImageBuildVersionsRef() *AllImageBuildVersionsReference {
	var returns *AllImageBuildVersionsReference
	_jsii_.Get(
		j,
		"allImageBuildVersionsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAllImageBuildVersionsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAllImageBuildVersionsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

