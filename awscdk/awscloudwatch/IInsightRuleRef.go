package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InsightRule.
// Experimental.
type IInsightRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InsightRule resource.
	// Experimental.
	InsightRuleRef() *InsightRuleReference
}

// The jsii proxy for IInsightRuleRef
type jsiiProxy_IInsightRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInsightRuleRef) InsightRuleRef() *InsightRuleReference {
	var returns *InsightRuleReference
	_jsii_.Get(
		j,
		"insightRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInsightRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInsightRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

