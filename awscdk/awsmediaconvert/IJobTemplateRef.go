package awsmediaconvert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobTemplate.
// Experimental.
type IJobTemplateRef interface {
	constructs.IConstruct
	// A reference to a JobTemplate resource.
	// Experimental.
	JobTemplateRef() *JobTemplateReference
}

// The jsii proxy for IJobTemplateRef
type jsiiProxy_IJobTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJobTemplateRef) JobTemplateRef() *JobTemplateReference {
	var returns *JobTemplateReference
	_jsii_.Get(
		j,
		"jobTemplateRef",
		&returns,
	)
	return returns
}

