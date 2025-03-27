package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnConnection_RedshiftCredentialsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftcredentials.html#cfn-datazone-connection-redshiftcredentials-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The username and password to be used for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-redshiftcredentials.html#cfn-datazone-connection-redshiftcredentials-usernamepassword
	//
	UsernamePassword interface{} `field:"optional" json:"usernamePassword" yaml:"usernamePassword"`
}

