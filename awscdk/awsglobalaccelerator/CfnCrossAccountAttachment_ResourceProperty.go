package awsglobalaccelerator


// A resource is one of the following: the ARN for an AWS resource that is supported by AWS Global Accelerator to be added as an endpoint, or a CIDR range that specifies a bring your own IP (BYOIP) address pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceProperty := &ResourceProperty{
//   	EndpointId: jsii.String("endpointId"),
//
//   	// the properties below are optional
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html
//
type CfnCrossAccountAttachment_ResourceProperty struct {
	// The endpoint ID for the endpoint that is specified as a AWS resource.
	//
	// An endpoint ID for the cross-account feature is the ARN of an AWS resource, such as a Network Load Balancer, that Global Accelerator supports as an endpoint for an accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html#cfn-globalaccelerator-crossaccountattachment-resource-endpointid
	//
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// The AWS Region where a shared endpoint resource is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-crossaccountattachment-resource.html#cfn-globalaccelerator-crossaccountattachment-resource-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

