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
//   awsCloudMapServiceDiscoveryProperty := &AwsCloudMapServiceDiscoveryProperty{
//   	NamespaceName: jsii.String("namespaceName"),
//   	ServiceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	Attributes: []interface{}{
//   		&AwsCloudMapInstanceAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	IpPreference: jsii.String("ipPreference"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapservicediscovery.html
//
type CfnVirtualNode_AwsCloudMapServiceDiscoveryProperty struct {
	// The HTTP name of the AWS Cloud Map namespace to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapservicediscovery.html#cfn-appmesh-virtualnode-awscloudmapservicediscovery-namespacename
	//
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// The name of the AWS Cloud Map service to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapservicediscovery.html#cfn-appmesh-virtualnode-awscloudmapservicediscovery-servicename
	//
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// A string map that contains attributes with values that you can use to filter instances by any custom attribute that you specified when you registered the instance.
	//
	// Only instances that match all of the specified key/value pairs will be returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapservicediscovery.html#cfn-appmesh-virtualnode-awscloudmapservicediscovery-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The preferred IP version that this virtual node uses.
	//
	// Setting the IP preference on the virtual node only overrides the IP preference set for the mesh on this specific node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-awscloudmapservicediscovery.html#cfn-appmesh-virtualnode-awscloudmapservicediscovery-ippreference
	//
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}

