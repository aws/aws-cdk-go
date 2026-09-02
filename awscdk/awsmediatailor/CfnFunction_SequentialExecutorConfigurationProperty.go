package awsmediatailor


// Configuration for sequential executor functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sequentialExecutorConfigurationProperty := &SequentialExecutorConfigurationProperty{
//   	FunctionList: []interface{}{
//   		&FunctionRefProperty{
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
	// The list of functions to execute sequentially.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-functionlist
	//
	FunctionList interface{} `field:"required" json:"functionList" yaml:"functionList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The timeout in milliseconds for the entire sequential execution chain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-timeoutmilliseconds
	//
	TimeoutMilliseconds *float64 `field:"required" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

