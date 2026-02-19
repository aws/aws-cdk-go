package awssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssigner"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Signer Profile.
type ISigningProfile interface {
	awscdk.IResource
	interfacesawssigner.ISigningProfileRef
	// The ARN of the signing profile.
	SigningProfileArn() *string
	// The name of signing profile.
	SigningProfileName() *string
	// The version of signing profile.
	SigningProfileVersion() *string
	// The ARN of signing profile version.
	SigningProfileVersionArn() *string
}

// The jsii proxy for ISigningProfile
type jsiiProxy_ISigningProfile struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawssignerISigningProfileRef
}

func (i *jsiiProxy_ISigningProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ISigningProfile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISigningProfile) SigningProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) SigningProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) SigningProfileVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) SigningProfileRef() *interfacesawssigner.SigningProfileReference {
	var returns *interfacesawssigner.SigningProfileReference
	_jsii_.Get(
		j,
		"signingProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

