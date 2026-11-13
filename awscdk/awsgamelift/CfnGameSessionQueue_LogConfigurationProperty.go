package awsgamelift


// A policy that determines where a game session queue's placement logs are delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	DeliveryDestinations: []interface{}{
//   		&DeliveryDestinationProperty{
//   			LogDestination: jsii.String("logDestination"),
//
//   			// the properties below are optional
//   			DestinationArn: jsii.String("destinationArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-logconfiguration.html
//
type CfnGameSessionQueue_LogConfigurationProperty struct {
	// The set of destinations that receive the game session queue's placement logs.
	//
	// A queue can deliver logs to between 1 and 3 destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-logconfiguration.html#cfn-gamelift-gamesessionqueue-logconfiguration-deliverydestinations
	//
	DeliveryDestinations interface{} `field:"required" json:"deliveryDestinations" yaml:"deliveryDestinations"`
}

