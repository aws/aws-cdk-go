package awspcaconnectorad


// Defines the attributes of the private key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateKeyAttributesV4Property := &PrivateKeyAttributesV4Property{
//   	KeySpec: jsii.String("keySpec"),
//   	MinimalKeyLength: jsii.Number(123),
//
//   	// the properties below are optional
//   	Algorithm: jsii.String("algorithm"),
//   	CryptoProviders: []*string{
//   		jsii.String("cryptoProviders"),
//   	},
//   	KeyUsageProperty: &KeyUsagePropertyProperty{
//   		PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   			Decrypt: jsii.Boolean(false),
//   			KeyAgreement: jsii.Boolean(false),
//   			Sign: jsii.Boolean(false),
//   		},
//   		PropertyType: jsii.String("propertyType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv4.html
//
type CfnTemplate_PrivateKeyAttributesV4Property struct {
	// Defines the purpose of the private key.
	//
	// Set it to "KEY_EXCHANGE" or "SIGNATURE" value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv4.html#cfn-pcaconnectorad-template-privatekeyattributesv4-keyspec
	//
	KeySpec *string `field:"required" json:"keySpec" yaml:"keySpec"`
	// Set the minimum key length of the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv4.html#cfn-pcaconnectorad-template-privatekeyattributesv4-minimalkeylength
	//
	MinimalKeyLength *float64 `field:"required" json:"minimalKeyLength" yaml:"minimalKeyLength"`
	// Defines the algorithm used to generate the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv4.html#cfn-pcaconnectorad-template-privatekeyattributesv4-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Defines the cryptographic providers used to generate the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv4.html#cfn-pcaconnectorad-template-privatekeyattributesv4-cryptoproviders
	//
	CryptoProviders *[]*string `field:"optional" json:"cryptoProviders" yaml:"cryptoProviders"`
	// The key usage property defines the purpose of the private key contained in the certificate.
	//
	// You can specify specific purposes using property flags or all by using property type ALL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv4.html#cfn-pcaconnectorad-template-privatekeyattributesv4-keyusageproperty
	//
	KeyUsageProperty interface{} `field:"optional" json:"keyUsageProperty" yaml:"keyUsageProperty"`
}

