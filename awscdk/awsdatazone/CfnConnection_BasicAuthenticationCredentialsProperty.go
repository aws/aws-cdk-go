package awsdatazone


// Basic Authentication Credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthenticationCredentialsProperty := &BasicAuthenticationCredentialsProperty{
//   	Password: jsii.String("password"),
//   	UserName: jsii.String("userName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-basicauthenticationcredentials.html
//
type CfnConnection_BasicAuthenticationCredentialsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-basicauthenticationcredentials.html#cfn-datazone-connection-basicauthenticationcredentials-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-basicauthenticationcredentials.html#cfn-datazone-connection-basicauthenticationcredentials-username
	//
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

