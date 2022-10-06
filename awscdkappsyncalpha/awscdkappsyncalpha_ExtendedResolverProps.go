// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Additional property for an AppSync resolver for data source reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var appsyncFunction appsyncFunction
//   var baseDataSource baseDataSource
//   var mappingTemplate mappingTemplate
//
//   extendedResolverProps := &extendedResolverProps{
//   	fieldName: jsii.String("fieldName"),
//   	typeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	cachingConfig: &cachingConfig{
//   		ttl: cdk.duration.minutes(jsii.Number(30)),
//
//   		// the properties below are optional
//   		cachingKeys: []*string{
//   			jsii.String("cachingKeys"),
//   		},
//   	},
//   	dataSource: baseDataSource,
//   	maxBatchSize: jsii.Number(123),
//   	pipelineConfig: []iAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	requestMappingTemplate: mappingTemplate,
//   	responseMappingTemplate: mappingTemplate,
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
	// The maximum number of elements per batch, when using batch invoke.
	// Experimental.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
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

