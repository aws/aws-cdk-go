package previewawsgameliftevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.gamelift@GameLiftMatchmakingEvent event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameLiftMatchmakingEventProps := &GameLiftMatchmakingEventProps{
//   	AcceptanceRequired: []*string{
//   		jsii.String("acceptanceRequired"),
//   	},
//   	AcceptanceTimeout: []*string{
//   		jsii.String("acceptanceTimeout"),
//   	},
//   	CustomEventData: []*string{
//   		jsii.String("customEventData"),
//   	},
//   	EstimatedWaitMillis: []*string{
//   		jsii.String("estimatedWaitMillis"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	GameSessionInfo: &GameSessionInfo{
//   		GameSessionArn: []*string{
//   			jsii.String("gameSessionArn"),
//   		},
//   		IpAddress: []*string{
//   			jsii.String("ipAddress"),
//   		},
//   		Players: []GameSessionInfoItem{
//   			&GameSessionInfoItem{
//   				PlayerId: []*string{
//   					jsii.String("playerId"),
//   				},
//   				Team: []*string{
//   					jsii.String("team"),
//   				},
//   			},
//   		},
//   		Port: []*string{
//   			jsii.String("port"),
//   		},
//   	},
//   	MatchId: []*string{
//   		jsii.String("matchId"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	RuleEvaluationMetrics: []RuleEvaluationMetrics{
//   		&RuleEvaluationMetrics{
//   			FailedCount: []*string{
//   				jsii.String("failedCount"),
//   			},
//   			PassedCount: []*string{
//   				jsii.String("passedCount"),
//   			},
//   			RuleName: []*string{
//   				jsii.String("ruleName"),
//   			},
//   		},
//   	},
//   	Tickets: []Ticket{
//   		&Ticket{
//   			Players: []GameSessionInfoItem{
//   				&GameSessionInfoItem{
//   					PlayerId: []*string{
//   						jsii.String("playerId"),
//   					},
//   					Team: []*string{
//   						jsii.String("team"),
//   					},
//   				},
//   			},
//   			StartTime: []*string{
//   				jsii.String("startTime"),
//   			},
//   			TicketId: []*string{
//   				jsii.String("ticketId"),
//   			},
//   		},
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type GameLiftMatchmakingEvent_GameLiftMatchmakingEventProps struct {
	// acceptanceRequired property.
	//
	// Specify an array of string values to match this event if the actual value of acceptanceRequired is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AcceptanceRequired *[]*string `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// acceptanceTimeout property.
	//
	// Specify an array of string values to match this event if the actual value of acceptanceTimeout is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AcceptanceTimeout *[]*string `field:"optional" json:"acceptanceTimeout" yaml:"acceptanceTimeout"`
	// customEventData property.
	//
	// Specify an array of string values to match this event if the actual value of customEventData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CustomEventData *[]*string `field:"optional" json:"customEventData" yaml:"customEventData"`
	// estimatedWaitMillis property.
	//
	// Specify an array of string values to match this event if the actual value of estimatedWaitMillis is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EstimatedWaitMillis *[]*string `field:"optional" json:"estimatedWaitMillis" yaml:"estimatedWaitMillis"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// gameSessionInfo property.
	//
	// Specify an array of string values to match this event if the actual value of gameSessionInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GameSessionInfo *GameLiftMatchmakingEvent_GameSessionInfo `field:"optional" json:"gameSessionInfo" yaml:"gameSessionInfo"`
	// matchId property.
	//
	// Specify an array of string values to match this event if the actual value of matchId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MatchId *[]*string `field:"optional" json:"matchId" yaml:"matchId"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// ruleEvaluationMetrics property.
	//
	// Specify an array of string values to match this event if the actual value of ruleEvaluationMetrics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RuleEvaluationMetrics *[]*GameLiftMatchmakingEvent_RuleEvaluationMetrics `field:"optional" json:"ruleEvaluationMetrics" yaml:"ruleEvaluationMetrics"`
	// tickets property.
	//
	// Specify an array of string values to match this event if the actual value of tickets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tickets *[]*GameLiftMatchmakingEvent_Ticket `field:"optional" json:"tickets" yaml:"tickets"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

