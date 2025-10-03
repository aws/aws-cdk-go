package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PullThroughCacheRule.
// Experimental.
type IPullThroughCacheRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPullThroughCacheRuleRef
type jsiiProxy_IPullThroughCacheRuleRef struct {
	internal.Type__constructsIConstruct
}

