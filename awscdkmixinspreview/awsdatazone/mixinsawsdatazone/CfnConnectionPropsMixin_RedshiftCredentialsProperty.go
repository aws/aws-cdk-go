package mixinsawsdatazone


// Amazon Redshift credentials of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftCredentialsProperty := &RedshiftCredentialsProperty{
//   	SecretArn: jsii.String("secretArn"),
//   	UsernamePassword: &UsernamePasswordProperty{
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftcredentials.html
//
type CfnConnectionPropsMixin_RedshiftCredentialsProperty struct {
	// The secret ARN of the Amazon Redshift credentials of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftcredentials.html#cfn-datazone-connection-redshiftcredentials-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The username and password of the Amazon Redshift credentials of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftcredentials.html#cfn-datazone-connection-redshiftcredentials-usernamepassword
	//
	UsernamePassword interface{} `field:"optional" json:"usernamePassword" yaml:"usernamePassword"`
}

