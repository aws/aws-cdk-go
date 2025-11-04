package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntelligentPromptRouter.
// Experimental.
type IIntelligentPromptRouterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IntelligentPromptRouter resource.
	// Experimental.
	IntelligentPromptRouterRef() *IntelligentPromptRouterReference
}

// The jsii proxy for IIntelligentPromptRouterRef
type jsiiProxy_IIntelligentPromptRouterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IIntelligentPromptRouterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntelligentPromptRouterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

