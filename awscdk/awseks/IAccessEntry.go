package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an access entry in an Amazon EKS cluster.
//
// An access entry defines the permissions and scope for a user or role to access an Amazon EKS cluster.
type IAccessEntry interface {
	interfacesawseks.IAccessEntryRef
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the access entry.
	AccessEntryArn() *string
	// The name of the access entry.
	AccessEntryName() *string
}

// The jsii proxy for IAccessEntry
type jsiiProxy_IAccessEntry struct {
	internal.Type__interfacesawseksIAccessEntryRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAccessEntry) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IAccessEntry) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAccessEntry) AccessEntryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessEntryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntry) AccessEntryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessEntryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntry) AccessEntryRef() *interfacesawseks.AccessEntryReference {
	var returns *interfacesawseks.AccessEntryReference
	_jsii_.Get(
		j,
		"accessEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntry) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessEntry) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

