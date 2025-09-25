package awskinesisanalytics


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationReference := &ApplicationReference{
//   	ApplicationId: jsii.String("applicationId"),
//   }
//
type ApplicationReference struct {
	// The Id of the Application resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
}

