package awsappconfig


// Properties for the Application construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationProps := &ApplicationProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	Description: jsii.String("description"),
//   }
//
type ApplicationProps struct {
	// The name of the application.
	// Default: - A name is generated.
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The description for the application.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

