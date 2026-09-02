package awseks


// The scoring strategy configuration for the NodeResourcesFit scheduler plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scoringStrategyProperty := &ScoringStrategyProperty{
//   	Resources: []interface{}{
//   		&ResourceWeightProperty{
//   			Name: jsii.String("name"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-scoringstrategy.html
//
type CfnClusterPropsMixin_ScoringStrategyProperty struct {
	// The resource weights used for scoring nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-scoringstrategy.html#cfn-eks-cluster-scoringstrategy-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// The scoring strategy type (LeastAllocated or MostAllocated).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-scoringstrategy.html#cfn-eks-cluster-scoringstrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

