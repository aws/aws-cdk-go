package awsemrcontainers


// The information about the EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksInfoProperty := &eksInfoProperty{
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnVirtualCluster_EksInfoProperty struct {
	// The namespaces of the EKS cluster.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 63
	//
	// *Pattern* : `[a-z0-9]([-a-z0-9]*[a-z0-9])?`
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

