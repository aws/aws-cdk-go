package awsredshift


// A reference to a ClusterSecurityGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSecurityGroupReference := &ClusterSecurityGroupReference{
//   	ClusterSecurityGroupId: jsii.String("clusterSecurityGroupId"),
//   }
//
type ClusterSecurityGroupReference struct {
	// The Id of the ClusterSecurityGroup resource.
	ClusterSecurityGroupId *string `field:"required" json:"clusterSecurityGroupId" yaml:"clusterSecurityGroupId"`
}

