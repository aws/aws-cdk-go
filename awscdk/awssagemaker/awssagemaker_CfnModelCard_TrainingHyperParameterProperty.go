package awssagemaker


// A hyper parameter that was configured in training the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingHyperParameterProperty := &trainingHyperParameterProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnModelCard_TrainingHyperParameterProperty struct {
	// The name of the hyper parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value specified for the hyper parameter.
	Value *string `field:"required" json:"value" yaml:"value"`
}

