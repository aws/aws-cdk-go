package previewawscustomerprofilesmixins


// The process of matching duplicate profiles using Rule-Based matching.
//
// If `RuleBasedMatching = true` , Amazon Connect Customer Profiles will start to match and merge your profiles according to your configuration in the `RuleBasedMatchingRequest` . You can use the `ListRuleBasedMatches` and `GetSimilarProfiles` API to return and review the results. Also, if you have configured `ExportingConfig` in the `RuleBasedMatchingRequest` , you can download the results from S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleBasedMatchingProperty := &RuleBasedMatchingProperty{
//   	AttributeTypesSelector: &AttributeTypesSelectorProperty{
//   		Address: []*string{
//   			jsii.String("address"),
//   		},
//   		AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   		EmailAddress: []*string{
//   			jsii.String("emailAddress"),
//   		},
//   		PhoneNumber: []*string{
//   			jsii.String("phoneNumber"),
//   		},
//   	},
//   	ConflictResolution: &ConflictResolutionProperty{
//   		ConflictResolvingModel: jsii.String("conflictResolvingModel"),
//   		SourceName: jsii.String("sourceName"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	ExportingConfig: &ExportingConfigProperty{
//   		S3Exporting: &S3ExportingConfigProperty{
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3KeyName: jsii.String("s3KeyName"),
//   		},
//   	},
//   	MatchingRules: []interface{}{
//   		&MatchingRuleProperty{
//   			Rule: []*string{
//   				jsii.String("rule"),
//   			},
//   		},
//   	},
//   	MaxAllowedRuleLevelForMatching: jsii.Number(123),
//   	MaxAllowedRuleLevelForMerging: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html
//
type CfnDomainPropsMixin_RuleBasedMatchingProperty struct {
	// Configures information about the `AttributeTypesSelector` where the rule-based identity resolution uses to match profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-attributetypesselector
	//
	AttributeTypesSelector interface{} `field:"optional" json:"attributeTypesSelector" yaml:"attributeTypesSelector"`
	// Determines how the auto-merging process should resolve conflicts between different profiles.
	//
	// For example, if Profile A and Profile B have the same `FirstName` and `LastName` , `ConflictResolution` specifies which `EmailAddress` should be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-conflictresolution
	//
	ConflictResolution interface{} `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// The flag that enables the matching process of duplicate profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 location where Identity Resolution Jobs write result files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-exportingconfig
	//
	ExportingConfig interface{} `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// Configures how the rule-based matching process should match profiles.
	//
	// You can have up to 15 `MatchingRule` in the `MatchingRules` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-matchingrules
	//
	MatchingRules interface{} `field:"optional" json:"matchingRules" yaml:"matchingRules"`
	// Indicates the maximum allowed rule level for matching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-maxallowedrulelevelformatching
	//
	MaxAllowedRuleLevelForMatching *float64 `field:"optional" json:"maxAllowedRuleLevelForMatching" yaml:"maxAllowedRuleLevelForMatching"`
	// Indicates the maximum allowed rule level for merging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-maxallowedrulelevelformerging
	//
	MaxAllowedRuleLevelForMerging *float64 `field:"optional" json:"maxAllowedRuleLevelForMerging" yaml:"maxAllowedRuleLevelForMerging"`
	// The status of rule-based matching rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-rulebasedmatching.html#cfn-customerprofiles-domain-rulebasedmatching-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

