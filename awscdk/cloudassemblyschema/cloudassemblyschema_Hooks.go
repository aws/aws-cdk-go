package cloudassemblyschema


// Commands to run at predefined points during the integration test workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hooks := &hooks{
//   	postDeploy: []*string{
//   		jsii.String("postDeploy"),
//   	},
//   	postDestroy: []*string{
//   		jsii.String("postDestroy"),
//   	},
//   	preDeploy: []*string{
//   		jsii.String("preDeploy"),
//   	},
//   	preDestroy: []*string{
//   		jsii.String("preDestroy"),
//   	},
//   }
//
type Hooks struct {
	// Commands to run prior after deploying the cdk stacks in the integration test.
	PostDeploy *[]*string `field:"optional" json:"postDeploy" yaml:"postDeploy"`
	// Commands to run after destroying the cdk stacks in the integration test.
	PostDestroy *[]*string `field:"optional" json:"postDestroy" yaml:"postDestroy"`
	// Commands to run prior to deploying the cdk stacks in the integration test.
	PreDeploy *[]*string `field:"optional" json:"preDeploy" yaml:"preDeploy"`
	// Commands to run prior to destroying the cdk stacks in the integration test.
	PreDestroy *[]*string `field:"optional" json:"preDestroy" yaml:"preDestroy"`
}

