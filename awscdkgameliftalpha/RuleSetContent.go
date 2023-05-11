// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The rule set determines the two key elements of a match: your game's team structure and size, and how to group players together for the best possible match.
//
// For example, a rule set might describe a match like this:
// - Create a match with two teams of five players each, one team is the defenders and the other team the invaders.
// - A team can have novice and experienced players, but the average skill of the two teams must be within 10 points of each other.
// - If no match is made after 30 seconds, gradually relax the skill requirements.
//
// Example:
//   gamelift.NewMatchmakingRuleSet(this, jsii.String("RuleSet"), &MatchmakingRuleSetProps{
//   	MatchmakingRuleSetName: jsii.String("my-test-ruleset"),
//   	Content: gamelift.RuleSetContent_FromJsonFile(path.join(__dirname, jsii.String("my-ruleset/ruleset.json"))),
//   })
//
// Experimental.
type RuleSetContent interface {
	IRuleSetContent
	// RuleSet body content.
	// Experimental.
	Content() IRuleSetBody
	// Called when the matchmaking ruleSet is initialized to allow this object to bind to the stack and add resources.
	// Experimental.
	Bind(_scope constructs.Construct) *RuleSetBodyConfig
}

// The jsii proxy struct for RuleSetContent
type jsiiProxy_RuleSetContent struct {
	jsiiProxy_IRuleSetContent
}

func (j *jsiiProxy_RuleSetContent) Content() IRuleSetBody {
	var returns IRuleSetBody
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}


// Experimental.
func NewRuleSetContent(props *RuleSetContentProps) RuleSetContent {
	_init_.Initialize()

	if err := validateNewRuleSetContentParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RuleSetContent{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.RuleSetContent",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewRuleSetContent_Override(r RuleSetContent, props *RuleSetContentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.RuleSetContent",
		[]interface{}{props},
		r,
	)
}

// Inline body for Matchmaking ruleSet.
//
// Returns: `RuleSetContent` with inline code.
// Experimental.
func RuleSetContent_FromInline(body *string) IRuleSetContent {
	_init_.Initialize()

	if err := validateRuleSetContent_FromInlineParameters(body); err != nil {
		panic(err)
	}
	var returns IRuleSetContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.RuleSetContent",
		"fromInline",
		[]interface{}{body},
		&returns,
	)

	return returns
}

// Matchmaking ruleSet body from a file.
//
// Returns: `RuleSetContentBase` based on JSON file content.
// Experimental.
func RuleSetContent_FromJsonFile(path *string) IRuleSetContent {
	_init_.Initialize()

	if err := validateRuleSetContent_FromJsonFileParameters(path); err != nil {
		panic(err)
	}
	var returns IRuleSetContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.RuleSetContent",
		"fromJsonFile",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleSetContent) Bind(_scope constructs.Construct) *RuleSetBodyConfig {
	if err := r.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *RuleSetBodyConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

