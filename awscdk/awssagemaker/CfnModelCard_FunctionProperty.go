package awssagemaker


// Function details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionProperty := &FunctionProperty{
//   	Condition: jsii.String("condition"),
//   	Facet: jsii.String("facet"),
//   	Function: jsii.String("function"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-function.html
//
type CfnModelCard_FunctionProperty struct {
	// An optional description of any conditions of your objective function metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-function.html#cfn-sagemaker-modelcard-function-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The metric of the model's objective function.
	//
	// For example, *loss* or *rmse* . The following list shows examples of the values that you can specify for the metric:
	//
	// - `ACCURACY`
	// - `AUC`
	// - `LOSS`
	// - `MAE`
	// - `RMSE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-function.html#cfn-sagemaker-modelcard-function-facet
	//
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// The optimization direction of the model's objective function. You must specify one of the following values:.
	//
	// - `Maximize`
	// - `Minimize`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-function.html#cfn-sagemaker-modelcard-function-function
	//
	Function *string `field:"optional" json:"function" yaml:"function"`
}

