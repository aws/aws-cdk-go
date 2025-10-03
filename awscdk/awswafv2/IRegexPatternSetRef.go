package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegexPatternSet.
// Experimental.
type IRegexPatternSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRegexPatternSetRef
type jsiiProxy_IRegexPatternSetRef struct {
	internal.Type__constructsIConstruct
}

