package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivacyBudgetTemplate.
// Experimental.
type IPrivacyBudgetTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPrivacyBudgetTemplateRef
type jsiiProxy_IPrivacyBudgetTemplateRef struct {
	internal.Type__constructsIConstruct
}

