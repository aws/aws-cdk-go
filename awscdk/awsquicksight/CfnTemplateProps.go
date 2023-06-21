package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTemplate`.
//
// Example:
//
//
type CfnTemplateProps struct {
	// The ID for the AWS account that the group is in.
	//
	// You use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// An ID for the template that you want to create.
	//
	// This template is unique per AWS Region ; in each AWS account.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// `AWS::QuickSight::Template.Definition`.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// A display name for the template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions to be set on the template.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// The entity that you are using as a source when you create the template.
	//
	// In `SourceEntity` , you specify the type of object you're using as source: `SourceTemplate` for a template or `SourceAnalysis` for an analysis. Both of these require an Amazon Resource Name (ARN). For `SourceTemplate` , specify the ARN of the source template. For `SourceAnalysis` , specify the ARN of the source analysis. The `SourceTemplate` ARN can contain any AWS account and any Amazon QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` or `SourceAnalysis` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	//
	// Either a `SourceEntity` or a `Definition` must be provided in order for the request to be valid.
	SourceEntity interface{} `field:"optional" json:"sourceEntity" yaml:"sourceEntity"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the current template version being created.
	//
	// This API operation creates the first version of the template. Every time `UpdateTemplate` is called, a new version is created. Each version of the template maintains a description of the version in the `VersionDescription` field.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

