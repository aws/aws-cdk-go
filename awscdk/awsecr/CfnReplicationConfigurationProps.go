package awsecr


// Properties for defining a `CfnReplicationConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationConfigurationProps := &CfnReplicationConfigurationProps{
//   	ReplicationConfiguration: &ReplicationConfigurationProperty{
//   		Rules: []interface{}{
//   			&ReplicationRuleProperty{
//   				Destinations: []interface{}{
//   					&ReplicationDestinationProperty{
//   						Region: jsii.String("region"),
//   						RegistryId: jsii.String("registryId"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				RepositoryFilters: []interface{}{
//   					&RepositoryFilterProperty{
//   						Filter: jsii.String("filter"),
//   						FilterType: jsii.String("filterType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-replicationconfiguration.html
//
type CfnReplicationConfigurationProps struct {
	// The replication configuration for a registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-replicationconfiguration.html#cfn-ecr-replicationconfiguration-replicationconfiguration
	//
	ReplicationConfiguration interface{} `field:"required" json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

