package awsec2


// Describes a route in a transit gateway route table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteTableRouteProperty := &transitGatewayRouteTableRouteProperty{
//   	attachmentId: jsii.String("attachmentId"),
//   	destinationCidr: jsii.String("destinationCidr"),
//   	prefixListId: jsii.String("prefixListId"),
//   	resourceId: jsii.String("resourceId"),
//   	resourceType: jsii.String("resourceType"),
//   	routeOrigin: jsii.String("routeOrigin"),
//   	state: jsii.String("state"),
//   }
//
type CfnNetworkInsightsAnalysis_TransitGatewayRouteTableRouteProperty struct {
	// The ID of the route attachment.
	AttachmentId *string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// The CIDR block used for destination matches.
	DestinationCidr *string `field:"optional" json:"destinationCidr" yaml:"destinationCidr"`
	// The ID of the prefix list.
	PrefixListId *string `field:"optional" json:"prefixListId" yaml:"prefixListId"`
	// The ID of the resource for the route attachment.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The resource type for the route attachment.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The route origin. The following are the possible values:.
	//
	// - static
	// - propagated.
	RouteOrigin *string `field:"optional" json:"routeOrigin" yaml:"routeOrigin"`
	// The state of the route.
	State *string `field:"optional" json:"state" yaml:"state"`
}

