package previewawsecrmixins


// An array of objects representing the details of a repository filter.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-repositoryfilter.html
//
type CfnSigningConfigurationPropsMixin_RepositoryFilterProperty struct {
	// Repository name pattern (supports '*' wildcard).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-repositoryfilter.html#cfn-ecr-signingconfiguration-repositoryfilter-filter
	//
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Type of repository filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-repositoryfilter.html#cfn-ecr-signingconfiguration-repositoryfilter-filtertype
	//
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

