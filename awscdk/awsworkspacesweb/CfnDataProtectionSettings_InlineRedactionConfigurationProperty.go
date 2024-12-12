package awsworkspacesweb


// The configuration for in-session inline redaction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineRedactionConfigurationProperty := &InlineRedactionConfigurationProperty{
//   	InlineRedactionPatterns: []interface{}{
//   		&InlineRedactionPatternProperty{
//   			RedactionPlaceHolder: &RedactionPlaceHolderProperty{
//   				RedactionPlaceHolderType: jsii.String("redactionPlaceHolderType"),
//
//   				// the properties below are optional
//   				RedactionPlaceHolderText: jsii.String("redactionPlaceHolderText"),
//   			},
//
//   			// the properties below are optional
//   			BuiltInPatternId: jsii.String("builtInPatternId"),
//   			ConfidenceLevel: jsii.Number(123),
//   			CustomPattern: &CustomPatternProperty{
//   				PatternName: jsii.String("patternName"),
//   				PatternRegex: jsii.String("patternRegex"),
//
//   				// the properties below are optional
//   				KeywordRegex: jsii.String("keywordRegex"),
//   				PatternDescription: jsii.String("patternDescription"),
//   			},
//   			EnforcedUrls: []*string{
//   				jsii.String("enforcedUrls"),
//   			},
//   			ExemptUrls: []*string{
//   				jsii.String("exemptUrls"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	GlobalConfidenceLevel: jsii.Number(123),
//   	GlobalEnforcedUrls: []*string{
//   		jsii.String("globalEnforcedUrls"),
//   	},
//   	GlobalExemptUrls: []*string{
//   		jsii.String("globalExemptUrls"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html
//
type CfnDataProtectionSettings_InlineRedactionConfigurationProperty struct {
	// The inline redaction patterns to be enabled for the inline redaction configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-inlineredactionpatterns
	//
	InlineRedactionPatterns interface{} `field:"required" json:"inlineRedactionPatterns" yaml:"inlineRedactionPatterns"`
	// The global confidence level for the inline redaction configuration.
	//
	// This indicates the certainty of data type matches in the redaction process. Confidence level 3 means high confidence, and requires a formatted text pattern match in order for content to be redacted. Confidence level 2 means medium confidence, and redaction considers both formatted and unformatted text, and adds keyword associate to the logic. Confidence level 1 means low confidence, and redaction is enforced for both formatted pattern + unformatted pattern without keyword. This is applied to patterns that do not have a pattern-level confidence level. Defaults to confidence level 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-globalconfidencelevel
	//
	GlobalConfidenceLevel *float64 `field:"optional" json:"globalConfidenceLevel" yaml:"globalConfidenceLevel"`
	// The global enforced URL configuration for the inline redaction configuration.
	//
	// This is applied to patterns that do not have a pattern-level enforced URL list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-globalenforcedurls
	//
	GlobalEnforcedUrls *[]*string `field:"optional" json:"globalEnforcedUrls" yaml:"globalEnforcedUrls"`
	// The global exempt URL configuration for the inline redaction configuration.
	//
	// This is applied to patterns that do not have a pattern-level exempt URL list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-globalexempturls
	//
	GlobalExemptUrls *[]*string `field:"optional" json:"globalExemptUrls" yaml:"globalExemptUrls"`
}

