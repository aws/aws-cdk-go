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
//   customDocumentEnrichmentConfigurationProperty := &customDocumentEnrichmentConfigurationProperty{
//   	inlineConfigurations: []interface{}{
//   		&inlineCustomDocumentEnrichmentConfigurationProperty{
//   			condition: &documentAttributeConditionProperty{
//   				conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   				operator: jsii.String("operator"),
//
//   				// the properties below are optional
//   				conditionOnValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			documentContentDeletion: jsii.Boolean(false),
//   			target: &documentAttributeTargetProperty{
//   				targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   				// the properties below are optional
//   				targetDocumentAttributeValue: &documentAttributeValueProperty{
//   					dateValue: jsii.String("dateValue"),
//   					longValue: jsii.Number(123),
//   					stringListValue: []*string{
//   						jsii.String("stringListValue"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   				targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	postExtractionHookConfiguration: &hookConfigurationProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		invocationCondition: &documentAttributeConditionProperty{
//   			conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   			operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			conditionOnValue: &documentAttributeValueProperty{
//   				dateValue: jsii.String("dateValue"),
//   				longValue: jsii.Number(123),
//   				stringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	preExtractionHookConfiguration: &hookConfigurationProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		invocationCondition: &documentAttributeConditionProperty{
//   			conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   			operator: jsii.String("operator"),
//
//   			// the properties below are optional
//   			conditionOnValue: &documentAttributeValueProperty{
//   				dateValue: jsii.String("dateValue"),
//   				longValue: jsii.Number(123),
//   				stringListValue: []*string{
//   					jsii.String("stringListValue"),
//   				},
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDataSource_CustomDocumentEnrichmentConfigurationProperty struct {
	// Configuration information to alter document attributes or metadata fields and content when ingesting documents into Amazon Kendra.
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation) .
	PostExtractionHookConfiguration interface{} `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and text.
	//
	// You can use a Lambda function to apply advanced logic for creating, modifying, or deleting document metadata and content. For more information, see [Advanced data manipulation](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#advanced-data-manipulation) .
	PreExtractionHookConfiguration interface{} `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
	// The Amazon Resource Name (ARN) of a role with permission to run `PreExtractionHookConfiguration` and `PostExtractionHookConfiguration` for altering document metadata and content during the document ingestion process.
	//
	// For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html) .
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

