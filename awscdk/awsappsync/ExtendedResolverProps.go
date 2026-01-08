package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

// Additional property for an AppSync resolver for data source reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var baseDataSource BaseDataSource
//   var code Code
//   var functionConfigurationRef IFunctionConfigurationRef
//   var functionRuntime FunctionRuntime
//   var mappingTemplate MappingTemplate
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
//   	PipelineConfig: []IFunctionConfigurationRef{
//   		functionConfigurationRef,
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
	// Default: - No caching configuration.
	//
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The function code.
	// Default: - no code is used.
	//
	Code Code `field:"optional" json:"code" yaml:"code"`
	// The maximum number of elements per batch, when using batch invoke.
	// Default: - No max batch size.
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// configuration of the pipeline resolver.
	// Default: - no pipeline resolver configuration
	// An empty array | undefined sets resolver to be of kind, unit.
	//
	PipelineConfig *[]interfacesawsappsync.IFunctionConfigurationRef `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Default: - No mapping template.
	//
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Default: - No mapping template.
	//
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	// Default: - no function runtime, VTL mapping templates used.
	//
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// The data source this resolver is using.
	// Default: - No datasource.
	//
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
}

