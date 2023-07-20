package awsgamelift


// A policy that limits the number of game sessions a player can create on the same fleet.
//
// This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: "An individual player can create a maximum number of new game sessions within a specified time period".
//
// The policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a `CreateGameSession` request, Amazon GameLift checks that the player (identified by `CreatorId` ) has created fewer than 10 game sessions in the past 60 minutes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceCreationLimitPolicyProperty := &ResourceCreationLimitPolicyProperty{
//   	NewGameSessionsPerCreator: jsii.Number(123),
//   	PolicyPeriodInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html
//
type CfnFleet_ResourceCreationLimitPolicyProperty struct {
	// A policy that puts limits on the number of game sessions that a player can create within a specified span of time.
	//
	// With this policy, you can control players' ability to consume available resources.
	//
	// The policy is evaluated when a player tries to create a new game session. On receiving a `CreateGameSession` request, Amazon GameLift checks that the player (identified by `CreatorId` ) has created fewer than game session limit in the specified time period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html#cfn-gamelift-fleet-resourcecreationlimitpolicy-newgamesessionspercreator
	//
	NewGameSessionsPerCreator *float64 `field:"optional" json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// The time span used in evaluating the resource creation limit policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html#cfn-gamelift-fleet-resourcecreationlimitpolicy-policyperiodinminutes
	//
	PolicyPeriodInMinutes *float64 `field:"optional" json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}

