package interfacesawsgamelift


// A reference to a MatchmakingRuleSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchmakingRuleSetReference := &MatchmakingRuleSetReference{
//   	MatchmakingRuleSetArn: jsii.String("matchmakingRuleSetArn"),
//   	MatchmakingRuleSetName: jsii.String("matchmakingRuleSetName"),
//   }
//
type MatchmakingRuleSetReference struct {
	// The ARN of the MatchmakingRuleSet resource.
	MatchmakingRuleSetArn *string `field:"required" json:"matchmakingRuleSetArn" yaml:"matchmakingRuleSetArn"`
	// The Name of the MatchmakingRuleSet resource.
	MatchmakingRuleSetName *string `field:"required" json:"matchmakingRuleSetName" yaml:"matchmakingRuleSetName"`
}

