package awscodedeploy


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReference := &ApplicationReference{
//   	ApplicationName: jsii.String("applicationName"),
//   }
//
type ApplicationReference struct {
	// The ApplicationName of the Application resource.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
}

