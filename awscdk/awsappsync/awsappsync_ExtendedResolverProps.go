package awsappsync


// Additional property for an AppSync resolver for data source reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appsyncFunction appsyncFunction
//   var baseDataSource baseDataSource
//   var duration duration
//   var mappingTemplate mappingTemplate
//
//   extendedResolverProps := &ExtendedResolverProps{
//   	FieldName: jsii.String("fieldName"),
//   	TypeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	CachingConfig: &CachingConfig{
//   		Ttl: duration,
//
//   		// the properties below are optional
//   		CachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   	},
//   	DataSource: baseDataSource,
//   	PipelineConfig: []iAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	RequestMappingTemplate: mappingTemplate,
//   	ResponseMappingTemplate: mappingTemplate,
//   }
//
// Experimental.
type ExtendedResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	// Experimental.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	// Experimental.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	// Experimental.
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The data source this resolver is using.
	// Experimental.
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
}

