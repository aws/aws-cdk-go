package previewawsmskmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rebalancingProperty := &RebalancingProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-rebalancing.html
//
type CfnClusterPropsMixin_RebalancingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-rebalancing.html#cfn-msk-cluster-rebalancing-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

