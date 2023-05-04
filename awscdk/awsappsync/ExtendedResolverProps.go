package awsappsync


// Additional property for an AppSync resolver for data source reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appsyncFunction appsyncFunction
//   var baseDataSource baseDataSource
//   var code code
//   var functionRuntime functionRuntime
//   var mappingTemplate mappingTemplate
//
//   extendedResolverProps := &ExtendedResolverProps{
//   	FieldName: jsii.String("fieldName"),
//   	TypeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	CachingConfig: &CachingConfig{
//   		Ttl: cdk.Duration_Minutes(jsii.Number(30)),
//
//   		// the properties below are optional
//   		CachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   	},
//   	Code: code,
//   	DataSource: baseDataSource,
//   	MaxBatchSize: jsii.Number(123),
//   	PipelineConfig: []iAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	RequestMappingTemplate: mappingTemplate,
//   	ResponseMappingTemplate: mappingTemplate,
//   	Runtime: functionRuntime,
//   }
//
type ExtendedResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The function code.
	Code Code `field:"optional" json:"code" yaml:"code"`
	// The maximum number of elements per batch, when using batch invoke.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// configuration of the pipeline resolver.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// The data source this resolver is using.
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
}

