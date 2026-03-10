package previewawsgameliftevents


// Type definition for GameSessionInfoItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameSessionInfoItem := &GameSessionInfoItem{
//   	PlayerId: []*string{
//   		jsii.String("playerId"),
//   	},
//   	Team: []*string{
//   		jsii.String("team"),
//   	},
//   }
//
// Experimental.
type GameLiftMatchmakingEvent_GameSessionInfoItem struct {
	// playerId property.
	//
	// Specify an array of string values to match this event if the actual value of playerId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlayerId *[]*string `field:"optional" json:"playerId" yaml:"playerId"`
	// team property.
	//
	// Specify an array of string values to match this event if the actual value of team is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Team *[]*string `field:"optional" json:"team" yaml:"team"`
}

