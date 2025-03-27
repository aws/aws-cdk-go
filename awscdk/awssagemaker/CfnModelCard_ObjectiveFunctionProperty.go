package awssagemaker


// The function that is optimized during model training.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectiveFunctionProperty := &ObjectiveFunctionProperty{
//   	Function: &FunctionProperty{
//   		Condition: jsii.String("condition"),
//   		Facet: jsii.String("facet"),
//   		Function: jsii.String("function"),
//   	},
//   	Notes: jsii.String("notes"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-objectivefunction.html
//
type CfnModelCard_ObjectiveFunctionProperty struct {
	// A function object that details optimization direction, metric, and additional descriptions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-objectivefunction.html#cfn-sagemaker-modelcard-objectivefunction-function
	//
	Function interface{} `field:"optional" json:"function" yaml:"function"`
	// Notes about the object function, including other considerations for possible objective functions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-objectivefunction.html#cfn-sagemaker-modelcard-objectivefunction-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
}

