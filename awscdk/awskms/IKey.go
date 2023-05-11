package awskms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms/internal"
)

// A KMS Key, either managed by this CDK app, or imported.
type IKey interface {
	awscdk.IResource
	// Defines a new alias for the key.
	AddAlias(alias *string) Alias
	// Adds a statement to the KMS key resource policy.
	AddToResourcePolicy(statement awsiam.PolicyStatement, allowNoOp *bool) *awsiam.AddToResourcePolicyResult
	// Grant the indicated permissions on this key to the given principal.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant decryption permissions using this key to the given principal.
	GrantDecrypt(grantee awsiam.IGrantable) awsiam.Grant
	// Grant encryption permissions using this key to the given principal.
	GrantEncrypt(grantee awsiam.IGrantable) awsiam.Grant
	// Grant encryption and decryption permissions using this key to the given principal.
	GrantEncryptDecrypt(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to generating MACs to the given principal.
	GrantGenerateMac(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to verifying MACs to the given principal.
	GrantVerifyMac(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the key.
	KeyArn() *string
	// The ID of the key (the part that looks something like: 1234abcd-12ab-34cd-56ef-1234567890ab).
	KeyId() *string
}

// The jsii proxy for IKey
type jsiiProxy_IKey struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKey) AddAlias(alias *string) Alias {
	if err := i.validateAddAliasParameters(alias); err != nil {
		panic(err)
	}
	var returns Alias

	_jsii_.Invoke(
		i,
		"addAlias",
		[]interface{}{alias},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) AddToResourcePolicy(statement awsiam.PolicyStatement, allowNoOp *bool) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement, allowNoOp},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) GrantDecrypt(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantDecryptParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDecrypt",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) GrantEncrypt(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantEncryptParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantEncrypt",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) GrantEncryptDecrypt(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantEncryptDecryptParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantEncryptDecrypt",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) GrantGenerateMac(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantGenerateMacParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantGenerateMac",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKey) GrantVerifyMac(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantVerifyMacParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantVerifyMac",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IKey) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKey) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

