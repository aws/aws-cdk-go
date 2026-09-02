package awsmediatailor


// Configuration for sequential executor functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sequentialExecutorConfigurationProperty := &SequentialExecutorConfigurationProperty{
//   	FunctionList: []interface{}{
//   		&FunctionRefProperty{
//   			FunctionId: jsii.String("functionId"),
//   			RunCondition: jsii.String("runCondition"),
//   		},
//   	},
//   	Output: map[string]*string{
//   		"outputKey": jsii.String("output"),
//   	},
//   	Runtime: jsii.String("runtime"),
//   	TimeoutMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html
//
type CfnFunctionPropsMixin_SequentialExecutorConfigurationProperty struct {
	// The list of functions to execute sequentially.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-functionlist
	//
	FunctionList interface{} `field:"optional" json:"functionList" yaml:"functionList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-output
	//
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The timeout in milliseconds for the entire sequential execution chain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-sequentialexecutorconfiguration.html#cfn-mediatailor-function-sequentialexecutorconfiguration-timeoutmilliseconds
	//
	TimeoutMilliseconds *float64 `field:"optional" json:"timeoutMilliseconds" yaml:"timeoutMilliseconds"`
}

