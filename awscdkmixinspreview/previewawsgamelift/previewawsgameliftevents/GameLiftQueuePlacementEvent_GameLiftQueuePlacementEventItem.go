package previewawsgameliftevents


// Type definition for GameLiftQueuePlacementEventItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameLiftQueuePlacementEventItem := &GameLiftQueuePlacementEventItem{
//   	PlayerId: []*string{
//   		jsii.String("playerId"),
//   	},
//   	PlayerSessionId: []*string{
//   		jsii.String("playerSessionId"),
//   	},
//   }
//
// Experimental.
type GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventItem struct {
	// playerId property.
	//
	// Specify an array of string values to match this event if the actual value of playerId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlayerId *[]*string `field:"optional" json:"playerId" yaml:"playerId"`
	// playerSessionId property.
	//
	// Specify an array of string values to match this event if the actual value of playerSessionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlayerSessionId *[]*string `field:"optional" json:"playerSessionId" yaml:"playerSessionId"`
}

