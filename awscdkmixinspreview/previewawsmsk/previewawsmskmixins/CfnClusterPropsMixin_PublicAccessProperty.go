package previewawsmskmixins


// Broker access controls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publicAccessProperty := &PublicAccessProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-publicaccess.html
//
type CfnClusterPropsMixin_PublicAccessProperty struct {
	// DISABLED means that public access is turned off.
	//
	// SERVICE_PROVIDED_EIPS means that public access is turned on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-publicaccess.html#cfn-msk-cluster-publicaccess-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

