package awsappsync


// The `OpenSearchServiceConfig` property type specifies the `AwsRegion` and `Endpoints` for an Amazon OpenSearch Service domain in your account for an AWS AppSync data source.
//
// `OpenSearchServiceConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openSearchServiceConfigProperty := &openSearchServiceConfigProperty{
//   	awsRegion: jsii.String("awsRegion"),
//   	endpoint: jsii.String("endpoint"),
//   }
//
type CfnDataSource_OpenSearchServiceConfigProperty struct {
	// The AWS Region.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

