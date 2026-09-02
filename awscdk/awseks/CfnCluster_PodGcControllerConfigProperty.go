package awseks


// The pod garbage collector controller configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   podGcControllerConfigProperty := &PodGcControllerConfigProperty{
//   	TerminatedPodGcThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-podgccontrollerconfig.html
//
type CfnCluster_PodGcControllerConfigProperty struct {
	// The number of terminated pods that can exist before the terminated pod garbage collector starts deleting them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-podgccontrollerconfig.html#cfn-eks-cluster-podgccontrollerconfig-terminatedpodgcthreshold
	//
	TerminatedPodGcThreshold *float64 `field:"optional" json:"terminatedPodGcThreshold" yaml:"terminatedPodGcThreshold"`
}

