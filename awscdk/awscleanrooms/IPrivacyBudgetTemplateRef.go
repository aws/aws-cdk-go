package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivacyBudgetTemplate.
// Experimental.
type IPrivacyBudgetTemplateRef interface {
	constructs.IConstruct
	// A reference to a PrivacyBudgetTemplate resource.
	// Experimental.
	PrivacyBudgetTemplateRef() *PrivacyBudgetTemplateReference
}

// The jsii proxy for IPrivacyBudgetTemplateRef
type jsiiProxy_IPrivacyBudgetTemplateRef struct {
	internal.Type__constructsIConstruct
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

