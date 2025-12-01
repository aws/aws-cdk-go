package previewawsecrmixins


// A signing rule that specifies a signing profile and optional repository filters.
//
// When an image is pushed to a matching repository, a signing job is created using the specified profile.
//
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
	// A list of repository filters that determine which repositories have their images signed on push.
	//
	// If no filters are specified, all images pushed to the registry are signed using the rule's signing profile. Maximum of 100 filters per rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html#cfn-ecr-signingconfiguration-rule-repositoryfilters
	//
	RepositoryFilters interface{} `field:"optional" json:"repositoryFilters" yaml:"repositoryFilters"`
	// The ARN of the AWS Signer signing profile to use for signing images that match this rule.
	//
	// For more information about signing profiles, see [Signing profiles](https://docs.aws.amazon.com/signer/latest/developerguide/signing-profiles.html) in the *AWS Signer Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-signingconfiguration-rule.html#cfn-ecr-signingconfiguration-rule-signingprofilearn
	//
	SigningProfileArn *string `field:"optional" json:"signingProfileArn" yaml:"signingProfileArn"`
}

