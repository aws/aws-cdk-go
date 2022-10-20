package awscodedeploy


// Construction properties for {@link EcsApplication}.
//
// Example:
//   application := codedeploy.NewEcsApplication(this, jsii.String("CodeDeployApplication"), &ecsApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
type EcsApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

