package previewawsglobalacceleratormixins


// A resource is one of the following: the ARN for an AWS resource that is supported by AWS Global Accelerator to be added as an endpoint, or a CIDR range that specifies a bring your own IP (BYOIP) address pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceProperty := &ResourceProperty{
//   	Cidr: jsii.String("cidr"),
//   	EndpointId: jsii.String("endpointId"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html
//
type CfnCrossAccountAttachmentPropsMixin_ResourceProperty struct {
	// An IP address range, in CIDR format, that is specified as resource.
	//
	// The address must be provisioned and advertised in AWS Global Accelerator by following the bring your own IP address (BYOIP) process for Global Accelerator
	//
	// For more information, see [Bring your own IP addresses (BYOIP)](https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in the AWS Global Accelerator Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html#cfn-globalaccelerator-crossaccountattachment-resource-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The endpoint ID for the endpoint that is specified as a AWS resource.
	//
	// An endpoint ID for the cross-account feature is the ARN of an AWS resource, such as a Network Load Balancer, that Global Accelerator supports as an endpoint for an accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html#cfn-globalaccelerator-crossaccountattachment-resource-endpointid
	//
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// The AWS Region where a shared endpoint resource is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html#cfn-globalaccelerator-crossaccountattachment-resource-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

