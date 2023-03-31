package awsecr


// An array of objects representing the destination for a replication rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationDestinationProperty := &replicationDestinationProperty{
//   	region: jsii.String("region"),
//   	registryId: jsii.String("registryId"),
//   }
//
type CfnReplicationConfiguration_ReplicationDestinationProperty struct {
	// The Region to replicate to.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The AWS account ID of the Amazon ECR private registry to replicate to.
	//
	// When configuring cross-Region replication within your own registry, specify your own account ID.
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

