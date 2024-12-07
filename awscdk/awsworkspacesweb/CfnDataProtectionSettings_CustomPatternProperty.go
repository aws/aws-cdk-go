package awsworkspacesweb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPatternProperty := &CustomPatternProperty{
//   	PatternName: jsii.String("patternName"),
//   	PatternRegex: jsii.String("patternRegex"),
//
//   	// the properties below are optional
//   	KeywordRegex: jsii.String("keywordRegex"),
//   	PatternDescription: jsii.String("patternDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html
//
type CfnDataProtectionSettings_CustomPatternProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-patternname
	//
	PatternName *string `field:"required" json:"patternName" yaml:"patternName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-patternregex
	//
	PatternRegex *string `field:"required" json:"patternRegex" yaml:"patternRegex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-keywordregex
	//
	KeywordRegex *string `field:"optional" json:"keywordRegex" yaml:"keywordRegex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-custompattern.html#cfn-workspacesweb-dataprotectionsettings-custompattern-patterndescription
	//
	PatternDescription *string `field:"optional" json:"patternDescription" yaml:"patternDescription"`
}

