package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageVersion.
// Experimental.
type IImageVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ImageVersion resource.
	// Experimental.
	ImageVersionRef() *ImageVersionReference
}

// The jsii proxy for IImageVersionRef
type jsiiProxy_IImageVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IImageVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IImageVersionRef) ImageVersionRef() *ImageVersionReference {
	var returns *ImageVersionReference
	_jsii_.Get(
		j,
		"imageVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

