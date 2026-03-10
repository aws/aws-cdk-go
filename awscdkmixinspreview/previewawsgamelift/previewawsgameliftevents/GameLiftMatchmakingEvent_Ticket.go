package previewawsgameliftevents


// Type definition for Ticket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ticket := &Ticket{
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
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	TicketId: []*string{
//   		jsii.String("ticketId"),
//   	},
//   }
//
// Experimental.
type GameLiftMatchmakingEvent_Ticket struct {
	// players property.
	//
	// Specify an array of string values to match this event if the actual value of players is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Players *[]*GameLiftMatchmakingEvent_GameSessionInfoItem `field:"optional" json:"players" yaml:"players"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// ticketId property.
	//
	// Specify an array of string values to match this event if the actual value of ticketId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TicketId *[]*string `field:"optional" json:"ticketId" yaml:"ticketId"`
}

