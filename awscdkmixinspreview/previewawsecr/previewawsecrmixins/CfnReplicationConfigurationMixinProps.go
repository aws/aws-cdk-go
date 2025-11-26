package previewawsecrmixins


// Properties for CfnReplicationConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReplicationConfigurationMixinProps := &CfnReplicationConfigurationMixinProps{
//   	ReplicationConfiguration: &ReplicationConfigurationProperty{
//   		Rules: []interface{}{
//   			&ReplicationRuleProperty{
//   				Destinations: []interface{}{
//   					&ReplicationDestinationProperty{
//   						Region: jsii.String("region"),
//   						RegistryId: jsii.String("registryId"),
//   					},
//   				},
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
type CfnReplicationConfigurationMixinProps struct {
	// The replication configuration for a registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-replicationconfiguration.html#cfn-ecr-replicationconfiguration-replicationconfiguration
	//
	ReplicationConfiguration interface{} `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

