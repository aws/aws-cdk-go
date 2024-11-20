package awsgamelift


// A policy that limits the number of game sessions a player can create on the same fleet.
//
// This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: "An individual player can create a maximum number of new game sessions within a specified time period".
//
// The policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a CreateGameSession request, Amazon GameLift checks that the player (identified by CreatorId) has created fewer than 10 game sessions in the past 60 minutes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gameSessionCreationLimitPolicyProperty := &GameSessionCreationLimitPolicyProperty{
//   	NewGameSessionsPerCreator: jsii.Number(123),
//   	PolicyPeriodInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html
//
type CfnContainerFleet_GameSessionCreationLimitPolicyProperty struct {
	// The maximum number of game sessions that an individual can create during the policy period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy-newgamesessionspercreator
	//
	NewGameSessionsPerCreator *float64 `field:"optional" json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// The time span used in evaluating the resource creation limit policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy-policyperiodinminutes
	//
	PolicyPeriodInMinutes *float64 `field:"optional" json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}

