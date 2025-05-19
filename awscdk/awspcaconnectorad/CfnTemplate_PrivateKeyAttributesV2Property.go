package awspcaconnectorad


// Defines the attributes of the private key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateKeyAttributesV2Property := &PrivateKeyAttributesV2Property{
//   	KeySpec: jsii.String("keySpec"),
//   	MinimalKeyLength: jsii.Number(123),
//
//   	// the properties below are optional
//   	CryptoProviders: []*string{
//   		jsii.String("cryptoProviders"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv2.html
//
type CfnTemplate_PrivateKeyAttributesV2Property struct {
	// Defines the purpose of the private key.
	//
	// Set it to "KEY_EXCHANGE" or "SIGNATURE" value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv2.html#cfn-pcaconnectorad-template-privatekeyattributesv2-keyspec
	//
	KeySpec *string `field:"required" json:"keySpec" yaml:"keySpec"`
	// Set the minimum key length of the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv2.html#cfn-pcaconnectorad-template-privatekeyattributesv2-minimalkeylength
	//
	MinimalKeyLength *float64 `field:"required" json:"minimalKeyLength" yaml:"minimalKeyLength"`
	// Defines the cryptographic providers used to generate the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv2.html#cfn-pcaconnectorad-template-privatekeyattributesv2-cryptoproviders
	//
	CryptoProviders *[]*string `field:"optional" json:"cryptoProviders" yaml:"cryptoProviders"`
}

