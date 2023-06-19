package awscodedeploy


// Construction properties for {@link LambdaApplication}.
//
// Example:
//   application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &LambdaApplicationProps{
//   	ApplicationName: jsii.String("MyApplication"),
//   })
//
// Experimental.
type LambdaApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	// Experimental.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

