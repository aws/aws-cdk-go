package awscdkclilibalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AWS CDK CLI operations.
// Deprecated.
type IAwsCdkCli interface {
	// cdk bootstrap.
	// Deprecated.
	Bootstrap(options *BootstrapOptions)
	// cdk deploy.
	// Deprecated.
	Deploy(options *DeployOptions)
	// cdk destroy.
	// Deprecated.
	Destroy(options *DestroyOptions)
	// cdk list.
	// Deprecated.
	List(options *ListOptions)
	// cdk synth.
	// Deprecated.
	Synth(options *SynthOptions)
}

// The jsii proxy for IAwsCdkCli
type jsiiProxy_IAwsCdkCli struct {
	_ byte // padding
}

func (i *jsiiProxy_IAwsCdkCli) Bootstrap(options *BootstrapOptions) {
	if err := i.validateBootstrapParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bootstrap",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IAwsCdkCli) Deploy(options *DeployOptions) {
	if err := i.validateDeployParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"deploy",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IAwsCdkCli) Destroy(options *DestroyOptions) {
	if err := i.validateDestroyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"destroy",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IAwsCdkCli) List(options *ListOptions) {
	if err := i.validateListParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"list",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IAwsCdkCli) Synth(options *SynthOptions) {
	if err := i.validateSynthParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"synth",
		[]interface{}{options},
	)
}

