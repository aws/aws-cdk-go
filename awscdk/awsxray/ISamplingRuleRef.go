package awsxray

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsxray/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SamplingRule.
// Experimental.
type ISamplingRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SamplingRule resource.
	// Experimental.
	SamplingRuleRef() *SamplingRuleReference
}

// The jsii proxy for ISamplingRuleRef
type jsiiProxy_ISamplingRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISamplingRuleRef) SamplingRuleRef() *SamplingRuleReference {
	var returns *SamplingRuleReference
	_jsii_.Get(
		j,
		"samplingRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISamplingRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISamplingRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

