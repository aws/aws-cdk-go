package previewawsquicksightmixins


// The combination of username, private key and passphrase that are used as credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyPairCredentialsProperty := &KeyPairCredentialsProperty{
//   	KeyPairUsername: jsii.String("keyPairUsername"),
//   	PrivateKey: jsii.String("privateKey"),
//   	PrivateKeyPassphrase: jsii.String("privateKeyPassphrase"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html
//
type CfnDataSourcePropsMixin_KeyPairCredentialsProperty struct {
	// Username.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html#cfn-quicksight-datasource-keypaircredentials-keypairusername
	//
	KeyPairUsername *string `field:"optional" json:"keyPairUsername" yaml:"keyPairUsername"`
	// PrivateKey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html#cfn-quicksight-datasource-keypaircredentials-privatekey
	//
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// PrivateKeyPassphrase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html#cfn-quicksight-datasource-keypaircredentials-privatekeypassphrase
	//
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
}

