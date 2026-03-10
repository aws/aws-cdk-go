package previewawsgameliftevents


// Type definition for GameSessionInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameSessionInfo := &GameSessionInfo{
//   	GameSessionArn: []*string{
//   		jsii.String("gameSessionArn"),
//   	},
//   	IpAddress: []*string{
//   		jsii.String("ipAddress"),
//   	},
//   	Players: []GameSessionInfoItem{
//   		&GameSessionInfoItem{
//   			PlayerId: []*string{
//   				jsii.String("playerId"),
//   			},
//   			Team: []*string{
//   				jsii.String("team"),
//   			},
//   		},
//   	},
//   	Port: []*string{
//   		jsii.String("port"),
//   	},
//   }
//
// Experimental.
type GameLiftMatchmakingEvent_GameSessionInfo struct {
	// gameSessionArn property.
	//
	// Specify an array of string values to match this event if the actual value of gameSessionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GameSessionArn *[]*string `field:"optional" json:"gameSessionArn" yaml:"gameSessionArn"`
	// ipAddress property.
	//
	// Specify an array of string values to match this event if the actual value of ipAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IpAddress *[]*string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// players property.
	//
	// Specify an array of string values to match this event if the actual value of players is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Players *[]*GameLiftMatchmakingEvent_GameSessionInfoItem `field:"optional" json:"players" yaml:"players"`
	// port property.
	//
	// Specify an array of string values to match this event if the actual value of port is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Port *[]*string `field:"optional" json:"port" yaml:"port"`
}

