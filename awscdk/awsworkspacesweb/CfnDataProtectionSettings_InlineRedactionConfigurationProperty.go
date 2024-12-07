package awsworkspacesweb


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-inlineredactionpatterns
	//
	InlineRedactionPatterns interface{} `field:"required" json:"inlineRedactionPatterns" yaml:"inlineRedactionPatterns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-globalconfidencelevel
	//
	GlobalConfidenceLevel *float64 `field:"optional" json:"globalConfidenceLevel" yaml:"globalConfidenceLevel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-globalenforcedurls
	//
	GlobalEnforcedUrls *[]*string `field:"optional" json:"globalEnforcedUrls" yaml:"globalEnforcedUrls"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html#cfn-workspacesweb-dataprotectionsettings-inlineredactionconfiguration-globalexempturls
	//
	GlobalExemptUrls *[]*string `field:"optional" json:"globalExemptUrls" yaml:"globalExemptUrls"`
}

