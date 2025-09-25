package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationInferenceProfile.
// Experimental.
type IApplicationInferenceProfileRef interface {
	constructs.IConstruct
	// A reference to a ApplicationInferenceProfile resource.
	// Experimental.
	ApplicationInferenceProfileRef() *ApplicationInferenceProfileReference
}

// The jsii proxy for IApplicationInferenceProfileRef
type jsiiProxy_IApplicationInferenceProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationInferenceProfileRef) ApplicationInferenceProfileRef() *ApplicationInferenceProfileReference {
	var returns *ApplicationInferenceProfileReference
	_jsii_.Get(
		j,
		"applicationInferenceProfileRef",
		&returns,
	)
	return returns
}

