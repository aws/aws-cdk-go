package awscodedeploy


// Construction properties for {@link EcsApplication}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsApplicationProps := &ecsApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   }
//
type EcsApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

