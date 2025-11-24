package mixinsawsglobalaccelerator


// Properties for CfnListenerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnListenerMixinProps := &CfnListenerMixinProps{
//   	AcceleratorArn: jsii.String("acceleratorArn"),
//   	ClientAffinity: jsii.String("clientAffinity"),
//   	PortRanges: []interface{}{
//   		&PortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html
//
type CfnListenerMixinProps struct {
	// The Amazon Resource Name (ARN) of your accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-acceleratorarn
	//
	AcceleratorArn *string `field:"optional" json:"acceleratorArn" yaml:"acceleratorArn"`
	// Client affinity lets you direct all requests from a user to the same endpoint, if you have stateful applications, regardless of the port and protocol of the client request.
	//
	// Client affinity gives you control over whether to always route each client to the same specific endpoint.
	//
	// AWS Global Accelerator uses a consistent-flow hashing algorithm to choose the optimal endpoint for a connection. If client affinity is `NONE` , Global Accelerator uses the "five-tuple" (5-tuple) properties—source IP address, source port, destination IP address, destination port, and protocol—to select the hash value, and then chooses the best endpoint. However, with this setting, if someone uses different ports to connect to Global Accelerator, their connections might not be always routed to the same endpoint because the hash value changes.
	//
	// If you want a given client to always be routed to the same endpoint, set client affinity to `SOURCE_IP` instead. When you use the `SOURCE_IP` setting, Global Accelerator uses the "two-tuple" (2-tuple) properties— source (client) IP address and destination IP address—to select the hash value.
	//
	// The default value is `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-clientaffinity
	//
	// Default: - "NONE".
	//
	ClientAffinity *string `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// The list of port ranges for the connections from clients to the accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-portranges
	//
	PortRanges interface{} `field:"optional" json:"portRanges" yaml:"portRanges"`
	// The protocol for the connections from clients to the accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-listener.html#cfn-globalaccelerator-listener-protocol
	//
	// Default: - "TCP".
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

