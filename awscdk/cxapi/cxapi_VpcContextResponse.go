package cxapi


// Properties of a discovered VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcContextResponse := &VpcContextResponse{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	IsolatedSubnetIds: []*string{
//   		jsii.String("isolatedSubnetIds"),
//   	},
//   	IsolatedSubnetNames: []*string{
//   		jsii.String("isolatedSubnetNames"),
//   	},
//   	IsolatedSubnetRouteTableIds: []*string{
//   		jsii.String("isolatedSubnetRouteTableIds"),
//   	},
//   	PrivateSubnetIds: []*string{
//   		jsii.String("privateSubnetIds"),
//   	},
//   	PrivateSubnetNames: []*string{
//   		jsii.String("privateSubnetNames"),
//   	},
//   	PrivateSubnetRouteTableIds: []*string{
//   		jsii.String("privateSubnetRouteTableIds"),
//   	},
//   	PublicSubnetIds: []*string{
//   		jsii.String("publicSubnetIds"),
//   	},
//   	PublicSubnetNames: []*string{
//   		jsii.String("publicSubnetNames"),
//   	},
//   	PublicSubnetRouteTableIds: []*string{
//   		jsii.String("publicSubnetRouteTableIds"),
//   	},
//   	SubnetGroups: []vpcSubnetGroup{
//   		&vpcSubnetGroup{
//   			Name: jsii.String("name"),
//   			Subnets: []vpcSubnet{
//   				&vpcSubnet{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					RouteTableId: jsii.String("routeTableId"),
//   					SubnetId: jsii.String("subnetId"),
//
//   					// the properties below are optional
//   					Cidr: jsii.String("cidr"),
//   				},
//   			},
//   			Type: awscdk.Cx_api.VpcSubnetGroupType_PUBLIC,
//   		},
//   	},
//   	VpcCidrBlock: jsii.String("vpcCidrBlock"),
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
// Experimental.
type VpcContextResponse struct {
	// AZs.
	// Experimental.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// VPC id.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// IDs of all isolated subnets.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	// Experimental.
	IsolatedSubnetIds *[]*string `field:"optional" json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// Name of isolated subnet groups.
	//
	// Element count: #(isolatedGroups).
	// Experimental.
	IsolatedSubnetNames *[]*string `field:"optional" json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// Route Table IDs of isolated subnet groups.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	// Experimental.
	IsolatedSubnetRouteTableIds *[]*string `field:"optional" json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// IDs of all private subnets.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	// Experimental.
	PrivateSubnetIds *[]*string `field:"optional" json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// Name of private subnet groups.
	//
	// Element count: #(privateGroups).
	// Experimental.
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// Route Table IDs of private subnet groups.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	// Experimental.
	PrivateSubnetRouteTableIds *[]*string `field:"optional" json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// IDs of all public subnets.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	// Experimental.
	PublicSubnetIds *[]*string `field:"optional" json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// Name of public subnet groups.
	//
	// Element count: #(publicGroups).
	// Experimental.
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// Route Table IDs of public subnet groups.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	// Experimental.
	PublicSubnetRouteTableIds *[]*string `field:"optional" json:"publicSubnetRouteTableIds" yaml:"publicSubnetRouteTableIds"`
	// The subnet groups discovered for the given VPC.
	//
	// Unlike the above properties, this will include asymmetric subnets,
	// if the VPC has any.
	// This property will only be populated if {@link VpcContextQuery.returnAsymmetricSubnets}
	// is true.
	// Experimental.
	SubnetGroups *[]*VpcSubnetGroup `field:"optional" json:"subnetGroups" yaml:"subnetGroups"`
	// VPC cidr.
	// Experimental.
	VpcCidrBlock *string `field:"optional" json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// The VPN gateway ID.
	// Experimental.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

