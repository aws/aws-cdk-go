package awssynthetics


// Replication status details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   replicaReplicationStatusProperty := &ReplicaReplicationStatusProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replicareplicationstatus.html
//
type CfnCanaryPropsMixin_ReplicaReplicationStatusProperty struct {
	// Replication state: InProgress, InSync, or Inconsistent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-replicareplicationstatus.html#cfn-synthetics-canary-replicareplicationstatus-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

