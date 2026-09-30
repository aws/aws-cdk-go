package awsmediatailor


// The configuration for a SEQUENTIAL_EXECUTOR function.
//
// A SEQUENTIAL_EXECUTOR runs an ordered list of child functions one at a time, passing data between them. For more information about functions, see Working with functions (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions.html) in the MediaTailor User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sequentialExecutorConfigurationProperty := &SequentialExecutorConfigurationProperty{
//   	FunctionList: []interface{}{
//   		&FunctionRefProperty{
//   			Alias: jsii.String("alias"),
//   			FunctionId: jsii.String("functionId"),
//   			RunCondition: jsii.String("runCondition"),
//   		},
//   	},
//   	Runtime: jsii.String("runtime"),
//   	TimeoutMilliseconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	Output: map[string]*string{
//   		"outputKey": jsii.String("output"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html
//
type CfnFunction_SequentialExecutorConfigurationProperty struct {
	// An ordered list of 1 to 10 steps.
	//
	// Each step specifies a child function to execute and an optional run condition expression that controls whether the step runs. MediaTailor executes the steps in order, passing data between steps through temporary data. Each step's resolved namespace must be unique across the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-functionlist
	//
	FunctionList interface{} `field:"required" json:"functionList" yaml:"functionList"`
	// The expression language used to evaluate expressions in the function configuration.
	//
	// Set this to JSONATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The maximum time, in milliseconds, for the entire sequence to complete.
	//
	// This timeout covers all steps, including any HTTP calls made by child functions. If the sequence exceeds this timeout, MediaTailor discards all output from the sequence and proceeds with default behavior. Valid values are 100 to 2000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-timeoutmilliseconds
	//
	TimeoutMilliseconds *float64 `field:"required" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
	// A map of output bindings that controls which bindings the sequence commits to the session state after all steps complete.
	//
	// Each key is a namespaced output path, and each value is an expression that MediaTailor evaluates against the accumulated results of the steps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

