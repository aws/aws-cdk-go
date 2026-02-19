package awskms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Collection of grant methods for an IKey.
//
// Example:
//   var principal IPrincipal
//   var key IKeyRef
//   // can be either an L1 or L2
//
//   kms.KeyGrants_FromKey(key).Sign(principal)
//
type KeyGrants interface {
	Resource() interfacesawskms.IKeyRef
	// Grant the indicated permissions on this key to the given principal.
	//
	// This modifies both the principal's policy as well as the resource policy,
	// since the default CloudFormation setup for KMS keys is that the policy
	// must not be empty and so default grants won't work.
	Actions(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant admins permissions using this key to the given principal.
	//
	// Key administrators have permissions to manage the key (e.g., change permissions, revoke), but do not have permissions
	// to use the key in cryptographic operations (e.g., encrypt, decrypt).
	Admin(grantee awsiam.IGrantable) awsiam.Grant
	// Grant decryption permissions using this key to the given principal.
	Decrypt(grantee awsiam.IGrantable) awsiam.Grant
	// Grant encryption permissions using this key to the given principal.
	Encrypt(grantee awsiam.IGrantable) awsiam.Grant
	// Grant encryption and decryption permissions using this key to the given principal.
	EncryptDecrypt(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to generating MACs to the given principal.
	GenerateMac(grantee awsiam.IGrantable) awsiam.Grant
	// Grant sign permissions using this key to the given principal.
	Sign(grantee awsiam.IGrantable) awsiam.Grant
	// Grant sign and verify permissions using this key to the given principal.
	SignVerify(grantee awsiam.IGrantable) awsiam.Grant
	// Grant verify permissions using this key to the given principal.
	Verify(grantee awsiam.IGrantable) awsiam.Grant
	// Grant permissions to verifying MACs to the given principal.
	VerifyMac(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for KeyGrants
type jsiiProxy_KeyGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_KeyGrants) Resource() interfacesawskms.IKeyRef {
	var returns interfacesawskms.IKeyRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for an IKeyRef.
func KeyGrants_FromKey(resource interfacesawskms.IKeyRef, trustAccountIdentities *bool) KeyGrants {
	_init_.Initialize()

	if err := validateKeyGrants_FromKeyParameters(resource); err != nil {
		panic(err)
	}
	var returns KeyGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kms.KeyGrants",
		"fromKey",
		[]interface{}{resource, trustAccountIdentities},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) Actions(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := k.validateActionsParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"actions",
		args,
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) Admin(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateAdminParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"admin",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) Decrypt(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateDecryptParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"decrypt",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) Encrypt(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateEncryptParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"encrypt",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) EncryptDecrypt(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateEncryptDecryptParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"encryptDecrypt",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) GenerateMac(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateGenerateMacParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"generateMac",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) Sign(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateSignParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"sign",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) SignVerify(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateSignVerifyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"signVerify",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) Verify(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateVerifyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"verify",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyGrants) VerifyMac(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateVerifyMacParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"verifyMac",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

