package awsgamelift


// A fleet or alias designated in a game session queue.
//
// Queues fulfill requests for new game sessions by placing a new game session on any of the queue's destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &destinationProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   }
//
type CfnGameSessionQueue_DestinationProperty struct {
	// The Amazon Resource Name (ARN) that is assigned to fleet or fleet alias.
	//
	// ARNs, which include a fleet ID or alias ID and a Region name, provide a unique identifier across all Regions.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
}

