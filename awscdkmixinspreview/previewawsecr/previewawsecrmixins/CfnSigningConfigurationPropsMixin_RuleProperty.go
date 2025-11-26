package previewawsecrmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleProperty := &RuleProperty{
//   	RepositoryFilters: []interface{}{
//   		&RepositoryFilterProperty{
//   			Filter: jsii.String("filter"),
//   			FilterType: jsii.String("filterType"),
//   		},
//   	},
//   	SigningProfileArn: jsii.String("signingProfileArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html
//
type CfnSigningConfigurationPropsMixin_RuleProperty struct {
	// Optional array of repository filters.
	//
	// If omitted, the rule matches all repositories. If provided, must contain at least one filter. Empty arrays are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html#cfn-ecr-signingconfiguration-rule-repositoryfilters
	//
	RepositoryFilters interface{} `field:"optional" json:"repositoryFilters" yaml:"repositoryFilters"`
	// AWS Signer signing profile ARN to use for matched repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html#cfn-ecr-signingconfiguration-rule-signingprofilearn
	//
	SigningProfileArn *string `field:"optional" json:"signingProfileArn" yaml:"signingProfileArn"`
}

