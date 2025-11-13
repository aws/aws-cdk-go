package interfacesawsconnect


// A reference to a Rule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleReference := &RuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type RuleReference struct {
	// The RuleArn of the Rule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

