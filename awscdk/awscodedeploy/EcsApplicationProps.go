package awscodedeploy


// Construction properties for `EcsApplication`.
//
// Example:
//   application := codedeploy.NewEcsApplication(this, jsii.String("CodeDeployApplication"), &EcsApplicationProps{
//   	ApplicationName: jsii.String("MyApplication"),
//   })
//
type EcsApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Default: an auto-generated name will be used.
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

