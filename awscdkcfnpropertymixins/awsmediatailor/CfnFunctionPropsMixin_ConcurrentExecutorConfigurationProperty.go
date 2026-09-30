package awsmediatailor


// The configuration for a CONCURRENT_EXECUTOR function.
//
// A CONCURRENT_EXECUTOR runs a set of child functions in parallel, up to a maximum concurrency, and combines their output when all functions complete. For more information about functions, see Working with functions (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions.html) in the MediaTailor User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   concurrentExecutorConfigurationProperty := &ConcurrentExecutorConfigurationProperty{
//   	FunctionList: []interface{}{
//   		&FunctionRefProperty{
//   			Alias: jsii.String("alias"),
//   			FunctionId: jsii.String("functionId"),
//   			RunCondition: jsii.String("runCondition"),
//   		},
//   	},
//   	MaxConcurrency: jsii.Number(123),
//   	Output: map[string]*string{
//   		"outputKey": jsii.String("output"),
//   	},
//   	Runtime: jsii.String("runtime"),
//   	TimeoutMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-concurrentexecutorconfiguration.html
//
type CfnFunctionPropsMixin_ConcurrentExecutorConfigurationProperty struct {
	// The list of 1 to 10 child functions that MediaTailor runs in parallel.
	//
	// Each entry specifies a child function to execute and an optional run condition expression that controls whether the function runs. Child functions cannot themselves be executors, and each child function's resolved namespace must be unique across the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-concurrentexecutorconfiguration.html#cfn-mediatailor-function-concurrentexecutorconfiguration-functionlist
	//
	FunctionList interface{} `field:"optional" json:"functionList" yaml:"functionList"`
	// The maximum number of child functions that MediaTailor runs simultaneously.
	//
	// When the list contains more functions than MaxConcurrency, MediaTailor starts additional functions as running ones complete, so that no more than MaxConcurrency functions run at the same time. Valid values are 1 to 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-concurrentexecutorconfiguration.html#cfn-mediatailor-function-concurrentexecutorconfiguration-maxconcurrency
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// A map of output bindings that controls which bindings the executor commits to the session state after all child functions complete.
	//
	// Each key is a namespaced output path, and each value is an expression that MediaTailor evaluates against the combined results of the child functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-concurrentexecutorconfiguration.html#cfn-mediatailor-function-concurrentexecutorconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// The expression language used to evaluate expressions in the function configuration.
	//
	// Set this to JSONATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-concurrentexecutorconfiguration.html#cfn-mediatailor-function-concurrentexecutorconfiguration-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The maximum time, in milliseconds, for all child functions to complete.
	//
	// This timeout covers every function in the list, including any HTTP calls the child functions make. If the executor exceeds this timeout, MediaTailor discards all output from the executor and proceeds with default behavior. Valid values are 100 to 2000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-concurrentexecutorconfiguration.html#cfn-mediatailor-function-concurrentexecutorconfiguration-timeoutmilliseconds
	//
	TimeoutMilliseconds *float64 `field:"optional" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
}

