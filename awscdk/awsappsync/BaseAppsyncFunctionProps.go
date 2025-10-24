package awsappsync


// the base properties for AppSync Functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var code Code
//   var functionRuntime FunctionRuntime
//   var mappingTemplate MappingTemplate
//
//   baseAppsyncFunctionProps := &BaseAppsyncFunctionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Code: code,
//   	Description: jsii.String("description"),
//   	MaxBatchSize: jsii.Number(123),
//   	RequestMappingTemplate: mappingTemplate,
//   	ResponseMappingTemplate: mappingTemplate,
//   	Runtime: functionRuntime,
//   }
//
type BaseAppsyncFunctionProps struct {
	// the name of the AppSync Function.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The function code.
	// Default: - no code is used.
	//
	Code Code `field:"optional" json:"code" yaml:"code"`
	// the description for this AppSync Function.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
	//
	// Can only be set when using LambdaDataSource.
	// Default: - No max batch size.
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// the request mapping template for the AppSync Function.
	// Default: - no request mapping template.
	//
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Default: - no response mapping template.
	//
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	// Default: - no function runtime, VTL mapping templates used.
	//
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

