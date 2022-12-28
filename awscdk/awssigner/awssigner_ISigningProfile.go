package awssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssigner/internal"
)

// A Signer Profile.
// Experimental.
type ISigningProfile interface {
	awscdk.IResource
	// The ARN of the signing profile.
	// Experimental.
	SigningProfileArn() *string
	// The name of signing profile.
	// Experimental.
	SigningProfileName() *string
	// The version of signing profile.
	// Experimental.
	SigningProfileVersion() *string
	// The ARN of signing profile version.
	// Experimental.
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

