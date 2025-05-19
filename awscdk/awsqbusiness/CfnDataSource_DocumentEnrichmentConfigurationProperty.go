package awsqbusiness


// Provides the configuration information for altering document metadata and content during the document ingestion process.
//
// For more information, see [Custom document enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentEnrichmentConfigurationProperty := &DocumentEnrichmentConfigurationProperty{
//   	InlineConfigurations: []interface{}{
//   		&InlineDocumentEnrichmentConfigurationProperty{
//   			Condition: &DocumentAttributeConditionProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			DocumentContentOperator: jsii.String("documentContentOperator"),
//   			Target: &DocumentAttributeTargetProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				AttributeValueOperator: jsii.String("attributeValueOperator"),
//   				Value: &DocumentAttributeValueProperty{
//   					DateValue: jsii.String("dateValue"),
//   					LongValue: jsii.Number(123),
//   					StringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   	},
//   	PostExtractionHookConfiguration: &HookConfigurationProperty{
//   		InvocationCondition: &DocumentAttributeConditionProperty{
//   			Key: jsii.String("key"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		LambdaArn: jsii.String("lambdaArn"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	PreExtractionHookConfiguration: &HookConfigurationProperty{
//   		InvocationCondition: &DocumentAttributeConditionProperty{
//   			Key: jsii.String("key"),
//   			Operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			Value: &DocumentAttributeValueProperty{
//   				DateValue: jsii.String("dateValue"),
//   				LongValue: jsii.Number(123),
//   				StringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		LambdaArn: jsii.String("lambdaArn"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html
//
type CfnDataSource_DocumentEnrichmentConfigurationProperty struct {
	// Configuration information to alter document attributes or metadata fields and content when ingesting documents into Amazon Q Business.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html#cfn-qbusiness-datasource-documentenrichmentconfiguration-inlineconfigurations
	//
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Using Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html#cfn-qbusiness-datasource-documentenrichmentconfiguration-postextractionhookconfiguration
	//
	PostExtractionHookConfiguration interface{} `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and text.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Using Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html#cfn-qbusiness-datasource-documentenrichmentconfiguration-preextractionhookconfiguration
	//
	PreExtractionHookConfiguration interface{} `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
}

