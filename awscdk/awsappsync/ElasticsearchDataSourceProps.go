package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

// Properties for the Elasticsearch Data Source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domain Domain
//   var graphQLApiRef IGraphQLApiRef
//   var role Role
//
//   elasticsearchDataSourceProps := &ElasticsearchDataSourceProps{
//   	Api: graphQLApiRef,
//   	Domain: domain,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MetricsConfig: awscdk.Aws_appsync.DataSourceMetricsConfig_ENABLED,
//   	Name: jsii.String("name"),
//   	ServiceRole: role,
//   }
//
// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
type ElasticsearchDataSourceProps struct {
	// The API to attach this data source to.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Api interfacesawsappsync.IGraphQLApiRef `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Default: - None.
	//
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable enhanced metrics of the data source Value will be ignored, if `enhancedMetricsConfig.dataSourceLevelMetricsBehavior` on AppSync GraphqlApi construct is set to `FULL_REQUEST_DATA_SOURCE_METRICS`.
	// Default: - no metrics configuration.
	//
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	MetricsConfig DataSourceMetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// The name of the data source.
	// Default: - id of data source.
	//
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Default: -  Create a new role.
	//
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The elasticsearch domain containing the endpoint for the data source.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Domain awselasticsearch.IDomain `field:"required" json:"domain" yaml:"domain"`
}

