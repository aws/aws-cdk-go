package awscdkgameliftalpha


// Properties for a new matchmaking ruleSet.
//
// Example:
//   gamelift.NewMatchmakingRuleSet(this, jsii.String("RuleSet"), &MatchmakingRuleSetProps{
//   	MatchmakingRuleSetName: jsii.String("my-test-ruleset"),
//   	Content: gamelift.RuleSetContent_FromJsonFile(path.join(__dirname, jsii.String("my-ruleset/ruleset.json"))),
//   })
//
// Experimental.
type MatchmakingRuleSetProps struct {
	// A collection of matchmaking rules.
	// Experimental.
	Content RuleSetContent `field:"required" json:"content" yaml:"content"`
	// A unique identifier for the matchmaking rule set.
	//
	// A matchmaking configuration identifies the rule set it uses by this name value.
	//
	// Note: the rule set name is different from the optional name field in the rule set body.
	// Experimental.
	MatchmakingRuleSetName *string `field:"required" json:"matchmakingRuleSetName" yaml:"matchmakingRuleSetName"`
}

