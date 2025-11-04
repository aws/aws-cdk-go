package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessEntry.
// Experimental.
type IAccessEntryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AccessEntry resource.
	// Experimental.
	AccessEntryRef() *AccessEntryReference
}

// The jsii proxy for IAccessEntryRef
type jsiiProxy_IAccessEntryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAccessEntryRef) AccessEntryRef() *AccessEntryReference {
	var returns *AccessEntryReference
	_jsii_.Get(
		j,
		"accessEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

