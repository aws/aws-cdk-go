package awsquicksight


// The combination of username, private key and passphrase that are used as credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyPairCredentialsProperty := &KeyPairCredentialsProperty{
//   	KeyPairUsername: jsii.String("keyPairUsername"),
//   	PrivateKey: jsii.String("privateKey"),
//
//   	// the properties below are optional
//   	PrivateKeyPassphrase: jsii.String("privateKeyPassphrase"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html
//
type CfnDataSource_KeyPairCredentialsProperty struct {
	// Username.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html#cfn-quicksight-datasource-keypaircredentials-keypairusername
	//
	KeyPairUsername *string `field:"required" json:"keyPairUsername" yaml:"keyPairUsername"`
	// PrivateKey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html#cfn-quicksight-datasource-keypaircredentials-privatekey
	//
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// PrivateKeyPassphrase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-keypaircredentials.html#cfn-quicksight-datasource-keypaircredentials-privatekeypassphrase
	//
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
}

