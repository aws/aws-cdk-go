package awsxray

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsxray/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SamplingRule.
// Experimental.
type ISamplingRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISamplingRuleRef
type jsiiProxy_ISamplingRuleRef struct {
	internal.Type__constructsIConstruct
}

