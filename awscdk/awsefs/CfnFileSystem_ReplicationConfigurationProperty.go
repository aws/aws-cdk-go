package awsefs


// Describes the replication configuration for a specific file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationProperty := &ReplicationConfigurationProperty{
//   	Destinations: []interface{}{
//   		&ReplicationDestinationProperty{
//   			AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   			FileSystemId: jsii.String("fileSystemId"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			Region: jsii.String("region"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationconfiguration.html
//
type CfnFileSystem_ReplicationConfigurationProperty struct {
	// An array of destination objects.
	//
	// Only one destination object is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationconfiguration.html#cfn-efs-filesystem-replicationconfiguration-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

