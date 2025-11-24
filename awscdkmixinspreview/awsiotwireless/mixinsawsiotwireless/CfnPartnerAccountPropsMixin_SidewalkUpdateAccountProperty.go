package mixinsawsiotwireless


// Sidewalk update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sidewalkUpdateAccountProperty := &SidewalkUpdateAccountProperty{
//   	AppServerPrivateKey: jsii.String("appServerPrivateKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-partneraccount-sidewalkupdateaccount.html
//
type CfnPartnerAccountPropsMixin_SidewalkUpdateAccountProperty struct {
	// The new Sidewalk application server private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-partneraccount-sidewalkupdateaccount.html#cfn-iotwireless-partneraccount-sidewalkupdateaccount-appserverprivatekey
	//
	AppServerPrivateKey *string `field:"optional" json:"appServerPrivateKey" yaml:"appServerPrivateKey"`
}

