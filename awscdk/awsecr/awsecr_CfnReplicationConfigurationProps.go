package awsecr


// Properties for defining a `CfnReplicationConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationConfigurationProps := &cfnReplicationConfigurationProps{
//   	replicationConfiguration: &replicationConfigurationProperty{
//   		rules: []interface{}{
//   			&replicationRuleProperty{
//   				destinations: []interface{}{
//   					&replicationDestinationProperty{
//   						region: jsii.String("region"),
//   						registryId: jsii.String("registryId"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				repositoryFilters: []interface{}{
//   					&repositoryFilterProperty{
//   						filter: jsii.String("filter"),
//   						filterType: jsii.String("filterType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnReplicationConfigurationProps struct {
	// The replication configuration for a registry.
	ReplicationConfiguration interface{} `field:"required" json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

