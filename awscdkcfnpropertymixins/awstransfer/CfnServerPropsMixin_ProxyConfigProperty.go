package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   proxyConfigProperty := &ProxyConfigProperty{
//   	SftpMode: jsii.String("sftpMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-proxyconfig.html
//
type CfnServerPropsMixin_ProxyConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-proxyconfig.html#cfn-transfer-server-proxyconfig-sftpmode
	//
	SftpMode *string `field:"optional" json:"sftpMode" yaml:"sftpMode"`
}

