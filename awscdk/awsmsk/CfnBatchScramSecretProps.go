package awsmsk


// Properties for defining a `CfnBatchScramSecret`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBatchScramSecretProps := &CfnBatchScramSecretProps{
//   	ClusterArn: jsii.String("clusterArn"),
//
//   	// the properties below are optional
//   	SecretArnList: []*string{
//   		jsii.String("secretArnList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html
//
type CfnBatchScramSecretProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html#cfn-msk-batchscramsecret-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-batchscramsecret.html#cfn-msk-batchscramsecret-secretarnlist
	//
	SecretArnList *[]*string `field:"optional" json:"secretArnList" yaml:"secretArnList"`
}

