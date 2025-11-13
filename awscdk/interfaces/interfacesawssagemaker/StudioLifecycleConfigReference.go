package interfacesawssagemaker


// A reference to a StudioLifecycleConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioLifecycleConfigReference := &StudioLifecycleConfigReference{
//   	StudioLifecycleConfigArn: jsii.String("studioLifecycleConfigArn"),
//   	StudioLifecycleConfigName: jsii.String("studioLifecycleConfigName"),
//   }
//
type StudioLifecycleConfigReference struct {
	// The ARN of the StudioLifecycleConfig resource.
	StudioLifecycleConfigArn *string `field:"required" json:"studioLifecycleConfigArn" yaml:"studioLifecycleConfigArn"`
	// The StudioLifecycleConfigName of the StudioLifecycleConfig resource.
	StudioLifecycleConfigName *string `field:"required" json:"studioLifecycleConfigName" yaml:"studioLifecycleConfigName"`
}

