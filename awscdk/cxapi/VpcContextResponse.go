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
//   	OwnerAccountId: jsii.String("ownerAccountId"),
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
//   	Region: jsii.String("region"),
//   	SubnetGroups: []VpcSubnetGroup{
//   		&VpcSubnetGroup{
//   			Name: jsii.String("name"),
//   			Subnets: []VpcSubnet{
//   				&VpcSubnet{
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
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type VpcContextResponse struct {
	// AZs.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// VPC id.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// IDs of all isolated subnets.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	IsolatedSubnetIds *[]*string `field:"optional" json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// Name of isolated subnet groups.
	//
	// Element count: #(isolatedGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	IsolatedSubnetNames *[]*string `field:"optional" json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// Route Table IDs of isolated subnet groups.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	IsolatedSubnetRouteTableIds *[]*string `field:"optional" json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// The ID of the AWS account that owns the VPC.
	// Default: the account id of the parent stack.
	//
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	OwnerAccountId *string `field:"optional" json:"ownerAccountId" yaml:"ownerAccountId"`
	// IDs of all private subnets.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	PrivateSubnetIds *[]*string `field:"optional" json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// Name of private subnet groups.
	//
	// Element count: #(privateGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	PrivateSubnetNames *[]*string `field:"optional" json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// Route Table IDs of private subnet groups.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	PrivateSubnetRouteTableIds *[]*string `field:"optional" json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// IDs of all public subnets.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	PublicSubnetIds *[]*string `field:"optional" json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// Name of public subnet groups.
	//
	// Element count: #(publicGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	PublicSubnetNames *[]*string `field:"optional" json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// Route Table IDs of public subnet groups.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	PublicSubnetRouteTableIds *[]*string `field:"optional" json:"publicSubnetRouteTableIds" yaml:"publicSubnetRouteTableIds"`
	// The region in which the VPC is in.
	// Default: - Region of the parent stack.
	//
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The subnet groups discovered for the given VPC.
	//
	// Unlike the above properties, this will include asymmetric subnets,
	// if the VPC has any.
	// This property will only be populated if `VpcContextQuery.returnAsymmetricSubnets`
	// is true.
	// Default: - no subnet groups will be returned unless `VpcContextQuery.returnAsymmetricSubnets` is true
	//
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	SubnetGroups *[]*VpcSubnetGroup `field:"optional" json:"subnetGroups" yaml:"subnetGroups"`
	// VPC cidr.
	// Default: - CIDR information not available.
	//
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpcCidrBlock *string `field:"optional" json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// The VPN gateway ID.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

