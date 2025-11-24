package mixinsawsglue


// For supplying basic auth credentials when not providing a `SecretArn` value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   basicAuthenticationCredentialsProperty := &BasicAuthenticationCredentialsProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-basicauthenticationcredentials.html
//
type CfnConnectionPropsMixin_BasicAuthenticationCredentialsProperty struct {
	// The password to connect to the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-basicauthenticationcredentials.html#cfn-glue-connection-basicauthenticationcredentials-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The username to connect to the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-basicauthenticationcredentials.html#cfn-glue-connection-basicauthenticationcredentials-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

