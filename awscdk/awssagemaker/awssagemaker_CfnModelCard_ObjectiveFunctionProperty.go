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
type CfnModelCard_ObjectiveFunctionProperty struct {
	// A function object that details optimization direction, metric, and additional descriptions.
	Function interface{} `field:"optional" json:"function" yaml:"function"`
	// Notes about the object function, including other considerations for possible objective functions.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
}

