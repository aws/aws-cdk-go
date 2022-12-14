package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsContainerOverrideProperty := &ecsContainerOverrideProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	cpu: jsii.Number(123),
//   	environment: []interface{}{
//   		&ecsEnvironmentVariableProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	environmentFiles: []interface{}{
//   		&ecsEnvironmentFileProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	memory: jsii.Number(123),
//   	memoryReservation: jsii.Number(123),
//   	name: jsii.String("name"),
//   	resourceRequirements: []interface{}{
//   		&ecsResourceRequirementProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipe_EcsContainerOverrideProperty struct {
	// `CfnPipe.EcsContainerOverrideProperty.Command`.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// `CfnPipe.EcsContainerOverrideProperty.Cpu`.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// `CfnPipe.EcsContainerOverrideProperty.Environment`.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// `CfnPipe.EcsContainerOverrideProperty.EnvironmentFiles`.
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// `CfnPipe.EcsContainerOverrideProperty.Memory`.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// `CfnPipe.EcsContainerOverrideProperty.MemoryReservation`.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// `CfnPipe.EcsContainerOverrideProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnPipe.EcsContainerOverrideProperty.ResourceRequirements`.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}

