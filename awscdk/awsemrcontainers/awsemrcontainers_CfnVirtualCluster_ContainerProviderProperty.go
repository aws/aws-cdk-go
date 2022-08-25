package awsemrcontainers


// The information about the container provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerProviderProperty := &containerProviderProperty{
//   	id: jsii.String("id"),
//   	info: &containerInfoProperty{
//   		eksInfo: &eksInfoProperty{
//   			namespace: jsii.String("namespace"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnVirtualCluster_ContainerProviderProperty struct {
	// The ID of the container cluster.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 100
	//
	// *Pattern* : `^[0-9A-Za-z][A-Za-z0-9\-_]*`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The information about the container cluster.
	Info interface{} `field:"required" json:"info" yaml:"info"`
	// The type of the container provider.
	//
	// EKS is the only supported type as of now.
	Type *string `field:"required" json:"type" yaml:"type"`
}

