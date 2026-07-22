package awscdk


// An acknowledgment of a validation rule, used to suppress it from output.
//
// Example:
//   awscdk.Annotations_Of(this).AcknowledgeWarning(jsii.String("IAM:Group:MaxPoliciesExceeded"), jsii.String("Account has quota increased to 20"))
//
//   // Because Annotations ultimately become Validations, you can also acknowledge the Validation
//   awscdk.Validations_Of(this).Acknowledge(&Acknowledgment{
//   	Id: jsii.String("Construct-Annotations::IAM:Group:MaxPoliciesExceeded"),
//   	Reason: jsii.String("Account has quota increased to 20"),
//   })
//
type Acknowledgment struct {
	// The rule ID to acknowledge.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The reason for acknowledging this rule.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
}

