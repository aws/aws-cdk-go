package awsappmesh


// Properties for VirtualNode Service Discovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceDiscoveryConfig := &ServiceDiscoveryConfig{
//   	Cloudmap: &AwsCloudMapServiceDiscoveryProperty{
//   		NamespaceName: jsii.String("namespaceName"),
//   		ServiceName: jsii.String("serviceName"),
//
//   		// the properties below are optional
//   		Attributes: []interface{}{
//   			&AwsCloudMapInstanceAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		IpPreference: jsii.String("ipPreference"),
//   	},
//   	Dns: &DnsServiceDiscoveryProperty{
//   		Hostname: jsii.String("hostname"),
//
//   		// the properties below are optional
//   		IpPreference: jsii.String("ipPreference"),
//   		ResponseType: jsii.String("responseType"),
//   	},
//   }
//
type ServiceDiscoveryConfig struct {
	// Cloud Map based Service Discovery.
	// Default: - no Cloud Map based service discovery.
	//
	Cloudmap *CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty `field:"optional" json:"cloudmap" yaml:"cloudmap"`
	// DNS based Service Discovery.
	// Default: - no DNS based service discovery.
	//
	Dns *CfnVirtualNode_DnsServiceDiscoveryProperty `field:"optional" json:"dns" yaml:"dns"`
}

