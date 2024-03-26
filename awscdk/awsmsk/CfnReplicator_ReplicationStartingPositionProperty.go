package awsmsk


// Configuration for specifying the position in the topics to start replicating from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationStartingPositionProperty := &ReplicationStartingPositionProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationstartingposition.html
//
type CfnReplicator_ReplicationStartingPositionProperty struct {
	// The type of replication starting position.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-replicationstartingposition.html#cfn-msk-replicator-replicationstartingposition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

