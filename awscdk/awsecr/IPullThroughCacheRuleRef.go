package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PullThroughCacheRule.
// Experimental.
type IPullThroughCacheRuleRef interface {
	constructs.IConstruct
	// A reference to a PullThroughCacheRule resource.
	// Experimental.
	PullThroughCacheRuleRef() *PullThroughCacheRuleReference
}

// The jsii proxy for IPullThroughCacheRuleRef
type jsiiProxy_IPullThroughCacheRuleRef struct {
	internal.Type__constructsIConstruct
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

