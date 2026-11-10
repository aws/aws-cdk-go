package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sftpPortWithOptionsProperty := &SftpPortWithOptionsProperty{
//   	SftpPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	CommunicationMode: jsii.String("communicationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-sftpportwithoptions.html
//
type CfnServer_SftpPortWithOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-sftpportwithoptions.html#cfn-transfer-server-sftpportwithoptions-sftpport
	//
	SftpPort *float64 `field:"required" json:"sftpPort" yaml:"sftpPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-sftpportwithoptions.html#cfn-transfer-server-sftpportwithoptions-communicationmode
	//
	CommunicationMode *string `field:"optional" json:"communicationMode" yaml:"communicationMode"`
}

