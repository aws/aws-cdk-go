package awsce

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CostCategory.
// Experimental.
type ICostCategoryRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICostCategoryRef
type jsiiProxy_ICostCategoryRef struct {
	internal.Type__constructsIConstruct
}

