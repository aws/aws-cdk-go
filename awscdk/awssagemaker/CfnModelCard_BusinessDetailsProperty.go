package awssagemaker


// Information about how the model supports business goals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   businessDetailsProperty := &BusinessDetailsProperty{
//   	BusinessProblem: jsii.String("businessProblem"),
//   	BusinessStakeholders: jsii.String("businessStakeholders"),
//   	LineOfBusiness: jsii.String("lineOfBusiness"),
//   }
//
type CfnModelCard_BusinessDetailsProperty struct {
	// The specific business problem that the model is trying to solve.
	BusinessProblem *string `field:"optional" json:"businessProblem" yaml:"businessProblem"`
	// The relevant stakeholders for the model.
	BusinessStakeholders *string `field:"optional" json:"businessStakeholders" yaml:"businessStakeholders"`
	// The broader business need that the model is serving.
	LineOfBusiness *string `field:"optional" json:"lineOfBusiness" yaml:"lineOfBusiness"`
}

