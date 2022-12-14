package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchContainerOverridesProperty := &batchContainerOverridesProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	environment: []interface{}{
//   		&batchEnvironmentVariableProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	instanceType: jsii.String("instanceType"),
//   	resourceRequirements: []interface{}{
//   		&batchResourceRequirementProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipe_BatchContainerOverridesProperty struct {
	// `CfnPipe.BatchContainerOverridesProperty.Command`.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// `CfnPipe.BatchContainerOverridesProperty.Environment`.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// `CfnPipe.BatchContainerOverridesProperty.InstanceType`.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// `CfnPipe.BatchContainerOverridesProperty.ResourceRequirements`.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}

