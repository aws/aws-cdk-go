package awspcs


// JWT authentication configuration for Slurm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnCluster_JwtAuthProperty struct {
	// JWT key configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-jwtauth.html#cfn-pcs-cluster-jwtauth-jwtkey
	//
	JwtKey interface{} `field:"optional" json:"jwtKey" yaml:"jwtKey"`
}

