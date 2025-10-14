package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InAppTemplate.
// Experimental.
type IInAppTemplateRef interface {
	constructs.IConstruct
	// A reference to a InAppTemplate resource.
	// Experimental.
	InAppTemplateRef() *InAppTemplateReference
}

// The jsii proxy for IInAppTemplateRef
type jsiiProxy_IInAppTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInAppTemplateRef) InAppTemplateRef() *InAppTemplateReference {
	var returns *InAppTemplateReference
	_jsii_.Get(
		j,
		"inAppTemplateRef",
		&returns,
	)
	return returns
}

