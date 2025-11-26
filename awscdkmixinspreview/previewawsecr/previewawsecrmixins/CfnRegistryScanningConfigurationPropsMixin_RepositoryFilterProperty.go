package previewawsecrmixins


// The filter settings used with image replication.
//
// Specifying a repository filter to a replication rule provides a method for controlling which repositories in a private registry are replicated. If no filters are added, the contents of all repositories are replicated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   repositoryFilterProperty := &RepositoryFilterProperty{
//   	Filter: jsii.String("filter"),
//   	FilterType: jsii.String("filterType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-repositoryfilter.html
//
type CfnRegistryScanningConfigurationPropsMixin_RepositoryFilterProperty struct {
	// The filter to use when scanning.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-repositoryfilter.html#cfn-ecr-registryscanningconfiguration-repositoryfilter-filter
	//
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// The type associated with the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-registryscanningconfiguration-repositoryfilter.html#cfn-ecr-registryscanningconfiguration-repositoryfilter-filtertype
	//
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

