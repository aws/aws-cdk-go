// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// the base properties for AppSync Functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   var mappingTemplate mappingTemplate
//
//   baseAppsyncFunctionProps := &baseAppsyncFunctionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	requestMappingTemplate: mappingTemplate,
//   	responseMappingTemplate: mappingTemplate,
//   }
//
// Experimental.
type BaseAppsyncFunctionProps struct {
	// the name of the AppSync Function.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// the description for this AppSync Function.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// the request mapping template for the AppSync Function.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
}

