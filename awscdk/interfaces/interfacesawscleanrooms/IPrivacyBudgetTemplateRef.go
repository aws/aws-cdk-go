package interfacesawscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivacyBudgetTemplate.
// Experimental.
type IPrivacyBudgetTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PrivacyBudgetTemplate resource.
	// Experimental.
	PrivacyBudgetTemplateRef() *PrivacyBudgetTemplateReference
}

// The jsii proxy for IPrivacyBudgetTemplateRef
type jsiiProxy_IPrivacyBudgetTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPrivacyBudgetTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPrivacyBudgetTemplateRef) PrivacyBudgetTemplateRef() *PrivacyBudgetTemplateReference {
	var returns *PrivacyBudgetTemplateReference
	_jsii_.Get(
		j,
		"privacyBudgetTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrivacyBudgetTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrivacyBudgetTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

