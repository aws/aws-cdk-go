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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-businessdetails.html
//
type CfnModelCard_BusinessDetailsProperty struct {
	// The specific business problem that the model is trying to solve.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-businessdetails.html#cfn-sagemaker-modelcard-businessdetails-businessproblem
	//
	BusinessProblem *string `field:"optional" json:"businessProblem" yaml:"businessProblem"`
	// The relevant stakeholders for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-businessdetails.html#cfn-sagemaker-modelcard-businessdetails-businessstakeholders
	//
	BusinessStakeholders *string `field:"optional" json:"businessStakeholders" yaml:"businessStakeholders"`
	// The broader business need that the model is serving.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-businessdetails.html#cfn-sagemaker-modelcard-businessdetails-lineofbusiness
	//
	LineOfBusiness *string `field:"optional" json:"lineOfBusiness" yaml:"lineOfBusiness"`
}

