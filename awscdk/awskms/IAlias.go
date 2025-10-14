package awskms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A KMS Key alias.
//
// An alias can be used in all places that expect a key.
type IAlias interface {
	IAliasRef
	IKey
	// The name of the alias.
	AliasName() *string
	// The Key to which the Alias refers.
	AliasTargetKey() IKey
}

// The jsii proxy for IAlias
type jsiiProxy_IAlias struct {
	jsiiProxy_IAliasRef
	jsiiProxy_IKey
}

func (i *jsiiProxy_IAlias) AddAlias(alias *string) Alias {
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

func (i *jsiiProxy_IAlias) AddToResourcePolicy(statement awsiam.PolicyStatement, allowNoOp *bool) *awsiam.AddToResourcePolicyResult {
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

func (i *jsiiProxy_IAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IAlias) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IAlias) GrantDecrypt(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IAlias) GrantEncrypt(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IAlias) GrantEncryptDecrypt(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IAlias) GrantGenerateMac(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IAlias) GrantSign(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantSignParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantSign",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAlias) GrantSignVerify(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantSignVerifyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantSignVerify",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAlias) GrantVerify(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantVerifyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantVerify",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAlias) GrantVerifyMac(grantee awsiam.IGrantable) awsiam.Grant {
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

func (j *jsiiProxy_IAlias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) AliasTargetKey() IKey {
	var returns IKey
	_jsii_.Get(
		j,
		"aliasTargetKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) AliasRef() *AliasReference {
	var returns *AliasReference
	_jsii_.Get(
		j,
		"aliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) KeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) KeyRef() *KeyReference {
	var returns *KeyReference
	_jsii_.Get(
		j,
		"keyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

