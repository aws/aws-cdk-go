package awsappmesh


// An object that represents the service discovery information for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceDiscoveryProperty := &serviceDiscoveryProperty{
//   	awsCloudMap: &awsCloudMapServiceDiscoveryProperty{
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
type CfnVirtualNode_ServiceDiscoveryProperty struct {
	// Specifies any AWS Cloud Map information for the virtual node.
	AwsCloudMap interface{} `field:"optional" json:"awsCloudMap" yaml:"awsCloudMap"`
	// Specifies the DNS information for the virtual node.
	Dns interface{} `field:"optional" json:"dns" yaml:"dns"`
}

