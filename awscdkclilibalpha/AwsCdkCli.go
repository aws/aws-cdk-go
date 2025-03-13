package awscdkclilibalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkclilibalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Provides a programmatic interface for interacting with the AWS CDK CLI.
// Experimental.
type AwsCdkCli interface {
	IAwsCdkCli
	// cdk bootstrap.
	// Experimental.
	Bootstrap(options *BootstrapOptions)
	// cdk deploy.
	// Experimental.
	Deploy(options *DeployOptions)
	// cdk destroy.
	// Experimental.
	Destroy(options *DestroyOptions)
	// cdk list.
	// Experimental.
	List(options *ListOptions)
	// cdk synth.
	// Experimental.
	Synth(options *SynthOptions)
}

// The jsii proxy struct for AwsCdkCli
type jsiiProxy_AwsCdkCli struct {
	jsiiProxy_IAwsCdkCli
}

// Create the CLI from a directory containing an AWS CDK app.
//
// Returns: an instance of `AwsCdkCli`.
// Experimental.
func AwsCdkCli_FromCdkAppDirectory(directory *string, props *CdkAppDirectoryProps) AwsCdkCli {
	_init_.Initialize()

	if err := validateAwsCdkCli_FromCdkAppDirectoryParameters(props); err != nil {
		panic(err)
	}
	var returns AwsCdkCli

	_jsii_.StaticInvoke(
		"@aws-cdk/cli-lib-alpha.AwsCdkCli",
		"fromCdkAppDirectory",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Create the CLI from a CloudAssemblyDirectoryProducer.
// Experimental.
func AwsCdkCli_FromCloudAssemblyDirectoryProducer(producer ICloudAssemblyDirectoryProducer) AwsCdkCli {
	_init_.Initialize()

	if err := validateAwsCdkCli_FromCloudAssemblyDirectoryProducerParameters(producer); err != nil {
		panic(err)
	}
	var returns AwsCdkCli

	_jsii_.StaticInvoke(
		"@aws-cdk/cli-lib-alpha.AwsCdkCli",
		"fromCloudAssemblyDirectoryProducer",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkCli) Bootstrap(options *BootstrapOptions) {
	if err := a.validateBootstrapParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bootstrap",
		[]interface{}{options},
	)
}

func (a *jsiiProxy_AwsCdkCli) Deploy(options *DeployOptions) {
	if err := a.validateDeployParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"deploy",
		[]interface{}{options},
	)
}

func (a *jsiiProxy_AwsCdkCli) Destroy(options *DestroyOptions) {
	if err := a.validateDestroyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"destroy",
		[]interface{}{options},
	)
}

func (a *jsiiProxy_AwsCdkCli) List(options *ListOptions) {
	if err := a.validateListParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"list",
		[]interface{}{options},
	)
}

func (a *jsiiProxy_AwsCdkCli) Synth(options *SynthOptions) {
	if err := a.validateSynthParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synth",
		[]interface{}{options},
	)
}

