package previewawsworkspaceswebmixins


// The set of patterns that determine the data types redacted in session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inlineRedactionPatternProperty := &InlineRedactionPatternProperty{
//   	BuiltInPatternId: jsii.String("builtInPatternId"),
//   	ConfidenceLevel: jsii.Number(123),
//   	CustomPattern: &CustomPatternProperty{
//   		KeywordRegex: jsii.String("keywordRegex"),
//   		PatternDescription: jsii.String("patternDescription"),
//   		PatternName: jsii.String("patternName"),
//   		PatternRegex: jsii.String("patternRegex"),
//   	},
//   	EnforcedUrls: []*string{
//   		jsii.String("enforcedUrls"),
//   	},
//   	ExemptUrls: []*string{
//   		jsii.String("exemptUrls"),
//   	},
//   	RedactionPlaceHolder: &RedactionPlaceHolderProperty{
//   		RedactionPlaceHolderText: jsii.String("redactionPlaceHolderText"),
//   		RedactionPlaceHolderType: jsii.String("redactionPlaceHolderType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html
//
type CfnDataProtectionSettingsPropsMixin_InlineRedactionPatternProperty struct {
	// The built-in pattern from the list of preconfigured patterns.
	//
	// Either a customPattern or builtInPatternId is required. To view the entire list of data types and their corresponding built-in pattern IDs, see [Base inline redaction](https://docs.aws.amazon.com/workspaces-web/latest/adminguide/base-inline-redaction.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-builtinpatternid
	//
	BuiltInPatternId *string `field:"optional" json:"builtInPatternId" yaml:"builtInPatternId"`
	// The confidence level for inline redaction pattern.
	//
	// This indicates the certainty of data type matches in the redaction process. Confidence level 3 means high confidence, and requires a formatted text pattern match in order for content to be redacted. Confidence level 2 means medium confidence, and redaction considers both formatted and unformatted text, and adds keyword associate to the logic. Confidence level 1 means low confidence, and redaction is enforced for both formatted pattern + unformatted pattern without keyword. This overrides the global confidence level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-confidencelevel
	//
	ConfidenceLevel *float64 `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
	// The configuration for a custom pattern.
	//
	// Either a customPattern or builtInPatternId is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-custompattern
	//
	CustomPattern interface{} `field:"optional" json:"customPattern" yaml:"customPattern"`
	// The enforced URL configuration for the inline redaction pattern.
	//
	// This will override the global enforced URL configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-enforcedurls
	//
	EnforcedUrls *[]*string `field:"optional" json:"enforcedUrls" yaml:"enforcedUrls"`
	// The exempt URL configuration for the inline redaction pattern.
	//
	// This will override the global exempt URL configuration for the inline redaction pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-exempturls
	//
	ExemptUrls *[]*string `field:"optional" json:"exemptUrls" yaml:"exemptUrls"`
	// The redaction placeholder that will replace the redacted text in session for the inline redaction pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionpattern-redactionplaceholder
	//
	RedactionPlaceHolder interface{} `field:"optional" json:"redactionPlaceHolder" yaml:"redactionPlaceHolder"`
}

