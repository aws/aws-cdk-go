package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAnalysisTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalysisTemplateProps := &CfnAnalysisTemplateProps{
//   	Format: jsii.String("format"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	Source: &AnalysisSourceProperty{
//   		Artifacts: &AnalysisTemplateArtifactsProperty{
//   			EntryPoint: &AnalysisTemplateArtifactProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			AdditionalArtifacts: []interface{}{
//   				&AnalysisTemplateArtifactProperty{
//   					Location: &S3LocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   			},
//   		},
//   		Text: jsii.String("text"),
//   	},
//
//   	// the properties below are optional
//   	AnalysisParameters: []interface{}{
//   		&AnalysisParameterProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			DefaultValue: jsii.String("defaultValue"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ErrorMessageConfiguration: &ErrorMessageConfigurationProperty{
//   		Type: jsii.String("type"),
//   	},
//   	Schema: &AnalysisSchemaProperty{
//   		ReferencedTables: []*string{
//   			jsii.String("referencedTables"),
//   		},
//   	},
//   	SourceMetadata: &AnalysisSourceMetadataProperty{
//   		Artifacts: &AnalysisTemplateArtifactMetadataProperty{
//   			EntryPointHash: &HashProperty{
//   				Sha256: jsii.String("sha256"),
//   			},
//
//   			// the properties below are optional
//   			AdditionalArtifactHashes: []interface{}{
//   				&HashProperty{
//   					Sha256: jsii.String("sha256"),
//   				},
//   			},
//   		},
//   	},
//   	SyntheticDataParameters: &SyntheticDataParametersProperty{
//   		MlSyntheticDataParameters: &MLSyntheticDataParametersProperty{
//   			ColumnClassification: &ColumnClassificationDetailsProperty{
//   				ColumnMapping: []interface{}{
//   					&SyntheticDataColumnPropertiesProperty{
//   						ColumnName: jsii.String("columnName"),
//   						ColumnType: jsii.String("columnType"),
//   						IsPredictiveValue: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			Epsilon: jsii.Number(123),
//   			MaxMembershipInferenceAttackScore: jsii.Number(123),
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
type CfnAnalysisTemplateProps struct {
	// The format of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// The identifier for a membership resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-membershipidentifier
	//
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The name of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
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
	// The entire schema object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// The source metadata for the analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-sourcemetadata
	//
	SourceMetadata interface{} `field:"optional" json:"sourceMetadata" yaml:"sourceMetadata"`
	// The parameters used to generate synthetic data for this analysis template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-syntheticdataparameters
	//
	SyntheticDataParameters interface{} `field:"optional" json:"syntheticDataParameters" yaml:"syntheticDataParameters"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-analysistemplate.html#cfn-cleanrooms-analysistemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

