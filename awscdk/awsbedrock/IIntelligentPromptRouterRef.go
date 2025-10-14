package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntelligentPromptRouter.
// Experimental.
type IIntelligentPromptRouterRef interface {
	constructs.IConstruct
	// A reference to a IntelligentPromptRouter resource.
	// Experimental.
	IntelligentPromptRouterRef() *IntelligentPromptRouterReference
}

// The jsii proxy for IIntelligentPromptRouterRef
type jsiiProxy_IIntelligentPromptRouterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIntelligentPromptRouterRef) IntelligentPromptRouterRef() *IntelligentPromptRouterReference {
	var returns *IntelligentPromptRouterReference
	_jsii_.Get(
		j,
		"intelligentPromptRouterRef",
		&returns,
	)
	return returns
}

