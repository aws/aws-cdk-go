package mixinsawsappsync


// The `AuthorizationConfig` property type specifies the authorization type and configuration for an AWS AppSync http data source.
//
// `AuthorizationConfig` is a property of the [AWS AppSync DataSource HttpConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authorizationConfigProperty := &AuthorizationConfigProperty{
//   	AuthorizationType: jsii.String("authorizationType"),
//   	AwsIamConfig: &AwsIamConfigProperty{
//   		SigningRegion: jsii.String("signingRegion"),
//   		SigningServiceName: jsii.String("signingServiceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html
//
type CfnDataSourcePropsMixin_AuthorizationConfigProperty struct {
	// The authorization type that the HTTP endpoint requires.
	//
	// - *AWS_IAM* : The authorization type is Signature Version 4 (SigV4).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html#cfn-appsync-datasource-authorizationconfig-authorizationtype
	//
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The AWS Identity and Access Management settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html#cfn-appsync-datasource-authorizationconfig-awsiamconfig
	//
	AwsIamConfig interface{} `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

