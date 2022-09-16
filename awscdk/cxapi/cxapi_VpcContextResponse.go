package cxapi


// Properties of a discovered VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcContextResponse := &vpcContextResponse{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	isolatedSubnetIds: []*string{
//   		jsii.String("isolatedSubnetIds"),
//   	},
//   	isolatedSubnetNames: []*string{
//   		jsii.String("isolatedSubnetNames"),
//   	},
//   	isolatedSubnetRouteTableIds: []*string{
//   		jsii.String("isolatedSubnetRouteTableIds"),
//   	},
//   	privateSubnetIds: []*string{
//   		jsii.String("privateSubnetIds"),
//   	},
//   	privateSubnetNames: []*string{
//   		jsii.String("privateSubnetNames"),
//   	},
//   	privateSubnetRouteTableIds: []*string{
//   		jsii.String("privateSubnetRouteTableIds"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("publicSubnetIds"),
//   	},
//   	publicSubnetNames: []*string{
//   		jsii.String("publicSubnetNames"),
//   	},
//   	publicSubnetRouteTableIds: []*string{
//   		jsii.String("publicSubnetRouteTableIds"),
//   	},
//   	region: jsii.String("region"),
//   	subnetGroups: []vpcSubnetGroup{
//   		&vpcSubnetGroup{
//   			name: jsii.String("name"),
//   			subnets: []vpcSubnet{
//   				&vpcSubnet{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					routeTableId: jsii.String("routeTableId"),
//   					subnetId: jsii.String("subnetId"),
//
//   					// the properties below are optional
//   					cidr: jsii.String("cidr"),
//   				},
//   			},
//   			type: awscdk.Cx_api.vpcSubnetGroupType_PUBLIC,
//   		},
//   	},
//   	vpcCidrBlock: jsii.String("vpcCidrBlock"),
//   	vpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
type VpcContextResponse struct {
	// AZs.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// VPC id.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// IDs of all isolated subnets.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	IsolatedSubnetIds *[]*string `field:"optional" json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// Name of isolated subnet groups.
	//
	// Element count: #(isolatedGroups).
	IsolatedSubnetNames *[]*string `field:"optional" json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// Route Table IDs of isolated subnet groups.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	IsolatedSubnetRouteTableIds *[]*string `field:"optional" json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// IDs of all private subnets.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	PrivateSubnetIds *[]*string `field:"optional" json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// Name of private subnet groups.
	//
	// Element count: #(privateGroups).
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// Route Table IDs of private subnet groups.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	PrivateSubnetRouteTableIds *[]*string `field:"optional" json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// IDs of all public subnets.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	PublicSubnetIds *[]*string `field:"optional" json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// Name of public subnet groups.
	//
	// Element count: #(publicGroups).
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// Route Table IDs of public subnet groups.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	PublicSubnetRouteTableIds *[]*string `field:"optional" json:"publicSubnetRouteTableIds" yaml:"publicSubnetRouteTableIds"`
	// The region in which the VPC is in.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The subnet groups discovered for the given VPC.
	//
	// Unlike the above properties, this will include asymmetric subnets,
	// if the VPC has any.
	// This property will only be populated if {@link VpcContextQuery.returnAsymmetricSubnets}
	// is true.
	SubnetGroups *[]*VpcSubnetGroup `field:"optional" json:"subnetGroups" yaml:"subnetGroups"`
	// VPC cidr.
	VpcCidrBlock *string `field:"optional" json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// The VPN gateway ID.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

