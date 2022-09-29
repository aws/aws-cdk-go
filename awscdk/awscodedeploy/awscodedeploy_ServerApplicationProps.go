package awscodedeploy


// Construction properties for {@link ServerApplication}.
//
// Example:
//   application := codedeploy.NewServerApplication(this, jsii.String("CodeDeployApplication"), &serverApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
// Experimental.
type ServerApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

