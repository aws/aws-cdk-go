package appstagingsynthesizeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Bootstrapped role specifier.
//
// These roles must exist already.
// This class does not create new IAM Roles.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := awscdk.NewApp(&AppProps{
//   	DefaultStackSynthesizer: appstagingsynthesizeralpha.AppStagingSynthesizer_DefaultResources(&DefaultResourcesOptions{
//   		AppId: jsii.String("my-app-id"),
//   		StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
//   		DeploymentIdentities: appstagingsynthesizeralpha.DeploymentIdentities_SpecifyRoles(&BootstrapRoles{
//   			CloudFormationExecutionRole: appstagingsynthesizeralpha.BootstrapRole_FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Execute")),
//   			DeploymentRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Deploy")),
//   			LookupRole: appstagingsynthesizeralpha.BootstrapRole_*FromRoleArn(jsii.String("arn:aws:iam::123456789012:role/Lookup")),
//   		}),
//   	}),
//   })
//
// Experimental.
type BootstrapRole interface {
	// Whether or not this is object was created using BootstrapRole.cliCredentials().
	// Experimental.
	IsCliCredentials() *bool
}

// The jsii proxy struct for BootstrapRole
type jsiiProxy_BootstrapRole struct {
	_ byte // padding
}

// Use the currently assumed role/credentials.
// Experimental.
func BootstrapRole_CliCredentials() BootstrapRole {
	_init_.Initialize()

	var returns BootstrapRole

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.BootstrapRole",
		"cliCredentials",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Specify an existing IAM Role to assume.
// Experimental.
func BootstrapRole_FromRoleArn(arn *string) BootstrapRole {
	_init_.Initialize()

	if err := validateBootstrapRole_FromRoleArnParameters(arn); err != nil {
		panic(err)
	}
	var returns BootstrapRole

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.BootstrapRole",
		"fromRoleArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstrapRole) IsCliCredentials() *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"isCliCredentials",
		nil, // no parameters
		&returns,
	)

	return returns
}

