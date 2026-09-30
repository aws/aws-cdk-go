package interfacesawselasticache


// A reference to a GlobalReplicationGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalReplicationGroupReference := &GlobalReplicationGroupReference{
//   	GlobalReplicationGroupArn: jsii.String("globalReplicationGroupArn"),
//   	GlobalReplicationGroupId: jsii.String("globalReplicationGroupId"),
//   }
//
type GlobalReplicationGroupReference struct {
	// The ARN of the GlobalReplicationGroup resource.
	GlobalReplicationGroupArn *string `field:"required" json:"globalReplicationGroupArn" yaml:"globalReplicationGroupArn"`
	// The GlobalReplicationGroupId of the GlobalReplicationGroup resource.
	GlobalReplicationGroupId *string `field:"required" json:"globalReplicationGroupId" yaml:"globalReplicationGroupId"`
}

