package previewawspcsmixins


// The shared Slurm key for authentication, also known as the *cluster secret* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authKeyProperty := &AuthKeyProperty{
//   	SecretArn: jsii.String("secretArn"),
//   	SecretVersion: jsii.String("secretVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-authkey.html
//
type CfnClusterPropsMixin_AuthKeyProperty struct {
	// The Amazon Resource Name (ARN) of the shared Slurm key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-authkey.html#cfn-pcs-cluster-authkey-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The version of the shared Slurm key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-authkey.html#cfn-pcs-cluster-authkey-secretversion
	//
	SecretVersion *string `field:"optional" json:"secretVersion" yaml:"secretVersion"`
}

