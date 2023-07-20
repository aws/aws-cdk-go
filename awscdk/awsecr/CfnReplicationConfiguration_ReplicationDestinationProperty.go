package awsecr


// An array of objects representing the destination for a replication rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationDestinationProperty := &ReplicationDestinationProperty{
//   	Region: jsii.String("region"),
//   	RegistryId: jsii.String("registryId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationdestination.html
//
type CfnReplicationConfiguration_ReplicationDestinationProperty struct {
	// The Region to replicate to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationdestination.html#cfn-ecr-replicationconfiguration-replicationdestination-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The AWS account ID of the Amazon ECR private registry to replicate to.
	//
	// When configuring cross-Region replication within your own registry, specify your own account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationdestination.html#cfn-ecr-replicationconfiguration-replicationdestination-registryid
	//
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

