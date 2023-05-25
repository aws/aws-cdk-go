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
type CfnModelCard_IntendedUsesProperty struct {
	// An explanation of why your organization categorizes the model with its risk rating.
	ExplanationsForRiskRating *string `field:"optional" json:"explanationsForRiskRating" yaml:"explanationsForRiskRating"`
	// Factors affecting model efficacy.
	FactorsAffectingModelEfficiency *string `field:"optional" json:"factorsAffectingModelEfficiency" yaml:"factorsAffectingModelEfficiency"`
	// The intended use cases for the model.
	IntendedUses *string `field:"optional" json:"intendedUses" yaml:"intendedUses"`
	// The general purpose of the model.
	PurposeOfModel *string `field:"optional" json:"purposeOfModel" yaml:"purposeOfModel"`
	// Your organization's risk rating. You can specify one the following values as the risk rating:.
	//
	// - High
	// - Medium
	// - Low
	// - Unknown.
	RiskRating *string `field:"optional" json:"riskRating" yaml:"riskRating"`
}

