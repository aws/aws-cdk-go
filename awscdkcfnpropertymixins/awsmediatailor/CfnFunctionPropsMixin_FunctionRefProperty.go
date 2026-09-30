package awsmediatailor


// A reference to a child function within an executor function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   functionRefProperty := &FunctionRefProperty{
//   	Alias: jsii.String("alias"),
//   	FunctionId: jsii.String("functionId"),
//   	RunCondition: jsii.String("runCondition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html
//
type CfnFunctionPropsMixin_FunctionRefProperty struct {
	// An optional alternate name for the child function within the executor.
	//
	// MediaTailor uses this value as the namespace for the child function's output. If omitted, MediaTailor uses the function identifier. The resolved namespace must be unique across all child functions in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html#cfn-mediatailor-function-functionref-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The identifier of the child function to execute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html#cfn-mediatailor-function-functionref-functionid
	//
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// An optional expression that evaluates to a boolean.
	//
	// MediaTailor evaluates this expression immediately before running the child function, using the accumulated state at that point. If the expression evaluates to false, MediaTailor skips the child function. If omitted, the child function always runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html#cfn-mediatailor-function-functionref-runcondition
	//
	RunCondition *string `field:"optional" json:"runCondition" yaml:"runCondition"`
}

