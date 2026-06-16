package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MlflowApp.
// Experimental.
type IMlflowAppRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MlflowApp resource.
	// Experimental.
	MlflowAppRef() *MlflowAppReference
}

// The jsii proxy for IMlflowAppRef
type jsiiProxy_IMlflowAppRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMlflowAppRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMlflowAppRef) MlflowAppRef() *MlflowAppReference {
	var returns *MlflowAppReference
	_jsii_.Get(
		j,
		"mlflowAppRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMlflowAppRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMlflowAppRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

