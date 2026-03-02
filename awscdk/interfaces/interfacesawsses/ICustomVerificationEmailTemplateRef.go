package interfacesawsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomVerificationEmailTemplate.
// Experimental.
type ICustomVerificationEmailTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CustomVerificationEmailTemplate resource.
	// Experimental.
	CustomVerificationEmailTemplateRef() *CustomVerificationEmailTemplateReference
}

// The jsii proxy for ICustomVerificationEmailTemplateRef
type jsiiProxy_ICustomVerificationEmailTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICustomVerificationEmailTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICustomVerificationEmailTemplateRef) CustomVerificationEmailTemplateRef() *CustomVerificationEmailTemplateReference {
	var returns *CustomVerificationEmailTemplateReference
	_jsii_.Get(
		j,
		"customVerificationEmailTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomVerificationEmailTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomVerificationEmailTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

