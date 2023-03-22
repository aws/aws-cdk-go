package awsappsync


// Use the `HttpConfig` property type to specify `HttpConfig` for an AWS AppSync data source.
//
// `HttpConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpConfigProperty := &httpConfigProperty{
//   	endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfigProperty{
//   		authorizationType: jsii.String("authorizationType"),
//
//   		// the properties below are optional
//   		awsIamConfig: &awsIamConfigProperty{
//   			signingRegion: jsii.String("signingRegion"),
//   			signingServiceName: jsii.String("signingServiceName"),
//   		},
//   	},
//   }
//
type CfnDataSource_HttpConfigProperty struct {
	// The endpoint.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The authorization configuration.
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

