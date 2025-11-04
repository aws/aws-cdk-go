package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PullThroughCacheRule.
// Experimental.
type IPullThroughCacheRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PullThroughCacheRule resource.
	// Experimental.
	PullThroughCacheRuleRef() *PullThroughCacheRuleReference
}

// The jsii proxy for IPullThroughCacheRuleRef
type jsiiProxy_IPullThroughCacheRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPullThroughCacheRuleRef) PullThroughCacheRuleRef() *PullThroughCacheRuleReference {
	var returns *PullThroughCacheRuleReference
	_jsii_.Get(
		j,
		"pullThroughCacheRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPullThroughCacheRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPullThroughCacheRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

