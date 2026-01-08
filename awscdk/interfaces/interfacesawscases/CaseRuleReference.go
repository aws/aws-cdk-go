package interfacesawscases


// A reference to a CaseRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caseRuleReference := &CaseRuleReference{
//   	CaseRuleArn: jsii.String("caseRuleArn"),
//   }
//
type CaseRuleReference struct {
	// The CaseRuleArn of the CaseRule resource.
	CaseRuleArn *string `field:"required" json:"caseRuleArn" yaml:"caseRuleArn"`
}

