package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplateProps := &cfnTemplateProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	sourceEntity: &templateSourceEntityProperty{
//   		sourceAnalysis: &templateSourceAnalysisProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   		sourceTemplate: &templateSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   	templateId: jsii.String("templateId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versionDescription: jsii.String("versionDescription"),
//   }
//
type CfnTemplateProps struct {
	// The ID for the AWS account that the group is in.
	//
	// You use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The entity that you are using as a source when you create the template.
	//
	// In `SourceEntity` , you specify the type of object you're using as source: `SourceTemplate` for a template or `SourceAnalysis` for an analysis. Both of these require an Amazon Resource Name (ARN). For `SourceTemplate` , specify the ARN of the source template. For `SourceAnalysis` , specify the ARN of the source analysis. The `SourceTemplate` ARN can contain any AWS account and any Amazon QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` or `SourceAnalysis` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	SourceEntity interface{} `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// An ID for the template that you want to create.
	//
	// This template is unique per AWS Region ; in each AWS account.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// A display name for the template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions to be set on the template.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the current template version being created.
	//
	// This API operation creates the first version of the template. Every time `UpdateTemplate` is called, a new version is created. Each version of the template maintains a description of the version in the `VersionDescription` field.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

