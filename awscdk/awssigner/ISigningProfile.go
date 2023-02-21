package awssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
)

// A Signer Profile.
type ISigningProfile interface {
	awscdk.IResource
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

