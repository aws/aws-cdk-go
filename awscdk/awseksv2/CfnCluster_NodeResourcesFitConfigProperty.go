package awseksv2


// The NodeResourcesFit plugin configuration for the Kubernetes scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeResourcesFitConfigProperty := &NodeResourcesFitConfigProperty{
//   	ScoringStrategy: &ScoringStrategyProperty{
//   		Resources: []interface{}{
//   			&ResourceWeightProperty{
//   				Name: jsii.String("name"),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-noderesourcesfitconfig.html
//
type CfnCluster_NodeResourcesFitConfigProperty struct {
	// The scoring strategy configuration for the NodeResourcesFit scheduler plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-noderesourcesfitconfig.html#cfn-eks-cluster-noderesourcesfitconfig-scoringstrategy
	//
	ScoringStrategy interface{} `field:"optional" json:"scoringStrategy" yaml:"scoringStrategy"`
}

