package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

// Base properties for an AppSync datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphQLApiRef IGraphQLApiRef
//
//   baseDataSourceProps := &BaseDataSourceProps{
//   	Api: graphQLApiRef,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MetricsConfig: awscdk.Aws_appsync.DataSourceMetricsConfig_ENABLED,
//   	Name: jsii.String("name"),
//   }
//
type BaseDataSourceProps struct {
	// The API to attach this data source to.
	Api interfacesawsappsync.IGraphQLApiRef `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable enhanced metrics of the data source Value will be ignored, if `enhancedMetricsConfig.dataSourceLevelMetricsBehavior` on AppSync GraphqlApi construct is set to `FULL_REQUEST_DATA_SOURCE_METRICS`.
	// Default: - no metrics configuration.
	//
	MetricsConfig DataSourceMetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// The name of the data source.
	// Default: - id of data source.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

