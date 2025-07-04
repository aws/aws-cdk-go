// The CDK Construct Library for AWS::EKS
package awscdkeks-v2alpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.AccessEntry",
		reflect.TypeOf((*AccessEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessEntryArn", GoGetter: "AccessEntryArn"},
			_jsii_.MemberProperty{JsiiProperty: "accessEntryName", GoGetter: "AccessEntryName"},
			_jsii_.MemberMethod{JsiiMethod: "addAccessPolicies", GoMethod: "AddAccessPolicies"},
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
			j := jsiiProxy_AccessEntry{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAccessEntry)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AccessEntryAttributes",
		reflect.TypeOf((*AccessEntryAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AccessEntryProps",
		reflect.TypeOf((*AccessEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.AccessEntryType",
		reflect.TypeOf((*AccessEntryType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": AccessEntryType_STANDARD,
			"FARGATE_LINUX": AccessEntryType_FARGATE_LINUX,
			"EC2_LINUX": AccessEntryType_EC2_LINUX,
			"EC2_WINDOWS": AccessEntryType_EC2_WINDOWS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicy",
		reflect.TypeOf((*AccessPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessScope", GoGetter: "AccessScope"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
		},
		func() interface{} {
			j := jsiiProxy_AccessPolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAccessPolicy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyArn",
		reflect.TypeOf((*AccessPolicyArn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "policyArn", GoGetter: "PolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyName", GoGetter: "PolicyName"},
		},
		func() interface{} {
			return &jsiiProxy_AccessPolicyArn{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyNameOptions",
		reflect.TypeOf((*AccessPolicyNameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AccessPolicyProps",
		reflect.TypeOf((*AccessPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AccessScope",
		reflect.TypeOf((*AccessScope)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.AccessScopeType",
		reflect.TypeOf((*AccessScopeType)(nil)).Elem(),
		map[string]interface{}{
			"NAMESPACE": AccessScopeType_NAMESPACE,
			"CLUSTER": AccessScopeType_CLUSTER,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.Addon",
		reflect.TypeOf((*Addon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addonArn", GoGetter: "AddonArn"},
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
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
			j := jsiiProxy_Addon{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAddon)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AddonAttributes",
		reflect.TypeOf((*AddonAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AddonProps",
		reflect.TypeOf((*AddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.AlbController",
		reflect.TypeOf((*AlbController)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AlbController{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AlbControllerOptions",
		reflect.TypeOf((*AlbControllerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AlbControllerProps",
		reflect.TypeOf((*AlbControllerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.AlbControllerVersion",
		reflect.TypeOf((*AlbControllerVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "custom", GoGetter: "Custom"},
			_jsii_.MemberProperty{JsiiProperty: "helmChartVersion", GoGetter: "HelmChartVersion"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_AlbControllerVersion{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.AlbScheme",
		reflect.TypeOf((*AlbScheme)(nil)).Elem(),
		map[string]interface{}{
			"INTERNAL": AlbScheme_INTERNAL,
			"INTERNET_FACING": AlbScheme_INTERNET_FACING,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AutoScalingGroupCapacityOptions",
		reflect.TypeOf((*AutoScalingGroupCapacityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.AutoScalingGroupOptions",
		reflect.TypeOf((*AutoScalingGroupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.BootstrapOptions",
		reflect.TypeOf((*BootstrapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.CapacityType",
		reflect.TypeOf((*CapacityType)(nil)).Elem(),
		map[string]interface{}{
			"SPOT": CapacityType_SPOT,
			"ON_DEMAND": CapacityType_ON_DEMAND,
			"CAPACITY_BLOCK": CapacityType_CAPACITY_BLOCK,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAutoScalingGroupCapacity", GoMethod: "AddAutoScalingGroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addCdk8sChart", GoMethod: "AddCdk8sChart"},
			_jsii_.MemberMethod{JsiiMethod: "addFargateProfile", GoMethod: "AddFargateProfile"},
			_jsii_.MemberMethod{JsiiMethod: "addHelmChart", GoMethod: "AddHelmChart"},
			_jsii_.MemberMethod{JsiiMethod: "addManifest", GoMethod: "AddManifest"},
			_jsii_.MemberMethod{JsiiMethod: "addNodegroupCapacity", GoMethod: "AddNodegroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccount", GoMethod: "AddServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "albController", GoGetter: "AlbController"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEncryptionConfigKeyArn", GoGetter: "ClusterEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterOpenIdConnectIssuerUrl", GoGetter: "ClusterOpenIdConnectIssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroup", GoGetter: "ClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "connectAutoScalingGroupCapacity", GoMethod: "ConnectAutoScalingGroupCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCapacity", GoGetter: "DefaultCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNodegroup", GoGetter: "DefaultNodegroup"},
			_jsii_.MemberProperty{JsiiProperty: "eksPodIdentityAgent", GoGetter: "EksPodIdentityAgent"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getIngressLoadBalancerAddress", GoMethod: "GetIngressLoadBalancerAddress"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getServiceLoadBalancerAddress", GoMethod: "GetServiceLoadBalancerAddress"},
			_jsii_.MemberMethod{JsiiMethod: "grantAccess", GoMethod: "GrantAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantClusterAdmin", GoMethod: "GrantClusterAdmin"},
			_jsii_.MemberProperty{JsiiProperty: "ipFamily", GoGetter: "IpFamily"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlProvider", GoGetter: "KubectlProvider"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProvider", GoGetter: "OpenIdConnectProvider"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "prune", GoGetter: "Prune"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ClusterAttributes",
		reflect.TypeOf((*ClusterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ClusterCommonOptions",
		reflect.TypeOf((*ClusterCommonOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.ClusterLoggingTypes",
		reflect.TypeOf((*ClusterLoggingTypes)(nil)).Elem(),
		map[string]interface{}{
			"API": ClusterLoggingTypes_API,
			"AUDIT": ClusterLoggingTypes_AUDIT,
			"AUTHENTICATOR": ClusterLoggingTypes_AUTHENTICATOR,
			"CONTROLLER_MANAGER": ClusterLoggingTypes_CONTROLLER_MANAGER,
			"SCHEDULER": ClusterLoggingTypes_SCHEDULER,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ComputeConfig",
		reflect.TypeOf((*ComputeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.CoreDnsComputeType",
		reflect.TypeOf((*CoreDnsComputeType)(nil)).Elem(),
		map[string]interface{}{
			"EC2": CoreDnsComputeType_EC2,
			"FARGATE": CoreDnsComputeType_FARGATE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.CpuArch",
		reflect.TypeOf((*CpuArch)(nil)).Elem(),
		map[string]interface{}{
			"ARM_64": CpuArch_ARM_64,
			"X86_64": CpuArch_X86_64,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.DefaultCapacityType",
		reflect.TypeOf((*DefaultCapacityType)(nil)).Elem(),
		map[string]interface{}{
			"NODEGROUP": DefaultCapacityType_NODEGROUP,
			"EC2": DefaultCapacityType_EC2,
			"AUTOMODE": DefaultCapacityType_AUTOMODE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.EksOptimizedImage",
		reflect.TypeOf((*EksOptimizedImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_EksOptimizedImage{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.EksOptimizedImageProps",
		reflect.TypeOf((*EksOptimizedImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.EndpointAccess",
		reflect.TypeOf((*EndpointAccess)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "onlyFrom", GoMethod: "OnlyFrom"},
		},
		func() interface{} {
			return &jsiiProxy_EndpointAccess{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.FargateCluster",
		reflect.TypeOf((*FargateCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAutoScalingGroupCapacity", GoMethod: "AddAutoScalingGroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addCdk8sChart", GoMethod: "AddCdk8sChart"},
			_jsii_.MemberMethod{JsiiMethod: "addFargateProfile", GoMethod: "AddFargateProfile"},
			_jsii_.MemberMethod{JsiiMethod: "addHelmChart", GoMethod: "AddHelmChart"},
			_jsii_.MemberMethod{JsiiMethod: "addManifest", GoMethod: "AddManifest"},
			_jsii_.MemberMethod{JsiiMethod: "addNodegroupCapacity", GoMethod: "AddNodegroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccount", GoMethod: "AddServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "albController", GoGetter: "AlbController"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEncryptionConfigKeyArn", GoGetter: "ClusterEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterOpenIdConnectIssuerUrl", GoGetter: "ClusterOpenIdConnectIssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroup", GoGetter: "ClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "connectAutoScalingGroupCapacity", GoMethod: "ConnectAutoScalingGroupCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCapacity", GoGetter: "DefaultCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNodegroup", GoGetter: "DefaultNodegroup"},
			_jsii_.MemberProperty{JsiiProperty: "defaultProfile", GoGetter: "DefaultProfile"},
			_jsii_.MemberProperty{JsiiProperty: "eksPodIdentityAgent", GoGetter: "EksPodIdentityAgent"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getIngressLoadBalancerAddress", GoMethod: "GetIngressLoadBalancerAddress"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getServiceLoadBalancerAddress", GoMethod: "GetServiceLoadBalancerAddress"},
			_jsii_.MemberMethod{JsiiMethod: "grantAccess", GoMethod: "GrantAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantClusterAdmin", GoMethod: "GrantClusterAdmin"},
			_jsii_.MemberProperty{JsiiProperty: "ipFamily", GoGetter: "IpFamily"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlProvider", GoGetter: "KubectlProvider"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProvider", GoGetter: "OpenIdConnectProvider"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "prune", GoGetter: "Prune"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_FargateCluster{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Cluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.FargateClusterProps",
		reflect.TypeOf((*FargateClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.FargateProfile",
		reflect.TypeOf((*FargateProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fargateProfileArn", GoGetter: "FargateProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "fargateProfileName", GoGetter: "FargateProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "podExecutionRole", GoGetter: "PodExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FargateProfile{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.FargateProfileOptions",
		reflect.TypeOf((*FargateProfileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.FargateProfileProps",
		reflect.TypeOf((*FargateProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.HelmChart",
		reflect.TypeOf((*HelmChart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "atomic", GoGetter: "Atomic"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "chartAsset", GoGetter: "ChartAsset"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_HelmChart{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.HelmChartOptions",
		reflect.TypeOf((*HelmChartOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.HelmChartProps",
		reflect.TypeOf((*HelmChartProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-eks-v2-alpha.IAccessEntry",
		reflect.TypeOf((*IAccessEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessEntryArn", GoGetter: "AccessEntryArn"},
			_jsii_.MemberProperty{JsiiProperty: "accessEntryName", GoGetter: "AccessEntryName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAccessEntry{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-eks-v2-alpha.IAccessPolicy",
		reflect.TypeOf((*IAccessPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessScope", GoGetter: "AccessScope"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
		},
		func() interface{} {
			return &jsiiProxy_IAccessPolicy{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-eks-v2-alpha.IAddon",
		reflect.TypeOf((*IAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addonArn", GoGetter: "AddonArn"},
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAddon{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-eks-v2-alpha.ICluster",
		reflect.TypeOf((*ICluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCdk8sChart", GoMethod: "AddCdk8sChart"},
			_jsii_.MemberMethod{JsiiMethod: "addHelmChart", GoMethod: "AddHelmChart"},
			_jsii_.MemberMethod{JsiiMethod: "addManifest", GoMethod: "AddManifest"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccount", GoMethod: "AddServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEncryptionConfigKeyArn", GoGetter: "ClusterEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroup", GoGetter: "ClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "connectAutoScalingGroupCapacity", GoMethod: "ConnectAutoScalingGroupCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "eksPodIdentityAgent", GoGetter: "EksPodIdentityAgent"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "ipFamily", GoGetter: "IpFamily"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlProvider", GoGetter: "KubectlProvider"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlProviderOptions", GoGetter: "KubectlProviderOptions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProvider", GoGetter: "OpenIdConnectProvider"},
			_jsii_.MemberProperty{JsiiProperty: "prune", GoGetter: "Prune"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ICluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-eks-v2-alpha.IKubectlProvider",
		reflect.TypeOf((*IKubectlProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
		},
		func() interface{} {
			j := jsiiProxy_IKubectlProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-eks-v2-alpha.INodegroup",
		reflect.TypeOf((*INodegroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupName", GoGetter: "NodegroupName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_INodegroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.IdentityType",
		reflect.TypeOf((*IdentityType)(nil)).Elem(),
		map[string]interface{}{
			"IRSA": IdentityType_IRSA,
			"POD_IDENTITY": IdentityType_POD_IDENTITY,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.IngressLoadBalancerAddressOptions",
		reflect.TypeOf((*IngressLoadBalancerAddressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.IpFamily",
		reflect.TypeOf((*IpFamily)(nil)).Elem(),
		map[string]interface{}{
			"IP_V4": IpFamily_IP_V4,
			"IP_V6": IpFamily_IP_V6,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProvider",
		reflect.TypeOf((*KubectlProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubectlProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKubectlProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProviderAttributes",
		reflect.TypeOf((*KubectlProviderAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProviderOptions",
		reflect.TypeOf((*KubectlProviderOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProviderProps",
		reflect.TypeOf((*KubectlProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifest",
		reflect.TypeOf((*KubernetesManifest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesManifest{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifestOptions",
		reflect.TypeOf((*KubernetesManifestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifestProps",
		reflect.TypeOf((*KubernetesManifestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesObjectValue",
		reflect.TypeOf((*KubernetesObjectValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesObjectValue{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesObjectValueProps",
		reflect.TypeOf((*KubernetesObjectValueProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesPatch",
		reflect.TypeOf((*KubernetesPatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesPatch{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesPatchProps",
		reflect.TypeOf((*KubernetesPatchProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		reflect.TypeOf((*KubernetesVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_KubernetesVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.LaunchTemplateSpec",
		reflect.TypeOf((*LaunchTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.MachineImageType",
		reflect.TypeOf((*MachineImageType)(nil)).Elem(),
		map[string]interface{}{
			"AMAZON_LINUX_2": MachineImageType_AMAZON_LINUX_2,
			"BOTTLEROCKET": MachineImageType_BOTTLEROCKET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.NodeType",
		reflect.TypeOf((*NodeType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": NodeType_STANDARD,
			"GPU": NodeType_GPU,
			"INFERENTIA": NodeType_INFERENTIA,
			"TRAINIUM": NodeType_TRAINIUM,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.Nodegroup",
		reflect.TypeOf((*Nodegroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupArn", GoGetter: "NodegroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupName", GoGetter: "NodegroupName"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Nodegroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INodegroup)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.NodegroupAmiType",
		reflect.TypeOf((*NodegroupAmiType)(nil)).Elem(),
		map[string]interface{}{
			"AL2_X86_64": NodegroupAmiType_AL2_X86_64,
			"AL2_X86_64_GPU": NodegroupAmiType_AL2_X86_64_GPU,
			"AL2_ARM_64": NodegroupAmiType_AL2_ARM_64,
			"BOTTLEROCKET_ARM_64": NodegroupAmiType_BOTTLEROCKET_ARM_64,
			"BOTTLEROCKET_X86_64": NodegroupAmiType_BOTTLEROCKET_X86_64,
			"BOTTLEROCKET_ARM_64_NVIDIA": NodegroupAmiType_BOTTLEROCKET_ARM_64_NVIDIA,
			"BOTTLEROCKET_X86_64_NVIDIA": NodegroupAmiType_BOTTLEROCKET_X86_64_NVIDIA,
			"BOTTLEROCKET_ARM_64_FIPS": NodegroupAmiType_BOTTLEROCKET_ARM_64_FIPS,
			"BOTTLEROCKET_X86_64_FIPS": NodegroupAmiType_BOTTLEROCKET_X86_64_FIPS,
			"WINDOWS_CORE_2019_X86_64": NodegroupAmiType_WINDOWS_CORE_2019_X86_64,
			"WINDOWS_CORE_2022_X86_64": NodegroupAmiType_WINDOWS_CORE_2022_X86_64,
			"WINDOWS_FULL_2019_X86_64": NodegroupAmiType_WINDOWS_FULL_2019_X86_64,
			"WINDOWS_FULL_2022_X86_64": NodegroupAmiType_WINDOWS_FULL_2022_X86_64,
			"AL2023_X86_64_STANDARD": NodegroupAmiType_AL2023_X86_64_STANDARD,
			"AL2023_X86_64_NEURON": NodegroupAmiType_AL2023_X86_64_NEURON,
			"AL2023_X86_64_NVIDIA": NodegroupAmiType_AL2023_X86_64_NVIDIA,
			"AL2023_ARM_64_STANDARD": NodegroupAmiType_AL2023_ARM_64_STANDARD,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.NodegroupOptions",
		reflect.TypeOf((*NodegroupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.NodegroupProps",
		reflect.TypeOf((*NodegroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.NodegroupRemoteAccess",
		reflect.TypeOf((*NodegroupRemoteAccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.OpenIdConnectProvider",
		reflect.TypeOf((*OpenIdConnectProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderArn", GoGetter: "OpenIdConnectProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderIssuer", GoGetter: "OpenIdConnectProviderIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderthumbprints", GoGetter: "OpenIdConnectProviderthumbprints"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OpenIdConnectProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamOpenIdConnectProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.OpenIdConnectProviderProps",
		reflect.TypeOf((*OpenIdConnectProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.PatchType",
		reflect.TypeOf((*PatchType)(nil)).Elem(),
		map[string]interface{}{
			"JSON": PatchType_JSON,
			"MERGE": PatchType_MERGE,
			"STRATEGIC": PatchType_STRATEGIC,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.Selector",
		reflect.TypeOf((*Selector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-eks-v2-alpha.ServiceAccount",
		reflect.TypeOf((*ServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleAction", GoGetter: "AssumeRoleAction"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyFragment", GoGetter: "PolicyFragment"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountName", GoGetter: "ServiceAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountNamespace", GoGetter: "ServiceAccountNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccount{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIPrincipal)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ServiceAccountOptions",
		reflect.TypeOf((*ServiceAccountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.ServiceLoadBalancerAddressOptions",
		reflect.TypeOf((*ServiceLoadBalancerAddressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-eks-v2-alpha.TaintEffect",
		reflect.TypeOf((*TaintEffect)(nil)).Elem(),
		map[string]interface{}{
			"NO_SCHEDULE": TaintEffect_NO_SCHEDULE,
			"PREFER_NO_SCHEDULE": TaintEffect_PREFER_NO_SCHEDULE,
			"NO_EXECUTE": TaintEffect_NO_EXECUTE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-eks-v2-alpha.TaintSpec",
		reflect.TypeOf((*TaintSpec)(nil)).Elem(),
	)
}
