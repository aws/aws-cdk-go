package mixinsawsgamelift


// The queue setting that determines the highest latency allowed for individual players when placing a game session.
//
// When a latency policy is in force, a game session cannot be placed with any fleet in a Region where a player reports latency higher than the cap. Latency policies are only enforced when the placement request contains player latency information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   playerLatencyPolicyProperty := &PlayerLatencyPolicyProperty{
//   	MaximumIndividualPlayerLatencyMilliseconds: jsii.Number(123),
//   	PolicyDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-playerlatencypolicy.html
//
type CfnGameSessionQueuePropsMixin_PlayerLatencyPolicyProperty struct {
	// The maximum latency value that is allowed for any player, in milliseconds.
	//
	// All policies must have a value set for this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-playerlatencypolicy.html#cfn-gamelift-gamesessionqueue-playerlatencypolicy-maximumindividualplayerlatencymilliseconds
	//
	MaximumIndividualPlayerLatencyMilliseconds *float64 `field:"optional" json:"maximumIndividualPlayerLatencyMilliseconds" yaml:"maximumIndividualPlayerLatencyMilliseconds"`
	// The length of time, in seconds, that the policy is enforced while placing a new game session.
	//
	// A null value for this property means that the policy is enforced until the queue times out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-playerlatencypolicy.html#cfn-gamelift-gamesessionqueue-playerlatencypolicy-policydurationseconds
	//
	PolicyDurationSeconds *float64 `field:"optional" json:"policyDurationSeconds" yaml:"policyDurationSeconds"`
}

