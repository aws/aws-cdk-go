package awseks


// The encryption configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &encryptionConfigProperty{
//   	provider: &providerProperty{
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	resources: []*string{
//   		jsii.String("resources"),
//   	},
//   }
//
type CfnCluster_EncryptionConfigProperty struct {
	// The encryption provider for the cluster.
	Provider interface{} `field:"optional" json:"provider" yaml:"provider"`
	// Specifies the resources to be encrypted.
	//
	// The only supported value is "secrets".
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

