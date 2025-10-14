package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CostCategory.
// Experimental.
type ICostCategoryRef interface {
	constructs.IConstruct
	// A reference to a CostCategory resource.
	// Experimental.
	CostCategoryRef() *CostCategoryReference
}

// The jsii proxy for ICostCategoryRef
type jsiiProxy_ICostCategoryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICostCategoryRef) CostCategoryRef() *CostCategoryReference {
	var returns *CostCategoryReference
	_jsii_.Get(
		j,
		"costCategoryRef",
		&returns,
	)
	return returns
}

