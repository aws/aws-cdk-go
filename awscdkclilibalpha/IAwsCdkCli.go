// AWS CDK Programmatic CLI library
package awscdkclilibalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AWS CDK CLI operations.
// Experimental.
type IAwsCdkCli interface {
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

// The jsii proxy for IAwsCdkCli
type jsiiProxy_IAwsCdkCli struct {
	_ byte // padding
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

