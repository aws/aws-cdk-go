package cloudassemblyschema


// Commands to run at predefined points during the integration test workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hooks := &Hooks{
//   	PostDeploy: []*string{
//   		jsii.String("postDeploy"),
//   	},
//   	PostDestroy: []*string{
//   		jsii.String("postDestroy"),
//   	},
//   	PreDeploy: []*string{
//   		jsii.String("preDeploy"),
//   	},
//   	PreDestroy: []*string{
//   		jsii.String("preDestroy"),
//   	},
//   }
//
type Hooks struct {
	// Commands to run prior after deploying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PostDeploy *[]*string `field:"optional" json:"postDeploy" yaml:"postDeploy"`
	// Commands to run after destroying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PostDestroy *[]*string `field:"optional" json:"postDestroy" yaml:"postDestroy"`
	// Commands to run prior to deploying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PreDeploy *[]*string `field:"optional" json:"preDeploy" yaml:"preDeploy"`
	// Commands to run prior to destroying the cdk stacks in the integration test.
	// Default: - no commands.
	//
	PreDestroy *[]*string `field:"optional" json:"preDestroy" yaml:"preDestroy"`
}

