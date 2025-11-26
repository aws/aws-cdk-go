package previewawspcsmixins


// JWT key configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jwtKeyProperty := &JwtKeyProperty{
//   	SecretArn: jsii.String("secretArn"),
//   	SecretVersion: jsii.String("secretVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-jwtkey.html
//
type CfnClusterPropsMixin_JwtKeyProperty struct {
	// The Amazon Resource Name (ARN) of the JWT key secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-jwtkey.html#cfn-pcs-cluster-jwtkey-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The version of the JWT key secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-jwtkey.html#cfn-pcs-cluster-jwtkey-secretversion
	//
	SecretVersion *string `field:"optional" json:"secretVersion" yaml:"secretVersion"`
}

