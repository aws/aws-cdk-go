package awssagemaker


// The value of a hyperparameter.
//
// Only one of StringValue or NumberValue can be specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trialComponentParameterValueProperty := &TrialComponentParameterValueProperty{
//   	NumberValue: jsii.Number(123),
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentparametervalue.html
//
type CfnTrialComponent_TrialComponentParameterValueProperty struct {
	// The numeric value of a numeric hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentparametervalue.html#cfn-sagemaker-trialcomponent-trialcomponentparametervalue-numbervalue
	//
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// The string value of a categorical hyperparameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-trialcomponent-trialcomponentparametervalue.html#cfn-sagemaker-trialcomponent-trialcomponentparametervalue-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

