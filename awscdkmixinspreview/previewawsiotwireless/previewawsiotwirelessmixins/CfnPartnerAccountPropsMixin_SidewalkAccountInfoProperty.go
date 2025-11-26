package previewawsiotwirelessmixins


// Information about a Sidewalk account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sidewalkAccountInfoProperty := &SidewalkAccountInfoProperty{
//   	AppServerPrivateKey: jsii.String("appServerPrivateKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-partneraccount-sidewalkaccountinfo.html
//
type CfnPartnerAccountPropsMixin_SidewalkAccountInfoProperty struct {
	// The Sidewalk application server private key.
	//
	// The application server private key is a secret key, which you should handle in a similar way as you would an application password. You can protect the application server private key by storing the value in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-partneraccount-sidewalkaccountinfo.html#cfn-iotwireless-partneraccount-sidewalkaccountinfo-appserverprivatekey
	//
	AppServerPrivateKey *string `field:"optional" json:"appServerPrivateKey" yaml:"appServerPrivateKey"`
}

