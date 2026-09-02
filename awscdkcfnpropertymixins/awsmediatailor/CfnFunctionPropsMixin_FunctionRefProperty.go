package awsmediatailor


// A reference to a function with an optional run condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   functionRefProperty := &FunctionRefProperty{
//   	FunctionId: jsii.String("functionId"),
//   	RunCondition: jsii.String("runCondition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html
//
type CfnFunctionPropsMixin_FunctionRefProperty struct {
	// The identifier of the function to execute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html#cfn-mediatailor-function-functionref-functionid
	//
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// A conditional expression that determines whether this function should execute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-function-functionref.html#cfn-mediatailor-function-functionref-runcondition
	//
	RunCondition *string `field:"optional" json:"runCondition" yaml:"runCondition"`
}

