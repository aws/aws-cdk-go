package awskendra


// Provides the configuration information for altering document metadata and content during the document ingestion process.
//
// For more information, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDocumentEnrichmentConfigurationProperty := &CustomDocumentEnrichmentConfigurationProperty{
//   	InlineConfigurations: []interface{}{
//   		&InlineCustomDocumentEnrichmentConfigurationProperty{
//   			Condition: &DocumentAttributeConditionProperty{
//   				ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				Operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				ConditionOnValue: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			DocumentContentDeletion: jsii.Boolean(false),
//   			Target: &DocumentAttributeTargetProperty{
//   				TargetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   				// the properties below are optional
//   				TargetDocumentAttributeValue: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   				TargetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	PostExtractionHookConfiguration: &HookConfigurationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		InvocationCondition: &DocumentAttributeConditionProperty{
//   			ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			ConditionOnValue: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	PreExtractionHookConfiguration: &HookConfigurationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		InvocationCondition: &DocumentAttributeConditionProperty{
//   			ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			ConditionOnValue: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html
//
type CfnDataSource_CustomDocumentEnrichmentConfigurationProperty struct {
	// Configuration information to alter document attributes or metadata fields and content when ingesting documents into Amazon Kendra.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-inlineconfigurations
	//
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-postextractionhookconfiguration
	//
	PostExtractionHookConfiguration interface{} `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and text.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-preextractionhookconfiguration
	//
	PreExtractionHookConfiguration interface{} `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
	// The Amazon Resource Name (ARN) of an IAM role with permission to run `PreExtractionHookConfiguration` and `PostExtractionHookConfiguration` for altering document metadata and content during the document ingestion process.
	//
	// For more information, see [an IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

