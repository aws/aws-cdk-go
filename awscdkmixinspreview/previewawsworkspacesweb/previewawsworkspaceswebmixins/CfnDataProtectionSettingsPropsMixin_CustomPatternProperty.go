package previewawsworkspaceswebmixins


// The pattern configuration for redacting custom data types in session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customPatternProperty := &CustomPatternProperty{
//   	KeywordRegex: jsii.String("keywordRegex"),
//   	PatternDescription: jsii.String("patternDescription"),
//   	PatternName: jsii.String("patternName"),
//   	PatternRegex: jsii.String("patternRegex"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html
//
type CfnDataProtectionSettingsPropsMixin_CustomPatternProperty struct {
	// The keyword regex for the customer pattern.
	//
	// After there is a match to the pattern regex, the keyword regex is used to search within the proximity of the match. If there is a keyword match, then the match is confirmed. If no keyword regex is provided, the pattern regex match will automatically be confirmed. The format must follow JavaScript regex format. The pattern must be enclosed between slashes, and can have flags behind the second slash. For example, “/ab+c/gi”
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-keywordregex
	//
	KeywordRegex *string `field:"optional" json:"keywordRegex" yaml:"keywordRegex"`
	// The pattern description for the customer pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-patterndescription
	//
	PatternDescription *string `field:"optional" json:"patternDescription" yaml:"patternDescription"`
	// The pattern name for the custom pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-patternname
	//
	PatternName *string `field:"optional" json:"patternName" yaml:"patternName"`
	// The pattern regex for the customer pattern.
	//
	// The format must follow JavaScript regex format. The pattern must be enclosed between slashes, and can have flags behind the second slash. For example: “/ab+c/gi”.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-patternregex
	//
	PatternRegex *string `field:"optional" json:"patternRegex" yaml:"patternRegex"`
}

