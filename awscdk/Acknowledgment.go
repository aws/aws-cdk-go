package awscdk


// An acknowledgment of a validation rule, used to suppress it from output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   acknowledgment := &Acknowledgment{
//   	Id: jsii.String("id"),
//   	Reason: jsii.String("reason"),
//   }
//
type Acknowledgment struct {
	// The rule ID to acknowledge.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The reason for acknowledging this rule.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
}

