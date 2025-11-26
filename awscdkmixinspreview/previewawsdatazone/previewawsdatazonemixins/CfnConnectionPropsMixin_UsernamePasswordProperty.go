package previewawsdatazonemixins


// The username and password of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   usernamePasswordProperty := &UsernamePasswordProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-usernamepassword.html
//
type CfnConnectionPropsMixin_UsernamePasswordProperty struct {
	// The password of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-usernamepassword.html#cfn-datazone-connection-usernamepassword-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The username of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-usernamepassword.html#cfn-datazone-connection-usernamepassword-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

