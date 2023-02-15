package awsappsync


// the base properties for AppSync Functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//   var functionRuntime functionRuntime
//   var mappingTemplate mappingTemplate
//
//   baseAppsyncFunctionProps := &baseAppsyncFunctionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	code: code,
//   	description: jsii.String("description"),
//   	requestMappingTemplate: mappingTemplate,
//   	responseMappingTemplate: mappingTemplate,
//   	runtime: functionRuntime,
//   }
//
type BaseAppsyncFunctionProps struct {
	// the name of the AppSync Function.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The function code.
	Code Code `field:"optional" json:"code" yaml:"code"`
	// the description for this AppSync Function.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// the request mapping template for the AppSync Function.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
}

