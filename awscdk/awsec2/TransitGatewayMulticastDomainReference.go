package awsec2


// A reference to a TransitGatewayMulticastDomain resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayMulticastDomainReference := &TransitGatewayMulticastDomainReference{
//   	TransitGatewayMulticastDomainArn: jsii.String("transitGatewayMulticastDomainArn"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type TransitGatewayMulticastDomainReference struct {
	// The ARN of the TransitGatewayMulticastDomain resource.
	TransitGatewayMulticastDomainArn *string `field:"required" json:"transitGatewayMulticastDomainArn" yaml:"transitGatewayMulticastDomainArn"`
	// The TransitGatewayMulticastDomainId of the TransitGatewayMulticastDomain resource.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

