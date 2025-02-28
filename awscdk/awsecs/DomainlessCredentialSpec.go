package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Credential specification for domainless gMSA.
//
// Example:
//   // Make sure the task definition's execution role has permissions to read from the S3 bucket or SSM parameter where the CredSpec file is stored.
//   var bucket bucket
//   var taskDefinition taskDefinition
//
//
//   // Domainless gMSA container from a S3 bucket object.
//   taskDefinition.AddContainer(jsii.String("gmsa-domainless-container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	Cpu: jsii.Number(128),
//   	MemoryLimitMiB: jsii.Number(256),
//   	CredentialSpecs: []credentialSpec{
//   		ecs.DomainlessCredentialSpec_FromS3Bucket(bucket, jsii.String("credSpec")),
//   	},
//   })
//
type DomainlessCredentialSpec interface {
	CredentialSpec
	// Location or ARN from where to retrieve the CredSpec file.
	FileLocation() *string
	// Prefix string based on the type of CredSpec.
	PrefixId() *string
	// Called when the container is initialized to allow this object to bind to the stack.
	Bind() *CredentialSpecConfig
}

// The jsii proxy struct for DomainlessCredentialSpec
type jsiiProxy_DomainlessCredentialSpec struct {
	jsiiProxy_CredentialSpec
}

func (j *jsiiProxy_DomainlessCredentialSpec) FileLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainlessCredentialSpec) PrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixId",
		&returns,
	)
	return returns
}


func NewDomainlessCredentialSpec(fileLocation *string) DomainlessCredentialSpec {
	_init_.Initialize()

	if err := validateNewDomainlessCredentialSpecParameters(fileLocation); err != nil {
		panic(err)
	}
	j := jsiiProxy_DomainlessCredentialSpec{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.DomainlessCredentialSpec",
		[]interface{}{fileLocation},
		&j,
	)

	return &j
}

func NewDomainlessCredentialSpec_Override(d DomainlessCredentialSpec, fileLocation *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.DomainlessCredentialSpec",
		[]interface{}{fileLocation},
		d,
	)
}

// Helper method to generate the ARN for a S3 object.
//
// Used to avoid duplication of logic in derived classes.
func DomainlessCredentialSpec_ArnForS3Object(bucket awss3.IBucket, key *string) *string {
	_init_.Initialize()

	if err := validateDomainlessCredentialSpec_ArnForS3ObjectParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.DomainlessCredentialSpec",
		"arnForS3Object",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Helper method to generate the ARN for a SSM parameter.
//
// Used to avoid duplication of logic in derived classes.
func DomainlessCredentialSpec_ArnForSsmParameter(parameter awsssm.IParameter) *string {
	_init_.Initialize()

	if err := validateDomainlessCredentialSpec_ArnForSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.DomainlessCredentialSpec",
		"arnForSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

// Loads the CredSpec from a S3 bucket object.
//
// Returns: CredSpec with it's locations set to the S3 object's ARN.
func DomainlessCredentialSpec_FromS3Bucket(bucket awss3.IBucket, key *string) DomainlessCredentialSpec {
	_init_.Initialize()

	if err := validateDomainlessCredentialSpec_FromS3BucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns DomainlessCredentialSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.DomainlessCredentialSpec",
		"fromS3Bucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Loads the CredSpec from a SSM parameter.
//
// Returns: CredSpec with it's locations set to the SSM parameter's ARN.
func DomainlessCredentialSpec_FromSsmParameter(parameter awsssm.IParameter) DomainlessCredentialSpec {
	_init_.Initialize()

	if err := validateDomainlessCredentialSpec_FromSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns DomainlessCredentialSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.DomainlessCredentialSpec",
		"fromSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainlessCredentialSpec) Bind() *CredentialSpecConfig {
	var returns *CredentialSpecConfig

	_jsii_.Invoke(
		d,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

