package awsglobalaccelerator


// A complex type for endpoints.
//
// A resource must be valid and active when you add it as an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &EndpointConfigurationProperty{
//   	EndpointId: jsii.String("endpointId"),
//
//   	// the properties below are optional
//   	AttachmentArn: jsii.String("attachmentArn"),
//   	ClientIpPreservationEnabled: jsii.Boolean(false),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html
//
type CfnEndpointGroup_EndpointConfigurationProperty struct {
	// An ID for the endpoint.
	//
	// If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID. For Amazon EC2 instances, this is the EC2 instance ID. A resource must be valid and active when you add it as an endpoint.
	//
	// For cross-account endpoints, this must be the ARN of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-endpointid
	//
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// The Amazon Resource Name (ARN) of the cross-account attachment that specifies the endpoints (resources) that can be added to accelerators and principals that have permission to add the endpoints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-attachmentarn
	//
	AttachmentArn *string `field:"optional" json:"attachmentArn" yaml:"attachmentArn"`
	// Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint.
	//
	// The value is true or false. The default value is true for new accelerators.
	//
	// If the value is set to true, the client's IP address is preserved in the `X-Forwarded-For` request header as traffic travels to applications on the Application Load Balancer endpoint fronted by the accelerator.
	//
	// For more information, see [Preserve Client IP Addresses](https://docs.aws.amazon.com/global-accelerator/latest/dg/preserve-client-ip-address.html) in the *AWS Global Accelerator Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-clientippreservationenabled
	//
	// Default: - true.
	//
	ClientIpPreservationEnabled interface{} `field:"optional" json:"clientIpPreservationEnabled" yaml:"clientIpPreservationEnabled"`
	// The weight associated with the endpoint.
	//
	// When you add weights to endpoints, you configure Global Accelerator to route traffic based on proportions that you specify. For example, you might specify endpoint weights of 4, 5, 5, and 6 (sum=20). The result is that 4/20 of your traffic, on average, is routed to the first endpoint, 5/20 is routed both to the second and third endpoints, and 6/20 is routed to the last endpoint. For more information, see [Endpoint Weights](https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html) in the *AWS Global Accelerator Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html#cfn-globalaccelerator-endpointgroup-endpointconfiguration-weight
	//
	// Default: - 100.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

