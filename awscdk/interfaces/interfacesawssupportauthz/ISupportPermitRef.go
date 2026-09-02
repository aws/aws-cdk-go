package interfacesawssupportauthz

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssupportauthz/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SupportPermit.
// Experimental.
type ISupportPermitRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SupportPermit resource.
	// Experimental.
	SupportPermitRef() *SupportPermitReference
}

// The jsii proxy for ISupportPermitRef
type jsiiProxy_ISupportPermitRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISupportPermitRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISupportPermitRef) SupportPermitRef() *SupportPermitReference {
	var returns *SupportPermitReference
	_jsii_.Get(
		j,
		"supportPermitRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISupportPermitRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISupportPermitRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

