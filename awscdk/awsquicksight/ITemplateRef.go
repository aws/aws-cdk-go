package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Template.
// Experimental.
type ITemplateRef interface {
	constructs.IConstruct
	// A reference to a Template resource.
	// Experimental.
	TemplateRef() *TemplateReference
}

// The jsii proxy for ITemplateRef
type jsiiProxy_ITemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITemplateRef) TemplateRef() *TemplateReference {
	var returns *TemplateReference
	_jsii_.Get(
		j,
		"templateRef",
		&returns,
	)
	return returns
}

