package interfacesawspanorama


// A reference to a ApplicationInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationInstanceReference := &ApplicationInstanceReference{
//   	ApplicationInstanceArn: jsii.String("applicationInstanceArn"),
//   	ApplicationInstanceId: jsii.String("applicationInstanceId"),
//   }
//
type ApplicationInstanceReference struct {
	// The ARN of the ApplicationInstance resource.
	ApplicationInstanceArn *string `field:"required" json:"applicationInstanceArn" yaml:"applicationInstanceArn"`
	// The ApplicationInstanceId of the ApplicationInstance resource.
	ApplicationInstanceId *string `field:"required" json:"applicationInstanceId" yaml:"applicationInstanceId"`
}

