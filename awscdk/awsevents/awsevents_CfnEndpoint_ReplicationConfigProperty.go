package awsevents


// Endpoints can replicate all events to the secondary Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigProperty := &replicationConfigProperty{
//   	state: jsii.String("state"),
//   }
//
type CfnEndpoint_ReplicationConfigProperty struct {
	// The state of event replication.
	State *string `field:"required" json:"state" yaml:"state"`
}

