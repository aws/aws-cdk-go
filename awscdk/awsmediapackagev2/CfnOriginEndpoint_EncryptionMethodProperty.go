package awsmediapackagev2


// The encryption method associated with the origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionMethodProperty := &EncryptionMethodProperty{
//   	CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   	TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptionmethod.html
//
type CfnOriginEndpoint_EncryptionMethodProperty struct {
	// The encryption method to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptionmethod.html#cfn-mediapackagev2-originendpoint-encryptionmethod-cmafencryptionmethod
	//
	CmafEncryptionMethod *string `field:"optional" json:"cmafEncryptionMethod" yaml:"cmafEncryptionMethod"`
	// The encryption method to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryptionmethod.html#cfn-mediapackagev2-originendpoint-encryptionmethod-tsencryptionmethod
	//
	TsEncryptionMethod *string `field:"optional" json:"tsEncryptionMethod" yaml:"tsEncryptionMethod"`
}

