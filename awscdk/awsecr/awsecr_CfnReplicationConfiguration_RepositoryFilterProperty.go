package awsecr


// The filter settings used with image replication.
//
// Specifying a repository filter to a replication rule provides a method for controlling which repositories in a private registry are replicated. If no repository filter is specified, all images in the repository are replicated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryFilterProperty := &repositoryFilterProperty{
//   	filter: jsii.String("filter"),
//   	filterType: jsii.String("filterType"),
//   }
//
type CfnReplicationConfiguration_RepositoryFilterProperty struct {
	// The repository filter details.
	//
	// When the `PREFIX_MATCH` filter type is specified, this value is required and should be the repository name prefix to configure replication for.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The repository filter type.
	//
	// The only supported value is `PREFIX_MATCH` , which is a repository name prefix specified with the `filter` parameter.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

