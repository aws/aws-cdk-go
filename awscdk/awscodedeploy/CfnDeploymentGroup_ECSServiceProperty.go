package awscodedeploy


// Contains the service and cluster names used to identify an Amazon ECS deployment's target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eCSServiceProperty := &ECSServiceProperty{
//   	ClusterName: jsii.String("clusterName"),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ecsservice.html
//
type CfnDeploymentGroup_ECSServiceProperty struct {
	// The name of the cluster that the Amazon ECS service is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ecsservice.html#cfn-codedeploy-deploymentgroup-ecsservice-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The name of the target Amazon ECS service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ecsservice.html#cfn-codedeploy-deploymentgroup-ecsservice-servicename
	//
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

