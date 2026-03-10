package previewawsgameliftevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.gamelift@GameLiftQueuePlacementEvent event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameLiftQueuePlacementEventProps := &GameLiftQueuePlacementEventProps{
//   	DnsName: []*string{
//   		jsii.String("dnsName"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
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
//   	GameSessionArn: []*string{
//   		jsii.String("gameSessionArn"),
//   	},
//   	GameSessionRegion: []*string{
//   		jsii.String("gameSessionRegion"),
//   	},
//   	IpAddress: []*string{
//   		jsii.String("ipAddress"),
//   	},
//   	PlacedPlayerSessions: []GameLiftQueuePlacementEventItem{
//   		&GameLiftQueuePlacementEventItem{
//   			PlayerId: []*string{
//   				jsii.String("playerId"),
//   			},
//   			PlayerSessionId: []*string{
//   				jsii.String("playerSessionId"),
//   			},
//   		},
//   	},
//   	PlacementId: []*string{
//   		jsii.String("placementId"),
//   	},
//   	Port: []*string{
//   		jsii.String("port"),
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventProps struct {
	// dnsName property.
	//
	// Specify an array of string values to match this event if the actual value of dnsName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DnsName *[]*string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// endTime property.
	//
	// Specify an array of string values to match this event if the actual value of endTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// gameSessionArn property.
	//
	// Specify an array of string values to match this event if the actual value of gameSessionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GameSessionArn *[]*string `field:"optional" json:"gameSessionArn" yaml:"gameSessionArn"`
	// gameSessionRegion property.
	//
	// Specify an array of string values to match this event if the actual value of gameSessionRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	GameSessionRegion *[]*string `field:"optional" json:"gameSessionRegion" yaml:"gameSessionRegion"`
	// ipAddress property.
	//
	// Specify an array of string values to match this event if the actual value of ipAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IpAddress *[]*string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// placedPlayerSessions property.
	//
	// Specify an array of string values to match this event if the actual value of placedPlayerSessions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlacedPlayerSessions *[]*GameLiftQueuePlacementEvent_GameLiftQueuePlacementEventItem `field:"optional" json:"placedPlayerSessions" yaml:"placedPlayerSessions"`
	// placementId property.
	//
	// Specify an array of string values to match this event if the actual value of placementId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PlacementId *[]*string `field:"optional" json:"placementId" yaml:"placementId"`
	// port property.
	//
	// Specify an array of string values to match this event if the actual value of port is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Port *[]*string `field:"optional" json:"port" yaml:"port"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

