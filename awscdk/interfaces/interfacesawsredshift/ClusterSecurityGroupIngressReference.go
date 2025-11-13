package interfacesawsredshift


// A reference to a ClusterSecurityGroupIngress resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSecurityGroupIngressReference := &ClusterSecurityGroupIngressReference{
//   	ClusterSecurityGroupIngressId: jsii.String("clusterSecurityGroupIngressId"),
//   }
//
type ClusterSecurityGroupIngressReference struct {
	// The Id of the ClusterSecurityGroupIngress resource.
	ClusterSecurityGroupIngressId *string `field:"required" json:"clusterSecurityGroupIngressId" yaml:"clusterSecurityGroupIngressId"`
}

