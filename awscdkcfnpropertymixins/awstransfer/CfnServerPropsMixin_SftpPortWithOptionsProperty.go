package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sftpPortWithOptionsProperty := &SftpPortWithOptionsProperty{
//   	CommunicationMode: jsii.String("communicationMode"),
//   	SftpPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-sftpportwithoptions.html
//
type CfnServerPropsMixin_SftpPortWithOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-sftpportwithoptions.html#cfn-transfer-server-sftpportwithoptions-communicationmode
	//
	CommunicationMode *string `field:"optional" json:"communicationMode" yaml:"communicationMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-sftpportwithoptions.html#cfn-transfer-server-sftpportwithoptions-sftpport
	//
	SftpPort *float64 `field:"optional" json:"sftpPort" yaml:"sftpPort"`
}

