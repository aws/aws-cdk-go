package previewawsqbusinessmixins


// Provides the configuration information for invoking a Lambda function in AWS Lambda to alter document metadata and content when ingesting documents into Amazon Q Business.
//
// You can configure your Lambda function using the `PreExtractionHookConfiguration` parameter if you want to apply advanced alterations on the original or raw documents.
//
// If you want to apply advanced alterations on the Amazon Q Business structured documents, you must configure your Lambda function using `PostExtractionHookConfiguration` .
//
// You can only invoke one Lambda function. However, this function can invoke other functions it requires.
//
// For more information, see [Custom document enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hookConfigurationProperty := &HookConfigurationProperty{
//   	InvocationCondition: &DocumentAttributeConditionProperty{
//   		Key: jsii.String("key"),
//   		Operator: jsii.String("operator"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	LambdaArn: jsii.String("lambdaArn"),
//   	RoleArn: jsii.String("roleArn"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-hookconfiguration.html
//
type CfnDataSourcePropsMixin_HookConfigurationProperty struct {
	// The condition used for when a Lambda function should be invoked.
	//
	// For example, you can specify a condition that if there are empty date-time values, then Amazon Q Business should invoke a function that inserts the current date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-hookconfiguration.html#cfn-qbusiness-datasource-hookconfiguration-invocationcondition
	//
	InvocationCondition interface{} `field:"optional" json:"invocationCondition" yaml:"invocationCondition"`
	// The Amazon Resource Name (ARN) of the Lambda function during ingestion.
	//
	// For more information, see [Using Lambda functions for Amazon Q Business document enrichment](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cde-lambda-operations.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-hookconfiguration.html#cfn-qbusiness-datasource-hookconfiguration-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// The Amazon Resource Name (ARN) of a role with permission to run `PreExtractionHookConfiguration` and `PostExtractionHookConfiguration` for altering document metadata and content during the document ingestion process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-hookconfiguration.html#cfn-qbusiness-datasource-hookconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Stores the original, raw documents or the structured, parsed documents before and after altering them.
	//
	// For more information, see [Data contracts for Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html#cde-lambda-operations-data-contracts) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-hookconfiguration.html#cfn-qbusiness-datasource-hookconfiguration-s3bucketname
	//
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
}

