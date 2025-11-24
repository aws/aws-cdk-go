package mixinsawspcaconnectorad


// Defines the attributes of the private key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateKeyAttributesV3Property := &PrivateKeyAttributesV3Property{
//   	Algorithm: jsii.String("algorithm"),
//   	CryptoProviders: []*string{
//   		jsii.String("cryptoProviders"),
//   	},
//   	KeySpec: jsii.String("keySpec"),
//   	KeyUsageProperty: &KeyUsagePropertyProperty{
//   		PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   			Decrypt: jsii.Boolean(false),
//   			KeyAgreement: jsii.Boolean(false),
//   			Sign: jsii.Boolean(false),
//   		},
//   		PropertyType: jsii.String("propertyType"),
//   	},
//   	MinimalKeyLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv3.html
//
type CfnTemplatePropsMixin_PrivateKeyAttributesV3Property struct {
	// Defines the algorithm used to generate the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv3.html#cfn-pcaconnectorad-template-privatekeyattributesv3-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Defines the cryptographic providers used to generate the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv3.html#cfn-pcaconnectorad-template-privatekeyattributesv3-cryptoproviders
	//
	CryptoProviders *[]*string `field:"optional" json:"cryptoProviders" yaml:"cryptoProviders"`
	// Defines the purpose of the private key.
	//
	// Set it to "KEY_EXCHANGE" or "SIGNATURE" value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv3.html#cfn-pcaconnectorad-template-privatekeyattributesv3-keyspec
	//
	KeySpec *string `field:"optional" json:"keySpec" yaml:"keySpec"`
	// The key usage property defines the purpose of the private key contained in the certificate.
	//
	// You can specify specific purposes using property flags or all by using property type ALL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv3.html#cfn-pcaconnectorad-template-privatekeyattributesv3-keyusageproperty
	//
	KeyUsageProperty interface{} `field:"optional" json:"keyUsageProperty" yaml:"keyUsageProperty"`
	// Set the minimum key length of the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-privatekeyattributesv3.html#cfn-pcaconnectorad-template-privatekeyattributesv3-minimalkeylength
	//
	MinimalKeyLength *float64 `field:"optional" json:"minimalKeyLength" yaml:"minimalKeyLength"`
}

