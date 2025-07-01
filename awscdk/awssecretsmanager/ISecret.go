package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
)

// A secret in AWS Secrets Manager.
type ISecret interface {
	awscdk.IResource
	// Adds a rotation schedule to the secret.
	AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule
	// Adds a statement to the IAM resource policy associated with this secret.
	//
	// If this secret was created in this stack, a resource policy will be
	// automatically created upon the first call to `addToResourcePolicy`. If
	// the secret is imported, then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Attach a target to this secret.
	//
	// Returns: An attached secret.
	Attach(target ISecretAttachmentTarget) ISecret
	// Denies the `DeleteSecret` action to all principals within the current account.
	DenyAccountRootDelete()
	// Grants reading the secret value to some role.
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	// Grants writing and updating the secret value to some role.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
	SecretValueFromJson(key *string) awscdk.SecretValue
	// The customer-managed encryption key that is used to encrypt this secret, if any.
	//
	// When not specified, the default
	// KMS key for the account and region is being used.
	EncryptionKey() awskms.IKey
	// The ARN of the secret in AWS Secrets Manager.
	//
	// Will return the full ARN if available, otherwise a partial arn.
	// For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.
	SecretArn() *string
	// The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.
	//
	// This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).
	SecretFullArn() *string
	// The name of the secret.
	//
	// For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
	// '@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.
	SecretName() *string
	// Retrieve the value of the stored secret as a `SecretValue`.
	SecretValue() awscdk.SecretValue
}

// The jsii proxy for ISecret
type jsiiProxy_ISecret struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ISecret) AddRotationSchedule(id *string, options *RotationScheduleOptions) RotationSchedule {
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

func (i *jsiiProxy_ISecret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
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

func (i *jsiiProxy_ISecret) Attach(target ISecretAttachmentTarget) ISecret {
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

func (i *jsiiProxy_ISecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		i,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ISecret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
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

func (i *jsiiProxy_ISecret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_ISecret) SecretValueFromJson(key *string) awscdk.SecretValue {
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

func (j *jsiiProxy_ISecret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

