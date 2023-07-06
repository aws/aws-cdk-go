package awsverifiedpermissions


// A structure that contains configuration information used when creating or updating an identity source that represents a connection to an Amazon Cognito user pool used as an identity provider for Verified Permissions .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoUserPoolConfigurationProperty := &CognitoUserPoolConfigurationProperty{
//   	UserPoolArn: jsii.String("userPoolArn"),
//
//   	// the properties below are optional
//   	ClientIds: []*string{
//   		jsii.String("clientIds"),
//   	},
//   }
//
type CfnIdentitySource_CognitoUserPoolConfigurationProperty struct {
	// The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) of the Amazon Cognito user pool that contains the identities to be authorized.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// The unique application client IDs that are associated with the specified Amazon Cognito user pool.
	//
	// Example: `"ClientIds": ["&ExampleCogClientId;"]`.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
}

