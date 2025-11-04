package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivacyBudgetTemplate.
// Experimental.
type IPrivacyBudgetTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PrivacyBudgetTemplate resource.
	// Experimental.
	PrivacyBudgetTemplateRef() *PrivacyBudgetTemplateReference
}

// The jsii proxy for IPrivacyBudgetTemplateRef
type jsiiProxy_IPrivacyBudgetTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IPrivacyBudgetTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

