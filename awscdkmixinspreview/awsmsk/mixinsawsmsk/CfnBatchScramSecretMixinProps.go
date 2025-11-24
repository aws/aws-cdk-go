package mixinsawsmsk


// Properties for CfnBatchScramSecretPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBatchScramSecretMixinProps := &CfnBatchScramSecretMixinProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	SecretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html
//
type CfnBatchScramSecretMixinProps struct {
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html#cfn-msk-batchscramsecret-clusterarn
	//
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// List of Amazon Resource Name (ARN)s of Secrets Manager secrets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html#cfn-msk-batchscramsecret-secretarnlist
	//
	SecretArnList *[]*string `field:"optional" json:"secretArnList" yaml:"secretArnList"`
}

