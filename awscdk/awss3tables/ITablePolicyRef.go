package awss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TablePolicy.
// Experimental.
type ITablePolicyRef interface {
	constructs.IConstruct
	// A reference to a TablePolicy resource.
	// Experimental.
	TablePolicyRef() *TablePolicyReference
}

// The jsii proxy for ITablePolicyRef
type jsiiProxy_ITablePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITablePolicyRef) TablePolicyRef() *TablePolicyReference {
	var returns *TablePolicyReference
	_jsii_.Get(
		j,
		"tablePolicyRef",
		&returns,
	)
	return returns
}

