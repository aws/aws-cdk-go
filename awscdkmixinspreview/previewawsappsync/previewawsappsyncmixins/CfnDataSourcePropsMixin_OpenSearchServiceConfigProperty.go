package previewawsappsyncmixins


// The `OpenSearchServiceConfig` property type specifies the `AwsRegion` and `Endpoints` for an Amazon OpenSearch Service domain in your account for an AWS AppSync data source.
//
// `OpenSearchServiceConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openSearchServiceConfigProperty := &OpenSearchServiceConfigProperty{
//   	AwsRegion: jsii.String("awsRegion"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-opensearchserviceconfig.html
//
type CfnDataSourcePropsMixin_OpenSearchServiceConfigProperty struct {
	// The AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-opensearchserviceconfig.html#cfn-appsync-datasource-opensearchserviceconfig-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-opensearchserviceconfig.html#cfn-appsync-datasource-opensearchserviceconfig-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

