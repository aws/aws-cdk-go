package interfacesawss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableBucketPolicy.
// Experimental.
type ITableBucketPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TableBucketPolicy resource.
	// Experimental.
	TableBucketPolicyRef() *TableBucketPolicyReference
}

// The jsii proxy for ITableBucketPolicyRef
type jsiiProxy_ITableBucketPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITableBucketPolicyRef) TableBucketPolicyRef() *TableBucketPolicyReference {
	var returns *TableBucketPolicyReference
	_jsii_.Get(
		j,
		"tableBucketPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableBucketPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableBucketPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

