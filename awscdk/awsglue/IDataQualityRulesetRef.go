package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataQualityRuleset.
// Experimental.
type IDataQualityRulesetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataQualityRulesetRef
type jsiiProxy_IDataQualityRulesetRef struct {
	internal.Type__constructsIConstruct
}

