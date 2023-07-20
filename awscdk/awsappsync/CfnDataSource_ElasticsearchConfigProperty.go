package awsappsync


// The `ElasticsearchConfig` property type specifies the `AwsRegion` and `Endpoints` for an Amazon OpenSearch Service domain in your account for an AWS AppSync data source.
//
// ElasticsearchConfig is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service . This property is deprecated. For new data sources, use *OpenSearchServiceConfig* to specify an OpenSearch Service data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticsearchConfigProperty := &ElasticsearchConfigProperty{
//   	AwsRegion: jsii.String("awsRegion"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html
//
type CfnDataSource_ElasticsearchConfigProperty struct {
	// The AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html#cfn-appsync-datasource-elasticsearchconfig-awsregion
	//
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-elasticsearchconfig.html#cfn-appsync-datasource-elasticsearchconfig-endpoint
	//
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

