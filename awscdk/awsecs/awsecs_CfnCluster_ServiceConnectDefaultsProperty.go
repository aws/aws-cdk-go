package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectDefaultsProperty := &serviceConnectDefaultsProperty{
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnCluster_ServiceConnectDefaultsProperty struct {
	// `CfnCluster.ServiceConnectDefaultsProperty.Namespace`.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

