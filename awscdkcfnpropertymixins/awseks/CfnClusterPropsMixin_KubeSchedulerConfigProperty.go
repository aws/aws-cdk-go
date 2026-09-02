package awseks


// The configuration for the Kubernetes scheduler on an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kubeSchedulerConfigProperty := &KubeSchedulerConfigProperty{
//   	NodeResourcesFit: &NodeResourcesFitConfigProperty{
//   		ScoringStrategy: &ScoringStrategyProperty{
//   			Resources: []interface{}{
//   				&ResourceWeightProperty{
//   					Name: jsii.String("name"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubeschedulerconfig.html
//
type CfnClusterPropsMixin_KubeSchedulerConfigProperty struct {
	// The NodeResourcesFit plugin configuration for the Kubernetes scheduler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubeschedulerconfig.html#cfn-eks-cluster-kubeschedulerconfig-noderesourcesfit
	//
	NodeResourcesFit interface{} `field:"optional" json:"nodeResourcesFit" yaml:"nodeResourcesFit"`
}

