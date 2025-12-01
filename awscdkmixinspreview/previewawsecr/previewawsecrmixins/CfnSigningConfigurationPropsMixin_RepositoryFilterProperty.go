package previewawsecrmixins


// A repository filter used to determine which repositories have their images automatically signed on push.
//
// Each filter consists of a filter type and filter value.
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
	// The filter value used to match repository names.
	//
	// When using `WILDCARD_MATCH` , the `*` character matches any sequence of characters.
	//
	// Examples:
	//
	// - `myapp/*` - Matches all repositories starting with `myapp/`
	// - `* /production` - Matches all repositories ending with `/production`
	// - `*prod*` - Matches all repositories containing `prod`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-repositoryfilter.html#cfn-ecr-signingconfiguration-repositoryfilter-filter
	//
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// The type of filter to apply.
	//
	// Currently, only `WILDCARD_MATCH` is supported, which uses wildcard patterns to match repository names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-repositoryfilter.html#cfn-ecr-signingconfiguration-repositoryfilter-filtertype
	//
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

