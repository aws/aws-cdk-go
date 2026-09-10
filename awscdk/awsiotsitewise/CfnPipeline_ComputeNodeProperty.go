package awsiotsitewise


// A single compute node in a pipeline DAG.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeNodeProperty := &ComputeNodeProperty{
//   	ComputeNodeName: jsii.String("computeNodeName"),
//   	TaskName: jsii.String("taskName"),
//
//   	// the properties below are optional
//   	DependsOn: []*string{
//   		jsii.String("dependsOn"),
//   	},
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html
//
type CfnPipeline_ComputeNodeProperty struct {
	// The unique name for this compute node within the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-computenodename
	//
	ComputeNodeName *string `field:"required" json:"computeNodeName" yaml:"computeNodeName"`
	// The name of the task to execute for this compute node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-taskname
	//
	TaskName *string `field:"required" json:"taskName" yaml:"taskName"`
	// A list of compute node names that must complete successfully before this node can start.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-dependson
	//
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// A map of environment variable key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-pipeline-computenode.html#cfn-iotsitewise-pipeline-computenode-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
}

