package previewawseksmixins


// The encryption configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	Provider: &ProviderProperty{
//   		KeyArn: jsii.String("keyArn"),
//   	},
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-encryptionconfig.html
//
type CfnClusterPropsMixin_EncryptionConfigProperty struct {
	// The encryption provider for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-encryptionconfig.html#cfn-eks-cluster-encryptionconfig-provider
	//
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
	// Specifies the resources to be encrypted.
	//
	// The only supported value is `secrets` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-encryptionconfig.html#cfn-eks-cluster-encryptionconfig-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

