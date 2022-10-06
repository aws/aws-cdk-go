package awseks


// Properties for a KubectlProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   kubectlProviderProps := &kubectlProviderProps{
//   	cluster: cluster,
//   }
//
type KubectlProviderProps struct {
	// The cluster to control.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

