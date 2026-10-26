package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConsentPortal.
// Experimental.
type IConsentPortalRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConsentPortal resource.
	// Experimental.
	ConsentPortalRef() *ConsentPortalReference
}

// The jsii proxy for IConsentPortalRef
type jsiiProxy_IConsentPortalRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IConsentPortalRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IConsentPortalRef) ConsentPortalRef() *ConsentPortalReference {
	var returns *ConsentPortalReference
	_jsii_.Get(
		j,
		"consentPortalRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsentPortalRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsentPortalRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

