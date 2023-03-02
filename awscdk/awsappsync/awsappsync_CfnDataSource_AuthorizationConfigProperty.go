package awsappsync


// The `AuthorizationConfig` property type specifies the authorization type and configuration for an AWS AppSync http data source.
//
// `AuthorizationConfig` is a property of the [AWS AppSync DataSource HttpConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-httpconfig.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfigProperty := &authorizationConfigProperty{
//   	authorizationType: jsii.String("authorizationType"),
//
//   	// the properties below are optional
//   	awsIamConfig: &awsIamConfigProperty{
//   		signingRegion: jsii.String("signingRegion"),
//   		signingServiceName: jsii.String("signingServiceName"),
//   	},
//   }
//
type CfnDataSource_AuthorizationConfigProperty struct {
	// The authorization type that the HTTP endpoint requires.
	//
	// - *AWS_IAM* : The authorization type is Signature Version 4 (SigV4).
	AuthorizationType *string `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// The AWS Identity and Access Management settings.
	AwsIamConfig interface{} `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

