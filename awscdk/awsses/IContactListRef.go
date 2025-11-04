package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactList.
// Experimental.
type IContactListRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ContactList resource.
	// Experimental.
	ContactListRef() *ContactListReference
}

// The jsii proxy for IContactListRef
type jsiiProxy_IContactListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IContactListRef) ContactListRef() *ContactListReference {
	var returns *ContactListReference
	_jsii_.Get(
		j,
		"contactListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactListRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

