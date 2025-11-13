package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM Access Key.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html
//
type IAccessKey interface {
	interfacesawsiam.IAccessKeyRef
	awscdk.IResource
	// The Access Key ID.
	AccessKeyId() *string
	// The Secret Access Key.
	SecretAccessKey() awscdk.SecretValue
}

// The jsii proxy for IAccessKey
type jsiiProxy_IAccessKey struct {
	internal.Type__interfacesawsiamIAccessKeyRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAccessKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IAccessKey) AccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessKey) SecretAccessKey() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessKey) AccessKeyRef() *interfacesawsiam.AccessKeyReference {
	var returns *interfacesawsiam.AccessKeyReference
	_jsii_.Get(
		j,
		"accessKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessKey) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

