package awsworkspacesweb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineRedactionPatternProperty := &InlineRedactionPatternProperty{
//   	RedactionPlaceHolder: &RedactionPlaceHolderProperty{
//   		RedactionPlaceHolderType: jsii.String("redactionPlaceHolderType"),
//
//   		// the properties below are optional
//   		RedactionPlaceHolderText: jsii.String("redactionPlaceHolderText"),
//   	},
//
//   	// the properties below are optional
//   	BuiltInPatternId: jsii.String("builtInPatternId"),
//   	ConfidenceLevel: jsii.Number(123),
//   	CustomPattern: &CustomPatternProperty{
//   		PatternName: jsii.String("patternName"),
//   		PatternRegex: jsii.String("patternRegex"),
//
//   		// the properties below are optional
//   		KeywordRegex: jsii.String("keywordRegex"),
//   		PatternDescription: jsii.String("patternDescription"),
//   	},
//   	EnforcedUrls: []*string{
//   		jsii.String("enforcedUrls"),
//   	},
//   	ExemptUrls: []*string{
//   		jsii.String("exemptUrls"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html
//
type CfnDataProtectionSettings_InlineRedactionPatternProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-redactionplaceholder
	//
	RedactionPlaceHolder interface{} `field:"required" json:"redactionPlaceHolder" yaml:"redactionPlaceHolder"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-builtinpatternid
	//
	BuiltInPatternId *string `field:"optional" json:"builtInPatternId" yaml:"builtInPatternId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-confidencelevel
	//
	ConfidenceLevel *float64 `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-custompattern
	//
	CustomPattern interface{} `field:"optional" json:"customPattern" yaml:"customPattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-enforcedurls
	//
	EnforcedUrls *[]*string `field:"optional" json:"enforcedUrls" yaml:"enforcedUrls"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-exempturls
	//
	ExemptUrls *[]*string `field:"optional" json:"exemptUrls" yaml:"exemptUrls"`
}

