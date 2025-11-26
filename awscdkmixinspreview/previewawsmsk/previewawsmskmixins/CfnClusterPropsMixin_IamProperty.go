package previewawsmskmixins


// Details for SASL/IAM client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iamProperty := &IamProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-iam.html
//
type CfnClusterPropsMixin_IamProperty struct {
	// SASL/IAM authentication is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-iam.html#cfn-msk-cluster-iam-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

