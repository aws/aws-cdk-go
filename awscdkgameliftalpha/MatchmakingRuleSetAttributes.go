package awscdkgameliftalpha


// A full specification of a matchmaking ruleSet that can be used to import it fluently into the CDK application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//
//   matchmakingRuleSetAttributes := &MatchmakingRuleSetAttributes{
//   	MatchmakingRuleSetArn: jsii.String("matchmakingRuleSetArn"),
//   	MatchmakingRuleSetName: jsii.String("matchmakingRuleSetName"),
//   }
//
// Experimental.
type MatchmakingRuleSetAttributes struct {
	// The ARN of the matchmaking ruleSet.
	//
	// At least one of `matchmakingRuleSetArn` and `matchmakingRuleSetName` must be provided.
	// Experimental.
	MatchmakingRuleSetArn *string `field:"optional" json:"matchmakingRuleSetArn" yaml:"matchmakingRuleSetArn"`
	// The unique name of the matchmaking ruleSet.
	//
	// At least one of `ruleSetName` and `matchmakingRuleSetArn`  must be provided.
	// Experimental.
	MatchmakingRuleSetName *string `field:"optional" json:"matchmakingRuleSetName" yaml:"matchmakingRuleSetName"`
}

