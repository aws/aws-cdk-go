package previewawspcsmixins


// The JWT authentication configuration for Slurm REST API access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jwtAuthProperty := &JwtAuthProperty{
//   	JwtKey: &JwtKeyProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		SecretVersion: jsii.String("secretVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-jwtauth.html
//
type CfnClusterPropsMixin_JwtAuthProperty struct {
	// The JWT key for Slurm REST API authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-jwtauth.html#cfn-pcs-cluster-jwtauth-jwtkey
	//
	JwtKey interface{} `field:"optional" json:"jwtKey" yaml:"jwtKey"`
}

