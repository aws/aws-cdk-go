package awsquicksight


// The configuration of loading animation in free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadingAnimationProperty := &LoadingAnimationProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_LoadingAnimationProperty struct {
	// The visibility configuration of `LoadingAnimation` .
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

