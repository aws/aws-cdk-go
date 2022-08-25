package awsdatabrew


// Represents a transformation and associated parameters that are used to apply a change to an AWS Glue DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	operation: jsii.String("operation"),
//
//   	// the properties below are optional
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   }
//
type CfnRecipe_ActionProperty struct {
	// The name of a valid DataBrew transformation to be performed on the data.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Contextual parameters for the transformation.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

