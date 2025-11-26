package previewawsquicksightmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAnalysisPropsMixin.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html
//
type CfnAnalysisMixinProps struct {
	// The ID for the analysis that you're creating.
	//
	// This ID displays in the URL of the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-analysisid
	//
	AnalysisId *string `field:"optional" json:"analysisId" yaml:"analysisId"`
	// The ID of the AWS account where you are creating an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// Errors associated with the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-errors
	//
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-folderarns
	//
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// A descriptive name for the analysis that you're creating.
	//
	// This name displays for the analysis in the Amazon Quick Sight console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameter names and override values that you want to use.
	//
	// An analysis can have any parameter type, and some parameters might accept multiple values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that describes the principals and the resource-level permissions on an analysis.
	//
	// You can use the `Permissions` structure to grant permissions by providing a list of AWS Identity and Access Management (IAM) action information for each principal listed by Amazon Resource Name (ARN).
	//
	// To specify no permissions, omit `Permissions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A list of the associated sheets with the unique identifier and name of each sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-sheets
	//
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// A source entity to use for the analysis that you're creating.
	//
	// This metadata structure contains details that describe a source template and one or more datasets.
	//
	// Either a `SourceEntity` or a `Definition` must be provided in order for the request to be valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-sourceentity
	//
	SourceEntity interface{} `field:"optional" json:"sourceEntity" yaml:"sourceEntity"`
	// Status associated with the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN for the theme to apply to the analysis that you're creating.
	//
	// To see the theme in the Amazon Quick Sight console, make sure that you have access to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-themearn
	//
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects.
	//
	// When you set this value to `LENIENT` , validation is skipped for specific errors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-analysis.html#cfn-quicksight-analysis-validationstrategy
	//
	ValidationStrategy interface{} `field:"optional" json:"validationStrategy" yaml:"validationStrategy"`
}

