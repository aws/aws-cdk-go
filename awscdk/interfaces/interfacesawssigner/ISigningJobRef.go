package interfacesawssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SigningJob.
// Experimental.
type ISigningJobRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SigningJob resource.
	// Experimental.
	SigningJobRef() *SigningJobReference
}

// The jsii proxy for ISigningJobRef
type jsiiProxy_ISigningJobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISigningJobRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISigningJobRef) SigningJobRef() *SigningJobReference {
	var returns *SigningJobReference
	_jsii_.Get(
		j,
		"signingJobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningJobRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningJobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

