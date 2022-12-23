package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksConfigurationProperty := &eksConfigurationProperty{
//   	eksClusterArn: jsii.String("eksClusterArn"),
//   	kubernetesNamespace: jsii.String("kubernetesNamespace"),
//   }
//
type CfnComputeEnvironment_EksConfigurationProperty struct {
	// `CfnComputeEnvironment.EksConfigurationProperty.EksClusterArn`.
	EksClusterArn *string `field:"required" json:"eksClusterArn" yaml:"eksClusterArn"`
	// `CfnComputeEnvironment.EksConfigurationProperty.KubernetesNamespace`.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

