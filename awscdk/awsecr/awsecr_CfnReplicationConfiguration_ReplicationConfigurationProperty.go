package awsecr


// The replication configuration for a registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationProperty := &replicationConfigurationProperty{
//   	rules: []interface{}{
//   		&replicationRuleProperty{
//   			destinations: []interface{}{
//   				&replicationDestinationProperty{
//   					region: jsii.String("region"),
//   					registryId: jsii.String("registryId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			repositoryFilters: []interface{}{
//   				&repositoryFilterProperty{
//   					filter: jsii.String("filter"),
//   					filterType: jsii.String("filterType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnReplicationConfiguration_ReplicationConfigurationProperty struct {
	// An array of objects representing the replication destinations and repository filters for a replication configuration.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

