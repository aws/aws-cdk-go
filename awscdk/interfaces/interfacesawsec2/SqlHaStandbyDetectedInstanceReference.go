package interfacesawsec2


// A reference to a SqlHaStandbyDetectedInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlHaStandbyDetectedInstanceReference := &SqlHaStandbyDetectedInstanceReference{
//   	InstanceId: jsii.String("instanceId"),
//   }
//
type SqlHaStandbyDetectedInstanceReference struct {
	// The InstanceId of the SqlHaStandbyDetectedInstance resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

