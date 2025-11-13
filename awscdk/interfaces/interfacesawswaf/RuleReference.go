package interfacesawswaf


// A reference to a Rule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleReference := &RuleReference{
//   	RuleId: jsii.String("ruleId"),
//   }
//
type RuleReference struct {
	// The Id of the Rule resource.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
}

