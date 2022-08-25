package awsappsync


// Properties for defining a `CfnFunctionConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionConfigurationProps := &cfnFunctionConfigurationProps{
//   	apiId: jsii.String("apiId"),
//   	dataSourceName: jsii.String("dataSourceName"),
//   	functionVersion: jsii.String("functionVersion"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	maxBatchSize: jsii.Number(123),
//   	requestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	requestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	responseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	responseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	syncConfig: &syncConfigProperty{
//   		conflictDetection: jsii.String("conflictDetection"),
//
//   		// the properties below are optional
//   		conflictHandler: jsii.String("conflictHandler"),
//   		lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   			lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   }
//
type CfnFunctionConfigurationProps struct {
	// The AWS AppSync GraphQL API that you want to attach using this function.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of data source this function will attach.
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// The version of the request mapping template.
	//
	// Currently, only the 2018-05-29 version of the template is supported.
	FunctionVersion *string `field:"required" json:"functionVersion" yaml:"functionVersion"`
	// The name of the function.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The `Function` description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a `BatchInvoke` operation.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The `Function` request mapping template.
	//
	// Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// Describes a Sync configuration for a resolver.
	//
	// Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
	RequestMappingTemplateS3Location *string `field:"optional" json:"requestMappingTemplateS3Location" yaml:"requestMappingTemplateS3Location"`
	// The `Function` response mapping template.
	ResponseMappingTemplate *string `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The location of a response mapping template in an Amazon S3 bucket.
	//
	// Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
	ResponseMappingTemplateS3Location *string `field:"optional" json:"responseMappingTemplateS3Location" yaml:"responseMappingTemplateS3Location"`
	// Describes a Sync configuration for a resolver.
	//
	// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
	SyncConfig interface{} `field:"optional" json:"syncConfig" yaml:"syncConfig"`
}

