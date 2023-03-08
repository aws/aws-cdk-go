package awsmsk


// Properties for defining a `CfnBatchScramSecret`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBatchScramSecretProps := &cfnBatchScramSecretProps{
//   	clusterArn: jsii.String("clusterArn"),
//
//   	// the properties below are optional
//   	secretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   }
//
type CfnBatchScramSecretProps struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// A list of Amazon Secrets Manager secret ARNs.
	SecretArnList *[]*string `field:"optional" json:"secretArnList" yaml:"secretArnList"`
}

