package interfacesawscognito


// A reference to a UserPoolReplica resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolReplicaReference := &UserPoolReplicaReference{
//   	RegionName: jsii.String("regionName"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolReplicaReference struct {
	// The RegionName of the UserPoolReplica resource.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// The UserPoolId of the UserPoolReplica resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

