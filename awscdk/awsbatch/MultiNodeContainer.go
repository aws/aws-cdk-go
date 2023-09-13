package awsbatch


// Runs the container on nodes [startNode, endNode].
//
// Example:
//   multiNodeJob := batch.NewMultiNodeJobDefinition(this, jsii.String("JobDefinition"), &MultiNodeJobDefinitionProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R4, ec2.InstanceSize_LARGE),
//   	Containers: []multiNodeContainer{
//   		&multiNodeContainer{
//   			Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("mainMPIContainer"), &EcsEc2ContainerDefinitionProps{
//   				Image: ecs.ContainerImage_FromRegistry(jsii.String("yourregsitry.com/yourMPIImage:latest")),
//   				Cpu: jsii.Number(256),
//   				Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   			StartNode: jsii.Number(0),
//   			EndNode: jsii.Number(5),
//   		},
//   	},
//   })
//   // convenience method
//   multiNodeJob.AddContainer(&multiNodeContainer{
//   	StartNode: jsii.Number(6),
//   	EndNode: jsii.Number(10),
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("multiContainer"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Cpu: jsii.Number(256),
//   		Memory: cdk.Size_*Mebibytes(jsii.Number(2048)),
//   	}),
//   })
//
type MultiNodeContainer struct {
	// The container that this node range will run.
	Container IEcsContainerDefinition `field:"required" json:"container" yaml:"container"`
	// The index of the last node to run this container.
	//
	// The container is run on all nodes in the range [startNode, endNode] (inclusive).
	EndNode *float64 `field:"required" json:"endNode" yaml:"endNode"`
	// The index of the first node to run this container.
	//
	// The container is run on all nodes in the range [startNode, endNode] (inclusive).
	StartNode *float64 `field:"required" json:"startNode" yaml:"startNode"`
}

