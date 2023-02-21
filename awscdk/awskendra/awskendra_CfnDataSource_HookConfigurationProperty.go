package awskendra


// Provides the configuration information for invoking a Lambda function in AWS Lambda to alter document metadata and content when ingesting documents into Amazon Kendra.
//
// You can configure your Lambda function using [PreExtractionHookConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_CustomDocumentEnrichmentConfiguration.html) if you want to apply advanced alterations on the original or raw documents. If you want to apply advanced alterations on the Amazon Kendra structured documents, you must configure your Lambda function using [PostExtractionHookConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_CustomDocumentEnrichmentConfiguration.html) . You can only invoke one Lambda function. However, this function can invoke other functions it requires.
//
// For more information, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hookConfigurationProperty := &hookConfigurationProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	s3Bucket: jsii.String("s3Bucket"),
//
//   	// the properties below are optional
//   	invocationCondition: &documentAttributeConditionProperty{
//   		conditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   		operator: jsii.String("operator"),
//
//   		// the properties below are optional
//   		conditionOnValue: &documentAttributeValueProperty{
//   			dateValue: jsii.String("dateValue"),
//   			longValue: jsii.Number(123),
//   			stringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
type CfnDataSource_HookConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of a role with permission to run a Lambda function during ingestion.
	//
	// For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html) .
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Stores the original, raw documents or the structured, parsed documents before and after altering them.
	//
	// For more information, see [Data contracts for Lambda functions](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html#cde-data-contracts-lambda) .
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The condition used for when a Lambda function should be invoked.
	//
	// For example, you can specify a condition that if there are empty date-time values, then Amazon Kendra should invoke a function that inserts the current date-time.
	InvocationCondition interface{} `field:"optional" json:"invocationCondition" yaml:"invocationCondition"`
}

