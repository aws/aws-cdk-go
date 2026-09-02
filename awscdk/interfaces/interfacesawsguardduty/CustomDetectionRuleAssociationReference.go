package interfacesawsguardduty


// A reference to a CustomDetectionRuleAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDetectionRuleAssociationReference := &CustomDetectionRuleAssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   	CustomDetectionRuleAssociationArn: jsii.String("customDetectionRuleAssociationArn"),
//   	RuleId: jsii.String("ruleId"),
//   }
//
type CustomDetectionRuleAssociationReference struct {
	// The AssociationId of the CustomDetectionRuleAssociation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
	// The ARN of the CustomDetectionRuleAssociation resource.
	CustomDetectionRuleAssociationArn *string `field:"required" json:"customDetectionRuleAssociationArn" yaml:"customDetectionRuleAssociationArn"`
	// The RuleId of the CustomDetectionRuleAssociation resource.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
}

