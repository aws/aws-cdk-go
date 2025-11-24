package mixinsawsbatch


// A list of containers that this task depends on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskContainerDependencyProperty := &TaskContainerDependencyProperty{
//   	Condition: jsii.String("condition"),
//   	ContainerName: jsii.String("containerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-taskcontainerdependency.html
//
type CfnJobDefinitionPropsMixin_TaskContainerDependencyProperty struct {
	// The dependency condition of the container. The following are the available conditions and their behavior:.
	//
	// - `START` - This condition emulates the behavior of links and volumes today. It validates that a dependent container is started before permitting other containers to start.
	// - `COMPLETE` - This condition validates that a dependent container runs to completion (exits) before permitting other containers to start. This can be useful for nonessential containers that run a script and then exit. This condition can't be set on an essential container.
	// - `SUCCESS` - This condition is the same as `COMPLETE` , but it also requires that the container exits with a zero status. This condition can't be set on an essential container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-taskcontainerdependency.html#cfn-batch-jobdefinition-taskcontainerdependency-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// A unique identifier for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-taskcontainerdependency.html#cfn-batch-jobdefinition-taskcontainerdependency-containername
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
}

