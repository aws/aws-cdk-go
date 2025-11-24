package mixinsawscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAnalysisTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnalysisTemplateMixinProps := &CfnAnalysisTemplateMixinProps{
//   	AnalysisParameters: []interface{}{
//   		&AnalysisParameterProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ErrorMessageConfiguration: &ErrorMessageConfigurationProperty{
//   		Type: jsii.String("type"),
//   	},
//   	Format: jsii.String("format"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	Schema: &AnalysisSchemaProperty{
//   		ReferencedTables: []*string{
//   			jsii.String("referencedTables"),
//   		},
//   	},
//   	Source: &AnalysisSourceProperty{
//   		Artifacts: &AnalysisTemplateArtifactsProperty{
//   			AdditionalArtifacts: []interface{}{
//   				&AnalysisTemplateArtifactProperty{
//   					Location: &S3LocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   			},
//   			EntryPoint: &AnalysisTemplateArtifactProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Text: jsii.String("text"),
//   	},
//   	SourceMetadata: &AnalysisSourceMetadataProperty{
//   		Artifacts: &AnalysisTemplateArtifactMetadataProperty{
//   			AdditionalArtifactHashes: []interface{}{
//   				&HashProperty{
//   					Sha256: jsii.String("sha256"),
//   				},
//   			},
//   			EntryPointHash: &HashProperty{
//   				Sha256: jsii.String("sha256"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html
//
type CfnAnalysisTemplateMixinProps struct {
	// The parameters of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-analysisparameters
	//
	AnalysisParameters interface{} `field:"optional" json:"analysisParameters" yaml:"analysisParameters"`
	// The description of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration that specifies the level of detail in error messages returned by analyses using this template.
	//
	// When set to `DETAILED` , error messages include more information to help troubleshoot issues with PySpark jobs. Detailed error messages may expose underlying data, including sensitive information. Recommended for faster troubleshooting in development and testing environments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-errormessageconfiguration
	//
	ErrorMessageConfiguration interface{} `field:"optional" json:"errorMessageConfiguration" yaml:"errorMessageConfiguration"`
	// The format of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The identifier for a membership resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-membershipidentifier
	//
	MembershipIdentifier *string `field:"optional" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The name of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The entire schema object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// The source of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The source metadata for the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-sourcemetadata
	//
	SourceMetadata interface{} `field:"optional" json:"sourceMetadata" yaml:"sourceMetadata"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

