package previewawsappsyncmixins


// Use the `HttpConfig` property type to specify `HttpConfig` for an AWS AppSync data source.
//
// `HttpConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpConfigProperty := &HttpConfigProperty{
//   	AuthorizationConfig: &AuthorizationConfigProperty{
//   		AuthorizationType: jsii.String("authorizationType"),
//   		AwsIamConfig: &AwsIamConfigProperty{
//   			SigningRegion: jsii.String("signingRegion"),
//   			SigningServiceName: jsii.String("signingServiceName"),
//   		},
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html
//
type CfnDataSourcePropsMixin_HttpConfigProperty struct {
	// The authorization configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html#cfn-appsync-datasource-httpconfig-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html#cfn-appsync-datasource-httpconfig-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

