package previewawsdaxmixins


// Represents the settings used to enable server-side encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sSESpecificationProperty := &SSESpecificationProperty{
//   	SseEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dax-cluster-ssespecification.html
//
type CfnClusterPropsMixin_SSESpecificationProperty struct {
	// Indicates whether server-side encryption is enabled (true) or disabled (false) on the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dax-cluster-ssespecification.html#cfn-dax-cluster-ssespecification-sseenabled
	//
	SseEnabled interface{} `field:"optional" json:"sseEnabled" yaml:"sseEnabled"`
}

