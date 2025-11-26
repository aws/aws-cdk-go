package awsecr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &RuleProperty{
//   	SigningProfileArn: jsii.String("signingProfileArn"),
//
//   	// the properties below are optional
//   	RepositoryFilters: []interface{}{
//   		&RepositoryFilterProperty{
//   			Filter: jsii.String("filter"),
//   			FilterType: jsii.String("filterType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html
//
type CfnSigningConfiguration_RuleProperty struct {
	// AWS Signer signing profile ARN to use for matched repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html#cfn-ecr-signingconfiguration-rule-signingprofilearn
	//
	SigningProfileArn *string `field:"required" json:"signingProfileArn" yaml:"signingProfileArn"`
	// Optional array of repository filters.
	//
	// If omitted, the rule matches all repositories. If provided, must contain at least one filter. Empty arrays are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html#cfn-ecr-signingconfiguration-rule-repositoryfilters
	//
	RepositoryFilters interface{} `field:"optional" json:"repositoryFilters" yaml:"repositoryFilters"`
}

