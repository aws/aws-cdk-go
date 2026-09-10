package awsiotsitewise


// A single compute node in a pipeline DAG.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   computeNodeProperty := &ComputeNodeProperty{
//   	ComputeNodeName: jsii.String("computeNodeName"),
//   	DependsOn: []*string{
//   		jsii.String("dependsOn"),
//   	},
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	TaskName: jsii.String("taskName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html
//
type CfnPipelinePropsMixin_ComputeNodeProperty struct {
	// The unique name for this compute node within the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-computenodename
	//
	ComputeNodeName *string `field:"optional" json:"computeNodeName" yaml:"computeNodeName"`
	// A list of compute node names that must complete successfully before this node can start.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-dependson
	//
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// A map of environment variable key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The name of the task to execute for this compute node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-taskname
	//
	TaskName *string `field:"optional" json:"taskName" yaml:"taskName"`
}

