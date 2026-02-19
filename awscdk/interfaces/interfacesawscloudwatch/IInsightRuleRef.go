package interfacesawscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InsightRule.
// Experimental.
type IInsightRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InsightRule resource.
	// Experimental.
	InsightRuleRef() *InsightRuleReference
}

// The jsii proxy for IInsightRuleRef
type jsiiProxy_IInsightRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInsightRuleRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IInsightRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

