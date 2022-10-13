package awscodedeploy


// Construction properties for {@link LambdaApplication}.
//
// Example:
//   application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &lambdaApplicationProps{
//   	applicationName: jsii.String("MyApplication"),
//   })
//
type LambdaApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

