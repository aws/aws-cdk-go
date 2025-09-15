package awsproton

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentTemplate.
// Experimental.
type IEnvironmentTemplateRef interface {
	constructs.IConstruct
	// A reference to a EnvironmentTemplate resource.
	// Experimental.
	EnvironmentTemplateRef() *EnvironmentTemplateReference
}

// The jsii proxy for IEnvironmentTemplateRef
type jsiiProxy_IEnvironmentTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnvironmentTemplateRef) EnvironmentTemplateRef() *EnvironmentTemplateReference {
	var returns *EnvironmentTemplateReference
	_jsii_.Get(
		j,
		"environmentTemplateRef",
		&returns,
	)
	return returns
}

