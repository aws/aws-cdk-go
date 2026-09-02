package awseks


// The configuration for the Kubernetes controller manager on an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubeControllerManagerConfigProperty := &KubeControllerManagerConfigProperty{
//   	HorizontalPodAutoscalerControllerConfig: &HorizontalPodAutoscalerControllerConfigProperty{
//   		HorizontalPodAutoscalerSyncPeriod: jsii.String("horizontalPodAutoscalerSyncPeriod"),
//   	},
//   	PodGcControllerConfig: &PodGcControllerConfigProperty{
//   		TerminatedPodGcThreshold: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubecontrollermanagerconfig.html
//
type CfnCluster_KubeControllerManagerConfigProperty struct {
	// The horizontal pod autoscaler controller configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubecontrollermanagerconfig.html#cfn-eks-cluster-kubecontrollermanagerconfig-horizontalpodautoscalercontrollerconfig
	//
	HorizontalPodAutoscalerControllerConfig interface{} `field:"optional" json:"horizontalPodAutoscalerControllerConfig" yaml:"horizontalPodAutoscalerControllerConfig"`
	// The pod garbage collector controller configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubecontrollermanagerconfig.html#cfn-eks-cluster-kubecontrollermanagerconfig-podgccontrollerconfig
	//
	PodGcControllerConfig interface{} `field:"optional" json:"podGcControllerConfig" yaml:"podGcControllerConfig"`
}

