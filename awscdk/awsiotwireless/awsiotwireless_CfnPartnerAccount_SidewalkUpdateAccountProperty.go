package awsiotwireless


// Sidewalk update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sidewalkUpdateAccountProperty := &sidewalkUpdateAccountProperty{
//   	appServerPrivateKey: jsii.String("appServerPrivateKey"),
//   }
//
type CfnPartnerAccount_SidewalkUpdateAccountProperty struct {
	// The new Sidewalk application server private key.
	AppServerPrivateKey *string `field:"optional" json:"appServerPrivateKey" yaml:"appServerPrivateKey"`
}

