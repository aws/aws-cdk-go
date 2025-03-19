package awssagemaker


// The intended uses of a model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intendedUsesProperty := &IntendedUsesProperty{
//   	ExplanationsForRiskRating: jsii.String("explanationsForRiskRating"),
//   	FactorsAffectingModelEfficiency: jsii.String("factorsAffectingModelEfficiency"),
//   	IntendedUses: jsii.String("intendedUses"),
//   	PurposeOfModel: jsii.String("purposeOfModel"),
//   	RiskRating: jsii.String("riskRating"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-intendeduses.html
//
type CfnModelCard_IntendedUsesProperty struct {
	// An explanation of why your organization categorizes the model with its risk rating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-intendeduses.html#cfn-sagemaker-modelcard-intendeduses-explanationsforriskrating
	//
	ExplanationsForRiskRating *string `field:"optional" json:"explanationsForRiskRating" yaml:"explanationsForRiskRating"`
	// Factors affecting model efficacy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-intendeduses.html#cfn-sagemaker-modelcard-intendeduses-factorsaffectingmodelefficiency
	//
	FactorsAffectingModelEfficiency *string `field:"optional" json:"factorsAffectingModelEfficiency" yaml:"factorsAffectingModelEfficiency"`
	// The intended use cases for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-intendeduses.html#cfn-sagemaker-modelcard-intendeduses-intendeduses
	//
	IntendedUses *string `field:"optional" json:"intendedUses" yaml:"intendedUses"`
	// The general purpose of the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-intendeduses.html#cfn-sagemaker-modelcard-intendeduses-purposeofmodel
	//
	PurposeOfModel *string `field:"optional" json:"purposeOfModel" yaml:"purposeOfModel"`
	// Your organization's risk rating. You can specify one the following values as the risk rating:.
	//
	// - High
	// - Medium
	// - Low
	// - Unknown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-intendeduses.html#cfn-sagemaker-modelcard-intendeduses-riskrating
	//
	RiskRating *string `field:"optional" json:"riskRating" yaml:"riskRating"`
}

