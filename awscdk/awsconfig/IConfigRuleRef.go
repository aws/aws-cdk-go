package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigRule.
// Experimental.
type IConfigRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigRuleRef
type jsiiProxy_IConfigRuleRef struct {
	internal.Type__constructsIConstruct
}

