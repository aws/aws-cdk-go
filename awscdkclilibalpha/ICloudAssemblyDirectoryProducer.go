// AWS CDK Programmatic CLI library
package awscdkclilibalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class returning the path to a Cloud Assembly Directory when its `produce` method is invoked with the current context  AWS CDK apps might need to be synthesized multiple times with additional context values before they are ready.
//
// When running the CLI from inside a directory, this is implemented by invoking the app multiple times.
// Here the `produce()` method provides this multi-pass ability.
// Experimental.
type ICloudAssemblyDirectoryProducer interface {
	// Synthesize a Cloud Assembly directory for a given context.
	//
	// For all features to work correctly, `cdk.App()` must be instantiated with the received context values in the method body.
	// Usually obtained similar to this:
	// ```ts fixture=imports
	// class MyProducer implements ICloudAssemblyDirectoryProducer {
	//    async produce(context: Record<string, any>) {
	//      const app = new cdk.App({ context });
	//      // create stacks here
	//      return app.synth().directory;
	//    }
	// }
	// ```.
	// Experimental.
	Produce(context *map[string]interface{}) *string
	// The working directory used to run the Cloud Assembly from.
	//
	// This is were a `cdk.context.json` files will be written.
	// Experimental.
	WorkingDirectory() *string
	// Experimental.
	SetWorkingDirectory(w *string)
}

// The jsii proxy for ICloudAssemblyDirectoryProducer
type jsiiProxy_ICloudAssemblyDirectoryProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_ICloudAssemblyDirectoryProducer) Produce(context *map[string]interface{}) *string {
	if err := i.validateProduceParameters(context); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICloudAssemblyDirectoryProducer) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudAssemblyDirectoryProducer)SetWorkingDirectory(val *string) {
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

