package interfacesawsec2


// A reference to a ApplicationStatusCheck resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationStatusCheckReference := &ApplicationStatusCheckReference{
//   	ApplicationStatusCheckArn: jsii.String("applicationStatusCheckArn"),
//   }
//
type ApplicationStatusCheckReference struct {
	// The Arn of the ApplicationStatusCheck resource.
	ApplicationStatusCheckArn *string `field:"required" json:"applicationStatusCheckArn" yaml:"applicationStatusCheckArn"`
}

