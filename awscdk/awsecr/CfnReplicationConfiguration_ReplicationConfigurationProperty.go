package awsecr


// The replication configuration for a registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationProperty := &ReplicationConfigurationProperty{
//   	Rules: []interface{}{
//   		&ReplicationRuleProperty{
//   			Destinations: []interface{}{
//   				&ReplicationDestinationProperty{
//   					Region: jsii.String("region"),
//   					RegistryId: jsii.String("registryId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			RepositoryFilters: []interface{}{
//   				&RepositoryFilterProperty{
//   					Filter: jsii.String("filter"),
//   					FilterType: jsii.String("filterType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationconfiguration.html
//
type CfnReplicationConfiguration_ReplicationConfigurationProperty struct {
	// An array of objects representing the replication destinations and repository filters for a replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-replicationconfiguration.html#cfn-ecr-replicationconfiguration-replicationconfiguration-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

