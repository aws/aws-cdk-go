package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Base construct for a credential specification (CredSpec).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialSpec := awscdk.Aws_ecs.NewCredentialSpec(jsii.String("prefixId"), jsii.String("fileLocation"))
//
type CredentialSpec interface {
	// Location or ARN from where to retrieve the CredSpec file.
	FileLocation() *string
	// Prefix string based on the type of CredSpec.
	PrefixId() *string
	// Called when the container is initialized to allow this object to bind to the stack.
	Bind() *CredentialSpecConfig
}

// The jsii proxy struct for CredentialSpec
type jsiiProxy_CredentialSpec struct {
	_ byte // padding
}

func (j *jsiiProxy_CredentialSpec) FileLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CredentialSpec) PrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixId",
		&returns,
	)
	return returns
}


func NewCredentialSpec(prefixId *string, fileLocation *string) CredentialSpec {
	_init_.Initialize()

	if err := validateNewCredentialSpecParameters(prefixId, fileLocation); err != nil {
		panic(err)
	}
	j := jsiiProxy_CredentialSpec{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CredentialSpec",
		[]interface{}{prefixId, fileLocation},
		&j,
	)

	return &j
}

func NewCredentialSpec_Override(c CredentialSpec, prefixId *string, fileLocation *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.CredentialSpec",
		[]interface{}{prefixId, fileLocation},
		c,
	)
}

// Helper method to generate the ARN for a S3 object.
//
// Used to avoid duplication of logic in derived classes.
func CredentialSpec_ArnForS3Object(bucket awss3.IBucket, key *string) *string {
	_init_.Initialize()

	if err := validateCredentialSpec_ArnForS3ObjectParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CredentialSpec",
		"arnForS3Object",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Helper method to generate the ARN for a SSM parameter.
//
// Used to avoid duplication of logic in derived classes.
func CredentialSpec_ArnForSsmParameter(parameter awsssm.IParameter) *string {
	_init_.Initialize()

	if err := validateCredentialSpec_ArnForSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.CredentialSpec",
		"arnForSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CredentialSpec) Bind() *CredentialSpecConfig {
	var returns *CredentialSpecConfig

	_jsii_.Invoke(
		c,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

