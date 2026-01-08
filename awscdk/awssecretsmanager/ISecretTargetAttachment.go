package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

type ISecretTargetAttachment interface {
	ISecret
	interfacesawssecretsmanager.ISecretTargetAttachmentRef
	// Same as `secretArn`.
	SecretTargetAttachmentSecretArn() *string
}

// The jsii proxy for ISecretTargetAttachment
type jsiiProxy_ISecretTargetAttachment struct {
	jsiiProxy_ISecret
	internal.Type__interfacesawssecretsmanagerISecretTargetAttachmentRef
}

func (i *jsiiProxy_ISecretTargetAttachment) AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule {
	if err := i.validateAddRotationScheduleParameters(id, options); err != nil {
		panic(err)
	}
	var returns RotationSchedule

	_jsii_.Invoke(
		i,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecretTargetAttachment) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecretTargetAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISecretTargetAttachment) Attach(target ISecretAttachmentTarget) ISecret {
	if err := i.validateAttachParameters(target); err != nil {
		panic(err)
	}
	var returns ISecret

	_jsii_.Invoke(
		i,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecretTargetAttachment) CfnDynamicReferenceKey(options *awscdk.SecretsManagerSecretOptions) *string {
	if err := i.validateCfnDynamicReferenceKeyParameters(options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"cfnDynamicReferenceKey",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecretTargetAttachment) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		i,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ISecretTargetAttachment) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecretTargetAttachment) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISecretTargetAttachment) SecretValueFromJson(key *string) awscdk.SecretValue {
	if err := i.validateSecretValueFromJsonParameters(key); err != nil {
		panic(err)
	}
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		i,
		"secretValueFromJson",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretTargetAttachmentSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretTargetAttachmentSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretRef() *interfacesawssecretsmanager.SecretReference {
	var returns *interfacesawssecretsmanager.SecretReference
	_jsii_.Get(
		j,
		"secretRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretTargetAttachmentRef() *interfacesawssecretsmanager.SecretTargetAttachmentReference {
	var returns *interfacesawssecretsmanager.SecretTargetAttachmentReference
	_jsii_.Get(
		j,
		"secretTargetAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

