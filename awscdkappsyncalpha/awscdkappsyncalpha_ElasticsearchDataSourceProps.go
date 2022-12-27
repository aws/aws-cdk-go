// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the Elasticsearch Data Source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domain domain
//   var graphqlApi graphqlApi
//   var role role
//
//   elasticsearchDataSourceProps := &elasticsearchDataSourceProps{
//   	api: graphqlApi,
//   	domain: domain,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	serviceRole: role,
//   }
//
// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
type ElasticsearchDataSourceProps struct {
	// The API to attach this data source to.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The elasticsearch domain containing the endpoint for the data source.
	// Deprecated: - use `OpenSearchDataSourceProps` with `OpenSearchDataSource`.
	Domain awselasticsearch.IDomain `field:"required" json:"domain" yaml:"domain"`
}

