package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAnalysis`.
//
// Example:
//
//
type CfnAnalysisProps struct {
	// The ID for the analysis that you're creating.
	//
	// This ID displays in the URL of the analysis.
	AnalysisId *string `field:"required" json:"analysisId" yaml:"analysisId"`
	// The ID of the AWS account where you are creating an analysis.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// A descriptive name for the analysis that you're creating.
	//
	// This name displays for the analysis in the Amazon QuickSight console.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::QuickSight::Analysis.Definition`.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The parameter names and override values that you want to use.
	//
	// An analysis can have any parameter type, and some parameters might accept multiple values.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that describes the principals and the resource-level permissions on an analysis.
	//
	// You can use the `Permissions` structure to grant permissions by providing a list of AWS Identity and Access Management (IAM) action information for each principal listed by Amazon Resource Name (ARN).
	//
	// To specify no permissions, omit `Permissions` .
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A source entity to use for the analysis that you're creating.
	//
	// This metadata structure contains details that describe a source template and one or more datasets.
	//
	// Either a `SourceEntity` or a `Definition` must be provided in order for the request to be valid.
	SourceEntity interface{} `field:"optional" json:"sourceEntity" yaml:"sourceEntity"`
	// Status associated with the analysis.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the analysis.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN for the theme to apply to the analysis that you're creating.
	//
	// To see the theme in the Amazon QuickSight console, make sure that you have access to it.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
}

