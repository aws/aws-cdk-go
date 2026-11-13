package awsgamelift


// The destination to deliver a game session queue's placement logs to.
//
// Omit DestinationArn for a CLOUDWATCH or S3 destination to have GameLift create one; FIREHOSE requires DestinationArn.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deliveryDestinationProperty := &DeliveryDestinationProperty{
//   	LogDestination: jsii.String("logDestination"),
//
//   	// the properties below are optional
//   	DestinationArn: jsii.String("destinationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-deliverydestination.html
//
type CfnGameSessionQueue_DeliveryDestinationProperty struct {
	// The type of service that receives the game session queue's placement logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-deliverydestination.html#cfn-gamelift-gamesessionqueue-deliverydestination-logdestination
	//
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// The Amazon Resource Name (ARN) of the resource that receives the placement logs.
	//
	// The resource type must match LogDestination (a CloudWatch Logs log group, S3 bucket, or Firehose delivery stream). Required for FIREHOSE; optional for CLOUDWATCH and S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-deliverydestination.html#cfn-gamelift-gamesessionqueue-deliverydestination-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
}

