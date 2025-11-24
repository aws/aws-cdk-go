package mixinsawspipes


// The overrides that are sent to a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   batchContainerOverridesProperty := &BatchContainerOverridesProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Environment: []interface{}{
//   		&BatchEnvironmentVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	ResourceRequirements: []interface{}{
//   		&BatchResourceRequirementProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html
//
type CfnPipePropsMixin_BatchContainerOverridesProperty struct {
	// The command to send to the container that overrides the default command from the Docker image or the task definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition.
	//
	// > Environment variables cannot start with " `AWS Batch` ". This naming convention is reserved for variables that AWS Batch sets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The instance type to use for a multi-node parallel job.
	//
	// > This parameter isn't applicable to single-node container jobs or jobs that run on Fargate resources, and shouldn't be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The type and amount of resources to assign to a container.
	//
	// This overrides the settings in the job definition. The supported resources include `GPU` , `MEMORY` , and `VCPU` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-resourcerequirements
	//
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}

