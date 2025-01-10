// The CDK construct library for VPC V2
package awsec2alpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ec2-alpha.AddressFamily",
		reflect.TypeOf((*AddressFamily)(nil)).Elem(),
		map[string]interface{}{
			"IP_V4": AddressFamily_IP_V4,
			"IP_V6": AddressFamily_IP_V6,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ec2-alpha.AwsServiceName",
		reflect.TypeOf((*AwsServiceName)(nil)).Elem(),
		map[string]interface{}{
			"EC2": AwsServiceName_EC2,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.EgressOnlyInternetGateway",
		reflect.TypeOf((*EgressOnlyInternetGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routerTargetId", GoGetter: "RouterTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "routerType", GoGetter: "RouterType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EgressOnlyInternetGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.EgressOnlyInternetGatewayOptions",
		reflect.TypeOf((*EgressOnlyInternetGatewayOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.EgressOnlyInternetGatewayProps",
		reflect.TypeOf((*EgressOnlyInternetGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IIpAddresses",
		reflect.TypeOf((*IIpAddresses)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allocateVpcCidr", GoMethod: "AllocateVpcCidr"},
		},
		func() interface{} {
			return &jsiiProxy_IIpAddresses{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IIpamPool",
		reflect.TypeOf((*IIpamPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipamCidrs", GoGetter: "IpamCidrs"},
			_jsii_.MemberProperty{JsiiProperty: "ipamIpv4Cidrs", GoGetter: "IpamIpv4Cidrs"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolId", GoGetter: "IpamPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "provisionCidr", GoMethod: "ProvisionCidr"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamPool{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IIpamScopeBase",
		reflect.TypeOf((*IIpamScopeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPool", GoMethod: "AddPool"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "scopeId", GoGetter: "ScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "scopeType", GoGetter: "ScopeType"},
		},
		func() interface{} {
			return &jsiiProxy_IIpamScopeBase{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IRouteTarget",
		reflect.TypeOf((*IRouteTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "routerTargetId", GoGetter: "RouterTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "routerType", GoGetter: "RouterType"},
		},
		func() interface{} {
			j := jsiiProxy_IRouteTarget{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIDependable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IRouteV2",
		reflect.TypeOf((*IRouteV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
		},
		func() interface{} {
			j := jsiiProxy_IRouteV2{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.ISubnetV2",
		reflect.TypeOf((*ISubnetV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetType", GoGetter: "SubnetType"},
		},
		func() interface{} {
			j := jsiiProxy_ISubnetV2{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2ISubnet)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IVPCCidrBlock",
		reflect.TypeOf((*IVPCCidrBlock)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amazonProvidedIpv6CidrBlock", GoGetter: "AmazonProvidedIpv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolId", GoGetter: "Ipv4IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolId", GoGetter: "Ipv6IpamPoolId"},
		},
		func() interface{} {
			return &jsiiProxy_IVPCCidrBlock{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-ec2-alpha.IVpcV2",
		reflect.TypeOf((*IVpcV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addEgressOnlyInternetGateway", GoMethod: "AddEgressOnlyInternetGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInternetGateway", GoMethod: "AddInternetGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addNatGateway", GoMethod: "AddNatGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberMethod{JsiiMethod: "createAcceptorVpcRole", GoMethod: "CreateAcceptorVpcRole"},
			_jsii_.MemberMethod{JsiiMethod: "createPeeringConnection", GoMethod: "CreatePeeringConnection"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGatewayV2", GoMethod: "EnableVpnGatewayV2"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamProvisionedCidrs", GoGetter: "Ipv4IpamProvisionedCidrs"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ownerAccountId", GoGetter: "OwnerAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryCidrBlock", GoGetter: "SecondaryCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcName", GoGetter: "VpcName"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcV2{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IVpc)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.InternetGateway",
		reflect.TypeOf((*InternetGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routerTargetId", GoGetter: "RouterTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "routerType", GoGetter: "RouterType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_InternetGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.InternetGatewayOptions",
		reflect.TypeOf((*InternetGatewayOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.InternetGatewayProps",
		reflect.TypeOf((*InternetGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		reflect.TypeOf((*IpAddresses)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IpAddresses{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.IpCidr",
		reflect.TypeOf((*IpCidr)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
		},
		func() interface{} {
			return &jsiiProxy_IpCidr{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.Ipam",
		reflect.TypeOf((*Ipam)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addScope", GoMethod: "AddScope"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipamId", GoGetter: "IpamId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamName", GoGetter: "IpamName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operatingRegions", GoGetter: "OperatingRegions"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateScope", GoGetter: "PrivateScope"},
			_jsii_.MemberProperty{JsiiProperty: "publicScope", GoGetter: "PublicScope"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ipam{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.IpamOptions",
		reflect.TypeOf((*IpamOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.IpamPoolCidrProvisioningOptions",
		reflect.TypeOf((*IpamPoolCidrProvisioningOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ec2-alpha.IpamPoolPublicIpSource",
		reflect.TypeOf((*IpamPoolPublicIpSource)(nil)).Elem(),
		map[string]interface{}{
			"BYOIP": IpamPoolPublicIpSource_BYOIP,
			"AMAZON": IpamPoolPublicIpSource_AMAZON,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.IpamProps",
		reflect.TypeOf((*IpamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.IpamScopeOptions",
		reflect.TypeOf((*IpamScopeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ec2-alpha.IpamScopeType",
		reflect.TypeOf((*IpamScopeType)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": IpamScopeType_DEFAULT,
			"CUSTOM": IpamScopeType_CUSTOM,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-ec2-alpha.NatConnectivityType",
		reflect.TypeOf((*NatConnectivityType)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC": NatConnectivityType_PUBLIC,
			"PRIVATE": NatConnectivityType_PRIVATE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.NatGateway",
		reflect.TypeOf((*NatGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connectivityType", GoGetter: "ConnectivityType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxDrainDuration", GoGetter: "MaxDrainDuration"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayId", GoGetter: "NatGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routerTargetId", GoGetter: "RouterTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "routerType", GoGetter: "RouterType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NatGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.NatGatewayOptions",
		reflect.TypeOf((*NatGatewayOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.NatGatewayProps",
		reflect.TypeOf((*NatGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.PoolOptions",
		reflect.TypeOf((*PoolOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.Route",
		reflect.TypeOf((*Route)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetRouterType", GoGetter: "TargetRouterType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Route{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.RouteProps",
		reflect.TypeOf((*RouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.RouteTable",
		reflect.TypeOf((*RouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RouteTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IRouteTable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.RouteTableProps",
		reflect.TypeOf((*RouteTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.RouteTargetProps",
		reflect.TypeOf((*RouteTargetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.RouteTargetType",
		reflect.TypeOf((*RouteTargetType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "gateway", GoGetter: "Gateway"},
		},
		func() interface{} {
			return &jsiiProxy_RouteTargetType{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.SecondaryAddressProps",
		reflect.TypeOf((*SecondaryAddressProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.SubnetV2",
		reflect.TypeOf((*SubnetV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetType", GoGetter: "SubnetType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SubnetV2{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubnetV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.SubnetV2Attributes",
		reflect.TypeOf((*SubnetV2Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.SubnetV2Props",
		reflect.TypeOf((*SubnetV2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VPCCidrBlockattributes",
		reflect.TypeOf((*VPCCidrBlockattributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.VPCPeeringConnection",
		reflect.TypeOf((*VPCPeeringConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routerTargetId", GoGetter: "RouterTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "routerType", GoGetter: "RouterType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VPCPeeringConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VPCPeeringConnectionOptions",
		reflect.TypeOf((*VPCPeeringConnectionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VPCPeeringConnectionProps",
		reflect.TypeOf((*VPCPeeringConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.VPNGatewayV2",
		reflect.TypeOf((*VPNGatewayV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "routerTargetId", GoGetter: "RouterTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "routerType", GoGetter: "RouterType"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_VPNGatewayV2{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VPNGatewayV2Options",
		reflect.TypeOf((*VPNGatewayV2Options)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VPNGatewayV2Props",
		reflect.TypeOf((*VPNGatewayV2Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VpcCidrOptions",
		reflect.TypeOf((*VpcCidrOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.VpcV2",
		reflect.TypeOf((*VpcV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addEgressOnlyInternetGateway", GoMethod: "AddEgressOnlyInternetGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInternetGateway", GoMethod: "AddInternetGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addNatGateway", GoMethod: "AddNatGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberMethod{JsiiMethod: "createAcceptorVpcRole", GoMethod: "CreateAcceptorVpcRole"},
			_jsii_.MemberMethod{JsiiMethod: "createPeeringConnection", GoMethod: "CreatePeeringConnection"},
			_jsii_.MemberProperty{JsiiProperty: "dnsHostnamesEnabled", GoGetter: "DnsHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupportEnabled", GoGetter: "DnsSupportEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGatewayV2", GoMethod: "EnableVpnGatewayV2"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "incompleteSubnetDefinition", GoGetter: "IncompleteSubnetDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddresses", GoGetter: "IpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamProvisionedCidrs", GoGetter: "Ipv4IpamProvisionedCidrs"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlocks", GoGetter: "Ipv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ownerAccountId", GoGetter: "OwnerAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryCidrBlock", GoGetter: "SecondaryCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnetObjects", GoMethod: "SelectSubnetObjects"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useIpv6", GoGetter: "UseIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcName", GoGetter: "VpcName"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_VpcV2{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VpcV2Base)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VpcV2Attributes",
		reflect.TypeOf((*VpcV2Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-ec2-alpha.VpcV2Base",
		reflect.TypeOf((*VpcV2Base)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addEgressOnlyInternetGateway", GoMethod: "AddEgressOnlyInternetGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInternetGateway", GoMethod: "AddInternetGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addNatGateway", GoMethod: "AddNatGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberMethod{JsiiMethod: "createAcceptorVpcRole", GoMethod: "CreateAcceptorVpcRole"},
			_jsii_.MemberMethod{JsiiMethod: "createPeeringConnection", GoMethod: "CreatePeeringConnection"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGatewayV2", GoMethod: "EnableVpnGatewayV2"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "incompleteSubnetDefinition", GoGetter: "IncompleteSubnetDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamProvisionedCidrs", GoGetter: "Ipv4IpamProvisionedCidrs"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ownerAccountId", GoGetter: "OwnerAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryCidrBlock", GoGetter: "SecondaryCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnetObjects", GoMethod: "SelectSubnetObjects"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcName", GoGetter: "VpcName"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_VpcV2Base{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-ec2-alpha.VpcV2Props",
		reflect.TypeOf((*VpcV2Props)(nil)).Elem(),
	)
}
