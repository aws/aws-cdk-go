package awsgamelift


// The routing configuration for a fleet alias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingStrategyProperty := &RoutingStrategyProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	FleetId: jsii.String("fleetId"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html
//
type CfnAlias_RoutingStrategyProperty struct {
	// A type of routing strategy.
	//
	// Possible routing types include the following:
	//
	// - *SIMPLE* - The alias resolves to one specific fleet. Use this type when routing to active fleets.
	// - *TERMINAL* - The alias does not resolve to a fleet but instead can be used to display a message to the user. A terminal alias throws a `TerminalRoutingStrategyException` with the message that you specified in the `Message` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A unique identifier for a fleet that the alias points to.
	//
	// If you specify `SIMPLE` for the `Type` property, you must specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-fleetid
	//
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// The message text to be used with a terminal routing strategy.
	//
	// If you specify `TERMINAL` for the `Type` property, you must specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

