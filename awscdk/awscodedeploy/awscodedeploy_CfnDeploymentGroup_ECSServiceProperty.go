package awscodedeploy


// Contains the service and cluster names used to identify an Amazon ECS deployment's target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eCSServiceProperty := &eCSServiceProperty{
//   	clusterName: jsii.String("clusterName"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type CfnDeploymentGroup_ECSServiceProperty struct {
	// The name of the cluster that the Amazon ECS service is associated with.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The name of the target Amazon ECS service.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

