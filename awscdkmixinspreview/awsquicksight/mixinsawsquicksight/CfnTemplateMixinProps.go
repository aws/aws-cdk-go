package mixinsawsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTemplatePropsMixin.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html
//
type CfnTemplateMixinProps struct {
	// The ID for the AWS account that the group is in.
	//
	// You use the ID for the AWS account that contains your Amazon Quick Sight account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// A display name for the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions to be set on the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// The entity that you are using as a source when you create the template.
	//
	// In `SourceEntity` , you specify the type of object you're using as source: `SourceTemplate` for a template or `SourceAnalysis` for an analysis. Both of these require an Amazon Resource Name (ARN). For `SourceTemplate` , specify the ARN of the source template. For `SourceAnalysis` , specify the ARN of the source analysis. The `SourceTemplate` ARN can contain any AWS account and any Quick Sight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` or `SourceAnalysis` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	//
	// Either a `SourceEntity` or a `Definition` must be provided in order for the request to be valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-sourceentity
	//
	SourceEntity interface{} `field:"optional" json:"sourceEntity" yaml:"sourceEntity"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An ID for the template that you want to create.
	//
	// This template is unique per AWS Region ; in each AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-templateid
	//
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects.
	//
	// When you set this value to `LENIENT` , validation is skipped for specific errors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-validationstrategy
	//
	ValidationStrategy interface{} `field:"optional" json:"validationStrategy" yaml:"validationStrategy"`
	// A description of the current template version being created.
	//
	// This API operation creates the first version of the template. Every time `UpdateTemplate` is called, a new version is created. Each version of the template maintains a description of the version in the `VersionDescription` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

