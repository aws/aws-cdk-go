package awscdkmskalpha


// The Amazon MSK configuration to use for the cluster.
//
// Note: There is currently no Cloudformation Resource to create a Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//
//   clusterConfigurationInfo := &ClusterConfigurationInfo{
//   	Arn: jsii.String("arn"),
//   	Revision: jsii.Number(123),
//   }
//
// Experimental.
type ClusterConfigurationInfo struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The revision of the Amazon MSK configuration to use.
	// Experimental.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

