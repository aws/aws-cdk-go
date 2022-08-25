package awsecs


// The configuration to use when setting an App Mesh proxy configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appMeshProxyConfigurationConfigProps := &appMeshProxyConfigurationConfigProps{
//   	containerName: jsii.String("containerName"),
//   	properties: &appMeshProxyConfigurationProps{
//   		appPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		proxyEgressPort: jsii.Number(123),
//   		proxyIngressPort: jsii.Number(123),
//
//   		// the properties below are optional
//   		egressIgnoredIPs: []*string{
//   			jsii.String("egressIgnoredIPs"),
//   		},
//   		egressIgnoredPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		ignoredGID: jsii.Number(123),
//   		ignoredUID: jsii.Number(123),
//   	},
//   }
//
type AppMeshProxyConfigurationConfigProps struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin.
	Properties *AppMeshProxyConfigurationProps `field:"required" json:"properties" yaml:"properties"`
}

