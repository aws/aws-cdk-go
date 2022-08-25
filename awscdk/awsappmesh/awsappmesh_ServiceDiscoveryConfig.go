package awsappmesh


// Properties for VirtualNode Service Discovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceDiscoveryConfig := &serviceDiscoveryConfig{
//   	cloudmap: &awsCloudMapServiceDiscoveryProperty{
//   		namespaceName: jsii.String("namespaceName"),
//   		serviceName: jsii.String("serviceName"),
//
//   		// the properties below are optional
//   		attributes: []interface{}{
//   			&awsCloudMapInstanceAttributeProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ipPreference: jsii.String("ipPreference"),
//   	},
//   	dns: &dnsServiceDiscoveryProperty{
//   		hostname: jsii.String("hostname"),
//
//   		// the properties below are optional
//   		ipPreference: jsii.String("ipPreference"),
//   		responseType: jsii.String("responseType"),
//   	},
//   }
//
type ServiceDiscoveryConfig struct {
	// Cloud Map based Service Discovery.
	Cloudmap *CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty `field:"optional" json:"cloudmap" yaml:"cloudmap"`
	// DNS based Service Discovery.
	Dns *CfnVirtualNode_DnsServiceDiscoveryProperty `field:"optional" json:"dns" yaml:"dns"`
}

