package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentRunWithProperty := &componentRunWithProperty{
//   	posixUser: jsii.String("posixUser"),
//   	systemResourceLimits: &systemResourceLimitsProperty{
//   		cpus: jsii.Number(123),
//   		memory: jsii.Number(123),
//   	},
//   	windowsUser: jsii.String("windowsUser"),
//   }
//
type CfnDeployment_ComponentRunWithProperty struct {
	// `CfnDeployment.ComponentRunWithProperty.PosixUser`.
	PosixUser *string `field:"optional" json:"posixUser" yaml:"posixUser"`
	// `CfnDeployment.ComponentRunWithProperty.SystemResourceLimits`.
	SystemResourceLimits interface{} `field:"optional" json:"systemResourceLimits" yaml:"systemResourceLimits"`
	// `CfnDeployment.ComponentRunWithProperty.WindowsUser`.
	WindowsUser *string `field:"optional" json:"windowsUser" yaml:"windowsUser"`
}

