package awsec2


// Describes a route in a transit gateway route table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteTableRouteProperty := &TransitGatewayRouteTableRouteProperty{
//   	AttachmentId: jsii.String("attachmentId"),
//   	DestinationCidr: jsii.String("destinationCidr"),
//   	PrefixListId: jsii.String("prefixListId"),
//   	ResourceId: jsii.String("resourceId"),
//   	ResourceType: jsii.String("resourceType"),
//   	RouteOrigin: jsii.String("routeOrigin"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html
//
type CfnNetworkInsightsAnalysis_TransitGatewayRouteTableRouteProperty struct {
	// The ID of the route attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-attachmentid
	//
	AttachmentId *string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// The CIDR block used for destination matches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-destinationcidr
	//
	DestinationCidr *string `field:"optional" json:"destinationCidr" yaml:"destinationCidr"`
	// The ID of the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-prefixlistid
	//
	PrefixListId *string `field:"optional" json:"prefixListId" yaml:"prefixListId"`
	// The ID of the resource for the route attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The resource type for the route attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The route origin. The following are the possible values:.
	//
	// - static
	// - propagated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-routeorigin
	//
	RouteOrigin *string `field:"optional" json:"routeOrigin" yaml:"routeOrigin"`
	// The state of the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-transitgatewayroutetableroute.html#cfn-ec2-networkinsightsanalysis-transitgatewayroutetableroute-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

