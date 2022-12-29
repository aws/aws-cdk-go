package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentDeploymentSpecificationProperty := &componentDeploymentSpecificationProperty{
//   	componentVersion: jsii.String("componentVersion"),
//   	configurationUpdate: &componentConfigurationUpdateProperty{
//   		merge: jsii.String("merge"),
//   		reset: []*string{
//   			jsii.String("reset"),
//   		},
//   	},
//   	runWith: &componentRunWithProperty{
//   		posixUser: jsii.String("posixUser"),
//   		systemResourceLimits: &systemResourceLimitsProperty{
//   			cpus: jsii.Number(123),
//   			memory: jsii.Number(123),
//   		},
//   		windowsUser: jsii.String("windowsUser"),
//   	},
//   }
//
type CfnDeployment_ComponentDeploymentSpecificationProperty struct {
	// `CfnDeployment.ComponentDeploymentSpecificationProperty.ComponentVersion`.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// `CfnDeployment.ComponentDeploymentSpecificationProperty.ConfigurationUpdate`.
	ConfigurationUpdate interface{} `field:"optional" json:"configurationUpdate" yaml:"configurationUpdate"`
	// `CfnDeployment.ComponentDeploymentSpecificationProperty.RunWith`.
	RunWith interface{} `field:"optional" json:"runWith" yaml:"runWith"`
}

