package awsecs


// The configuration to use when setting an App Mesh proxy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfigurationConfigProps := &AppMeshProxyConfigurationConfigProps{
//   	ContainerName: jsii.String("containerName"),
//   	Properties: &AppMeshProxyConfigurationProps{
//   		AppPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		ProxyEgressPort: jsii.Number(123),
//   		ProxyIngressPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		EgressIgnoredIPs: []*string{
//   			jsii.String("egressIgnoredIPs"),
//   		},
//   		EgressIgnoredPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		IgnoredGID: jsii.Number(123),
//   		IgnoredUID: jsii.Number(123),
//   	},
//   }
//
type AppMeshProxyConfigurationConfigProps struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin.
	Properties *AppMeshProxyConfigurationProps `field:"required" json:"properties" yaml:"properties"`
}

