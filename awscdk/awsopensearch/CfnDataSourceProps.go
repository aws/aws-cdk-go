package awsopensearch


// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
//   	DataSourceType: &DataSourceTypeProperty{
//   		S3GlueDataCatalog: &S3GlueDataCatalogProperty{
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	DomainName: jsii.String("domainName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearch-datasource.html
//
type CfnDataSourceProps struct {
	// The type of data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearch-datasource.html#cfn-opensearch-datasource-datasourcetype
	//
	DataSourceType interface{} `field:"required" json:"dataSourceType" yaml:"dataSourceType"`
	// The name of the OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearch-datasource.html#cfn-opensearch-datasource-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearch-datasource.html#cfn-opensearch-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearch-datasource.html#cfn-opensearch-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

