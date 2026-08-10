package interfacesawsidentitystore


// A reference to a AllGroupMemberships resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allGroupMembershipsReference := &AllGroupMembershipsReference{
//   	AllGroupMembershipsArn: jsii.String("allGroupMembershipsArn"),
//   }
//
type AllGroupMembershipsReference struct {
	// The Arn of the AllGroupMemberships resource.
	AllGroupMembershipsArn *string `field:"required" json:"allGroupMembershipsArn" yaml:"allGroupMembershipsArn"`
}

