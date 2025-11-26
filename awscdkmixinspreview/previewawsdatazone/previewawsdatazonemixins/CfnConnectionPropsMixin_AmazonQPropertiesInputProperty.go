package previewawsdatazonemixins


// Amazon Q properties of the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   amazonQPropertiesInputProperty := &AmazonQPropertiesInputProperty{
//   	AuthMode: jsii.String("authMode"),
//   	IsEnabled: jsii.Boolean(false),
//   	ProfileArn: jsii.String("profileArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-amazonqpropertiesinput.html
//
type CfnConnectionPropsMixin_AmazonQPropertiesInputProperty struct {
	// The authentication mode of the connection's AmazonQ properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-amazonqpropertiesinput.html#cfn-datazone-connection-amazonqpropertiesinput-authmode
	//
	AuthMode *string `field:"optional" json:"authMode" yaml:"authMode"`
	// Specifies whether Amazon Q is enabled for the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-amazonqpropertiesinput.html#cfn-datazone-connection-amazonqpropertiesinput-isenabled
	//
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-amazonqpropertiesinput.html#cfn-datazone-connection-amazonqpropertiesinput-profilearn
	//
	ProfileArn *string `field:"optional" json:"profileArn" yaml:"profileArn"`
}

