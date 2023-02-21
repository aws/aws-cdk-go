package awsecs


// The `ProxyConfiguration` property specifies the details for the App Mesh proxy.
//
// For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// For tasks using the Fargate launch type, the task or service requires platform version 1.3.0 or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxyConfigurationProperty := &proxyConfigurationProperty{
//   	containerName: jsii.String("containerName"),
//
//   	// the properties below are optional
//   	proxyConfigurationProperties: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnTaskDefinition_ProxyConfigurationProperty struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified as key-value pairs.
	//
	// - `IgnoredUID` - (Required) The user ID (UID) of the proxy container as defined by the `user` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If `IgnoredGID` is specified, this field can be empty.
	// - `IgnoredGID` - (Required) The group ID (GID) of the proxy container as defined by the `user` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If `IgnoredUID` is specified, this field can be empty.
	// - `AppPorts` - (Required) The list of ports that the application uses. Network traffic to these ports is forwarded to the `ProxyIngressPort` and `ProxyEgressPort` .
	// - `ProxyIngressPort` - (Required) Specifies the port that incoming traffic to the `AppPorts` is directed to.
	// - `ProxyEgressPort` - (Required) Specifies the port that outgoing traffic from the `AppPorts` is directed to.
	// - `EgressIgnoredPorts` - (Required) The egress traffic going to the specified ports is ignored and not redirected to the `ProxyEgressPort` . It can be an empty list.
	// - `EgressIgnoredIPs` - (Required) The egress traffic going to the specified IP addresses is ignored and not redirected to the `ProxyEgressPort` . It can be an empty list.
	ProxyConfigurationProperties interface{} `field:"optional" json:"proxyConfigurationProperties" yaml:"proxyConfigurationProperties"`
	// The proxy type.
	//
	// The only supported value is `APPMESH` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

