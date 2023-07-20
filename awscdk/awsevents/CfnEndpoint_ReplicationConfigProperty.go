package awsevents


// Endpoints can replicate all events to the secondary Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigProperty := &ReplicationConfigProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-replicationconfig.html
//
type CfnEndpoint_ReplicationConfigProperty struct {
	// The state of event replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-replicationconfig.html#cfn-events-endpoint-replicationconfig-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

