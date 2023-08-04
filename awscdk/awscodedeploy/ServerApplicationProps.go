package awscodedeploy


// Construction properties for `ServerApplication`.
//
// Example:
//   application := codedeploy.NewServerApplication(this, jsii.String("CodeDeployApplication"), &ServerApplicationProps{
//   	ApplicationName: jsii.String("MyApplication"),
//   })
//
type ServerApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Default: an auto-generated name will be used.
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

