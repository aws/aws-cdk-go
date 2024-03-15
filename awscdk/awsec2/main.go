package awsec2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AclCidr",
		reflect.TypeOf((*AclCidr)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toCidrConfig", GoMethod: "ToCidrConfig"},
		},
		func() interface{} {
			return &jsiiProxy_AclCidr{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AclCidrConfig",
		reflect.TypeOf((*AclCidrConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AclIcmp",
		reflect.TypeOf((*AclIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AclPortRange",
		reflect.TypeOf((*AclPortRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AclTraffic",
		reflect.TypeOf((*AclTraffic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toTrafficConfig", GoMethod: "ToTrafficConfig"},
		},
		func() interface{} {
			return &jsiiProxy_AclTraffic{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AclTrafficConfig",
		reflect.TypeOf((*AclTrafficConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
		map[string]interface{}{
			"ALLOW": Action_ALLOW,
			"DENY": Action_DENY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AddRouteOptions",
		reflect.TypeOf((*AddRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AddressFamily",
		reflect.TypeOf((*AddressFamily)(nil)).Elem(),
		map[string]interface{}{
			"IP_V4": AddressFamily_IP_V4,
			"IP_V6": AddressFamily_IP_V6,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AllocateCidrRequest",
		reflect.TypeOf((*AllocateCidrRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AllocateIpv6CidrRequest",
		reflect.TypeOf((*AllocateIpv6CidrRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AllocateVpcIpv6CidrRequest",
		reflect.TypeOf((*AllocateVpcIpv6CidrRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AllocatedSubnet",
		reflect.TypeOf((*AllocatedSubnet)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022ImageSsmParameter",
		reflect.TypeOf((*AmazonLinux2022ImageSsmParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonLinux2022ImageSsmParameter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AmazonLinuxImageSsmParameterBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022ImageSsmParameterProps",
		reflect.TypeOf((*AmazonLinux2022ImageSsmParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022Kernel",
		reflect.TypeOf((*AmazonLinux2022Kernel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_AmazonLinux2022Kernel{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023ImageSsmParameter",
		reflect.TypeOf((*AmazonLinux2023ImageSsmParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonLinux2023ImageSsmParameter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AmazonLinuxImageSsmParameterBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023ImageSsmParameterProps",
		reflect.TypeOf((*AmazonLinux2023ImageSsmParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023Kernel",
		reflect.TypeOf((*AmazonLinux2023Kernel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_AmazonLinux2023Kernel{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinux2ImageSsmParameter",
		reflect.TypeOf((*AmazonLinux2ImageSsmParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonLinux2ImageSsmParameter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AmazonLinuxImageSsmParameterBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinux2ImageSsmParameterProps",
		reflect.TypeOf((*AmazonLinux2ImageSsmParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinux2Kernel",
		reflect.TypeOf((*AmazonLinux2Kernel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_AmazonLinux2Kernel{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AmazonLinuxCpuType",
		reflect.TypeOf((*AmazonLinuxCpuType)(nil)).Elem(),
		map[string]interface{}{
			"ARM_64": AmazonLinuxCpuType_ARM_64,
			"X86_64": AmazonLinuxCpuType_X86_64,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AmazonLinuxEdition",
		reflect.TypeOf((*AmazonLinuxEdition)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": AmazonLinuxEdition_STANDARD,
			"MINIMAL": AmazonLinuxEdition_MINIMAL,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AmazonLinuxGeneration",
		reflect.TypeOf((*AmazonLinuxGeneration)(nil)).Elem(),
		map[string]interface{}{
			"AMAZON_LINUX": AmazonLinuxGeneration_AMAZON_LINUX,
			"AMAZON_LINUX_2": AmazonLinuxGeneration_AMAZON_LINUX_2,
			"AMAZON_LINUX_2022": AmazonLinuxGeneration_AMAZON_LINUX_2022,
			"AMAZON_LINUX_2023": AmazonLinuxGeneration_AMAZON_LINUX_2023,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImage",
		reflect.TypeOf((*AmazonLinuxImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonLinuxImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GenericSSMParameterImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImageProps",
		reflect.TypeOf((*AmazonLinuxImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImageSsmParameterBase",
		reflect.TypeOf((*AmazonLinuxImageSsmParameterBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonLinuxImageSsmParameterBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImageSsmParameterBaseOptions",
		reflect.TypeOf((*AmazonLinuxImageSsmParameterBaseOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImageSsmParameterBaseProps",
		reflect.TypeOf((*AmazonLinuxImageSsmParameterBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImageSsmParameterCommonOptions",
		reflect.TypeOf((*AmazonLinuxImageSsmParameterCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AmazonLinuxKernel",
		reflect.TypeOf((*AmazonLinuxKernel)(nil)).Elem(),
		map[string]interface{}{
			"KERNEL5_X": AmazonLinuxKernel_KERNEL5_X,
			"KERNEL6_1": AmazonLinuxKernel_KERNEL6_1,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AmazonLinuxStorage",
		reflect.TypeOf((*AmazonLinuxStorage)(nil)).Elem(),
		map[string]interface{}{
			"EBS": AmazonLinuxStorage_EBS,
			"S3": AmazonLinuxStorage_S3,
			"GENERAL_PURPOSE": AmazonLinuxStorage_GENERAL_PURPOSE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.AmazonLinuxVirt",
		reflect.TypeOf((*AmazonLinuxVirt)(nil)).Elem(),
		map[string]interface{}{
			"HVM": AmazonLinuxVirt_HVM,
			"PV": AmazonLinuxVirt_PV,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ApplyCloudFormationInitOptions",
		reflect.TypeOf((*ApplyCloudFormationInitOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AttachInitOptions",
		reflect.TypeOf((*AttachInitOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.AwsIpamProps",
		reflect.TypeOf((*AwsIpamProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.BastionHostLinux",
		reflect.TypeOf((*BastionHostLinux)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allowSshAccessFrom", GoMethod: "AllowSshAccessFrom"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "instanceAvailabilityZone", GoGetter: "InstanceAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateDnsName", GoGetter: "InstancePrivateDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateIp", GoGetter: "InstancePrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicDnsName", GoGetter: "InstancePublicDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicIp", GoGetter: "InstancePublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BastionHostLinux{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInstance)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.BastionHostLinuxProps",
		reflect.TypeOf((*BastionHostLinuxProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.BlockDevice",
		reflect.TypeOf((*BlockDevice)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.BlockDeviceVolume",
		reflect.TypeOf((*BlockDeviceVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ebsDevice", GoGetter: "EbsDevice"},
			_jsii_.MemberProperty{JsiiProperty: "virtualName", GoGetter: "VirtualName"},
		},
		func() interface{} {
			return &jsiiProxy_BlockDeviceVolume{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservation",
		reflect.TypeOf((*CfnCapacityReservation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailabilityZone", GoGetter: "AttrAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailableInstanceCount", GoGetter: "AttrAvailableInstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInstanceType", GoGetter: "AttrInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrTenancy", GoGetter: "AttrTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalInstanceCount", GoGetter: "AttrTotalInstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "endDate", GoGetter: "EndDate"},
			_jsii_.MemberProperty{JsiiProperty: "endDateType", GoGetter: "EndDateType"},
			_jsii_.MemberProperty{JsiiProperty: "ephemeralStorage", GoGetter: "EphemeralStorage"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMatchCriteria", GoGetter: "InstanceMatchCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "instancePlatform", GoGetter: "InstancePlatform"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "outPostArn", GoGetter: "OutPostArn"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroupArn", GoGetter: "PlacementGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "tenancy", GoGetter: "Tenancy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapacityReservation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservation.TagSpecificationProperty",
		reflect.TypeOf((*CfnCapacityReservation_TagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservationFleet",
		reflect.TypeOf((*CfnCapacityReservationFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allocationStrategy", GoGetter: "AllocationStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCapacityReservationFleetId", GoGetter: "AttrCapacityReservationFleetId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endDate", GoGetter: "EndDate"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMatchCriteria", GoGetter: "InstanceMatchCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeSpecifications", GoGetter: "InstanceTypeSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noRemoveEndDate", GoGetter: "NoRemoveEndDate"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberProperty{JsiiProperty: "removeEndDate", GoGetter: "RemoveEndDate"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "tenancy", GoGetter: "Tenancy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalTargetCapacity", GoGetter: "TotalTargetCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCapacityReservationFleet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservationFleet.InstanceTypeSpecificationProperty",
		reflect.TypeOf((*CfnCapacityReservationFleet_InstanceTypeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservationFleet.TagSpecificationProperty",
		reflect.TypeOf((*CfnCapacityReservationFleet_TagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservationFleetProps",
		reflect.TypeOf((*CfnCapacityReservationFleetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCapacityReservationProps",
		reflect.TypeOf((*CfnCapacityReservationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnCarrierGateway",
		reflect.TypeOf((*CfnCarrierGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCarrierGatewayId", GoGetter: "AttrCarrierGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwnerId", GoGetter: "AttrOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCarrierGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCarrierGatewayProps",
		reflect.TypeOf((*CfnCarrierGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnClientVpnAuthorizationRule",
		reflect.TypeOf((*CfnClientVpnAuthorizationRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessGroupId", GoGetter: "AccessGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "authorizeAllGroups", GoGetter: "AuthorizeAllGroups"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientVpnEndpointId", GoGetter: "ClientVpnEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworkCidr", GoGetter: "TargetNetworkCidr"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClientVpnAuthorizationRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnAuthorizationRuleProps",
		reflect.TypeOf((*CfnClientVpnAuthorizationRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint",
		reflect.TypeOf((*CfnClientVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationOptions", GoGetter: "AuthenticationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientCidrBlock", GoGetter: "ClientCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "clientConnectOptions", GoGetter: "ClientConnectOptions"},
			_jsii_.MemberProperty{JsiiProperty: "clientLoginBannerOptions", GoGetter: "ClientLoginBannerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "connectionLogOptions", GoGetter: "ConnectionLogOptions"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServers", GoGetter: "DnsServers"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePortal", GoGetter: "SelfServicePortal"},
			_jsii_.MemberProperty{JsiiProperty: "serverCertificateArn", GoGetter: "ServerCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTimeoutHours", GoGetter: "SessionTimeoutHours"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "splitTunnel", GoGetter: "SplitTunnel"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transportProtocol", GoGetter: "TransportProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnPort", GoGetter: "VpnPort"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClientVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.CertificateAuthenticationRequestProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_CertificateAuthenticationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.ClientAuthenticationRequestProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_ClientAuthenticationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.ClientConnectOptionsProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_ClientConnectOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.ClientLoginBannerOptionsProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_ClientLoginBannerOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.ConnectionLogOptionsProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_ConnectionLogOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.DirectoryServiceAuthenticationRequestProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_DirectoryServiceAuthenticationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.FederatedAuthenticationRequestProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_FederatedAuthenticationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpoint.TagSpecificationProperty",
		reflect.TypeOf((*CfnClientVpnEndpoint_TagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnEndpointProps",
		reflect.TypeOf((*CfnClientVpnEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnClientVpnRoute",
		reflect.TypeOf((*CfnClientVpnRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientVpnEndpointId", GoGetter: "ClientVpnEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetVpcSubnetId", GoGetter: "TargetVpcSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClientVpnRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnRouteProps",
		reflect.TypeOf((*CfnClientVpnRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnClientVpnTargetNetworkAssociation",
		reflect.TypeOf((*CfnClientVpnTargetNetworkAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientVpnEndpointId", GoGetter: "ClientVpnEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClientVpnTargetNetworkAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnClientVpnTargetNetworkAssociationProps",
		reflect.TypeOf((*CfnClientVpnTargetNetworkAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnCustomerGateway",
		reflect.TypeOf((*CfnCustomerGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCustomerGatewayId", GoGetter: "AttrCustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "bgpAsn", GoGetter: "BgpAsn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceName", GoGetter: "DeviceName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomerGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnCustomerGatewayProps",
		reflect.TypeOf((*CfnCustomerGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnDHCPOptions",
		reflect.TypeOf((*CfnDHCPOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDhcpOptionsId", GoGetter: "AttrDhcpOptionsId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameServers", GoGetter: "DomainNameServers"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "netbiosNameServers", GoGetter: "NetbiosNameServers"},
			_jsii_.MemberProperty{JsiiProperty: "netbiosNodeType", GoGetter: "NetbiosNodeType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ntpServers", GoGetter: "NtpServers"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDHCPOptions{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnDHCPOptionsProps",
		reflect.TypeOf((*CfnDHCPOptionsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet",
		reflect.TypeOf((*CfnEC2Fleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrFleetId", GoGetter: "AttrFleetId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excessCapacityTerminationPolicy", GoGetter: "ExcessCapacityTerminationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateConfigs", GoGetter: "LaunchTemplateConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandOptions", GoGetter: "OnDemandOptions"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "replaceUnhealthyInstances", GoGetter: "ReplaceUnhealthyInstances"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spotOptions", GoGetter: "SpotOptions"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "targetCapacitySpecification", GoGetter: "TargetCapacitySpecification"},
			_jsii_.MemberProperty{JsiiProperty: "terminateInstancesWithExpiration", GoGetter: "TerminateInstancesWithExpiration"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "validFrom", GoGetter: "ValidFrom"},
			_jsii_.MemberProperty{JsiiProperty: "validUntil", GoGetter: "ValidUntil"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEC2Fleet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.AcceleratorCountRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_AcceleratorCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.AcceleratorTotalMemoryMiBRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_AcceleratorTotalMemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.BaselineEbsBandwidthMbpsRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_BaselineEbsBandwidthMbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.CapacityRebalanceProperty",
		reflect.TypeOf((*CfnEC2Fleet_CapacityRebalanceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.CapacityReservationOptionsRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_CapacityReservationOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.FleetLaunchTemplateConfigRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_FleetLaunchTemplateConfigRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.FleetLaunchTemplateOverridesRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_FleetLaunchTemplateOverridesRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.FleetLaunchTemplateSpecificationRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_FleetLaunchTemplateSpecificationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.InstanceRequirementsRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_InstanceRequirementsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.MaintenanceStrategiesProperty",
		reflect.TypeOf((*CfnEC2Fleet_MaintenanceStrategiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.MemoryGiBPerVCpuRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_MemoryGiBPerVCpuRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.MemoryMiBRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_MemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.NetworkBandwidthGbpsRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_NetworkBandwidthGbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.NetworkInterfaceCountRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_NetworkInterfaceCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.OnDemandOptionsRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_OnDemandOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.PlacementProperty",
		reflect.TypeOf((*CfnEC2Fleet_PlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.SpotOptionsRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_SpotOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.TagSpecificationProperty",
		reflect.TypeOf((*CfnEC2Fleet_TagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.TargetCapacitySpecificationRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_TargetCapacitySpecificationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.TotalLocalStorageGBRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_TotalLocalStorageGBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2Fleet.VCpuCountRangeRequestProperty",
		reflect.TypeOf((*CfnEC2Fleet_VCpuCountRangeRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEC2FleetProps",
		reflect.TypeOf((*CfnEC2FleetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnEIP",
		reflect.TypeOf((*CfnEIP)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAllocationId", GoGetter: "AttrAllocationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublicIp", GoGetter: "AttrPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkBorderGroup", GoGetter: "NetworkBorderGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpv4Pool", GoGetter: "PublicIpv4Pool"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transferAddress", GoGetter: "TransferAddress"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEIP{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnEIPAssociation",
		reflect.TypeOf((*CfnEIPAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allocationId", GoGetter: "AllocationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eip", GoGetter: "Eip"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEIPAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEIPAssociationProps",
		reflect.TypeOf((*CfnEIPAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEIPProps",
		reflect.TypeOf((*CfnEIPProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnEgressOnlyInternetGateway",
		reflect.TypeOf((*CfnEgressOnlyInternetGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEgressOnlyInternetGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEgressOnlyInternetGatewayProps",
		reflect.TypeOf((*CfnEgressOnlyInternetGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnEnclaveCertificateIamRoleAssociation",
		reflect.TypeOf((*CfnEnclaveCertificateIamRoleAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCertificateS3BucketName", GoGetter: "AttrCertificateS3BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "attrCertificateS3ObjectKey", GoGetter: "AttrCertificateS3ObjectKey"},
			_jsii_.MemberProperty{JsiiProperty: "attrEncryptionKmsKeyId", GoGetter: "AttrEncryptionKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnclaveCertificateIamRoleAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnEnclaveCertificateIamRoleAssociationProps",
		reflect.TypeOf((*CfnEnclaveCertificateIamRoleAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
		reflect.TypeOf((*CfnFlowLog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deliverCrossAccountRole", GoGetter: "DeliverCrossAccountRole"},
			_jsii_.MemberProperty{JsiiProperty: "deliverLogsPermissionArn", GoGetter: "DeliverLogsPermissionArn"},
			_jsii_.MemberProperty{JsiiProperty: "destinationOptions", GoGetter: "DestinationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logDestination", GoGetter: "LogDestination"},
			_jsii_.MemberProperty{JsiiProperty: "logDestinationType", GoGetter: "LogDestinationType"},
			_jsii_.MemberProperty{JsiiProperty: "logFormat", GoGetter: "LogFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxAggregationInterval", GoGetter: "MaxAggregationInterval"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resourceId", GoGetter: "ResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficType", GoGetter: "TrafficType"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowLog{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnFlowLog.DestinationOptionsProperty",
		reflect.TypeOf((*CfnFlowLog_DestinationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnFlowLogProps",
		reflect.TypeOf((*CfnFlowLogProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnGatewayRouteTableAssociation",
		reflect.TypeOf((*CfnGatewayRouteTableAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAssociationId", GoGetter: "AttrAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayRouteTableAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnGatewayRouteTableAssociationProps",
		reflect.TypeOf((*CfnGatewayRouteTableAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnHost",
		reflect.TypeOf((*CfnHost)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assetId", GoGetter: "AssetId"},
			_jsii_.MemberProperty{JsiiProperty: "attrHostId", GoGetter: "AttrHostId"},
			_jsii_.MemberProperty{JsiiProperty: "autoPlacement", GoGetter: "AutoPlacement"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "hostMaintenance", GoGetter: "HostMaintenance"},
			_jsii_.MemberProperty{JsiiProperty: "hostRecovery", GoGetter: "HostRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceFamily", GoGetter: "InstanceFamily"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "outpostArn", GoGetter: "OutpostArn"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHost{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnHostProps",
		reflect.TypeOf((*CfnHostProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAM",
		reflect.TypeOf((*CfnIPAM)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultResourceDiscoveryAssociationId", GoGetter: "AttrDefaultResourceDiscoveryAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultResourceDiscoveryId", GoGetter: "AttrDefaultResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamId", GoGetter: "AttrIpamId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrivateDefaultScopeId", GoGetter: "AttrPrivateDefaultScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublicDefaultScopeId", GoGetter: "AttrPublicDefaultScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceDiscoveryAssociationCount", GoGetter: "AttrResourceDiscoveryAssociationCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrScopeCount", GoGetter: "AttrScopeCount"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "operatingRegions", GoGetter: "OperatingRegions"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberProperty{JsiiProperty: "tier", GoGetter: "Tier"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAM{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAM.IpamOperatingRegionProperty",
		reflect.TypeOf((*CfnIPAM_IpamOperatingRegionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAMAllocation",
		reflect.TypeOf((*CfnIPAMAllocation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamPoolAllocationId", GoGetter: "AttrIpamPoolAllocationId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolId", GoGetter: "IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "netmaskLength", GoGetter: "NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAMAllocation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMAllocationProps",
		reflect.TypeOf((*CfnIPAMAllocationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool",
		reflect.TypeOf((*CfnIPAMPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "addressFamily", GoGetter: "AddressFamily"},
			_jsii_.MemberProperty{JsiiProperty: "allocationDefaultNetmaskLength", GoGetter: "AllocationDefaultNetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "allocationMaxNetmaskLength", GoGetter: "AllocationMaxNetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "allocationMinNetmaskLength", GoGetter: "AllocationMinNetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "allocationResourceTags", GoGetter: "AllocationResourceTags"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamArn", GoGetter: "AttrIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamPoolId", GoGetter: "AttrIpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamScopeArn", GoGetter: "AttrIpamScopeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamScopeType", GoGetter: "AttrIpamScopeType"},
			_jsii_.MemberProperty{JsiiProperty: "attrPoolDepth", GoGetter: "AttrPoolDepth"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateMessage", GoGetter: "AttrStateMessage"},
			_jsii_.MemberProperty{JsiiProperty: "autoImport", GoGetter: "AutoImport"},
			_jsii_.MemberProperty{JsiiProperty: "awsService", GoGetter: "AwsService"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipamScopeId", GoGetter: "IpamScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "locale", GoGetter: "Locale"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedCidrs", GoGetter: "ProvisionedCidrs"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpSource", GoGetter: "PublicIpSource"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAdvertisable", GoGetter: "PubliclyAdvertisable"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIpamPoolId", GoGetter: "SourceIpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceResource", GoGetter: "SourceResource"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAMPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool.ProvisionedCidrProperty",
		reflect.TypeOf((*CfnIPAMPool_ProvisionedCidrProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMPool.SourceResourceProperty",
		reflect.TypeOf((*CfnIPAMPool_SourceResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAMPoolCidr",
		reflect.TypeOf((*CfnIPAMPoolCidr)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamPoolCidrId", GoGetter: "AttrIpamPoolCidrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipamPoolId", GoGetter: "IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "netmaskLength", GoGetter: "NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAMPoolCidr{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMPoolCidrProps",
		reflect.TypeOf((*CfnIPAMPoolCidrProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMPoolProps",
		reflect.TypeOf((*CfnIPAMPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMProps",
		reflect.TypeOf((*CfnIPAMProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAMResourceDiscovery",
		reflect.TypeOf((*CfnIPAMResourceDiscovery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamResourceDiscoveryArn", GoGetter: "AttrIpamResourceDiscoveryArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamResourceDiscoveryId", GoGetter: "AttrIpamResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamResourceDiscoveryRegion", GoGetter: "AttrIpamResourceDiscoveryRegion"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefault", GoGetter: "AttrIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwnerId", GoGetter: "AttrOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "operatingRegions", GoGetter: "OperatingRegions"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAMResourceDiscovery{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMResourceDiscovery.IpamOperatingRegionProperty",
		reflect.TypeOf((*CfnIPAMResourceDiscovery_IpamOperatingRegionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAMResourceDiscoveryAssociation",
		reflect.TypeOf((*CfnIPAMResourceDiscoveryAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamArn", GoGetter: "AttrIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamRegion", GoGetter: "AttrIpamRegion"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamResourceDiscoveryAssociationArn", GoGetter: "AttrIpamResourceDiscoveryAssociationArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamResourceDiscoveryAssociationId", GoGetter: "AttrIpamResourceDiscoveryAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefault", GoGetter: "AttrIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwnerId", GoGetter: "AttrOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceDiscoveryStatus", GoGetter: "AttrResourceDiscoveryStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipamId", GoGetter: "IpamId"},
			_jsii_.MemberProperty{JsiiProperty: "ipamResourceDiscoveryId", GoGetter: "IpamResourceDiscoveryId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAMResourceDiscoveryAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMResourceDiscoveryAssociationProps",
		reflect.TypeOf((*CfnIPAMResourceDiscoveryAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMResourceDiscoveryProps",
		reflect.TypeOf((*CfnIPAMResourceDiscoveryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnIPAMScope",
		reflect.TypeOf((*CfnIPAMScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamArn", GoGetter: "AttrIpamArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamScopeId", GoGetter: "AttrIpamScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpamScopeType", GoGetter: "AttrIpamScopeType"},
			_jsii_.MemberProperty{JsiiProperty: "attrIsDefault", GoGetter: "AttrIsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "attrPoolCount", GoGetter: "AttrPoolCount"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipamId", GoGetter: "IpamId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIPAMScope{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnIPAMScopeProps",
		reflect.TypeOf((*CfnIPAMScopeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnInstance",
		reflect.TypeOf((*CfnInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "additionalInfo", GoGetter: "AdditionalInfo"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "affinity", GoGetter: "Affinity"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailabilityZone", GoGetter: "AttrAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrivateDnsName", GoGetter: "AttrPrivateDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrivateIp", GoGetter: "AttrPrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublicDnsName", GoGetter: "AttrPublicDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "attrPublicIp", GoGetter: "AttrPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappings", GoGetter: "BlockDeviceMappings"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cpuOptions", GoGetter: "CpuOptions"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "creditSpecification", GoGetter: "CreditSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "disableApiTermination", GoGetter: "DisableApiTermination"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "elasticGpuSpecifications", GoGetter: "ElasticGpuSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "elasticInferenceAccelerators", GoGetter: "ElasticInferenceAccelerators"},
			_jsii_.MemberProperty{JsiiProperty: "enclaveOptions", GoGetter: "EnclaveOptions"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "hibernationOptions", GoGetter: "HibernationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "hostId", GoGetter: "HostId"},
			_jsii_.MemberProperty{JsiiProperty: "hostResourceGroupArn", GoGetter: "HostResourceGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "iamInstanceProfile", GoGetter: "IamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceInitiatedShutdownBehavior", GoGetter: "InstanceInitiatedShutdownBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressCount", GoGetter: "Ipv6AddressCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Addresses", GoGetter: "Ipv6Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "kernelId", GoGetter: "KernelId"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "licenseSpecifications", GoGetter: "LicenseSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "monitoring", GoGetter: "Monitoring"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaces", GoGetter: "NetworkInterfaces"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroupName", GoGetter: "PlacementGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsNameOptions", GoGetter: "PrivateDnsNameOptions"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "propagateTagsToVolumeOnCreation", GoGetter: "PropagateTagsToVolumeOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "ramdiskId", GoGetter: "RamdiskId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDestCheck", GoGetter: "SourceDestCheck"},
			_jsii_.MemberProperty{JsiiProperty: "ssmAssociations", GoGetter: "SsmAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberProperty{JsiiProperty: "tenancy", GoGetter: "Tenancy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.AssociationParameterProperty",
		reflect.TypeOf((*CfnInstance_AssociationParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.BlockDeviceMappingProperty",
		reflect.TypeOf((*CfnInstance_BlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.CpuOptionsProperty",
		reflect.TypeOf((*CfnInstance_CpuOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.CreditSpecificationProperty",
		reflect.TypeOf((*CfnInstance_CreditSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.EbsProperty",
		reflect.TypeOf((*CfnInstance_EbsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.ElasticGpuSpecificationProperty",
		reflect.TypeOf((*CfnInstance_ElasticGpuSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.ElasticInferenceAcceleratorProperty",
		reflect.TypeOf((*CfnInstance_ElasticInferenceAcceleratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.EnclaveOptionsProperty",
		reflect.TypeOf((*CfnInstance_EnclaveOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.HibernationOptionsProperty",
		reflect.TypeOf((*CfnInstance_HibernationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.InstanceIpv6AddressProperty",
		reflect.TypeOf((*CfnInstance_InstanceIpv6AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnInstance_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.LicenseSpecificationProperty",
		reflect.TypeOf((*CfnInstance_LicenseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.NetworkInterfaceProperty",
		reflect.TypeOf((*CfnInstance_NetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.NoDeviceProperty",
		reflect.TypeOf((*CfnInstance_NoDeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.PrivateDnsNameOptionsProperty",
		reflect.TypeOf((*CfnInstance_PrivateDnsNameOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.PrivateIpAddressSpecificationProperty",
		reflect.TypeOf((*CfnInstance_PrivateIpAddressSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.SsmAssociationProperty",
		reflect.TypeOf((*CfnInstance_SsmAssociationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstance.VolumeProperty",
		reflect.TypeOf((*CfnInstance_VolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnInstanceConnectEndpoint",
		reflect.TypeOf((*CfnInstanceConnectEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clientToken", GoGetter: "ClientToken"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preserveClientIp", GoGetter: "PreserveClientIp"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceConnectEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstanceConnectEndpointProps",
		reflect.TypeOf((*CfnInstanceConnectEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInstanceProps",
		reflect.TypeOf((*CfnInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnInternetGateway",
		reflect.TypeOf((*CfnInternetGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrInternetGatewayId", GoGetter: "AttrInternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInternetGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnInternetGatewayProps",
		reflect.TypeOf((*CfnInternetGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnKeyPair",
		reflect.TypeOf((*CfnKeyPair)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrKeyFingerprint", GoGetter: "AttrKeyFingerprint"},
			_jsii_.MemberProperty{JsiiProperty: "attrKeyPairId", GoGetter: "AttrKeyPairId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "keyFormat", GoGetter: "KeyFormat"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyMaterial", GoGetter: "PublicKeyMaterial"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeyPair{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnKeyPairProps",
		reflect.TypeOf((*CfnKeyPairProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate",
		reflect.TypeOf((*CfnLaunchTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultVersionNumber", GoGetter: "AttrDefaultVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "attrLatestVersionNumber", GoGetter: "AttrLatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "attrLaunchTemplateId", GoGetter: "AttrLaunchTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateData", GoGetter: "LaunchTemplateData"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateName", GoGetter: "LaunchTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "versionDescription", GoGetter: "VersionDescription"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.AcceleratorCountProperty",
		reflect.TypeOf((*CfnLaunchTemplate_AcceleratorCountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.AcceleratorTotalMemoryMiBProperty",
		reflect.TypeOf((*CfnLaunchTemplate_AcceleratorTotalMemoryMiBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.BaselineEbsBandwidthMbpsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_BaselineEbsBandwidthMbpsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.BlockDeviceMappingProperty",
		reflect.TypeOf((*CfnLaunchTemplate_BlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.CapacityReservationSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_CapacityReservationSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.CapacityReservationTargetProperty",
		reflect.TypeOf((*CfnLaunchTemplate_CapacityReservationTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.ConnectionTrackingSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_ConnectionTrackingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.CpuOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_CpuOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.CreditSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_CreditSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.EbsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_EbsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.ElasticGpuSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_ElasticGpuSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.EnaSrdSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_EnaSrdSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.EnaSrdUdpSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_EnaSrdUdpSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.EnclaveOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_EnclaveOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.HibernationOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_HibernationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.IamInstanceProfileProperty",
		reflect.TypeOf((*CfnLaunchTemplate_IamInstanceProfileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.InstanceMarketOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_InstanceMarketOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.InstanceRequirementsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_InstanceRequirementsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.Ipv4PrefixSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_Ipv4PrefixSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.Ipv6AddProperty",
		reflect.TypeOf((*CfnLaunchTemplate_Ipv6AddProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.Ipv6PrefixSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_Ipv6PrefixSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.LaunchTemplateDataProperty",
		reflect.TypeOf((*CfnLaunchTemplate_LaunchTemplateDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.LaunchTemplateElasticInferenceAcceleratorProperty",
		reflect.TypeOf((*CfnLaunchTemplate_LaunchTemplateElasticInferenceAcceleratorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.LaunchTemplateTagSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_LaunchTemplateTagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.LicenseSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_LicenseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.MaintenanceOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_MaintenanceOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.MemoryGiBPerVCpuProperty",
		reflect.TypeOf((*CfnLaunchTemplate_MemoryGiBPerVCpuProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.MemoryMiBProperty",
		reflect.TypeOf((*CfnLaunchTemplate_MemoryMiBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.MetadataOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_MetadataOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.MonitoringProperty",
		reflect.TypeOf((*CfnLaunchTemplate_MonitoringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.NetworkBandwidthGbpsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_NetworkBandwidthGbpsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.NetworkInterfaceCountProperty",
		reflect.TypeOf((*CfnLaunchTemplate_NetworkInterfaceCountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.NetworkInterfaceProperty",
		reflect.TypeOf((*CfnLaunchTemplate_NetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.PlacementProperty",
		reflect.TypeOf((*CfnLaunchTemplate_PlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.PrivateDnsNameOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_PrivateDnsNameOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.PrivateIpAddProperty",
		reflect.TypeOf((*CfnLaunchTemplate_PrivateIpAddProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.SpotOptionsProperty",
		reflect.TypeOf((*CfnLaunchTemplate_SpotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.TagSpecificationProperty",
		reflect.TypeOf((*CfnLaunchTemplate_TagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.TotalLocalStorageGBProperty",
		reflect.TypeOf((*CfnLaunchTemplate_TotalLocalStorageGBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplate.VCpuCountProperty",
		reflect.TypeOf((*CfnLaunchTemplate_VCpuCountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLaunchTemplateProps",
		reflect.TypeOf((*CfnLaunchTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRoute",
		reflect.TypeOf((*CfnLocalGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrType", GoGetter: "AttrType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayRouteTableId", GoGetter: "LocalGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayVirtualInterfaceGroupId", GoGetter: "LocalGatewayVirtualInterfaceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocalGatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteProps",
		reflect.TypeOf((*CfnLocalGatewayRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteTable",
		reflect.TypeOf((*CfnLocalGatewayRouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayRouteTableArn", GoGetter: "AttrLocalGatewayRouteTableArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayRouteTableId", GoGetter: "AttrLocalGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "attrOutpostArn", GoGetter: "AttrOutpostArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwnerId", GoGetter: "AttrOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayId", GoGetter: "LocalGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocalGatewayRouteTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteTableProps",
		reflect.TypeOf((*CfnLocalGatewayRouteTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteTableVPCAssociation",
		reflect.TypeOf((*CfnLocalGatewayRouteTableVPCAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayId", GoGetter: "AttrLocalGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayRouteTableVpcAssociationId", GoGetter: "AttrLocalGatewayRouteTableVpcAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayRouteTableId", GoGetter: "LocalGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocalGatewayRouteTableVPCAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteTableVPCAssociationProps",
		reflect.TypeOf((*CfnLocalGatewayRouteTableVPCAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociation",
		reflect.TypeOf((*CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayId", GoGetter: "AttrLocalGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayRouteTableArn", GoGetter: "AttrLocalGatewayRouteTableArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrLocalGatewayRouteTableVirtualInterfaceGroupAssociationId", GoGetter: "AttrLocalGatewayRouteTableVirtualInterfaceGroupAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwnerId", GoGetter: "AttrOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayRouteTableId", GoGetter: "LocalGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayVirtualInterfaceGroupId", GoGetter: "LocalGatewayVirtualInterfaceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps",
		reflect.TypeOf((*CfnLocalGatewayRouteTableVirtualInterfaceGroupAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNatGateway",
		reflect.TypeOf((*CfnNatGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allocationId", GoGetter: "AllocationId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrNatGatewayId", GoGetter: "AttrNatGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectivityType", GoGetter: "ConnectivityType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxDrainDurationSeconds", GoGetter: "MaxDrainDurationSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryAllocationIds", GoGetter: "SecondaryAllocationIds"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryPrivateIpAddressCount", GoGetter: "SecondaryPrivateIpAddressCount"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryPrivateIpAddresses", GoGetter: "SecondaryPrivateIpAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNatGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNatGatewayProps",
		reflect.TypeOf((*CfnNatGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkAcl",
		reflect.TypeOf((*CfnNetworkAcl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkAcl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkAclEntry",
		reflect.TypeOf((*CfnNetworkAclEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "egress", GoGetter: "Egress"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "icmp", GoGetter: "Icmp"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkAclId", GoGetter: "NetworkAclId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "portRange", GoGetter: "PortRange"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "ruleAction", GoGetter: "RuleAction"},
			_jsii_.MemberProperty{JsiiProperty: "ruleNumber", GoGetter: "RuleNumber"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkAclEntry{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkAclEntry.IcmpProperty",
		reflect.TypeOf((*CfnNetworkAclEntry_IcmpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkAclEntry.PortRangeProperty",
		reflect.TypeOf((*CfnNetworkAclEntry_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkAclEntryProps",
		reflect.TypeOf((*CfnNetworkAclEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkAclProps",
		reflect.TypeOf((*CfnNetworkAclProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScope",
		reflect.TypeOf((*CfnNetworkInsightsAccessScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedDate", GoGetter: "AttrCreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsAccessScopeArn", GoGetter: "AttrNetworkInsightsAccessScopeArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsAccessScopeId", GoGetter: "AttrNetworkInsightsAccessScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdatedDate", GoGetter: "AttrUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludePaths", GoGetter: "ExcludePaths"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "matchPaths", GoGetter: "MatchPaths"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInsightsAccessScope{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScope.AccessScopePathRequestProperty",
		reflect.TypeOf((*CfnNetworkInsightsAccessScope_AccessScopePathRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScope.PacketHeaderStatementRequestProperty",
		reflect.TypeOf((*CfnNetworkInsightsAccessScope_PacketHeaderStatementRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScope.PathStatementRequestProperty",
		reflect.TypeOf((*CfnNetworkInsightsAccessScope_PathStatementRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScope.ResourceStatementRequestProperty",
		reflect.TypeOf((*CfnNetworkInsightsAccessScope_ResourceStatementRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScope.ThroughResourcesStatementRequestProperty",
		reflect.TypeOf((*CfnNetworkInsightsAccessScope_ThroughResourcesStatementRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScopeAnalysis",
		reflect.TypeOf((*CfnNetworkInsightsAccessScopeAnalysis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAnalyzedEniCount", GoGetter: "AttrAnalyzedEniCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndDate", GoGetter: "AttrEndDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrFindingsFound", GoGetter: "AttrFindingsFound"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsAccessScopeAnalysisArn", GoGetter: "AttrNetworkInsightsAccessScopeAnalysisArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsAccessScopeAnalysisId", GoGetter: "AttrNetworkInsightsAccessScopeAnalysisId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStartDate", GoGetter: "AttrStartDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusMessage", GoGetter: "AttrStatusMessage"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInsightsAccessScopeId", GoGetter: "NetworkInsightsAccessScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInsightsAccessScopeAnalysis{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScopeAnalysisProps",
		reflect.TypeOf((*CfnNetworkInsightsAccessScopeAnalysisProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAccessScopeProps",
		reflect.TypeOf((*CfnNetworkInsightsAccessScopeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "additionalAccounts", GoGetter: "AdditionalAccounts"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAlternatePathHints", GoGetter: "AttrAlternatePathHints"},
			_jsii_.MemberProperty{JsiiProperty: "attrExplanations", GoGetter: "AttrExplanations"},
			_jsii_.MemberProperty{JsiiProperty: "attrForwardPathComponents", GoGetter: "AttrForwardPathComponents"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsAnalysisArn", GoGetter: "AttrNetworkInsightsAnalysisArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsAnalysisId", GoGetter: "AttrNetworkInsightsAnalysisId"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkPathFound", GoGetter: "AttrNetworkPathFound"},
			_jsii_.MemberProperty{JsiiProperty: "attrReturnPathComponents", GoGetter: "AttrReturnPathComponents"},
			_jsii_.MemberProperty{JsiiProperty: "attrStartDate", GoGetter: "AttrStartDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusMessage", GoGetter: "AttrStatusMessage"},
			_jsii_.MemberProperty{JsiiProperty: "attrSuggestedAccounts", GoGetter: "AttrSuggestedAccounts"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "filterInArns", GoGetter: "FilterInArns"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInsightsPathId", GoGetter: "NetworkInsightsPathId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInsightsAnalysis{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AdditionalDetailProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AdditionalDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AlternatePathHintProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AlternatePathHintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisAclRuleProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisAclRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisComponentProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisLoadBalancerListenerProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisLoadBalancerListenerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisLoadBalancerTargetProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisLoadBalancerTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisPacketHeaderProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisPacketHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisRouteTableRouteProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisRouteTableRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.AnalysisSecurityGroupRuleProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_AnalysisSecurityGroupRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.ExplanationProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_ExplanationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.PathComponentProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_PathComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.PortRangeProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_PortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysis.TransitGatewayRouteTableRouteProperty",
		reflect.TypeOf((*CfnNetworkInsightsAnalysis_TransitGatewayRouteTableRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsAnalysisProps",
		reflect.TypeOf((*CfnNetworkInsightsAnalysisProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsPath",
		reflect.TypeOf((*CfnNetworkInsightsPath)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedDate", GoGetter: "AttrCreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrDestinationArn", GoGetter: "AttrDestinationArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsPathArn", GoGetter: "AttrNetworkInsightsPathArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInsightsPathId", GoGetter: "AttrNetworkInsightsPathId"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceArn", GoGetter: "AttrSourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationIp", GoGetter: "DestinationIp"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPort", GoGetter: "DestinationPort"},
			_jsii_.MemberProperty{JsiiProperty: "filterAtDestination", GoGetter: "FilterAtDestination"},
			_jsii_.MemberProperty{JsiiProperty: "filterAtSource", GoGetter: "FilterAtSource"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIp", GoGetter: "SourceIp"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInsightsPath{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsPath.FilterPortRangeProperty",
		reflect.TypeOf((*CfnNetworkInsightsPath_FilterPortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsPath.PathFilterProperty",
		reflect.TypeOf((*CfnNetworkInsightsPath_PathFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInsightsPathProps",
		reflect.TypeOf((*CfnNetworkInsightsPathProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterface",
		reflect.TypeOf((*CfnNetworkInterface)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrimaryIpv6Address", GoGetter: "AttrPrimaryIpv6Address"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrimaryPrivateIpAddress", GoGetter: "AttrPrimaryPrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "attrSecondaryPrivateIpAddresses", GoGetter: "AttrSecondaryPrivateIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTrackingSpecification", GoGetter: "ConnectionTrackingSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enablePrimaryIpv6", GoGetter: "EnablePrimaryIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupSet", GoGetter: "GroupSet"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceType", GoGetter: "InterfaceType"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixCount", GoGetter: "Ipv4PrefixCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Prefixes", GoGetter: "Ipv4Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressCount", GoGetter: "Ipv6AddressCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Addresses", GoGetter: "Ipv6Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixCount", GoGetter: "Ipv6PrefixCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Prefixes", GoGetter: "Ipv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddresses", GoGetter: "PrivateIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryPrivateIpAddressCount", GoGetter: "SecondaryPrivateIpAddressCount"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceDestCheck", GoGetter: "SourceDestCheck"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInterface{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterface.ConnectionTrackingSpecificationProperty",
		reflect.TypeOf((*CfnNetworkInterface_ConnectionTrackingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterface.InstanceIpv6AddressProperty",
		reflect.TypeOf((*CfnNetworkInterface_InstanceIpv6AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterface.Ipv4PrefixSpecificationProperty",
		reflect.TypeOf((*CfnNetworkInterface_Ipv4PrefixSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterface.Ipv6PrefixSpecificationProperty",
		reflect.TypeOf((*CfnNetworkInterface_Ipv6PrefixSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterface.PrivateIpAddressSpecificationProperty",
		reflect.TypeOf((*CfnNetworkInterface_PrivateIpAddressSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfaceAttachment",
		reflect.TypeOf((*CfnNetworkInterfaceAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAttachmentId", GoGetter: "AttrAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTermination", GoGetter: "DeleteOnTermination"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndex", GoGetter: "DeviceIndex"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdSpecification", GoGetter: "EnaSrdSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInterfaceAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfaceAttachment.EnaSrdSpecificationProperty",
		reflect.TypeOf((*CfnNetworkInterfaceAttachment_EnaSrdSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfaceAttachment.EnaSrdUdpSpecificationProperty",
		reflect.TypeOf((*CfnNetworkInterfaceAttachment_EnaSrdUdpSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfaceAttachmentProps",
		reflect.TypeOf((*CfnNetworkInterfaceAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfacePermission",
		reflect.TypeOf((*CfnNetworkInterfacePermission)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permission", GoGetter: "Permission"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkInterfacePermission{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfacePermissionProps",
		reflect.TypeOf((*CfnNetworkInterfacePermissionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkInterfaceProps",
		reflect.TypeOf((*CfnNetworkInterfaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnNetworkPerformanceMetricSubscription",
		reflect.TypeOf((*CfnNetworkPerformanceMetricSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metric", GoGetter: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkPerformanceMetricSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnNetworkPerformanceMetricSubscriptionProps",
		reflect.TypeOf((*CfnNetworkPerformanceMetricSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnPlacementGroup",
		reflect.TypeOf((*CfnPlacementGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupName", GoGetter: "AttrGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "partitionCount", GoGetter: "PartitionCount"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spreadLevel", GoGetter: "SpreadLevel"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPlacementGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnPlacementGroupProps",
		reflect.TypeOf((*CfnPlacementGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnPrefixList",
		reflect.TypeOf((*CfnPrefixList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "addressFamily", GoGetter: "AddressFamily"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwnerId", GoGetter: "AttrOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPrefixListId", GoGetter: "AttrPrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVersion", GoGetter: "AttrVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "entries", GoGetter: "Entries"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "maxEntries", GoGetter: "MaxEntries"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "prefixListName", GoGetter: "PrefixListName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrefixList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnPrefixList.EntryProperty",
		reflect.TypeOf((*CfnPrefixList_EntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnPrefixListProps",
		reflect.TypeOf((*CfnPrefixListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnRoute",
		reflect.TypeOf((*CfnRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCidrBlock", GoGetter: "AttrCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "carrierGatewayId", GoGetter: "CarrierGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "coreNetworkArn", GoGetter: "CoreNetworkArn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "destinationIpv6CidrBlock", GoGetter: "DestinationIpv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPrefixListId", GoGetter: "DestinationPrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "egressOnlyInternetGatewayId", GoGetter: "EgressOnlyInternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "localGatewayId", GoGetter: "LocalGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayId", GoGetter: "NatGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcPeeringConnectionId", GoGetter: "VpcPeeringConnectionId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnRouteProps",
		reflect.TypeOf((*CfnRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnRouteTable",
		reflect.TypeOf((*CfnRouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrRouteTableId", GoGetter: "AttrRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRouteTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnRouteTableProps",
		reflect.TypeOf((*CfnRouteTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroup",
		reflect.TypeOf((*CfnSecurityGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupId", GoGetter: "AttrGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcId", GoGetter: "AttrVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupDescription", GoGetter: "GroupDescription"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupEgress", GoGetter: "SecurityGroupEgress"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIngress", GoGetter: "SecurityGroupIngress"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroup.EgressProperty",
		reflect.TypeOf((*CfnSecurityGroup_EgressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroup.IngressProperty",
		reflect.TypeOf((*CfnSecurityGroup_IngressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroupEgress",
		reflect.TypeOf((*CfnSecurityGroupEgress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIp", GoGetter: "CidrIp"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIpv6", GoGetter: "CidrIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPrefixListId", GoGetter: "DestinationPrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "destinationSecurityGroupId", GoGetter: "DestinationSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "fromPort", GoGetter: "FromPort"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipProtocol", GoGetter: "IpProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "toPort", GoGetter: "ToPort"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityGroupEgress{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroupEgressProps",
		reflect.TypeOf((*CfnSecurityGroupEgressProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroupIngress",
		reflect.TypeOf((*CfnSecurityGroupIngress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIp", GoGetter: "CidrIp"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIpv6", GoGetter: "CidrIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fromPort", GoGetter: "FromPort"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipProtocol", GoGetter: "IpProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePrefixListId", GoGetter: "SourcePrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupId", GoGetter: "SourceSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupName", GoGetter: "SourceSecurityGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupOwnerId", GoGetter: "SourceSecurityGroupOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "toPort", GoGetter: "ToPort"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityGroupIngress{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroupIngressProps",
		reflect.TypeOf((*CfnSecurityGroupIngressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSecurityGroupProps",
		reflect.TypeOf((*CfnSecurityGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSnapshotBlockPublicAccess",
		reflect.TypeOf((*CfnSnapshotBlockPublicAccess)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAccountId", GoGetter: "AttrAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSnapshotBlockPublicAccess{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSnapshotBlockPublicAccessProps",
		reflect.TypeOf((*CfnSnapshotBlockPublicAccessProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet",
		reflect.TypeOf((*CfnSpotFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "spotFleetRequestConfigData", GoGetter: "SpotFleetRequestConfigData"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSpotFleet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.AcceleratorCountRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_AcceleratorCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.AcceleratorTotalMemoryMiBRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_AcceleratorTotalMemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.BaselineEbsBandwidthMbpsRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_BaselineEbsBandwidthMbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.BlockDeviceMappingProperty",
		reflect.TypeOf((*CfnSpotFleet_BlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.ClassicLoadBalancerProperty",
		reflect.TypeOf((*CfnSpotFleet_ClassicLoadBalancerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.ClassicLoadBalancersConfigProperty",
		reflect.TypeOf((*CfnSpotFleet_ClassicLoadBalancersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.EbsBlockDeviceProperty",
		reflect.TypeOf((*CfnSpotFleet_EbsBlockDeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.FleetLaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnSpotFleet_FleetLaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.GroupIdentifierProperty",
		reflect.TypeOf((*CfnSpotFleet_GroupIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.IamInstanceProfileSpecificationProperty",
		reflect.TypeOf((*CfnSpotFleet_IamInstanceProfileSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.InstanceIpv6AddressProperty",
		reflect.TypeOf((*CfnSpotFleet_InstanceIpv6AddressProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.InstanceNetworkInterfaceSpecificationProperty",
		reflect.TypeOf((*CfnSpotFleet_InstanceNetworkInterfaceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.InstanceRequirementsRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_InstanceRequirementsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.LaunchTemplateConfigProperty",
		reflect.TypeOf((*CfnSpotFleet_LaunchTemplateConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.LaunchTemplateOverridesProperty",
		reflect.TypeOf((*CfnSpotFleet_LaunchTemplateOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.LoadBalancersConfigProperty",
		reflect.TypeOf((*CfnSpotFleet_LoadBalancersConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.MemoryGiBPerVCpuRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_MemoryGiBPerVCpuRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.MemoryMiBRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_MemoryMiBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.NetworkBandwidthGbpsRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_NetworkBandwidthGbpsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.NetworkInterfaceCountRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_NetworkInterfaceCountRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.PrivateIpAddressSpecificationProperty",
		reflect.TypeOf((*CfnSpotFleet_PrivateIpAddressSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotCapacityRebalanceProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotCapacityRebalanceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotFleetLaunchSpecificationProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotFleetLaunchSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotFleetMonitoringProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotFleetMonitoringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotFleetRequestConfigDataProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotFleetRequestConfigDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotFleetTagSpecificationProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotFleetTagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotMaintenanceStrategiesProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotMaintenanceStrategiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.SpotPlacementProperty",
		reflect.TypeOf((*CfnSpotFleet_SpotPlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.TargetGroupProperty",
		reflect.TypeOf((*CfnSpotFleet_TargetGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.TargetGroupsConfigProperty",
		reflect.TypeOf((*CfnSpotFleet_TargetGroupsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.TotalLocalStorageGBRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_TotalLocalStorageGBRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleet.VCpuCountRangeRequestProperty",
		reflect.TypeOf((*CfnSpotFleet_VCpuCountRangeRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSpotFleetProps",
		reflect.TypeOf((*CfnSpotFleetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSubnet",
		reflect.TypeOf((*CfnSubnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assignIpv6AddressOnCreation", GoGetter: "AssignIpv6AddressOnCreation"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailabilityZone", GoGetter: "AttrAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "attrAvailabilityZoneId", GoGetter: "AttrAvailabilityZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "attrCidrBlock", GoGetter: "AttrCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpv6CidrBlocks", GoGetter: "AttrIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkAclAssociationId", GoGetter: "AttrNetworkAclAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "attrOutpostArn", GoGetter: "AttrOutpostArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrSubnetId", GoGetter: "AttrSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcId", GoGetter: "AttrVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneId", GoGetter: "AvailabilityZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableDns64", GoGetter: "EnableDns64"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolId", GoGetter: "Ipv4IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4NetmaskLength", GoGetter: "Ipv4NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlocks", GoGetter: "Ipv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolId", GoGetter: "Ipv6IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Native", GoGetter: "Ipv6Native"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6NetmaskLength", GoGetter: "Ipv6NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mapPublicIpOnLaunch", GoGetter: "MapPublicIpOnLaunch"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "outpostArn", GoGetter: "OutpostArn"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsNameOptionsOnLaunch", GoGetter: "PrivateDnsNameOptionsOnLaunch"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubnet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSubnet.PrivateDnsNameOptionsOnLaunchProperty",
		reflect.TypeOf((*CfnSubnet_PrivateDnsNameOptionsOnLaunchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSubnetCidrBlock",
		reflect.TypeOf((*CfnSubnetCidrBlock)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolId", GoGetter: "Ipv6IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6NetmaskLength", GoGetter: "Ipv6NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubnetCidrBlock{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSubnetCidrBlockProps",
		reflect.TypeOf((*CfnSubnetCidrBlockProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSubnetNetworkAclAssociation",
		reflect.TypeOf((*CfnSubnetNetworkAclAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAssociationId", GoGetter: "AttrAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkAclId", GoGetter: "NetworkAclId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubnetNetworkAclAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSubnetNetworkAclAssociationProps",
		reflect.TypeOf((*CfnSubnetNetworkAclAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSubnetProps",
		reflect.TypeOf((*CfnSubnetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnSubnetRouteTableAssociation",
		reflect.TypeOf((*CfnSubnetRouteTableAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSubnetRouteTableAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnSubnetRouteTableAssociationProps",
		reflect.TypeOf((*CfnSubnetRouteTableAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorFilter",
		reflect.TypeOf((*CfnTrafficMirrorFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkServices", GoGetter: "NetworkServices"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrafficMirrorFilter{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorFilterProps",
		reflect.TypeOf((*CfnTrafficMirrorFilterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorFilterRule",
		reflect.TypeOf((*CfnTrafficMirrorFilterRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPortRange", GoGetter: "DestinationPortRange"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "ruleAction", GoGetter: "RuleAction"},
			_jsii_.MemberProperty{JsiiProperty: "ruleNumber", GoGetter: "RuleNumber"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCidrBlock", GoGetter: "SourceCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePortRange", GoGetter: "SourcePortRange"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficDirection", GoGetter: "TrafficDirection"},
			_jsii_.MemberProperty{JsiiProperty: "trafficMirrorFilterId", GoGetter: "TrafficMirrorFilterId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrafficMirrorFilterRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorFilterRule.TrafficMirrorPortRangeProperty",
		reflect.TypeOf((*CfnTrafficMirrorFilterRule_TrafficMirrorPortRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorFilterRuleProps",
		reflect.TypeOf((*CfnTrafficMirrorFilterRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorSession",
		reflect.TypeOf((*CfnTrafficMirrorSession)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "packetLength", GoGetter: "PacketLength"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "sessionNumber", GoGetter: "SessionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficMirrorFilterId", GoGetter: "TrafficMirrorFilterId"},
			_jsii_.MemberProperty{JsiiProperty: "trafficMirrorTargetId", GoGetter: "TrafficMirrorTargetId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrafficMirrorSession{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorSessionProps",
		reflect.TypeOf((*CfnTrafficMirrorSessionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorTarget",
		reflect.TypeOf((*CfnTrafficMirrorTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayLoadBalancerEndpointId", GoGetter: "GatewayLoadBalancerEndpointId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "networkLoadBalancerArn", GoGetter: "NetworkLoadBalancerArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrafficMirrorTarget{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTrafficMirrorTargetProps",
		reflect.TypeOf((*CfnTrafficMirrorTargetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGateway",
		reflect.TypeOf((*CfnTransitGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonSideAsn", GoGetter: "AmazonSideAsn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "associationDefaultRouteTableId", GoGetter: "AssociationDefaultRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayArn", GoGetter: "AttrTransitGatewayArn"},
			_jsii_.MemberProperty{JsiiProperty: "autoAcceptSharedAttachments", GoGetter: "AutoAcceptSharedAttachments"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTableAssociation", GoGetter: "DefaultRouteTableAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRouteTablePropagation", GoGetter: "DefaultRouteTablePropagation"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupport", GoGetter: "DnsSupport"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "multicastSupport", GoGetter: "MulticastSupport"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "propagationDefaultRouteTableId", GoGetter: "PropagationDefaultRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayCidrBlocks", GoGetter: "TransitGatewayCidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpnEcmpSupport", GoGetter: "VpnEcmpSupport"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayAttachment",
		reflect.TypeOf((*CfnTransitGatewayAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayAttachment.OptionsProperty",
		reflect.TypeOf((*CfnTransitGatewayAttachment_OptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayAttachmentProps",
		reflect.TypeOf((*CfnTransitGatewayAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayConnect",
		reflect.TypeOf((*CfnTransitGatewayConnect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayAttachmentId", GoGetter: "AttrTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayId", GoGetter: "AttrTransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transportTransitGatewayAttachmentId", GoGetter: "TransportTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayConnect{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayConnect.TransitGatewayConnectOptionsProperty",
		reflect.TypeOf((*CfnTransitGatewayConnect_TransitGatewayConnectOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayConnectProps",
		reflect.TypeOf((*CfnTransitGatewayConnectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastDomain",
		reflect.TypeOf((*CfnTransitGatewayMulticastDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayMulticastDomainArn", GoGetter: "AttrTransitGatewayMulticastDomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayMulticastDomainId", GoGetter: "AttrTransitGatewayMulticastDomainId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayMulticastDomain{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastDomain.OptionsProperty",
		reflect.TypeOf((*CfnTransitGatewayMulticastDomain_OptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastDomainAssociation",
		reflect.TypeOf((*CfnTransitGatewayMulticastDomainAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceId", GoGetter: "AttrResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceType", GoGetter: "AttrResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayMulticastDomainId", GoGetter: "TransitGatewayMulticastDomainId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayMulticastDomainAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastDomainAssociationProps",
		reflect.TypeOf((*CfnTransitGatewayMulticastDomainAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastDomainProps",
		reflect.TypeOf((*CfnTransitGatewayMulticastDomainProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastGroupMember",
		reflect.TypeOf((*CfnTransitGatewayMulticastGroupMember)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupMember", GoGetter: "AttrGroupMember"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupSource", GoGetter: "AttrGroupSource"},
			_jsii_.MemberProperty{JsiiProperty: "attrMemberType", GoGetter: "AttrMemberType"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceId", GoGetter: "AttrResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceType", GoGetter: "AttrResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceType", GoGetter: "AttrSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrSubnetId", GoGetter: "AttrSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayAttachmentId", GoGetter: "AttrTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupIpAddress", GoGetter: "GroupIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayMulticastDomainId", GoGetter: "TransitGatewayMulticastDomainId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayMulticastGroupMember{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastGroupMemberProps",
		reflect.TypeOf((*CfnTransitGatewayMulticastGroupMemberProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastGroupSource",
		reflect.TypeOf((*CfnTransitGatewayMulticastGroupSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupMember", GoGetter: "AttrGroupMember"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupSource", GoGetter: "AttrGroupSource"},
			_jsii_.MemberProperty{JsiiProperty: "attrMemberType", GoGetter: "AttrMemberType"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceId", GoGetter: "AttrResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "attrResourceType", GoGetter: "AttrResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrSourceType", GoGetter: "AttrSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "attrSubnetId", GoGetter: "AttrSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayAttachmentId", GoGetter: "AttrTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupIpAddress", GoGetter: "GroupIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayMulticastDomainId", GoGetter: "TransitGatewayMulticastDomainId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayMulticastGroupSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayMulticastGroupSourceProps",
		reflect.TypeOf((*CfnTransitGatewayMulticastGroupSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayPeeringAttachment",
		reflect.TypeOf((*CfnTransitGatewayPeeringAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusCode", GoGetter: "AttrStatusCode"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusMessage", GoGetter: "AttrStatusMessage"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayAttachmentId", GoGetter: "AttrTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "peerAccountId", GoGetter: "PeerAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "peerRegion", GoGetter: "PeerRegion"},
			_jsii_.MemberProperty{JsiiProperty: "peerTransitGatewayId", GoGetter: "PeerTransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayPeeringAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayPeeringAttachment.PeeringAttachmentStatusProperty",
		reflect.TypeOf((*CfnTransitGatewayPeeringAttachment_PeeringAttachmentStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayPeeringAttachmentProps",
		reflect.TypeOf((*CfnTransitGatewayPeeringAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayProps",
		reflect.TypeOf((*CfnTransitGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRoute",
		reflect.TypeOf((*CfnTransitGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "blackhole", GoGetter: "Blackhole"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableId", GoGetter: "TransitGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteProps",
		reflect.TypeOf((*CfnTransitGatewayRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteTable",
		reflect.TypeOf((*CfnTransitGatewayRouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrTransitGatewayRouteTableId", GoGetter: "AttrTransitGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRouteTable{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteTableAssociation",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableId", GoGetter: "TransitGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRouteTableAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteTableAssociationProps",
		reflect.TypeOf((*CfnTransitGatewayRouteTableAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteTablePropagation",
		reflect.TypeOf((*CfnTransitGatewayRouteTablePropagation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayAttachmentId", GoGetter: "TransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayRouteTableId", GoGetter: "TransitGatewayRouteTableId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayRouteTablePropagation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteTablePropagationProps",
		reflect.TypeOf((*CfnTransitGatewayRouteTablePropagationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayRouteTableProps",
		reflect.TypeOf((*CfnTransitGatewayRouteTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayVpcAttachment",
		reflect.TypeOf((*CfnTransitGatewayVpcAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "addSubnetIds", GoGetter: "AddSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberProperty{JsiiProperty: "removeSubnetIds", GoGetter: "RemoveSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransitGatewayVpcAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayVpcAttachment.OptionsProperty",
		reflect.TypeOf((*CfnTransitGatewayVpcAttachment_OptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnTransitGatewayVpcAttachmentProps",
		reflect.TypeOf((*CfnTransitGatewayVpcAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPC",
		reflect.TypeOf((*CfnVPC)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCidrBlock", GoGetter: "AttrCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "attrCidrBlockAssociations", GoGetter: "AttrCidrBlockAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultNetworkAcl", GoGetter: "AttrDefaultNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "attrDefaultSecurityGroup", GoGetter: "AttrDefaultSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "attrIpv6CidrBlocks", GoGetter: "AttrIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcId", GoGetter: "AttrVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsHostnames", GoGetter: "EnableDnsHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "enableDnsSupport", GoGetter: "EnableDnsSupport"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTenancy", GoGetter: "InstanceTenancy"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolId", GoGetter: "Ipv4IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4NetmaskLength", GoGetter: "Ipv4NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPC{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCCidrBlock",
		reflect.TypeOf((*CfnVPCCidrBlock)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonProvidedIpv6CidrBlock", GoGetter: "AmazonProvidedIpv6CidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolId", GoGetter: "Ipv4IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4NetmaskLength", GoGetter: "Ipv4NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolId", GoGetter: "Ipv6IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6NetmaskLength", GoGetter: "Ipv6NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Pool", GoGetter: "Ipv6Pool"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCCidrBlock{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCCidrBlockProps",
		reflect.TypeOf((*CfnVPCCidrBlockProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCDHCPOptionsAssociation",
		reflect.TypeOf((*CfnVPCDHCPOptionsAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dhcpOptionsId", GoGetter: "DhcpOptionsId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCDHCPOptionsAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCDHCPOptionsAssociationProps",
		reflect.TypeOf((*CfnVPCDHCPOptionsAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpoint",
		reflect.TypeOf((*CfnVPCEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTimestamp", GoGetter: "AttrCreationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "attrDnsEntries", GoGetter: "AttrDnsEntries"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrNetworkInterfaceIds", GoGetter: "AttrNetworkInterfaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsEnabled", GoGetter: "PrivateDnsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableIds", GoGetter: "RouteTableIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointType", GoGetter: "VpcEndpointType"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointConnectionNotification",
		reflect.TypeOf((*CfnVPCEndpointConnectionNotification)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpcEndpointConnectionNotificationId", GoGetter: "AttrVpcEndpointConnectionNotificationId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionEvents", GoGetter: "ConnectionEvents"},
			_jsii_.MemberProperty{JsiiProperty: "connectionNotificationArn", GoGetter: "ConnectionNotificationArn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceId", GoGetter: "ServiceId"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCEndpointConnectionNotification{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointConnectionNotificationProps",
		reflect.TypeOf((*CfnVPCEndpointConnectionNotificationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointProps",
		reflect.TypeOf((*CfnVPCEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointService",
		reflect.TypeOf((*CfnVPCEndpointService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptanceRequired", GoGetter: "AcceptanceRequired"},
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrServiceId", GoGetter: "AttrServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "contributorInsightsEnabled", GoGetter: "ContributorInsightsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayLoadBalancerArns", GoGetter: "GatewayLoadBalancerArns"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkLoadBalancerArns", GoGetter: "NetworkLoadBalancerArns"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "payerResponsibility", GoGetter: "PayerResponsibility"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCEndpointService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointServicePermissions",
		reflect.TypeOf((*CfnVPCEndpointServicePermissions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedPrincipals", GoGetter: "AllowedPrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceId", GoGetter: "ServiceId"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCEndpointServicePermissions{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointServicePermissionsProps",
		reflect.TypeOf((*CfnVPCEndpointServicePermissionsProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCEndpointServiceProps",
		reflect.TypeOf((*CfnVPCEndpointServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCGatewayAttachment",
		reflect.TypeOf((*CfnVPCGatewayAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAttachmentType", GoGetter: "AttrAttachmentType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCGatewayAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCGatewayAttachmentProps",
		reflect.TypeOf((*CfnVPCGatewayAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPCPeeringConnection",
		reflect.TypeOf((*CfnVPCPeeringConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "peerOwnerId", GoGetter: "PeerOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "peerRegion", GoGetter: "PeerRegion"},
			_jsii_.MemberProperty{JsiiProperty: "peerRoleArn", GoGetter: "PeerRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "peerVpcId", GoGetter: "PeerVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPCPeeringConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCPeeringConnectionProps",
		reflect.TypeOf((*CfnVPCPeeringConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPCProps",
		reflect.TypeOf((*CfnVPCProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection",
		reflect.TypeOf((*CfnVPNConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpnConnectionId", GoGetter: "AttrVpnConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "staticRoutesOnly", GoGetter: "StaticRoutesOnly"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnTunnelOptionsSpecifications", GoGetter: "VpnTunnelOptionsSpecifications"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPNConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPNConnection.VpnTunnelOptionsSpecificationProperty",
		reflect.TypeOf((*CfnVPNConnection_VpnTunnelOptionsSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPNConnectionProps",
		reflect.TypeOf((*CfnVPNConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPNConnectionRoute",
		reflect.TypeOf((*CfnVPNConnectionRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpnConnectionId", GoGetter: "VpnConnectionId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPNConnectionRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPNConnectionRouteProps",
		reflect.TypeOf((*CfnVPNConnectionRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPNGateway",
		reflect.TypeOf((*CfnVPNGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonSideAsn", GoGetter: "AmazonSideAsn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrVpnGatewayId", GoGetter: "AttrVpnGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPNGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPNGatewayProps",
		reflect.TypeOf((*CfnVPNGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVPNGatewayRoutePropagation",
		reflect.TypeOf((*CfnVPNGatewayRoutePropagation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableIds", GoGetter: "RouteTableIds"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVPNGatewayRoutePropagation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVPNGatewayRoutePropagationProps",
		reflect.TypeOf((*CfnVPNGatewayRoutePropagationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessEndpoint",
		reflect.TypeOf((*CfnVerifiedAccessEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationDomain", GoGetter: "ApplicationDomain"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attachmentType", GoGetter: "AttachmentType"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrDeviceValidationDomain", GoGetter: "AttrDeviceValidationDomain"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpointDomain", GoGetter: "AttrEndpointDomain"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrVerifiedAccessEndpointId", GoGetter: "AttrVerifiedAccessEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "attrVerifiedAccessInstanceId", GoGetter: "AttrVerifiedAccessInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "domainCertificateArn", GoGetter: "DomainCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "endpointDomainPrefix", GoGetter: "EndpointDomainPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "endpointType", GoGetter: "EndpointType"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerOptions", GoGetter: "LoadBalancerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceOptions", GoGetter: "NetworkInterfaceOptions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "policyEnabled", GoGetter: "PolicyEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "verifiedAccessGroupId", GoGetter: "VerifiedAccessGroupId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVerifiedAccessEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessEndpoint.LoadBalancerOptionsProperty",
		reflect.TypeOf((*CfnVerifiedAccessEndpoint_LoadBalancerOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessEndpoint.NetworkInterfaceOptionsProperty",
		reflect.TypeOf((*CfnVerifiedAccessEndpoint_NetworkInterfaceOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessEndpoint.SseSpecificationProperty",
		reflect.TypeOf((*CfnVerifiedAccessEndpoint_SseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessEndpointProps",
		reflect.TypeOf((*CfnVerifiedAccessEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessGroup",
		reflect.TypeOf((*CfnVerifiedAccessGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrOwner", GoGetter: "AttrOwner"},
			_jsii_.MemberProperty{JsiiProperty: "attrVerifiedAccessGroupArn", GoGetter: "AttrVerifiedAccessGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrVerifiedAccessGroupId", GoGetter: "AttrVerifiedAccessGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "policyEnabled", GoGetter: "PolicyEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "verifiedAccessInstanceId", GoGetter: "VerifiedAccessInstanceId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVerifiedAccessGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessGroup.SseSpecificationProperty",
		reflect.TypeOf((*CfnVerifiedAccessGroup_SseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessGroupProps",
		reflect.TypeOf((*CfnVerifiedAccessGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstance",
		reflect.TypeOf((*CfnVerifiedAccessInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVerifiedAccessInstanceId", GoGetter: "AttrVerifiedAccessInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fipsEnabled", GoGetter: "FipsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurations", GoGetter: "LoggingConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "verifiedAccessTrustProviderIds", GoGetter: "VerifiedAccessTrustProviderIds"},
			_jsii_.MemberProperty{JsiiProperty: "verifiedAccessTrustProviders", GoGetter: "VerifiedAccessTrustProviders"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVerifiedAccessInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstance.CloudWatchLogsProperty",
		reflect.TypeOf((*CfnVerifiedAccessInstance_CloudWatchLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstance.KinesisDataFirehoseProperty",
		reflect.TypeOf((*CfnVerifiedAccessInstance_KinesisDataFirehoseProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstance.S3Property",
		reflect.TypeOf((*CfnVerifiedAccessInstance_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstance.VerifiedAccessLogsProperty",
		reflect.TypeOf((*CfnVerifiedAccessInstance_VerifiedAccessLogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstance.VerifiedAccessTrustProviderProperty",
		reflect.TypeOf((*CfnVerifiedAccessInstance_VerifiedAccessTrustProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessInstanceProps",
		reflect.TypeOf((*CfnVerifiedAccessInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessTrustProvider",
		reflect.TypeOf((*CfnVerifiedAccessTrustProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdatedTime", GoGetter: "AttrLastUpdatedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrVerifiedAccessTrustProviderId", GoGetter: "AttrVerifiedAccessTrustProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "deviceOptions", GoGetter: "DeviceOptions"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTrustProviderType", GoGetter: "DeviceTrustProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "oidcOptions", GoGetter: "OidcOptions"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyReferenceName", GoGetter: "PolicyReferenceName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sseSpecification", GoGetter: "SseSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustProviderType", GoGetter: "TrustProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "userTrustProviderType", GoGetter: "UserTrustProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVerifiedAccessTrustProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessTrustProvider.DeviceOptionsProperty",
		reflect.TypeOf((*CfnVerifiedAccessTrustProvider_DeviceOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessTrustProvider.OidcOptionsProperty",
		reflect.TypeOf((*CfnVerifiedAccessTrustProvider_OidcOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessTrustProvider.SseSpecificationProperty",
		reflect.TypeOf((*CfnVerifiedAccessTrustProvider_SseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVerifiedAccessTrustProviderProps",
		reflect.TypeOf((*CfnVerifiedAccessTrustProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVolume",
		reflect.TypeOf((*CfnVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrVolumeId", GoGetter: "AttrVolumeId"},
			_jsii_.MemberProperty{JsiiProperty: "autoEnableIo", GoGetter: "AutoEnableIo"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encrypted", GoGetter: "Encrypted"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "multiAttachEnabled", GoGetter: "MultiAttachEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "outpostArn", GoGetter: "OutpostArn"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotId", GoGetter: "SnapshotId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberProperty{JsiiProperty: "throughput", GoGetter: "Throughput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "volumeType", GoGetter: "VolumeType"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVolume{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CfnVolumeAttachment",
		reflect.TypeOf((*CfnVolumeAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "device", GoGetter: "Device"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "volumeId", GoGetter: "VolumeId"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVolumeAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVolumeAttachmentProps",
		reflect.TypeOf((*CfnVolumeAttachmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CfnVolumeProps",
		reflect.TypeOf((*CfnVolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.ClientVpnAuthorizationRule",
		reflect.TypeOf((*ClientVpnAuthorizationRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClientVpnAuthorizationRule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnAuthorizationRuleOptions",
		reflect.TypeOf((*ClientVpnAuthorizationRuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnAuthorizationRuleProps",
		reflect.TypeOf((*ClientVpnAuthorizationRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.ClientVpnEndpoint",
		reflect.TypeOf((*ClientVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAuthorizationRule", GoMethod: "AddAuthorizationRule"},
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "endpointId", GoGetter: "EndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworksAssociated", GoGetter: "TargetNetworksAssociated"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClientVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IClientVpnEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnEndpointAttributes",
		reflect.TypeOf((*ClientVpnEndpointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnEndpointOptions",
		reflect.TypeOf((*ClientVpnEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnEndpointProps",
		reflect.TypeOf((*ClientVpnEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.ClientVpnRoute",
		reflect.TypeOf((*ClientVpnRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClientVpnRoute{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnRouteOptions",
		reflect.TypeOf((*ClientVpnRouteOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ClientVpnRouteProps",
		reflect.TypeOf((*ClientVpnRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.ClientVpnRouteTarget",
		reflect.TypeOf((*ClientVpnRouteTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
		},
		func() interface{} {
			return &jsiiProxy_ClientVpnRouteTarget{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.ClientVpnSessionTimeout",
		reflect.TypeOf((*ClientVpnSessionTimeout)(nil)).Elem(),
		map[string]interface{}{
			"EIGHT_HOURS": ClientVpnSessionTimeout_EIGHT_HOURS,
			"TEN_HOURS": ClientVpnSessionTimeout_TEN_HOURS,
			"TWELVE_HOURS": ClientVpnSessionTimeout_TWELVE_HOURS,
			"TWENTY_FOUR_HOURS": ClientVpnSessionTimeout_TWENTY_FOUR_HOURS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.ClientVpnUserBasedAuthentication",
		reflect.TypeOf((*ClientVpnUserBasedAuthentication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_ClientVpnUserBasedAuthentication{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.CloudFormationInit",
		reflect.TypeOf((*CloudFormationInit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfig", GoMethod: "AddConfig"},
			_jsii_.MemberMethod{JsiiMethod: "addConfigSet", GoMethod: "AddConfigSet"},
			_jsii_.MemberMethod{JsiiMethod: "attach", GoMethod: "Attach"},
		},
		func() interface{} {
			return &jsiiProxy_CloudFormationInit{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CommonNetworkAclEntryOptions",
		reflect.TypeOf((*CommonNetworkAclEntryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ConfigSetProps",
		reflect.TypeOf((*ConfigSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ConfigureNatOptions",
		reflect.TypeOf((*ConfigureNatOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ConnectionRule",
		reflect.TypeOf((*ConnectionRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Connections",
		reflect.TypeOf((*Connections)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "allowDefaultPortFrom", GoMethod: "AllowDefaultPortFrom"},
			_jsii_.MemberMethod{JsiiMethod: "allowDefaultPortFromAnyIpv4", GoMethod: "AllowDefaultPortFromAnyIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "allowDefaultPortInternally", GoMethod: "AllowDefaultPortInternally"},
			_jsii_.MemberMethod{JsiiMethod: "allowDefaultPortTo", GoMethod: "AllowDefaultPortTo"},
			_jsii_.MemberMethod{JsiiMethod: "allowFrom", GoMethod: "AllowFrom"},
			_jsii_.MemberMethod{JsiiMethod: "allowFromAnyIpv4", GoMethod: "AllowFromAnyIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "allowInternally", GoMethod: "AllowInternally"},
			_jsii_.MemberMethod{JsiiMethod: "allowTo", GoMethod: "AllowTo"},
			_jsii_.MemberMethod{JsiiMethod: "allowToAnyIpv4", GoMethod: "AllowToAnyIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "allowToDefaultPort", GoMethod: "AllowToDefaultPort"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPort", GoGetter: "DefaultPort"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
		},
		func() interface{} {
			j := jsiiProxy_Connections{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ConnectionsProps",
		reflect.TypeOf((*ConnectionsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.CpuCredits",
		reflect.TypeOf((*CpuCredits)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": CpuCredits_STANDARD,
			"UNLIMITED": CpuCredits_UNLIMITED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.CreateIpv6CidrBlocksRequest",
		reflect.TypeOf((*CreateIpv6CidrBlocksRequest)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.DefaultInstanceTenancy",
		reflect.TypeOf((*DefaultInstanceTenancy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DefaultInstanceTenancy_DEFAULT,
			"DEDICATED": DefaultInstanceTenancy_DEDICATED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.DestinationOptions",
		reflect.TypeOf((*DestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.EbsDeviceOptions",
		reflect.TypeOf((*EbsDeviceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.EbsDeviceOptionsBase",
		reflect.TypeOf((*EbsDeviceOptionsBase)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.EbsDeviceProps",
		reflect.TypeOf((*EbsDeviceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.EbsDeviceSnapshotOptions",
		reflect.TypeOf((*EbsDeviceSnapshotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.EbsDeviceVolumeType",
		reflect.TypeOf((*EbsDeviceVolumeType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": EbsDeviceVolumeType_STANDARD,
			"IO1": EbsDeviceVolumeType_IO1,
			"IO2": EbsDeviceVolumeType_IO2,
			"GP2": EbsDeviceVolumeType_GP2,
			"GP3": EbsDeviceVolumeType_GP3,
			"ST1": EbsDeviceVolumeType_ST1,
			"SC1": EbsDeviceVolumeType_SC1,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.EnableVpnGatewayOptions",
		reflect.TypeOf((*EnableVpnGatewayOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.ExecuteFileOptions",
		reflect.TypeOf((*ExecuteFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.FlowLog",
		reflect.TypeOf((*FlowLog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogId", GoGetter: "FlowLogId"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamRole", GoGetter: "IamRole"},
			_jsii_.MemberProperty{JsiiProperty: "keyPrefix", GoGetter: "KeyPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLog{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFlowLog)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.FlowLogDestination",
		reflect.TypeOf((*FlowLogDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_FlowLogDestination{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.FlowLogDestinationConfig",
		reflect.TypeOf((*FlowLogDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.FlowLogDestinationType",
		reflect.TypeOf((*FlowLogDestinationType)(nil)).Elem(),
		map[string]interface{}{
			"CLOUD_WATCH_LOGS": FlowLogDestinationType_CLOUD_WATCH_LOGS,
			"S3": FlowLogDestinationType_S3,
			"KINESIS_DATA_FIREHOSE": FlowLogDestinationType_KINESIS_DATA_FIREHOSE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.FlowLogFileFormat",
		reflect.TypeOf((*FlowLogFileFormat)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN_TEXT": FlowLogFileFormat_PLAIN_TEXT,
			"PARQUET": FlowLogFileFormat_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.FlowLogMaxAggregationInterval",
		reflect.TypeOf((*FlowLogMaxAggregationInterval)(nil)).Elem(),
		map[string]interface{}{
			"ONE_MINUTE": FlowLogMaxAggregationInterval_ONE_MINUTE,
			"TEN_MINUTES": FlowLogMaxAggregationInterval_TEN_MINUTES,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.FlowLogOptions",
		reflect.TypeOf((*FlowLogOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.FlowLogProps",
		reflect.TypeOf((*FlowLogProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.FlowLogResourceType",
		reflect.TypeOf((*FlowLogResourceType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "resourceId", GoGetter: "ResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
		},
		func() interface{} {
			return &jsiiProxy_FlowLogResourceType{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.FlowLogTrafficType",
		reflect.TypeOf((*FlowLogTrafficType)(nil)).Elem(),
		map[string]interface{}{
			"ACCEPT": FlowLogTrafficType_ACCEPT,
			"ALL": FlowLogTrafficType_ALL,
			"REJECT": FlowLogTrafficType_REJECT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.GatewayConfig",
		reflect.TypeOf((*GatewayConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpoint",
		reflect.TypeOf((*GatewayVpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointCreationTimestamp", GoGetter: "VpcEndpointCreationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointDnsEntries", GoGetter: "VpcEndpointDnsEntries"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointNetworkInterfaceIds", GoGetter: "VpcEndpointNetworkInterfaceIds"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayVpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VpcEndpoint)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayVpcEndpoint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointAwsService",
		reflect.TypeOf((*GatewayVpcEndpointAwsService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayVpcEndpointAwsService{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGatewayVpcEndpointService)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointOptions",
		reflect.TypeOf((*GatewayVpcEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.GatewayVpcEndpointProps",
		reflect.TypeOf((*GatewayVpcEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.GenericLinuxImage",
		reflect.TypeOf((*GenericLinuxImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_GenericLinuxImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.GenericLinuxImageProps",
		reflect.TypeOf((*GenericLinuxImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.GenericSSMParameterImage",
		reflect.TypeOf((*GenericSSMParameterImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
		},
		func() interface{} {
			j := jsiiProxy_GenericSSMParameterImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.GenericWindowsImage",
		reflect.TypeOf((*GenericWindowsImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_GenericWindowsImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.GenericWindowsImageProps",
		reflect.TypeOf((*GenericWindowsImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IClientVpnConnectionHandler",
		reflect.TypeOf((*IClientVpnConnectionHandler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "functionArn", GoGetter: "FunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
		},
		func() interface{} {
			return &jsiiProxy_IClientVpnConnectionHandler{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IClientVpnEndpoint",
		reflect.TypeOf((*IClientVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "endpointId", GoGetter: "EndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworksAssociated", GoGetter: "TargetNetworksAssociated"},
		},
		func() interface{} {
			j := jsiiProxy_IClientVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IConnectable",
		reflect.TypeOf((*IConnectable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
		},
		func() interface{} {
			return &jsiiProxy_IConnectable{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IFlowLog",
		reflect.TypeOf((*IFlowLog)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "flowLogId", GoGetter: "FlowLogId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IFlowLog{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IGatewayVpcEndpoint",
		reflect.TypeOf((*IGatewayVpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
		},
		func() interface{} {
			j := jsiiProxy_IGatewayVpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcEndpoint)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IGatewayVpcEndpointService",
		reflect.TypeOf((*IGatewayVpcEndpointService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IGatewayVpcEndpointService{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IInstance",
		reflect.TypeOf((*IInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instanceAvailabilityZone", GoGetter: "InstanceAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateDnsName", GoGetter: "InstancePrivateDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateIp", GoGetter: "InstancePrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicDnsName", GoGetter: "InstancePublicDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicIp", GoGetter: "InstancePublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IInstance{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IInterfaceVpcEndpoint",
		reflect.TypeOf((*IInterfaceVpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
		},
		func() interface{} {
			j := jsiiProxy_IInterfaceVpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcEndpoint)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IInterfaceVpcEndpointService",
		reflect.TypeOf((*IInterfaceVpcEndpointService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsDefault", GoGetter: "PrivateDnsDefault"},
		},
		func() interface{} {
			return &jsiiProxy_IInterfaceVpcEndpointService{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IIpAddresses",
		reflect.TypeOf((*IIpAddresses)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allocateSubnetsCidr", GoMethod: "AllocateSubnetsCidr"},
			_jsii_.MemberMethod{JsiiMethod: "allocateVpcCidr", GoMethod: "AllocateVpcCidr"},
		},
		func() interface{} {
			return &jsiiProxy_IIpAddresses{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IIpv6Addresses",
		reflect.TypeOf((*IIpv6Addresses)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allocateSubnetsIpv6Cidr", GoMethod: "AllocateSubnetsIpv6Cidr"},
			_jsii_.MemberMethod{JsiiMethod: "allocateVpcIpv6Cidr", GoMethod: "AllocateVpcIpv6Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "amazonProvided", GoGetter: "AmazonProvided"},
			_jsii_.MemberMethod{JsiiMethod: "createIpv6CidrBlocks", GoMethod: "CreateIpv6CidrBlocks"},
		},
		func() interface{} {
			return &jsiiProxy_IIpv6Addresses{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IKeyPair",
		reflect.TypeOf((*IKeyPair)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairName", GoGetter: "KeyPairName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_IKeyPair{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.ILaunchTemplate",
		reflect.TypeOf((*ILaunchTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateId", GoGetter: "LaunchTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateName", GoGetter: "LaunchTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "versionNumber", GoGetter: "VersionNumber"},
		},
		func() interface{} {
			j := jsiiProxy_ILaunchTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IMachineImage",
		reflect.TypeOf((*IMachineImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			return &jsiiProxy_IMachineImage{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.INetworkAcl",
		reflect.TypeOf((*INetworkAcl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEntry", GoMethod: "AddEntry"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "networkAclId", GoGetter: "NetworkAclId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_INetworkAcl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.INetworkAclEntry",
		reflect.TypeOf((*INetworkAclEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_INetworkAclEntry{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IPeer",
		reflect.TypeOf((*IPeer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canInlineRule", GoGetter: "CanInlineRule"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberMethod{JsiiMethod: "toEgressRuleConfig", GoMethod: "ToEgressRuleConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toIngressRuleConfig", GoMethod: "ToIngressRuleConfig"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueId", GoGetter: "UniqueId"},
		},
		func() interface{} {
			j := jsiiProxy_IPeer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IPlacementGroup",
		reflect.TypeOf((*IPlacementGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partitions", GoGetter: "Partitions"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroupName", GoGetter: "PlacementGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "spreadLevel", GoGetter: "SpreadLevel"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
		},
		func() interface{} {
			j := jsiiProxy_IPlacementGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IPrefixList",
		reflect.TypeOf((*IPrefixList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "prefixListId", GoGetter: "PrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPrefixList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IPrivateSubnet",
		reflect.TypeOf((*IPrivateSubnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
		},
		func() interface{} {
			j := jsiiProxy_IPrivateSubnet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubnet)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IPublicSubnet",
		reflect.TypeOf((*IPublicSubnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
		},
		func() interface{} {
			j := jsiiProxy_IPublicSubnet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubnet)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IRouteTable",
		reflect.TypeOf((*IRouteTable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "routeTableId", GoGetter: "RouteTableId"},
		},
		func() interface{} {
			return &jsiiProxy_IRouteTable{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.ISecurityGroup",
		reflect.TypeOf((*ISecurityGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEgressRule", GoMethod: "AddEgressRule"},
			_jsii_.MemberMethod{JsiiMethod: "addIngressRule", GoMethod: "AddIngressRule"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllOutbound", GoGetter: "AllowAllOutbound"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "canInlineRule", GoGetter: "CanInlineRule"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupId", GoGetter: "SecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toEgressRuleConfig", GoMethod: "ToEgressRuleConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toIngressRuleConfig", GoMethod: "ToIngressRuleConfig"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueId", GoGetter: "UniqueId"},
		},
		func() interface{} {
			j := jsiiProxy_ISecurityGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPeer)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.ISubnet",
		reflect.TypeOf((*ISubnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
		},
		func() interface{} {
			j := jsiiProxy_ISubnet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.ISubnetNetworkAclAssociation",
		reflect.TypeOf((*ISubnetNetworkAclAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetNetworkAclAssociationAssociationId", GoGetter: "SubnetNetworkAclAssociationAssociationId"},
		},
		func() interface{} {
			j := jsiiProxy_ISubnetNetworkAclAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVolume",
		reflect.TypeOf((*IVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantAttachVolume", GoMethod: "GrantAttachVolume"},
			_jsii_.MemberMethod{JsiiMethod: "grantAttachVolumeByResourceTag", GoMethod: "GrantAttachVolumeByResourceTag"},
			_jsii_.MemberMethod{JsiiMethod: "grantDetachVolume", GoMethod: "GrantDetachVolume"},
			_jsii_.MemberMethod{JsiiMethod: "grantDetachVolumeByResourceTag", GoMethod: "GrantDetachVolumeByResourceTag"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "volumeId", GoGetter: "VolumeId"},
		},
		func() interface{} {
			j := jsiiProxy_IVolume{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVpc",
		reflect.TypeOf((*IVpc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_IVpc{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVpcEndpoint",
		reflect.TypeOf((*IVpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVpcEndpointService",
		reflect.TypeOf((*IVpcEndpointService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceId", GoGetter: "VpcEndpointServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceName", GoGetter: "VpcEndpointServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_IVpcEndpointService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVpcEndpointServiceLoadBalancer",
		reflect.TypeOf((*IVpcEndpointServiceLoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerArn", GoGetter: "LoadBalancerArn"},
		},
		func() interface{} {
			return &jsiiProxy_IVpcEndpointServiceLoadBalancer{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVpnConnection",
		reflect.TypeOf((*IVpnConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayAsn", GoGetter: "CustomerGatewayAsn"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayIp", GoGetter: "CustomerGatewayIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataIn", GoMethod: "MetricTunnelDataIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataOut", GoMethod: "MetricTunnelDataOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelState", GoMethod: "MetricTunnelState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpnId", GoGetter: "VpnId"},
		},
		func() interface{} {
			j := jsiiProxy_IVpnConnection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_ec2.IVpnGateway",
		reflect.TypeOf((*IVpnGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IVpnGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitCommand",
		reflect.TypeOf((*InitCommand)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			j := jsiiProxy_InitCommand{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitCommandOptions",
		reflect.TypeOf((*InitCommandOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitCommandWaitDuration",
		reflect.TypeOf((*InitCommandWaitDuration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InitCommandWaitDuration{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitConfig",
		reflect.TypeOf((*InitConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberMethod{JsiiMethod: "isEmpty", GoMethod: "IsEmpty"},
		},
		func() interface{} {
			return &jsiiProxy_InitConfig{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitElement",
		reflect.TypeOf((*InitElement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			return &jsiiProxy_InitElement{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitFile",
		reflect.TypeOf((*InitFile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			j := jsiiProxy_InitFile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitFileAssetOptions",
		reflect.TypeOf((*InitFileAssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitFileOptions",
		reflect.TypeOf((*InitFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitGroup",
		reflect.TypeOf((*InitGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			j := jsiiProxy_InitGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitPackage",
		reflect.TypeOf((*InitPackage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
			_jsii_.MemberMethod{JsiiMethod: "renderPackageVersions", GoMethod: "RenderPackageVersions"},
		},
		func() interface{} {
			j := jsiiProxy_InitPackage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitService",
		reflect.TypeOf((*InitService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			j := jsiiProxy_InitService{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitServiceOptions",
		reflect.TypeOf((*InitServiceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitServiceRestartHandle",
		reflect.TypeOf((*InitServiceRestartHandle)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InitServiceRestartHandle{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitSource",
		reflect.TypeOf((*InitSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			j := jsiiProxy_InitSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitSourceAssetOptions",
		reflect.TypeOf((*InitSourceAssetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitSourceOptions",
		reflect.TypeOf((*InitSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InitUser",
		reflect.TypeOf((*InitUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elementType", GoGetter: "ElementType"},
		},
		func() interface{} {
			j := jsiiProxy_InitUser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InitElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InitUserOptions",
		reflect.TypeOf((*InitUserOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Instance",
		reflect.TypeOf((*Instance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addUserData", GoMethod: "AddUserData"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "instanceAvailabilityZone", GoGetter: "InstanceAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateDnsName", GoGetter: "InstancePrivateDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateIp", GoGetter: "InstancePrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicDnsName", GoGetter: "InstancePublicDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicIp", GoGetter: "InstancePublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
		},
		func() interface{} {
			j := jsiiProxy_Instance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInstance)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.InstanceArchitecture",
		reflect.TypeOf((*InstanceArchitecture)(nil)).Elem(),
		map[string]interface{}{
			"ARM_64": InstanceArchitecture_ARM_64,
			"X86_64": InstanceArchitecture_X86_64,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.InstanceClass",
		reflect.TypeOf((*InstanceClass)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD3": InstanceClass_STANDARD3,
			"M3": InstanceClass_M3,
			"STANDARD4": InstanceClass_STANDARD4,
			"M4": InstanceClass_M4,
			"STANDARD5": InstanceClass_STANDARD5,
			"M5": InstanceClass_M5,
			"STANDARD5_NVME_DRIVE": InstanceClass_STANDARD5_NVME_DRIVE,
			"M5D": InstanceClass_M5D,
			"STANDARD5_AMD": InstanceClass_STANDARD5_AMD,
			"M5A": InstanceClass_M5A,
			"STANDARD5_AMD_NVME_DRIVE": InstanceClass_STANDARD5_AMD_NVME_DRIVE,
			"M5AD": InstanceClass_M5AD,
			"STANDARD5_HIGH_PERFORMANCE": InstanceClass_STANDARD5_HIGH_PERFORMANCE,
			"M5N": InstanceClass_M5N,
			"STANDARD5_NVME_DRIVE_HIGH_PERFORMANCE": InstanceClass_STANDARD5_NVME_DRIVE_HIGH_PERFORMANCE,
			"M5DN": InstanceClass_M5DN,
			"STANDARD5_HIGH_COMPUTE": InstanceClass_STANDARD5_HIGH_COMPUTE,
			"M5ZN": InstanceClass_M5ZN,
			"MEMORY3": InstanceClass_MEMORY3,
			"R3": InstanceClass_R3,
			"MEMORY4": InstanceClass_MEMORY4,
			"R4": InstanceClass_R4,
			"MEMORY5": InstanceClass_MEMORY5,
			"R5": InstanceClass_R5,
			"MEMORY6_AMD": InstanceClass_MEMORY6_AMD,
			"R6A": InstanceClass_R6A,
			"MEMORY6_INTEL": InstanceClass_MEMORY6_INTEL,
			"R6I": InstanceClass_R6I,
			"MEMORY6_INTEL_NVME_DRIVE": InstanceClass_MEMORY6_INTEL_NVME_DRIVE,
			"R6ID": InstanceClass_R6ID,
			"MEMORY6_INTEL_HIGH_PERFORMANCE": InstanceClass_MEMORY6_INTEL_HIGH_PERFORMANCE,
			"R6IN": InstanceClass_R6IN,
			"MEMORY6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE": InstanceClass_MEMORY6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE,
			"R6IDN": InstanceClass_R6IDN,
			"MEMORY5_HIGH_PERFORMANCE": InstanceClass_MEMORY5_HIGH_PERFORMANCE,
			"R5N": InstanceClass_R5N,
			"MEMORY5_NVME_DRIVE": InstanceClass_MEMORY5_NVME_DRIVE,
			"R5D": InstanceClass_R5D,
			"MEMORY5_NVME_DRIVE_HIGH_PERFORMANCE": InstanceClass_MEMORY5_NVME_DRIVE_HIGH_PERFORMANCE,
			"R5DN": InstanceClass_R5DN,
			"MEMORY5_AMD": InstanceClass_MEMORY5_AMD,
			"R5A": InstanceClass_R5A,
			"MEMORY5_AMD_NVME_DRIVE": InstanceClass_MEMORY5_AMD_NVME_DRIVE,
			"HIGH_MEMORY_3TB_1": InstanceClass_HIGH_MEMORY_3TB_1,
			"U_3TB1": InstanceClass_U_3TB1,
			"HIGH_MEMORY_6TB_1": InstanceClass_HIGH_MEMORY_6TB_1,
			"U_6TB1": InstanceClass_U_6TB1,
			"HIGH_MEMORY_9TB_1": InstanceClass_HIGH_MEMORY_9TB_1,
			"U_9TB1": InstanceClass_U_9TB1,
			"HIGH_MEMORY_12TB_1": InstanceClass_HIGH_MEMORY_12TB_1,
			"U_12TB1": InstanceClass_U_12TB1,
			"HIGH_MEMORY_18TB_1": InstanceClass_HIGH_MEMORY_18TB_1,
			"U_18TB1": InstanceClass_U_18TB1,
			"HIGH_MEMORY_24TB_1": InstanceClass_HIGH_MEMORY_24TB_1,
			"U_24TB1": InstanceClass_U_24TB1,
			"R5AD": InstanceClass_R5AD,
			"MEMORY5_EBS_OPTIMIZED": InstanceClass_MEMORY5_EBS_OPTIMIZED,
			"R5B": InstanceClass_R5B,
			"MEMORY6_GRAVITON": InstanceClass_MEMORY6_GRAVITON,
			"R6G": InstanceClass_R6G,
			"MEMORY6_GRAVITON2_NVME_DRIVE": InstanceClass_MEMORY6_GRAVITON2_NVME_DRIVE,
			"R6GD": InstanceClass_R6GD,
			"MEMORY7_GRAVITON": InstanceClass_MEMORY7_GRAVITON,
			"R7G": InstanceClass_R7G,
			"MEMORY7_GRAVITON3_NVME_DRIVE": InstanceClass_MEMORY7_GRAVITON3_NVME_DRIVE,
			"R7GD": InstanceClass_R7GD,
			"MEMORY7_INTEL_BASE": InstanceClass_MEMORY7_INTEL_BASE,
			"R7I": InstanceClass_R7I,
			"MEMORY7_INTEL": InstanceClass_MEMORY7_INTEL,
			"R7IZ": InstanceClass_R7IZ,
			"MEMORY7_AMD": InstanceClass_MEMORY7_AMD,
			"R7A": InstanceClass_R7A,
			"COMPUTE3": InstanceClass_COMPUTE3,
			"C3": InstanceClass_C3,
			"COMPUTE4": InstanceClass_COMPUTE4,
			"C4": InstanceClass_C4,
			"COMPUTE5": InstanceClass_COMPUTE5,
			"C5": InstanceClass_C5,
			"COMPUTE5_NVME_DRIVE": InstanceClass_COMPUTE5_NVME_DRIVE,
			"C5D": InstanceClass_C5D,
			"COMPUTE5_AMD": InstanceClass_COMPUTE5_AMD,
			"C5A": InstanceClass_C5A,
			"COMPUTE5_AMD_NVME_DRIVE": InstanceClass_COMPUTE5_AMD_NVME_DRIVE,
			"C5AD": InstanceClass_C5AD,
			"COMPUTE5_HIGH_PERFORMANCE": InstanceClass_COMPUTE5_HIGH_PERFORMANCE,
			"C5N": InstanceClass_C5N,
			"COMPUTE6_INTEL": InstanceClass_COMPUTE6_INTEL,
			"C6I": InstanceClass_C6I,
			"COMPUTE6_INTEL_NVME_DRIVE": InstanceClass_COMPUTE6_INTEL_NVME_DRIVE,
			"C6ID": InstanceClass_C6ID,
			"COMPUTE6_INTEL_HIGH_PERFORMANCE": InstanceClass_COMPUTE6_INTEL_HIGH_PERFORMANCE,
			"C6IN": InstanceClass_C6IN,
			"COMPUTE6_AMD": InstanceClass_COMPUTE6_AMD,
			"C6A": InstanceClass_C6A,
			"COMPUTE6_GRAVITON2": InstanceClass_COMPUTE6_GRAVITON2,
			"C6G": InstanceClass_C6G,
			"COMPUTE7_GRAVITON3": InstanceClass_COMPUTE7_GRAVITON3,
			"C7G": InstanceClass_C7G,
			"COMPUTE6_GRAVITON2_NVME_DRIVE": InstanceClass_COMPUTE6_GRAVITON2_NVME_DRIVE,
			"C6GD": InstanceClass_C6GD,
			"COMPUTE7_GRAVITON3_NVME_DRIVE": InstanceClass_COMPUTE7_GRAVITON3_NVME_DRIVE,
			"C7GD": InstanceClass_C7GD,
			"COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWIDTH": InstanceClass_COMPUTE6_GRAVITON2_HIGH_NETWORK_BANDWIDTH,
			"C6GN": InstanceClass_C6GN,
			"COMPUTE7_GRAVITON3_HIGH_NETWORK_BANDWIDTH": InstanceClass_COMPUTE7_GRAVITON3_HIGH_NETWORK_BANDWIDTH,
			"C7GN": InstanceClass_C7GN,
			"COMPUTE7_INTEL": InstanceClass_COMPUTE7_INTEL,
			"C7I": InstanceClass_C7I,
			"COMPUTE7_AMD": InstanceClass_COMPUTE7_AMD,
			"C7A": InstanceClass_C7A,
			"STORAGE2": InstanceClass_STORAGE2,
			"D2": InstanceClass_D2,
			"STORAGE3": InstanceClass_STORAGE3,
			"D3": InstanceClass_D3,
			"STORAGE3_ENHANCED_NETWORK": InstanceClass_STORAGE3_ENHANCED_NETWORK,
			"D3EN": InstanceClass_D3EN,
			"STORAGE_COMPUTE_1": InstanceClass_STORAGE_COMPUTE_1,
			"TRN1": InstanceClass_TRN1,
			"TRN1N": InstanceClass_TRN1N,
			"H1": InstanceClass_H1,
			"IO3": InstanceClass_IO3,
			"I3": InstanceClass_I3,
			"IO3_DENSE_NVME_DRIVE": InstanceClass_IO3_DENSE_NVME_DRIVE,
			"I3EN": InstanceClass_I3EN,
			"IO4_INTEL": InstanceClass_IO4_INTEL,
			"I4I": InstanceClass_I4I,
			"STORAGE4_GRAVITON": InstanceClass_STORAGE4_GRAVITON,
			"I4G": InstanceClass_I4G,
			"STORAGE4_GRAVITON_NETWORK_OPTIMIZED": InstanceClass_STORAGE4_GRAVITON_NETWORK_OPTIMIZED,
			"IM4GN": InstanceClass_IM4GN,
			"STORAGE4_GRAVITON_NETWORK_STORAGE_OPTIMIZED": InstanceClass_STORAGE4_GRAVITON_NETWORK_STORAGE_OPTIMIZED,
			"IS4GEN": InstanceClass_IS4GEN,
			"BURSTABLE2": InstanceClass_BURSTABLE2,
			"T2": InstanceClass_T2,
			"BURSTABLE3": InstanceClass_BURSTABLE3,
			"T3": InstanceClass_T3,
			"BURSTABLE3_AMD": InstanceClass_BURSTABLE3_AMD,
			"T3A": InstanceClass_T3A,
			"BURSTABLE4_GRAVITON": InstanceClass_BURSTABLE4_GRAVITON,
			"T4G": InstanceClass_T4G,
			"MEMORY_INTENSIVE_1": InstanceClass_MEMORY_INTENSIVE_1,
			"X1": InstanceClass_X1,
			"MEMORY_INTENSIVE_1_EXTENDED": InstanceClass_MEMORY_INTENSIVE_1_EXTENDED,
			"X1E": InstanceClass_X1E,
			"MEMORY_INTENSIVE_2_GRAVITON2": InstanceClass_MEMORY_INTENSIVE_2_GRAVITON2,
			"X2G": InstanceClass_X2G,
			"MEMORY_INTENSIVE_2_GRAVITON2_NVME_DRIVE": InstanceClass_MEMORY_INTENSIVE_2_GRAVITON2_NVME_DRIVE,
			"X2GD": InstanceClass_X2GD,
			"MEMORY_INTENSIVE_2_XT_INTEL": InstanceClass_MEMORY_INTENSIVE_2_XT_INTEL,
			"X2IEDN": InstanceClass_X2IEDN,
			"MEMORY_INTENSIVE_2_INTEL": InstanceClass_MEMORY_INTENSIVE_2_INTEL,
			"X2IDN": InstanceClass_X2IDN,
			"MEMORY_INTENSIVE_2_XTZ_INTEL": InstanceClass_MEMORY_INTENSIVE_2_XTZ_INTEL,
			"X2IEZN": InstanceClass_X2IEZN,
			"FPGA1": InstanceClass_FPGA1,
			"F1": InstanceClass_F1,
			"GRAPHICS3_SMALL": InstanceClass_GRAPHICS3_SMALL,
			"G3S": InstanceClass_G3S,
			"GRAPHICS3": InstanceClass_GRAPHICS3,
			"G3": InstanceClass_G3,
			"GRAPHICS4_NVME_DRIVE_HIGH_PERFORMANCE": InstanceClass_GRAPHICS4_NVME_DRIVE_HIGH_PERFORMANCE,
			"G4DN": InstanceClass_G4DN,
			"GRAPHICS4_AMD_NVME_DRIVE": InstanceClass_GRAPHICS4_AMD_NVME_DRIVE,
			"G4AD": InstanceClass_G4AD,
			"GRAPHICS5": InstanceClass_GRAPHICS5,
			"G5": InstanceClass_G5,
			"GRAPHICS5_GRAVITON2": InstanceClass_GRAPHICS5_GRAVITON2,
			"G5G": InstanceClass_G5G,
			"PARALLEL2": InstanceClass_PARALLEL2,
			"P2": InstanceClass_P2,
			"PARALLEL3": InstanceClass_PARALLEL3,
			"P3": InstanceClass_P3,
			"PARALLEL3_NVME_DRIVE_HIGH_PERFORMANCE": InstanceClass_PARALLEL3_NVME_DRIVE_HIGH_PERFORMANCE,
			"P3DN": InstanceClass_P3DN,
			"PARALLEL4_NVME_DRIVE_EXTENDED": InstanceClass_PARALLEL4_NVME_DRIVE_EXTENDED,
			"P4DE": InstanceClass_P4DE,
			"PARALLEL4": InstanceClass_PARALLEL4,
			"P4D": InstanceClass_P4D,
			"PARALLEL5": InstanceClass_PARALLEL5,
			"P5": InstanceClass_P5,
			"ARM1": InstanceClass_ARM1,
			"A1": InstanceClass_A1,
			"STANDARD6_GRAVITON": InstanceClass_STANDARD6_GRAVITON,
			"M6G": InstanceClass_M6G,
			"STANDARD6_INTEL": InstanceClass_STANDARD6_INTEL,
			"M6I": InstanceClass_M6I,
			"STANDARD6_INTEL_NVME_DRIVE": InstanceClass_STANDARD6_INTEL_NVME_DRIVE,
			"M6ID": InstanceClass_M6ID,
			"STANDARD6_INTEL_HIGH_PERFORMANCE": InstanceClass_STANDARD6_INTEL_HIGH_PERFORMANCE,
			"M6IN": InstanceClass_M6IN,
			"STANDARD6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE": InstanceClass_STANDARD6_INTEL_NVME_DRIVE_HIGH_PERFORMANCE,
			"M6IDN": InstanceClass_M6IDN,
			"STANDARD6_AMD": InstanceClass_STANDARD6_AMD,
			"M6A": InstanceClass_M6A,
			"STANDARD6_GRAVITON2_NVME_DRIVE": InstanceClass_STANDARD6_GRAVITON2_NVME_DRIVE,
			"M6GD": InstanceClass_M6GD,
			"STANDARD7_GRAVITON": InstanceClass_STANDARD7_GRAVITON,
			"M7G": InstanceClass_M7G,
			"STANDARD7_GRAVITON3_NVME_DRIVE": InstanceClass_STANDARD7_GRAVITON3_NVME_DRIVE,
			"M7GD": InstanceClass_M7GD,
			"STANDARD7_INTEL": InstanceClass_STANDARD7_INTEL,
			"M7I": InstanceClass_M7I,
			"STANDARD7_INTEL_FLEX": InstanceClass_STANDARD7_INTEL_FLEX,
			"M7I_FLEX": InstanceClass_M7I_FLEX,
			"STANDARD7_AMD": InstanceClass_STANDARD7_AMD,
			"M7A": InstanceClass_M7A,
			"HIGH_COMPUTE_MEMORY1": InstanceClass_HIGH_COMPUTE_MEMORY1,
			"Z1D": InstanceClass_Z1D,
			"INFERENCE1": InstanceClass_INFERENCE1,
			"INF1": InstanceClass_INF1,
			"INFERENCE2": InstanceClass_INFERENCE2,
			"INF2": InstanceClass_INF2,
			"MACINTOSH1_INTEL": InstanceClass_MACINTOSH1_INTEL,
			"MAC1": InstanceClass_MAC1,
			"MACINTOSH2_M1": InstanceClass_MACINTOSH2_M1,
			"MAC2": InstanceClass_MAC2,
			"MACINTOSH2_M2": InstanceClass_MACINTOSH2_M2,
			"MAC2_M2": InstanceClass_MAC2_M2,
			"MACINTOSH2_M2_PRO": InstanceClass_MACINTOSH2_M2_PRO,
			"MAC2_M2PRO": InstanceClass_MAC2_M2PRO,
			"VIDEO_TRANSCODING1": InstanceClass_VIDEO_TRANSCODING1,
			"VT1": InstanceClass_VT1,
			"HIGH_PERFORMANCE_COMPUTING6_AMD": InstanceClass_HIGH_PERFORMANCE_COMPUTING6_AMD,
			"HPC6A": InstanceClass_HPC6A,
			"HIGH_PERFORMANCE_COMPUTING6_INTEL_NVME_DRIVE": InstanceClass_HIGH_PERFORMANCE_COMPUTING6_INTEL_NVME_DRIVE,
			"HPC6ID": InstanceClass_HPC6ID,
			"HIGH_PERFORMANCE_COMPUTING7_AMD": InstanceClass_HIGH_PERFORMANCE_COMPUTING7_AMD,
			"HPC7A": InstanceClass_HPC7A,
			"HIGH_PERFORMANCE_COMPUTING7_GRAVITON": InstanceClass_HIGH_PERFORMANCE_COMPUTING7_GRAVITON,
			"HPC7G": InstanceClass_HPC7G,
			"DEEP_LEARNING1": InstanceClass_DEEP_LEARNING1,
			"DL1": InstanceClass_DL1,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.InstanceInitiatedShutdownBehavior",
		reflect.TypeOf((*InstanceInitiatedShutdownBehavior)(nil)).Elem(),
		map[string]interface{}{
			"STOP": InstanceInitiatedShutdownBehavior_STOP,
			"TERMINATE": InstanceInitiatedShutdownBehavior_TERMINATE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InstanceProps",
		reflect.TypeOf((*InstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InstanceRequireImdsv2Aspect",
		reflect.TypeOf((*InstanceRequireImdsv2Aspect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "suppressWarnings", GoGetter: "SuppressWarnings"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "warn", GoMethod: "Warn"},
		},
		func() interface{} {
			j := jsiiProxy_InstanceRequireImdsv2Aspect{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InstanceRequireImdsv2AspectProps",
		reflect.TypeOf((*InstanceRequireImdsv2AspectProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.InstanceSize",
		reflect.TypeOf((*InstanceSize)(nil)).Elem(),
		map[string]interface{}{
			"NANO": InstanceSize_NANO,
			"MICRO": InstanceSize_MICRO,
			"SMALL": InstanceSize_SMALL,
			"MEDIUM": InstanceSize_MEDIUM,
			"LARGE": InstanceSize_LARGE,
			"XLARGE": InstanceSize_XLARGE,
			"XLARGE2": InstanceSize_XLARGE2,
			"XLARGE3": InstanceSize_XLARGE3,
			"XLARGE4": InstanceSize_XLARGE4,
			"XLARGE6": InstanceSize_XLARGE6,
			"XLARGE8": InstanceSize_XLARGE8,
			"XLARGE9": InstanceSize_XLARGE9,
			"XLARGE10": InstanceSize_XLARGE10,
			"XLARGE12": InstanceSize_XLARGE12,
			"XLARGE16": InstanceSize_XLARGE16,
			"XLARGE18": InstanceSize_XLARGE18,
			"XLARGE24": InstanceSize_XLARGE24,
			"XLARGE32": InstanceSize_XLARGE32,
			"XLARGE48": InstanceSize_XLARGE48,
			"XLARGE56": InstanceSize_XLARGE56,
			"XLARGE96": InstanceSize_XLARGE96,
			"XLARGE112": InstanceSize_XLARGE112,
			"METAL": InstanceSize_METAL,
			"XLARGE16METAL": InstanceSize_XLARGE16METAL,
			"XLARGE24METAL": InstanceSize_XLARGE24METAL,
			"XLARGE32METAL": InstanceSize_XLARGE32METAL,
			"XLARGE48METAL": InstanceSize_XLARGE48METAL,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InstanceType",
		reflect.TypeOf((*InstanceType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "architecture", GoGetter: "Architecture"},
			_jsii_.MemberMethod{JsiiMethod: "isBurstable", GoMethod: "IsBurstable"},
			_jsii_.MemberMethod{JsiiMethod: "sameInstanceClassAs", GoMethod: "SameInstanceClassAs"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_InstanceType{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		reflect.TypeOf((*InterfaceVpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointCreationTimestamp", GoGetter: "VpcEndpointCreationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointDnsEntries", GoGetter: "VpcEndpointDnsEntries"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointNetworkInterfaceIds", GoGetter: "VpcEndpointNetworkInterfaceIds"},
		},
		func() interface{} {
			j := jsiiProxy_InterfaceVpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VpcEndpoint)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInterfaceVpcEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAttributes",
		reflect.TypeOf((*InterfaceVpcEndpointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointAwsService",
		reflect.TypeOf((*InterfaceVpcEndpointAwsService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsDefault", GoGetter: "PrivateDnsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "shortName", GoGetter: "ShortName"},
		},
		func() interface{} {
			j := jsiiProxy_InterfaceVpcEndpointAwsService{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInterfaceVpcEndpointService)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointOptions",
		reflect.TypeOf((*InterfaceVpcEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointProps",
		reflect.TypeOf((*InterfaceVpcEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpointService",
		reflect.TypeOf((*InterfaceVpcEndpointService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsDefault", GoGetter: "PrivateDnsDefault"},
		},
		func() interface{} {
			j := jsiiProxy_InterfaceVpcEndpointService{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInterfaceVpcEndpointService)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.IpAddresses",
		reflect.TypeOf((*IpAddresses)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IpAddresses{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.IpProtocol",
		reflect.TypeOf((*IpProtocol)(nil)).Elem(),
		map[string]interface{}{
			"IPV4_ONLY": IpProtocol_IPV4_ONLY,
			"DUAL_STACK": IpProtocol_DUAL_STACK,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Ipv6Addresses",
		reflect.TypeOf((*Ipv6Addresses)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Ipv6Addresses{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.KeyPair",
		reflect.TypeOf((*KeyPair)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hasImportedMaterial", GoGetter: "HasImportedMaterial"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairFingerprint", GoGetter: "KeyPairFingerprint"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairId", GoGetter: "KeyPairId"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairName", GoGetter: "KeyPairName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_KeyPair{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKeyPair)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.KeyPairAttributes",
		reflect.TypeOf((*KeyPairAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.KeyPairFormat",
		reflect.TypeOf((*KeyPairFormat)(nil)).Elem(),
		map[string]interface{}{
			"PPK": KeyPairFormat_PPK,
			"PEM": KeyPairFormat_PEM,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.KeyPairProps",
		reflect.TypeOf((*KeyPairProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.KeyPairType",
		reflect.TypeOf((*KeyPairType)(nil)).Elem(),
		map[string]interface{}{
			"RSA": KeyPairType_RSA,
			"ED25519": KeyPairType_ED25519,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.LaunchTemplate",
		reflect.TypeOf((*LaunchTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVersionNumber", GoGetter: "DefaultVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateId", GoGetter: "LaunchTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateName", GoGetter: "LaunchTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
			_jsii_.MemberProperty{JsiiProperty: "versionNumber", GoGetter: "VersionNumber"},
		},
		func() interface{} {
			j := jsiiProxy_LaunchTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILaunchTemplate)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LaunchTemplateAttributes",
		reflect.TypeOf((*LaunchTemplateAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.LaunchTemplateHttpTokens",
		reflect.TypeOf((*LaunchTemplateHttpTokens)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": LaunchTemplateHttpTokens_OPTIONAL,
			"REQUIRED": LaunchTemplateHttpTokens_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LaunchTemplateProps",
		reflect.TypeOf((*LaunchTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.LaunchTemplateRequireImdsv2Aspect",
		reflect.TypeOf((*LaunchTemplateRequireImdsv2Aspect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "suppressWarnings", GoGetter: "SuppressWarnings"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "warn", GoMethod: "Warn"},
		},
		func() interface{} {
			j := jsiiProxy_LaunchTemplateRequireImdsv2Aspect{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LaunchTemplateRequireImdsv2AspectProps",
		reflect.TypeOf((*LaunchTemplateRequireImdsv2AspectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.LaunchTemplateSpecialVersions",
		reflect.TypeOf((*LaunchTemplateSpecialVersions)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LaunchTemplateSpecialVersions{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LaunchTemplateSpotOptions",
		reflect.TypeOf((*LaunchTemplateSpotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LinuxUserDataOptions",
		reflect.TypeOf((*LinuxUserDataOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LocationPackageOptions",
		reflect.TypeOf((*LocationPackageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.LogFormat",
		reflect.TypeOf((*LogFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_LogFormat{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.LookupMachineImage",
		reflect.TypeOf((*LookupMachineImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_LookupMachineImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.LookupMachineImageProps",
		reflect.TypeOf((*LookupMachineImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.MachineImage",
		reflect.TypeOf((*MachineImage)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MachineImage{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.MachineImageConfig",
		reflect.TypeOf((*MachineImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.MultipartBody",
		reflect.TypeOf((*MultipartBody)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "renderBodyPart", GoMethod: "RenderBodyPart"},
		},
		func() interface{} {
			return &jsiiProxy_MultipartBody{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.MultipartBodyOptions",
		reflect.TypeOf((*MultipartBodyOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.MultipartUserData",
		reflect.TypeOf((*MultipartUserData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCommands", GoMethod: "AddCommands"},
			_jsii_.MemberMethod{JsiiMethod: "addExecuteFileCommand", GoMethod: "AddExecuteFileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addOnExitCommands", GoMethod: "AddOnExitCommands"},
			_jsii_.MemberMethod{JsiiMethod: "addPart", GoMethod: "AddPart"},
			_jsii_.MemberMethod{JsiiMethod: "addS3DownloadCommand", GoMethod: "AddS3DownloadCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addSignalOnExitCommand", GoMethod: "AddSignalOnExitCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addUserDataPart", GoMethod: "AddUserDataPart"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			j := jsiiProxy_MultipartUserData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UserData)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.MultipartUserDataOptions",
		reflect.TypeOf((*MultipartUserDataOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.NamedPackageOptions",
		reflect.TypeOf((*NamedPackageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.NatGatewayProps",
		reflect.TypeOf((*NatGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.NatInstanceImage",
		reflect.TypeOf((*NatInstanceImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_NatInstanceImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_LookupMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.NatInstanceProps",
		reflect.TypeOf((*NatInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.NatInstanceProvider",
		reflect.TypeOf((*NatInstanceProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuredGateways", GoGetter: "ConfiguredGateways"},
			_jsii_.MemberMethod{JsiiMethod: "configureNat", GoMethod: "ConfigureNat"},
			_jsii_.MemberMethod{JsiiMethod: "configureSubnet", GoMethod: "ConfigureSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
		},
		func() interface{} {
			j := jsiiProxy_NatInstanceProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NatProvider)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.NatInstanceProviderV2",
		reflect.TypeOf((*NatInstanceProviderV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuredGateways", GoGetter: "ConfiguredGateways"},
			_jsii_.MemberMethod{JsiiMethod: "configureNat", GoMethod: "ConfigureNat"},
			_jsii_.MemberMethod{JsiiMethod: "configureSubnet", GoMethod: "ConfigureSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
		},
		func() interface{} {
			j := jsiiProxy_NatInstanceProviderV2{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NatProvider)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConnectable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.NatProvider",
		reflect.TypeOf((*NatProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuredGateways", GoGetter: "ConfiguredGateways"},
			_jsii_.MemberMethod{JsiiMethod: "configureNat", GoMethod: "ConfigureNat"},
			_jsii_.MemberMethod{JsiiMethod: "configureSubnet", GoMethod: "ConfigureSubnet"},
		},
		func() interface{} {
			return &jsiiProxy_NatProvider{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.NatTrafficDirection",
		reflect.TypeOf((*NatTrafficDirection)(nil)).Elem(),
		map[string]interface{}{
			"OUTBOUND_ONLY": NatTrafficDirection_OUTBOUND_ONLY,
			"INBOUND_AND_OUTBOUND": NatTrafficDirection_INBOUND_AND_OUTBOUND,
			"NONE": NatTrafficDirection_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.NetworkAcl",
		reflect.TypeOf((*NetworkAcl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEntry", GoMethod: "AddEntry"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateWithSubnet", GoMethod: "AssociateWithSubnet"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "networkAclId", GoGetter: "NetworkAclId"},
			_jsii_.MemberProperty{JsiiProperty: "networkAclVpcId", GoGetter: "NetworkAclVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkAcl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkAcl)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.NetworkAclEntry",
		reflect.TypeOf((*NetworkAclEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkAclEntry{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INetworkAclEntry)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.NetworkAclEntryProps",
		reflect.TypeOf((*NetworkAclEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.NetworkAclProps",
		reflect.TypeOf((*NetworkAclProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.OperatingSystemType",
		reflect.TypeOf((*OperatingSystemType)(nil)).Elem(),
		map[string]interface{}{
			"LINUX": OperatingSystemType_LINUX,
			"WINDOWS": OperatingSystemType_WINDOWS,
			"UNKNOWN": OperatingSystemType_UNKNOWN,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Peer",
		reflect.TypeOf((*Peer)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Peer{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.PlacementGroup",
		reflect.TypeOf((*PlacementGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partitions", GoGetter: "Partitions"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroupName", GoGetter: "PlacementGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "spreadLevel", GoGetter: "SpreadLevel"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PlacementGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPlacementGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PlacementGroupProps",
		reflect.TypeOf((*PlacementGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.PlacementGroupSpreadLevel",
		reflect.TypeOf((*PlacementGroupSpreadLevel)(nil)).Elem(),
		map[string]interface{}{
			"HOST": PlacementGroupSpreadLevel_HOST,
			"RACK": PlacementGroupSpreadLevel_RACK,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.PlacementGroupStrategy",
		reflect.TypeOf((*PlacementGroupStrategy)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER": PlacementGroupStrategy_CLUSTER,
			"PARTITION": PlacementGroupStrategy_PARTITION,
			"SPREAD": PlacementGroupStrategy_SPREAD,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Port",
		reflect.TypeOf((*Port)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canInlineRule", GoGetter: "CanInlineRule"},
			_jsii_.MemberMethod{JsiiMethod: "toRuleJson", GoMethod: "ToRuleJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_Port{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PortProps",
		reflect.TypeOf((*PortProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.PrefixList",
		reflect.TypeOf((*PrefixList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addressFamily", GoGetter: "AddressFamily"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ownerId", GoGetter: "OwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "prefixListArn", GoGetter: "PrefixListArn"},
			_jsii_.MemberProperty{JsiiProperty: "prefixListId", GoGetter: "PrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "prefixListName", GoGetter: "PrefixListName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_PrefixList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrefixList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PrefixListOptions",
		reflect.TypeOf((*PrefixListOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PrefixListProps",
		reflect.TypeOf((*PrefixListProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.PrivateSubnet",
		reflect.TypeOf((*PrivateSubnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultInternetRoute", GoMethod: "AddDefaultInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addDefaultNatRoute", GoMethod: "AddDefaultNatRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6DefaultEgressOnlyInternetRoute", GoMethod: "AddIpv6DefaultEgressOnlyInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6DefaultInternetRoute", GoMethod: "AddIpv6DefaultInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6Nat64Route", GoMethod: "AddIpv6Nat64Route"},
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyElements", GoGetter: "DependencyElements"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetAvailabilityZone", GoGetter: "SubnetAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIpv6CidrBlocks", GoGetter: "SubnetIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "subnetNetworkAclAssociationId", GoGetter: "SubnetNetworkAclAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetOutpostArn", GoGetter: "SubnetOutpostArn"},
			_jsii_.MemberProperty{JsiiProperty: "subnetVpcId", GoGetter: "SubnetVpcId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PrivateSubnet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Subnet)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrivateSubnet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PrivateSubnetAttributes",
		reflect.TypeOf((*PrivateSubnetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PrivateSubnetProps",
		reflect.TypeOf((*PrivateSubnetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"ALL": Protocol_ALL,
			"HOPOPT": Protocol_HOPOPT,
			"ICMP": Protocol_ICMP,
			"IGMP": Protocol_IGMP,
			"GGP": Protocol_GGP,
			"IPV4": Protocol_IPV4,
			"ST": Protocol_ST,
			"TCP": Protocol_TCP,
			"CBT": Protocol_CBT,
			"EGP": Protocol_EGP,
			"IGP": Protocol_IGP,
			"BBN_RCC_MON": Protocol_BBN_RCC_MON,
			"NVP_II": Protocol_NVP_II,
			"PUP": Protocol_PUP,
			"EMCON": Protocol_EMCON,
			"XNET": Protocol_XNET,
			"CHAOS": Protocol_CHAOS,
			"UDP": Protocol_UDP,
			"MUX": Protocol_MUX,
			"DCN_MEAS": Protocol_DCN_MEAS,
			"HMP": Protocol_HMP,
			"PRM": Protocol_PRM,
			"XNS_IDP": Protocol_XNS_IDP,
			"TRUNK_1": Protocol_TRUNK_1,
			"TRUNK_2": Protocol_TRUNK_2,
			"LEAF_1": Protocol_LEAF_1,
			"LEAF_2": Protocol_LEAF_2,
			"RDP": Protocol_RDP,
			"IRTP": Protocol_IRTP,
			"ISO_TP4": Protocol_ISO_TP4,
			"NETBLT": Protocol_NETBLT,
			"MFE_NSP": Protocol_MFE_NSP,
			"MERIT_INP": Protocol_MERIT_INP,
			"DCCP": Protocol_DCCP,
			"THREEPC": Protocol_THREEPC,
			"IDPR": Protocol_IDPR,
			"XTP": Protocol_XTP,
			"DDP": Protocol_DDP,
			"IDPR_CMTP": Protocol_IDPR_CMTP,
			"TPPLUSPLUS": Protocol_TPPLUSPLUS,
			"IL": Protocol_IL,
			"IPV6": Protocol_IPV6,
			"SDRP": Protocol_SDRP,
			"IPV6_ROUTE": Protocol_IPV6_ROUTE,
			"IPV6_FRAG": Protocol_IPV6_FRAG,
			"IDRP": Protocol_IDRP,
			"RSVP": Protocol_RSVP,
			"GRE": Protocol_GRE,
			"DSR": Protocol_DSR,
			"BNA": Protocol_BNA,
			"ESP": Protocol_ESP,
			"AH": Protocol_AH,
			"I_NLSP": Protocol_I_NLSP,
			"SWIPE": Protocol_SWIPE,
			"NARP": Protocol_NARP,
			"MOBILE": Protocol_MOBILE,
			"TLSP": Protocol_TLSP,
			"SKIP": Protocol_SKIP,
			"ICMPV6": Protocol_ICMPV6,
			"IPV6_NONXT": Protocol_IPV6_NONXT,
			"IPV6_OPTS": Protocol_IPV6_OPTS,
			"CFTP": Protocol_CFTP,
			"ANY_LOCAL": Protocol_ANY_LOCAL,
			"SAT_EXPAK": Protocol_SAT_EXPAK,
			"KRYPTOLAN": Protocol_KRYPTOLAN,
			"RVD": Protocol_RVD,
			"IPPC": Protocol_IPPC,
			"ANY_DFS": Protocol_ANY_DFS,
			"SAT_MON": Protocol_SAT_MON,
			"VISA": Protocol_VISA,
			"IPCV": Protocol_IPCV,
			"CPNX": Protocol_CPNX,
			"CPHB": Protocol_CPHB,
			"WSN": Protocol_WSN,
			"PVP": Protocol_PVP,
			"BR_SAT_MON": Protocol_BR_SAT_MON,
			"SUN_ND": Protocol_SUN_ND,
			"WB_MON": Protocol_WB_MON,
			"WB_EXPAK": Protocol_WB_EXPAK,
			"ISO_IP": Protocol_ISO_IP,
			"VMTP": Protocol_VMTP,
			"SECURE_VMTP": Protocol_SECURE_VMTP,
			"VINES": Protocol_VINES,
			"TTP": Protocol_TTP,
			"IPTM": Protocol_IPTM,
			"NSFNET_IGP": Protocol_NSFNET_IGP,
			"DGP": Protocol_DGP,
			"TCF": Protocol_TCF,
			"EIGRP": Protocol_EIGRP,
			"OSPFIGP": Protocol_OSPFIGP,
			"SPRITE_RPC": Protocol_SPRITE_RPC,
			"LARP": Protocol_LARP,
			"MTP": Protocol_MTP,
			"AX_25": Protocol_AX_25,
			"IPIP": Protocol_IPIP,
			"MICP": Protocol_MICP,
			"SCC_SP": Protocol_SCC_SP,
			"ETHERIP": Protocol_ETHERIP,
			"ENCAP": Protocol_ENCAP,
			"ANY_ENC": Protocol_ANY_ENC,
			"GMTP": Protocol_GMTP,
			"IFMP": Protocol_IFMP,
			"PNNI": Protocol_PNNI,
			"PIM": Protocol_PIM,
			"ARIS": Protocol_ARIS,
			"SCPS": Protocol_SCPS,
			"QNX": Protocol_QNX,
			"A_N": Protocol_A_N,
			"IPCOMP": Protocol_IPCOMP,
			"SNP": Protocol_SNP,
			"COMPAQ_PEER": Protocol_COMPAQ_PEER,
			"IPX_IN_IP": Protocol_IPX_IN_IP,
			"VRRP": Protocol_VRRP,
			"PGM": Protocol_PGM,
			"ANY_0_HOP": Protocol_ANY_0_HOP,
			"L2_T_P": Protocol_L2_T_P,
			"DDX": Protocol_DDX,
			"IATP": Protocol_IATP,
			"STP": Protocol_STP,
			"SRP": Protocol_SRP,
			"UTI": Protocol_UTI,
			"SMP": Protocol_SMP,
			"SM": Protocol_SM,
			"PTP": Protocol_PTP,
			"ISIS_IPV4": Protocol_ISIS_IPV4,
			"FIRE": Protocol_FIRE,
			"CRTP": Protocol_CRTP,
			"CRUDP": Protocol_CRUDP,
			"SSCOPMCE": Protocol_SSCOPMCE,
			"IPLT": Protocol_IPLT,
			"SPS": Protocol_SPS,
			"PIPE": Protocol_PIPE,
			"SCTP": Protocol_SCTP,
			"FC": Protocol_FC,
			"RSVP_E2E_IGNORE": Protocol_RSVP_E2E_IGNORE,
			"MOBILITY_HEADER": Protocol_MOBILITY_HEADER,
			"UDPLITE": Protocol_UDPLITE,
			"MPLS_IN_IP": Protocol_MPLS_IN_IP,
			"MANET": Protocol_MANET,
			"HIP": Protocol_HIP,
			"SHIM6": Protocol_SHIM6,
			"WESP": Protocol_WESP,
			"ROHC": Protocol_ROHC,
			"ETHERNET": Protocol_ETHERNET,
			"EXPERIMENT_1": Protocol_EXPERIMENT_1,
			"EXPERIMENT_2": Protocol_EXPERIMENT_2,
			"RESERVED": Protocol_RESERVED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		reflect.TypeOf((*PublicSubnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultInternetRoute", GoMethod: "AddDefaultInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addDefaultNatRoute", GoMethod: "AddDefaultNatRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6DefaultEgressOnlyInternetRoute", GoMethod: "AddIpv6DefaultEgressOnlyInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6DefaultInternetRoute", GoMethod: "AddIpv6DefaultInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6Nat64Route", GoMethod: "AddIpv6Nat64Route"},
			_jsii_.MemberMethod{JsiiMethod: "addNatGateway", GoMethod: "AddNatGateway"},
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyElements", GoGetter: "DependencyElements"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetAvailabilityZone", GoGetter: "SubnetAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIpv6CidrBlocks", GoGetter: "SubnetIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "subnetNetworkAclAssociationId", GoGetter: "SubnetNetworkAclAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetOutpostArn", GoGetter: "SubnetOutpostArn"},
			_jsii_.MemberProperty{JsiiProperty: "subnetVpcId", GoGetter: "SubnetVpcId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PublicSubnet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Subnet)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPublicSubnet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PublicSubnetAttributes",
		reflect.TypeOf((*PublicSubnetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.PublicSubnetProps",
		reflect.TypeOf((*PublicSubnetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.RequestedSubnet",
		reflect.TypeOf((*RequestedSubnet)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.ResolveSsmParameterAtLaunchImage",
		reflect.TypeOf((*ResolveSsmParameterAtLaunchImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
		},
		func() interface{} {
			j := jsiiProxy_ResolveSsmParameterAtLaunchImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.RouterType",
		reflect.TypeOf((*RouterType)(nil)).Elem(),
		map[string]interface{}{
			"CARRIER_GATEWAY": RouterType_CARRIER_GATEWAY,
			"EGRESS_ONLY_INTERNET_GATEWAY": RouterType_EGRESS_ONLY_INTERNET_GATEWAY,
			"GATEWAY": RouterType_GATEWAY,
			"INSTANCE": RouterType_INSTANCE,
			"LOCAL_GATEWAY": RouterType_LOCAL_GATEWAY,
			"NAT_GATEWAY": RouterType_NAT_GATEWAY,
			"NETWORK_INTERFACE": RouterType_NETWORK_INTERFACE,
			"TRANSIT_GATEWAY": RouterType_TRANSIT_GATEWAY,
			"VPC_PEERING_CONNECTION": RouterType_VPC_PEERING_CONNECTION,
			"VPC_ENDPOINT": RouterType_VPC_ENDPOINT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.RuleScope",
		reflect.TypeOf((*RuleScope)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.S3DestinationOptions",
		reflect.TypeOf((*S3DestinationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.S3DownloadOptions",
		reflect.TypeOf((*S3DownloadOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		reflect.TypeOf((*SecurityGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEgressRule", GoMethod: "AddEgressRule"},
			_jsii_.MemberMethod{JsiiMethod: "addIngressRule", GoMethod: "AddIngressRule"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllIpv6Outbound", GoGetter: "AllowAllIpv6Outbound"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllOutbound", GoGetter: "AllowAllOutbound"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "canInlineRule", GoGetter: "CanInlineRule"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPort", GoGetter: "DefaultPort"},
			_jsii_.MemberMethod{JsiiMethod: "determineRuleScope", GoMethod: "DetermineRuleScope"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupId", GoGetter: "SecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupVpcId", GoGetter: "SecurityGroupVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toEgressRuleConfig", GoMethod: "ToEgressRuleConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toIngressRuleConfig", GoMethod: "ToIngressRuleConfig"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueId", GoGetter: "UniqueId"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISecurityGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SecurityGroupImportOptions",
		reflect.TypeOf((*SecurityGroupImportOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SecurityGroupProps",
		reflect.TypeOf((*SecurityGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SelectedSubnets",
		reflect.TypeOf((*SelectedSubnets)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.ServiceManager",
		reflect.TypeOf((*ServiceManager)(nil)).Elem(),
		map[string]interface{}{
			"SYSVINIT": ServiceManager_SYSVINIT,
			"WINDOWS": ServiceManager_WINDOWS,
			"SYSTEMD": ServiceManager_SYSTEMD,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.SpotInstanceInterruption",
		reflect.TypeOf((*SpotInstanceInterruption)(nil)).Elem(),
		map[string]interface{}{
			"STOP": SpotInstanceInterruption_STOP,
			"TERMINATE": SpotInstanceInterruption_TERMINATE,
			"HIBERNATE": SpotInstanceInterruption_HIBERNATE,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.SpotRequestType",
		reflect.TypeOf((*SpotRequestType)(nil)).Elem(),
		map[string]interface{}{
			"ONE_TIME": SpotRequestType_ONE_TIME,
			"PERSISTENT": SpotRequestType_PERSISTENT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SsmParameterImageOptions",
		reflect.TypeOf((*SsmParameterImageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Subnet",
		reflect.TypeOf((*Subnet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultInternetRoute", GoMethod: "AddDefaultInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addDefaultNatRoute", GoMethod: "AddDefaultNatRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6DefaultEgressOnlyInternetRoute", GoMethod: "AddIpv6DefaultEgressOnlyInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6DefaultInternetRoute", GoMethod: "AddIpv6DefaultInternetRoute"},
			_jsii_.MemberMethod{JsiiMethod: "addIpv6Nat64Route", GoMethod: "AddIpv6Nat64Route"},
			_jsii_.MemberMethod{JsiiMethod: "addRoute", GoMethod: "AddRoute"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateNetworkAcl", GoMethod: "AssociateNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyElements", GoGetter: "DependencyElements"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4CidrBlock", GoGetter: "Ipv4CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeTable", GoGetter: "RouteTable"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnetAvailabilityZone", GoGetter: "SubnetAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIpv6CidrBlocks", GoGetter: "SubnetIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "subnetNetworkAclAssociationId", GoGetter: "SubnetNetworkAclAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetOutpostArn", GoGetter: "SubnetOutpostArn"},
			_jsii_.MemberProperty{JsiiProperty: "subnetVpcId", GoGetter: "SubnetVpcId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Subnet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubnet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SubnetAttributes",
		reflect.TypeOf((*SubnetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SubnetConfiguration",
		reflect.TypeOf((*SubnetConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.SubnetFilter",
		reflect.TypeOf((*SubnetFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
		},
		func() interface{} {
			return &jsiiProxy_SubnetFilter{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SubnetIpamOptions",
		reflect.TypeOf((*SubnetIpamOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.SubnetNetworkAclAssociation",
		reflect.TypeOf((*SubnetNetworkAclAssociation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "networkAcl", GoGetter: "NetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnet", GoGetter: "Subnet"},
			_jsii_.MemberProperty{JsiiProperty: "subnetNetworkAclAssociationAssociationId", GoGetter: "SubnetNetworkAclAssociationAssociationId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SubnetNetworkAclAssociation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubnetNetworkAclAssociation)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SubnetNetworkAclAssociationProps",
		reflect.TypeOf((*SubnetNetworkAclAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SubnetProps",
		reflect.TypeOf((*SubnetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SubnetSelection",
		reflect.TypeOf((*SubnetSelection)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.SubnetType",
		reflect.TypeOf((*SubnetType)(nil)).Elem(),
		map[string]interface{}{
			"PRIVATE_ISOLATED": SubnetType_PRIVATE_ISOLATED,
			"PRIVATE_WITH_EGRESS": SubnetType_PRIVATE_WITH_EGRESS,
			"PRIVATE_WITH_NAT": SubnetType_PRIVATE_WITH_NAT,
			"PUBLIC": SubnetType_PUBLIC,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.SystemdConfigFileOptions",
		reflect.TypeOf((*SystemdConfigFileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.TrafficDirection",
		reflect.TypeOf((*TrafficDirection)(nil)).Elem(),
		map[string]interface{}{
			"EGRESS": TrafficDirection_EGRESS,
			"INGRESS": TrafficDirection_INGRESS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.TransportProtocol",
		reflect.TypeOf((*TransportProtocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": TransportProtocol_TCP,
			"UDP": TransportProtocol_UDP,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.UserData",
		reflect.TypeOf((*UserData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCommands", GoMethod: "AddCommands"},
			_jsii_.MemberMethod{JsiiMethod: "addExecuteFileCommand", GoMethod: "AddExecuteFileCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addOnExitCommands", GoMethod: "AddOnExitCommands"},
			_jsii_.MemberMethod{JsiiMethod: "addS3DownloadCommand", GoMethod: "AddS3DownloadCommand"},
			_jsii_.MemberMethod{JsiiMethod: "addSignalOnExitCommand", GoMethod: "AddSignalOnExitCommand"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_UserData{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantAttachVolume", GoMethod: "GrantAttachVolume"},
			_jsii_.MemberMethod{JsiiMethod: "grantAttachVolumeByResourceTag", GoMethod: "GrantAttachVolumeByResourceTag"},
			_jsii_.MemberMethod{JsiiMethod: "grantDetachVolume", GoMethod: "GrantDetachVolume"},
			_jsii_.MemberMethod{JsiiMethod: "grantDetachVolumeByResourceTag", GoMethod: "GrantDetachVolumeByResourceTag"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateProps", GoMethod: "ValidateProps"},
			_jsii_.MemberProperty{JsiiProperty: "volumeId", GoGetter: "VolumeId"},
		},
		func() interface{} {
			j := jsiiProxy_Volume{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VolumeAttributes",
		reflect.TypeOf((*VolumeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VolumeProps",
		reflect.TypeOf((*VolumeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.Vpc",
		reflect.TypeOf((*Vpc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addClientVpnEndpoint", GoMethod: "AddClientVpnEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addFlowLog", GoMethod: "AddFlowLog"},
			_jsii_.MemberMethod{JsiiMethod: "addGatewayEndpoint", GoMethod: "AddGatewayEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addInterfaceEndpoint", GoMethod: "AddInterfaceEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addVpnConnection", GoMethod: "AddVpnConnection"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "dnsHostnamesEnabled", GoGetter: "DnsHostnamesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSupportEnabled", GoGetter: "DnsSupportEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "enableVpnGateway", GoMethod: "EnableVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "incompleteSubnetDefinition", GoGetter: "IncompleteSubnetDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "internetConnectivityEstablished", GoGetter: "InternetConnectivityEstablished"},
			_jsii_.MemberProperty{JsiiProperty: "internetGatewayId", GoGetter: "InternetGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "isolatedSubnets", GoGetter: "IsolatedSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnets", GoGetter: "PrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnets", GoGetter: "PublicSubnets"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnetObjects", GoMethod: "SelectSubnetObjects"},
			_jsii_.MemberMethod{JsiiMethod: "selectSubnets", GoMethod: "SelectSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcArn", GoGetter: "VpcArn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlock", GoGetter: "VpcCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlockAssociations", GoGetter: "VpcCidrBlockAssociations"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultNetworkAcl", GoGetter: "VpcDefaultNetworkAcl"},
			_jsii_.MemberProperty{JsiiProperty: "vpcDefaultSecurityGroup", GoGetter: "VpcDefaultSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIpv6CidrBlocks", GoGetter: "VpcIpv6CidrBlocks"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayId", GoGetter: "VpnGatewayId"},
		},
		func() interface{} {
			j := jsiiProxy_Vpc{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpc)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpcAttributes",
		reflect.TypeOf((*VpcAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.VpcEndpoint",
		reflect.TypeOf((*VpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
		},
		func() interface{} {
			j := jsiiProxy_VpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcEndpoint)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.VpcEndpointService",
		reflect.TypeOf((*VpcEndpointService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptanceRequired", GoGetter: "AcceptanceRequired"},
			_jsii_.MemberProperty{JsiiProperty: "allowedPrincipals", GoGetter: "AllowedPrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "contributorInsightsEnabled", GoGetter: "ContributorInsightsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceId", GoGetter: "VpcEndpointServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceLoadBalancers", GoGetter: "VpcEndpointServiceLoadBalancers"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceName", GoGetter: "VpcEndpointServiceName"},
		},
		func() interface{} {
			j := jsiiProxy_VpcEndpointService{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpcEndpointService)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpcEndpointServiceProps",
		reflect.TypeOf((*VpcEndpointServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.VpcEndpointType",
		reflect.TypeOf((*VpcEndpointType)(nil)).Elem(),
		map[string]interface{}{
			"INTERFACE": VpcEndpointType_INTERFACE,
			"GATEWAY": VpcEndpointType_GATEWAY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpcIpamOptions",
		reflect.TypeOf((*VpcIpamOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpcLookupOptions",
		reflect.TypeOf((*VpcLookupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpcProps",
		reflect.TypeOf((*VpcProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.VpnConnection",
		reflect.TypeOf((*VpnConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayAsn", GoGetter: "CustomerGatewayAsn"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayIp", GoGetter: "CustomerGatewayIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataIn", GoMethod: "MetricTunnelDataIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataOut", GoMethod: "MetricTunnelDataOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelState", GoMethod: "MetricTunnelState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpnId", GoGetter: "VpnId"},
		},
		func() interface{} {
			j := jsiiProxy_VpnConnection{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VpnConnectionBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpnConnectionAttributes",
		reflect.TypeOf((*VpnConnectionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.VpnConnectionBase",
		reflect.TypeOf((*VpnConnectionBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayAsn", GoGetter: "CustomerGatewayAsn"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayId", GoGetter: "CustomerGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "customerGatewayIp", GoGetter: "CustomerGatewayIp"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataIn", GoMethod: "MetricTunnelDataIn"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelDataOut", GoMethod: "MetricTunnelDataOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricTunnelState", GoMethod: "MetricTunnelState"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpnId", GoGetter: "VpnId"},
		},
		func() interface{} {
			j := jsiiProxy_VpnConnectionBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpnConnection)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpnConnectionOptions",
		reflect.TypeOf((*VpnConnectionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpnConnectionProps",
		reflect.TypeOf((*VpnConnectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.VpnConnectionType",
		reflect.TypeOf((*VpnConnectionType)(nil)).Elem(),
		map[string]interface{}{
			"IPSEC_1": VpnConnectionType_IPSEC_1,
			"DUMMY": VpnConnectionType_DUMMY,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.VpnGateway",
		reflect.TypeOf((*VpnGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gatewayId", GoGetter: "GatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VpnGateway{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVpnGateway)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpnGatewayProps",
		reflect.TypeOf((*VpnGatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.VpnPort",
		reflect.TypeOf((*VpnPort)(nil)).Elem(),
		map[string]interface{}{
			"HTTPS": VpnPort_HTTPS,
			"OPENVPN": VpnPort_OPENVPN,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.VpnTunnelOption",
		reflect.TypeOf((*VpnTunnelOption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_ec2.WindowsImage",
		reflect.TypeOf((*WindowsImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
			_jsii_.MemberProperty{JsiiProperty: "parameterName", GoGetter: "ParameterName"},
		},
		func() interface{} {
			j := jsiiProxy_WindowsImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GenericSSMParameterImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.WindowsImageProps",
		reflect.TypeOf((*WindowsImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_ec2.WindowsUserDataOptions",
		reflect.TypeOf((*WindowsUserDataOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_ec2.WindowsVersion",
		reflect.TypeOf((*WindowsVersion)(nil)).Elem(),
		map[string]interface{}{
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_22": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_22,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_22": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_22,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_23": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_23,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_23": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_23,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_24": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_24,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_24": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_24,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_25": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_25,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_25": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_25,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_26": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_26,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_26": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_26,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_27": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_27,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_27": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_27,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_28": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_EKS_OPTIMIZED_1_28,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_28": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_EKS_OPTIMIZED_1_28,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_23": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_23,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_23": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_23,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_24": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_24,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_24": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_24,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_25": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_25,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_25": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_25,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_26": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_26,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_26": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_26,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_27": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_27,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_27": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_27,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_28": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_EKS_OPTIMIZED_1_28,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_28": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_EKS_OPTIMIZED_1_28,
			"WINDOWS_SERVER_1709_ENGLISH_CORE_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_1709_ENGLISH_CORE_CONTAINERSLATEST,
			"WINDOWS_SERVER_1709_ENGLISH_CORE_BASE": WindowsVersion_WINDOWS_SERVER_1709_ENGLISH_CORE_BASE,
			"WINDOWS_SERVER_1803_ENGLISH_CORE_BASE": WindowsVersion_WINDOWS_SERVER_1803_ENGLISH_CORE_BASE,
			"WINDOWS_SERVER_1803_ENGLISH_CORE_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_1803_ENGLISH_CORE_CONTAINERSLATEST,
			"WINDOWS_SERVER_1809_ENGLISH_CORE_BASE": WindowsVersion_WINDOWS_SERVER_1809_ENGLISH_CORE_BASE,
			"WINDOWS_SERVER_1809_ENGLISH_CORE_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_1809_ENGLISH_CORE_CONTAINERSLATEST,
			"WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_32BIT_BASE": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_32BIT_BASE,
			"WINDOWS_SERVER_2003_R2_SP2_ENGLISH_64BIT_SQL_2005_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_ENGLISH_64BIT_SQL_2005_SP4_EXPRESS,
			"WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_64BIT_SQL_2005_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_64BIT_SQL_2005_SP4_STANDARD,
			"WINDOWS_SERVER_2003_R2_SP2_ENGLISH_32BIT_BASE": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_ENGLISH_32BIT_BASE,
			"WINDOWS_SERVER_2003_R2_SP2_ENGLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_ENGLISH_64BIT_BASE,
			"WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_64BIT_SQL_2005_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_64BIT_SQL_2005_SP4_EXPRESS,
			"WINDOWS_SERVER_2003_R2_SP2_ENGLISH_64BIT_SQL_2005_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_ENGLISH_64BIT_SQL_2005_SP4_STANDARD,
			"WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2003_R2_SP2_LANGUAGE_PACKS_64BIT_BASE,
			"WINDOWS_SERVER_2007_R2_SP1_LANGUAGE_PACKS_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2007_R2_SP1_LANGUAGE_PACKS_64BIT_BASE,
			"WINDOWS_SERVER_2008_SP2_ENGLISH_64BIT_SQL_2008_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2008_SP2_ENGLISH_64BIT_SQL_2008_SP4_EXPRESS,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2008_R2_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2008_R2_SP3_WEB,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_EXPRESS,
			"WINDOWS_SERVER_2008_R2_SP1_KOREAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_KOREAN_64BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_CHINESE_HONG_KONG_SAR_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_CHINESE_HONG_KONG_SAR_64BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_CHINESE_PRC_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_CHINESE_PRC_64BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2008_R2_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2008_R2_SP3_EXPRESS,
			"WINDOWS_SERVER_2008_SP2_ENGLISH_32BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_SP2_ENGLISH_32BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_WEB": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_WEB,
			"WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_BASE,
			"WINDOWS_SERVER_2008_SP2_ENGLISH_64BIT_SQL_2008_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_SP2_ENGLISH_64BIT_SQL_2008_SP4_STANDARD,
			"WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2012_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2012_SP4_EXPRESS,
			"WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2008_R2_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2008_R2_SP3_WEB,
			"WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2012_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2012_SP4_STANDARD,
			"WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2008_R2_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2008_R2_SP3_STANDARD,
			"WINDOWS_SERVER_2008_SP2_ENGLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_SP2_ENGLISH_64BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_ENTERPRISE,
			"WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2008_R2_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_JAPANESE_64BIT_SQL_2008_R2_SP3_EXPRESS,
			"WINDOWS_SERVER_2008_R2_SP1_PORTUGUESE_BRAZIL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_PORTUGUESE_BRAZIL_64BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_LANGUAGE_PACKS_64BIT_SQL_2008_R2_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_LANGUAGE_PACKS_64BIT_SQL_2008_R2_SP3_STANDARD,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_61BIT_SQL_2012_RTM_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_61BIT_SQL_2012_RTM_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2012_SP4_STANDARD,
			"WINDOWS_SERVER_2008_SP2_PORTUGUESE_BRAZIL_32BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_SP2_PORTUGUESE_BRAZIL_32BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_BASE,
			"WINDOWS_SERVER_2008_R2_SP1_LANGUAGE_PACKS_64BIT_SQL_2008_R2_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_LANGUAGE_PACKS_64BIT_SQL_2008_R2_SP3_EXPRESS,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_CORE_SQL_2012_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_CORE_SQL_2012_SP4_STANDARD,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_CORE": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_CORE,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2008_R2_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SQL_2008_R2_SP3_STANDARD,
			"WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SHAREPOINT_2010_SP2_FOUNDATION": WindowsVersion_WINDOWS_SERVER_2008_R2_SP1_ENGLISH_64BIT_SHAREPOINT_2010_SP2_FOUNDATION,
			"WINDOWS_SERVER_2012_R2_RTM_CHINESE_SIMPLIFIED_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_CHINESE_SIMPLIFIED_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_CHINESE_TRADITIONAL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_CHINESE_TRADITIONAL_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_DUTCH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_DUTCH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2012_R2_RTM_HUNGARIAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_HUNGARIAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2008_R2_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2008_R2_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_SP1_PORTUGUESE_BRAZIL_64BIT_CORE": WindowsVersion_WINDOWS_SERVER_2012_R2_SP1_PORTUGUESE_BRAZIL_64BIT_CORE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP2_EXPRESS,
			"WINDOWS_SERVER_2012_RTM_ITALIAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_ITALIAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP2_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_DEEP_LEARNING": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_DEEP_LEARNING,
			"WINDOWS_SERVER_2012_R2_RTM_GERMAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_GERMAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_RUSSIAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_RUSSIAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_CHINESE_TRADITIONAL_HONG_KONG_SAR_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_CHINESE_TRADITIONAL_HONG_KONG_SAR_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_HUNGARIAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_HUNGARIAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP3_STANDARD,
			"WINDOWS_SERVER_2012_RTM_FRENCH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_FRENCH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_FRENCH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_FRENCH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_POLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_POLISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2012_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2012_SP4_EXPRESS,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP3_STANDARD,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_2012_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_2012_SP4_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_CORE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_CORE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_2014_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_2014_SP3_WEB,
			"WINDOWS_SERVER_2012_RTM_SWEDISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_SWEDISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_PORTUGUESE_BRAZIL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_PORTUGUESE_BRAZIL_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_PORTUGUESE_PORTUGAL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_PORTUGUESE_PORTUGAL_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_SWEDISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_SWEDISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2012_SP4_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2012_SP4_ENTERPRISE,
			"WINDOWS_SERVER_2012_RTM_CHINESE_TRADITIONAL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_CHINESE_TRADITIONAL_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP2_STANDARD,
			"WINDOWS_SERVER_2012_RTM_CZECH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_CZECH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_TURKISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_TURKISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_HYPERV": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_HYPERV,
			"WINDOWS_SERVER_2012_RTM_KOREAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_KOREAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_RUSSIAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_RUSSIAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_FULL_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP2_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_ITALIAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ITALIAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2008_R2_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2008_R2_SP3_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_STANDARD,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2007_R2_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2007_R2_SP3_WEB,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP2_WEB,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP2_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_CZECH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_CZECH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP2_EXPRESS,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2012_SP4_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2012_SP4_STANDARD,
			"WINDOWS_SERVER_2012_SP2_PORTUGUESE_BRAZIL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_SP2_PORTUGUESE_BRAZIL_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP1_ENTERPRISE,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2016_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2016_SP2_EXPRESS,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP2_STANDARD,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_FULL_BASE,
			"WINDOWS_SERVER_2012_R2_ENGLISH_STIG_FULL": WindowsVersion_WINDOWS_SERVER_2012_R2_ENGLISH_STIG_FULL,
			"WINDOWS_SERVER_2012_RTM_PORTUGUESE_PORTUGAL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_PORTUGUESE_PORTUGAL_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP1_ENTERPRISE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2014_SP2_WEB,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2008_R2_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2008_R2_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_PORTUGUESE_BRAZIL_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_PORTUGUESE_BRAZIL_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_P3": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_P3,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP3_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_SPANISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_SPANISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2012_R2_ENGLISH_STIG_CORE": WindowsVersion_WINDOWS_SERVER_2012_R2_ENGLISH_STIG_CORE,
			"WINDOWS_SERVER_2012_R2_RTM_TURKISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_TURKISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2012_SP4_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2012_SP4_WEB,
			"WINDOWS_SERVER_2012_RTM_POLISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_POLISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_SPANISH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_SPANISH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_KOREAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_KOREAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_DUTCH_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_DUTCH_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_CHINESE_TRADITIONAL_HONG_KONG_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_CHINESE_TRADITIONAL_HONG_KONG_64BIT_BASE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2014_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2012_RTM_CHINESE_SIMPLIFIED_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_CHINESE_SIMPLIFIED_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2012_SP4_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_ENGLISH_64BIT_SQL_2012_SP4_WEB,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2014_SP3_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_WEB,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2014_SP2_STANDARD,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2012_SP4_EXPRESS": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2012_SP4_EXPRESS,
			"WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_JAPANESE_64BIT_SQL_2016_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2012_R2_RTM_ENGLISH_64BIT_SQL_2016_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_CONTAINERS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_CONTAINERS,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_WEB,
			"WINDOWS_SERVER_2016_GERMAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_GERMAN_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_EXPRESS,
			"WINDOWS_SERVER_2016_ENGLISH_DEEP_LEARNING": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_DEEP_LEARNING,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_FQL_2016_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_FQL_2016_SP2_WEB,
			"WINDOWS_SERVER_2016_KOREAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_KOREAN_FULL_BASE,
			"WINDOWS_SERVER_2016_KOREAN_FULL_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_KOREAN_FULL_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2016_POLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_POLISH_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_CONTAINERS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_CONTAINERS,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_STANDARD,
			"WINDOWS_SERVER_2016_RUSSIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_RUSSIAN_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_EXPRESS,
			"WINDOWS_SERVER_2016_ITALIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_ITALIAN_FULL_BASE,
			"WINDOWS_SERVER_2016_SPANISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_SPANISH_FULL_BASE,
			"WINDOWS_SERVER_2012_RTM_GERMAN_64BIT_BASE": WindowsVersion_WINDOWS_SERVER_2012_RTM_GERMAN_64BIT_BASE,
			"WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2008_R2_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2012_RTM_JAPANESE_64BIT_SQL_2008_R2_SP3_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_HYPERV": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_HYPERV,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_CONTAINERSLATEST,
			"WINDOWS_SERVER_2016_DUTCH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_DUTCH_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_EXPRESS,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_ENTERPRISE,
			"WINDOWS_SERVER_2016_HUNGARIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_HUNGARIAN_FULL_BASE,
			"WINDOWS_SERVER_2016_KOREAN_FULL_SQL_2016_SP1_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_KOREAN_FULL_SQL_2016_SP1_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_BASE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_ENTERPRISE,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_WEB": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_WEB,
			"WINDOWS_SERVER_2016_SWEDISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_SWEDISH_FULL_BASE,
			"WINDOWS_SERVER_2016_TURKISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_TURKISH_FULL_BASE,
			"WINDOWS_SERVER_2016_PORTUGUESE_BRAZIL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_PORTUGUESE_BRAZIL_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2014_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2014_SP3_STANDARD,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_64BIT_SQL_2012_SP4_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_64BIT_SQL_2012_SP4_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP1_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_EXPRESS,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_STANDARD,
			"WINDOWS_SERVER_2016_PORTUGUESE_PORTUGAL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_PORTUGUESE_PORTUGAL_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2014_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2014_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_ENTERPRISE,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2017_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2017_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_EXPRESS,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2017_WEB": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2017_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_STIG_CORE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_STIG_CORE,
			"WINDOWS_SERVER_2016_KOREAN_FULL_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_KOREAN_FULL_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_ECS_OPTIMIZED": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_ECS_OPTIMIZED,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2017_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2017_ENTERPRISE,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP2_EXPRESS,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2019_WEB": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2019_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_P3": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_P3,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_ENTERPRISE,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_BASE,
			"WINDOWS_SERVER_2016_CHINESE_SIMPLIFIED_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_CHINESE_SIMPLIFIED_FULL_BASE,
			"WINDOWS_SERVER_2016_FRENCH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_FRENCH_FULL_BASE,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2016_CZECH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_CZECH_FULL_BASE,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2016_CHINESE_TRADITIONAL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2016_CHINESE_TRADITIONAL_FULL_BASE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP2_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2017_EXPRESS,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_WEB": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2019_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_EXPRESS,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2019_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2019_STANDARD,
			"WINDOWS_SERVER_2016_ENGLISH_TESLA": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_TESLA,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_EXPRESS,
			"WINDOWS_SERVER_2016_ENGLISH_STIG_FULL": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_STIG_FULL,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2019_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2019_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP3_EXPRESS,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP3_WEB,
			"WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_CORE_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_EXPRESS": WindowsVersion_WINDOWS_SERVER_2016_JAPANESE_FULL_SQL_2016_SP1_EXPRESS,
			"WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2016_ENGLISH_FULL_SQL_2016_SP1_ENTERPRISE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_WEB": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_WEB,
			"WINDOWS_SERVER_2019_FRENCH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_FRENCH_FULL_BASE,
			"WINDOWS_SERVER_2019_KOREAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_KOREAN_FULL_BASE,
			"WINDOWS_SERVER_2019_ITALIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_ITALIAN_FULL_BASE,
			"WINDOWS_SERVER_2019_CHINESE_SIMPLIFIED_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_CHINESE_SIMPLIFIED_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_WEB": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_WEB,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_HYPERV": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_HYPERV,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_STANDARD,
			"WINDOWS_SERVER_2019_HUNGARIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_HUNGARIAN_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_EXPRESS": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_EXPRESS,
			"WINDOWS_SERVER_2019_TURKISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_TURKISH_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_STANDARD,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_STANDARD,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_CONTAINERSLATEST,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_EXPRESS": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_EXPRESS,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_BASE,
			"WINDOWS_SERVER_2019_RUSSIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_RUSSIAN_FULL_BASE,
			"WINDOWS_SERVER_2019_CHINESE_TRADITIONAL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_CHINESE_TRADITIONAL_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_BASE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2022_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2022_STANDARD,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2022_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2022_ENTERPRISE,
			"WINDOWS_SERVER_2019_ENGLISH_TESLA": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_TESLA,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_ENTERPRISE,
			"WINDOWS_SERVER_2019_SPANISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_SPANISH_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_ENTERPRISE,
			"WINDOWS_SERVER_2019_ENGLISH_STIG_FULL": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_STIG_FULL,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_WEB": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_WEB,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_STANDARD,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2017_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2017_ENTERPRISE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP2_ENTERPRISE,
			"WINDOWS_SERVER_2019_PORTUGUESE_PORTUGAL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_PORTUGUESE_PORTUGAL_FULL_BASE,
			"WINDOWS_SERVER_2019_SWEDISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_SWEDISH_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_EXPRESS": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_EXPRESS,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2022_WEB": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2022_WEB,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_ENTERPRISE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_WEB": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_WEB,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_WEB": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_WEB,
			"WINDOWS_SERVER_2019_PORTUGUESE_BRAZIL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_PORTUGUESE_BRAZIL_FULL_BASE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_CONTAINERSLATEST,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2017_ENTERPRISE,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2019_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2019_ENTERPRISE,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_EXPRESS": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2019_EXPRESS,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2017_WEB": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2017_WEB,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_EXPRESS": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2016_SP3_EXPRESS,
			"WINDOWS_SERVER_2019_ENGLISH_STIG_CORE": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_STIG_CORE,
			"WINDOWS_SERVER_2019_ENGLISH_CORE_ECS_OPTIMIZED": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_CORE_ECS_OPTIMIZED,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_SQL_2022_STANDARD,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2017_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2017_STANDARD,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2019_WEB": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2019_WEB,
			"WINDOWS_SERVER_2019_ENGLISH_FULL_ECS_OPTIMIZED": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_ECS_OPTIMIZED,
			"WINDOWS_SERVER_2019_ENGLISH_DEEP_LEARNING": WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_DEEP_LEARNING,
			"WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2019_STANDARD": WindowsVersion_WINDOWS_SERVER_2019_JAPANESE_FULL_SQL_2019_STANDARD,
			"WINDOWS_SERVER_2019_CZECH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_CZECH_FULL_BASE,
			"WINDOWS_SERVER_2019_POLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_POLISH_FULL_BASE,
			"WINDOWS_SERVER_2019_GERMAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_GERMAN_FULL_BASE,
			"WINDOWS_SERVER_2019_DUTCH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2019_DUTCH_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_STIG_FULL": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_STIG_FULL,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2022_WEB": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2022_WEB,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_WEB": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_WEB,
			"WINDOWS_SERVER_2022_ENGLISH_STIG_CORE": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_STIG_CORE,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2019_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2019_ENTERPRISE,
			"WINDOWS_SERVER_2022_PORTUGUESE_BRAZIL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_PORTUGUESE_BRAZIL_FULL_BASE,
			"WINDOWS_SERVER_2022_ITALIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_ITALIAN_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_CONTAINERSLATEST,
			"WINDOWS_SERVER_2022_RUSSIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_RUSSIAN_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_EXPRESS": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_EXPRESS,
			"WINDOWS_SERVER_2022_POLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_POLISH_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_BASE": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_BASE,
			"WINDOWS_SERVER_2022_HUNGARIAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_HUNGARIAN_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_EXPRESS": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_EXPRESS,
			"WINDOWS_SERVER_2022_GERMAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_GERMAN_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_CONTAINERSLATEST": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_CONTAINERSLATEST,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_STANDARD": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_STANDARD,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2017_WEB": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2017_WEB,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_WEB": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_WEB,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_BASE,
			"WINDOWS_SERVER_2022_KOREAN_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_KOREAN_FULL_BASE,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2017_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2017_ENTERPRISE,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2019_STANDARD": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2019_STANDARD,
			"WINDOWS_SERVER_2022_CHINESE_SIMPLIFIED_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_CHINESE_SIMPLIFIED_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_WEB": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_WEB,
			"WINDOWS_SERVER_2022_SPANISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_SPANISH_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_CORE_ECS_OPTIMIZED": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_CORE_ECS_OPTIMIZED,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_STANDARD": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_STANDARD,
			"WINDOWS_SERVER_2022_CHINESE_TRADITIONAL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_CHINESE_TRADITIONAL_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2019_ENTERPRISE,
			"WINDOWS_SERVER_2022_FRENCH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_FRENCH_FULL_BASE,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2017_STANDARD": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2017_STANDARD,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_BASE,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2019_WEB": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2019_WEB,
			"WINDOWS_SERVER_2022_TURKISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_TURKISH_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2017_ENTERPRISE,
			"WINDOWS_SERVER_2022_PORTUGUESE_PORTUGAL_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_PORTUGUESE_PORTUGAL_FULL_BASE,
			"WINDOWS_SERVER_2022_CZECH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_CZECH_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_ECS_OPTIMIZED": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_ECS_OPTIMIZED,
			"WINDOWS_SERVER_2022_DUTCH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_DUTCH_FULL_BASE,
			"WINDOWS_SERVER_2022_SWEDISH_FULL_BASE": WindowsVersion_WINDOWS_SERVER_2022_SWEDISH_FULL_BASE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_ENTERPRISE,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_EXPRESS": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_EXPRESS,
			"WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_STANDARD": WindowsVersion_WINDOWS_SERVER_2022_ENGLISH_FULL_SQL_2022_STANDARD,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2022_STANDARD": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2022_STANDARD,
			"WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2022_ENTERPRISE": WindowsVersion_WINDOWS_SERVER_2022_JAPANESE_FULL_SQL_2022_ENTERPRISE,
		},
	)
}
