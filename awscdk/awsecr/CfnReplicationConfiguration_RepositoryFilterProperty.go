package awsecr


// The filter settings used with image replication.
//
// Specifying a repository filter to a replication rule provides a method for controlling which repositories in a private registry are replicated. If no filters are added, the contents of all repositories are replicated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryFilterProperty := &RepositoryFilterProperty{
//   	Filter: jsii.String("filter"),
//   	FilterType: jsii.String("filterType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-repositoryfilter.html
//
type CfnReplicationConfiguration_RepositoryFilterProperty struct {
	// The repository filter details.
	//
	// When the `PREFIX_MATCH` filter type is specified, this value is required and should be the repository name prefix to configure replication for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-repositoryfilter.html#cfn-ecr-replicationconfiguration-repositoryfilter-filter
	//
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The repository filter type.
	//
	// The only supported value is `PREFIX_MATCH` , which is a repository name prefix specified with the `filter` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-replicationconfiguration-repositoryfilter.html#cfn-ecr-replicationconfiguration-repositoryfilter-filtertype
	//
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

