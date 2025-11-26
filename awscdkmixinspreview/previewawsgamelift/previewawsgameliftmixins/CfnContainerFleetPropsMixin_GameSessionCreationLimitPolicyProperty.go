package previewawsgameliftmixins


// A policy that puts limits on the number of game sessions that a player can create within a specified span of time.
//
// With this policy, you can control players' ability to consume available resources.
//
// The policy is evaluated when a player tries to create a new game session. On receiving a `CreateGameSession` request, Amazon GameLift Servers checks that the player (identified by `CreatorId` ) has created fewer than game session limit in the specified time period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameSessionCreationLimitPolicyProperty := &GameSessionCreationLimitPolicyProperty{
//   	NewGameSessionsPerCreator: jsii.Number(123),
//   	PolicyPeriodInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html
//
type CfnContainerFleetPropsMixin_GameSessionCreationLimitPolicyProperty struct {
	// A policy that puts limits on the number of game sessions that a player can create within a specified span of time.
	//
	// With this policy, you can control players' ability to consume available resources.
	//
	// The policy evaluates when a player tries to create a new game session. On receiving a `CreateGameSession` request, Amazon GameLift Servers checks that the player (identified by `CreatorId` ) has created fewer than game session limit in the specified time period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy-newgamesessionspercreator
	//
	NewGameSessionsPerCreator *float64 `field:"optional" json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// The time span used in evaluating the resource creation limit policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy-policyperiodinminutes
	//
	PolicyPeriodInMinutes *float64 `field:"optional" json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}

