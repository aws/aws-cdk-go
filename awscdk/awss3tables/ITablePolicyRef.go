package awss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TablePolicy.
// Experimental.
type ITablePolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TablePolicy resource.
	// Experimental.
	TablePolicyRef() *TablePolicyReference
}

// The jsii proxy for ITablePolicyRef
type jsiiProxy_ITablePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITablePolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITablePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

