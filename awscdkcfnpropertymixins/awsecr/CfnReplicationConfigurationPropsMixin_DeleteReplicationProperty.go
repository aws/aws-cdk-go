package awsecr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   deleteReplicationProperty := &DeleteReplicationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-deletereplication.html
//
type CfnReplicationConfigurationPropsMixin_DeleteReplicationProperty struct {
	// Whether image deletions and image tag removals in the source registry are replicated to the replication destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-deletereplication.html#cfn-ecr-replicationconfiguration-deletereplication-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

