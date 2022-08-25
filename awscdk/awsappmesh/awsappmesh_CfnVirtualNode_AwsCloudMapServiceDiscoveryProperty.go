package awsappmesh


// An object that represents the AWS Cloud Map service discovery information for your virtual node.
//
// > AWS Cloud Map is not available in the eu-south-1 Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsCloudMapServiceDiscoveryProperty := &awsCloudMapServiceDiscoveryProperty{
//   	namespaceName: jsii.String("namespaceName"),
//   	serviceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	attributes: []interface{}{
//   		&awsCloudMapInstanceAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ipPreference: jsii.String("ipPreference"),
//   }
//
type CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty struct {
	// The name of the AWS Cloud Map namespace to use.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// The name of the AWS Cloud Map service to use.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// A string map that contains attributes with values that you can use to filter instances by any custom attribute that you specified when you registered the instance.
	//
	// Only instances that match all of the specified key/value pairs will be returned.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The preferred IP version that this virtual node uses.
	//
	// Setting the IP preference on the virtual node only overrides the IP preference set for the mesh on this specific node.
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}

