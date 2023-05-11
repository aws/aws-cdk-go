package awscdkbatchalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
		map[string]interface{}{
			"EXIT": Action_EXIT,
			"RETRY": Action_RETRY,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.AllocationStrategy",
		reflect.TypeOf((*AllocationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"BEST_FIT": AllocationStrategy_BEST_FIT,
			"BEST_FIT_PROGRESSIVE": AllocationStrategy_BEST_FIT_PROGRESSIVE,
			"SPOT_CAPACITY_OPTIMIZED": AllocationStrategy_SPOT_CAPACITY_OPTIMIZED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.ComputeEnvironmentProps",
		reflect.TypeOf((*ComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.CustomReason",
		reflect.TypeOf((*CustomReason)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.Device",
		reflect.TypeOf((*Device)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.DevicePermission",
		reflect.TypeOf((*DevicePermission)(nil)).Elem(),
		map[string]interface{}{
			"READ": DevicePermission_READ,
			"WRITE": DevicePermission_WRITE,
			"MKNOD": DevicePermission_MKNOD,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.DnsPolicy",
		reflect.TypeOf((*DnsPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DnsPolicy_DEFAULT,
			"CLUSTER_FIRST": DnsPolicy_CLUSTER_FIRST,
			"CLUSTER_FIRST_WITH_HOST_NET": DnsPolicy_CLUSTER_FIRST_WITH_HOST_NET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EcsContainerDefinitionProps",
		reflect.TypeOf((*EcsContainerDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EcsEc2ContainerDefinition",
		reflect.TypeOf((*EcsEc2ContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUlimit", GoMethod: "AddUlimit"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "gpu", GoGetter: "Gpu"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "jobRole", GoGetter: "JobRole"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParameters", GoGetter: "LinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logDriverConfig", GoGetter: "LogDriverConfig"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ulimits", GoGetter: "Ulimits"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_EcsEc2ContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsContainerDefinition)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsEc2ContainerDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EcsEc2ContainerDefinitionProps",
		reflect.TypeOf((*EcsEc2ContainerDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EcsFargateContainerDefinition",
		reflect.TypeOf((*EcsFargateContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "assignPublicIp", GoGetter: "AssignPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "fargatePlatformVersion", GoGetter: "FargatePlatformVersion"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "jobRole", GoGetter: "JobRole"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParameters", GoGetter: "LinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logDriverConfig", GoGetter: "LogDriverConfig"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_EcsFargateContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsContainerDefinition)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsFargateContainerDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EcsFargateContainerDefinitionProps",
		reflect.TypeOf((*EcsFargateContainerDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EcsJobDefinition",
		reflect.TypeOf((*EcsJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRetryStrategy", GoMethod: "AddRetryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "propagateTags", GoGetter: "PropagateTags"},
			_jsii_.MemberProperty{JsiiProperty: "retryAttempts", GoGetter: "RetryAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "retryStrategies", GoGetter: "RetryStrategies"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPriority", GoGetter: "SchedulingPriority"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsJobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EcsJobDefinitionProps",
		reflect.TypeOf((*EcsJobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EcsMachineImage",
		reflect.TypeOf((*EcsMachineImage)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.EcsMachineImageType",
		reflect.TypeOf((*EcsMachineImageType)(nil)).Elem(),
		map[string]interface{}{
			"ECS_AL2": EcsMachineImageType_ECS_AL2,
			"ECS_AL2_NVIDIA": EcsMachineImageType_ECS_AL2_NVIDIA,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EcsVolume",
		reflect.TypeOf((*EcsVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
		},
		func() interface{} {
			return &jsiiProxy_EcsVolume{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EcsVolumeOptions",
		reflect.TypeOf((*EcsVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EfsVolume",
		reflect.TypeOf((*EfsVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessPointId", GoGetter: "AccessPointId"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "enableTransitEncryption", GoGetter: "EnableTransitEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystem", GoGetter: "FileSystem"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberProperty{JsiiProperty: "rootDirectory", GoGetter: "RootDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "transitEncryptionPort", GoGetter: "TransitEncryptionPort"},
			_jsii_.MemberProperty{JsiiProperty: "useJobRole", GoGetter: "UseJobRole"},
		},
		func() interface{} {
			j := jsiiProxy_EfsVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EcsVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EfsVolumeOptions",
		reflect.TypeOf((*EfsVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EksContainerDefinition",
		reflect.TypeOf((*EksContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpuLimit", GoGetter: "CpuLimit"},
			_jsii_.MemberProperty{JsiiProperty: "cpuReservation", GoGetter: "CpuReservation"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gpuLimit", GoGetter: "GpuLimit"},
			_jsii_.MemberProperty{JsiiProperty: "gpuReservation", GoGetter: "GpuReservation"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imagePullPolicy", GoGetter: "ImagePullPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservation", GoGetter: "MemoryReservation"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "runAsGroup", GoGetter: "RunAsGroup"},
			_jsii_.MemberProperty{JsiiProperty: "runAsRoot", GoGetter: "RunAsRoot"},
			_jsii_.MemberProperty{JsiiProperty: "runAsUser", GoGetter: "RunAsUser"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_EksContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEksContainerDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EksContainerDefinitionProps",
		reflect.TypeOf((*EksContainerDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EksJobDefinition",
		reflect.TypeOf((*EksJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRetryStrategy", GoMethod: "AddRetryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "dnsPolicy", GoGetter: "DnsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "retryAttempts", GoGetter: "RetryAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "retryStrategies", GoGetter: "RetryStrategies"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPriority", GoGetter: "SchedulingPriority"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useHostNetwork", GoGetter: "UseHostNetwork"},
		},
		func() interface{} {
			j := jsiiProxy_EksJobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEksJobDefinition)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EksJobDefinitionProps",
		reflect.TypeOf((*EksJobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EksMachineImage",
		reflect.TypeOf((*EksMachineImage)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.EksMachineImageType",
		reflect.TypeOf((*EksMachineImageType)(nil)).Elem(),
		map[string]interface{}{
			"EKS_AL2": EksMachineImageType_EKS_AL2,
			"EKS_AL2_NVIDIA": EksMachineImageType_EKS_AL2_NVIDIA,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EksVolume",
		reflect.TypeOf((*EksVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
		},
		func() interface{} {
			return &jsiiProxy_EksVolume{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EksVolumeOptions",
		reflect.TypeOf((*EksVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.EmptyDirMediumType",
		reflect.TypeOf((*EmptyDirMediumType)(nil)).Elem(),
		map[string]interface{}{
			"DISK": EmptyDirMediumType_DISK,
			"MEMORY": EmptyDirMediumType_MEMORY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolume",
		reflect.TypeOf((*EmptyDirVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "medium", GoGetter: "Medium"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberProperty{JsiiProperty: "sizeLimit", GoGetter: "SizeLimit"},
		},
		func() interface{} {
			j := jsiiProxy_EmptyDirVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EksVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.EmptyDirVolumeOptions",
		reflect.TypeOf((*EmptyDirVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.FairshareSchedulingPolicy",
		reflect.TypeOf((*FairshareSchedulingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addShare", GoMethod: "AddShare"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeReservation", GoGetter: "ComputeReservation"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyArn", GoGetter: "SchedulingPolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyName", GoGetter: "SchedulingPolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "shareDecay", GoGetter: "ShareDecay"},
			_jsii_.MemberProperty{JsiiProperty: "shares", GoGetter: "Shares"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FairshareSchedulingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFairshareSchedulingPolicy)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISchedulingPolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.FairshareSchedulingPolicyProps",
		reflect.TypeOf((*FairshareSchedulingPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.FargateComputeEnvironment",
		reflect.TypeOf((*FargateComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxvCpus", GoGetter: "MaxvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "replaceComputeEnvironment", GoGetter: "ReplaceComputeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "terminateOnUpdate", GoGetter: "TerminateOnUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateTimeout", GoGetter: "UpdateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "updateToLatestImageVersion", GoGetter: "UpdateToLatestImageVersion"},
		},
		func() interface{} {
			j := jsiiProxy_FargateComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFargateComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManagedComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.FargateComputeEnvironmentProps",
		reflect.TypeOf((*FargateComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		reflect.TypeOf((*HostPathVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
		},
		func() interface{} {
			j := jsiiProxy_HostPathVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EksVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.HostPathVolumeOptions",
		reflect.TypeOf((*HostPathVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.HostVolume",
		reflect.TypeOf((*HostVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "hostPath", GoGetter: "HostPath"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
		},
		func() interface{} {
			j := jsiiProxy_HostVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EcsVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.HostVolumeOptions",
		reflect.TypeOf((*HostVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IComputeEnvironment",
		reflect.TypeOf((*IComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IEcsContainerDefinition",
		reflect.TypeOf((*IEcsContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "jobRole", GoGetter: "JobRole"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParameters", GoGetter: "LinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logDriverConfig", GoGetter: "LogDriverConfig"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_IEcsContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IEcsEc2ContainerDefinition",
		reflect.TypeOf((*IEcsEc2ContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addUlimit", GoMethod: "AddUlimit"},
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "gpu", GoGetter: "Gpu"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "jobRole", GoGetter: "JobRole"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParameters", GoGetter: "LinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logDriverConfig", GoGetter: "LogDriverConfig"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberProperty{JsiiProperty: "ulimits", GoGetter: "Ulimits"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_IEcsEc2ContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsContainerDefinition)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IEcsFargateContainerDefinition",
		reflect.TypeOf((*IEcsFargateContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "assignPublicIp", GoGetter: "AssignPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "fargatePlatformVersion", GoGetter: "FargatePlatformVersion"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "jobRole", GoGetter: "JobRole"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParameters", GoGetter: "LinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logDriverConfig", GoGetter: "LogDriverConfig"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_IEcsFargateContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsContainerDefinition)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IEksContainerDefinition",
		reflect.TypeOf((*IEksContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVolume", GoMethod: "AddVolume"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "cpuLimit", GoGetter: "CpuLimit"},
			_jsii_.MemberProperty{JsiiProperty: "cpuReservation", GoGetter: "CpuReservation"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gpuLimit", GoGetter: "GpuLimit"},
			_jsii_.MemberProperty{JsiiProperty: "gpuReservation", GoGetter: "GpuReservation"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imagePullPolicy", GoGetter: "ImagePullPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "memoryLimit", GoGetter: "MemoryLimit"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservation", GoGetter: "MemoryReservation"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "runAsGroup", GoGetter: "RunAsGroup"},
			_jsii_.MemberProperty{JsiiProperty: "runAsRoot", GoGetter: "RunAsRoot"},
			_jsii_.MemberProperty{JsiiProperty: "runAsUser", GoGetter: "RunAsUser"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
		},
		func() interface{} {
			j := jsiiProxy_IEksContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IEksJobDefinition",
		reflect.TypeOf((*IEksJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRetryStrategy", GoMethod: "AddRetryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "dnsPolicy", GoGetter: "DnsPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "retryAttempts", GoGetter: "RetryAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "retryStrategies", GoGetter: "RetryStrategies"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPriority", GoGetter: "SchedulingPriority"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "useHostNetwork", GoGetter: "UseHostNetwork"},
		},
		func() interface{} {
			j := jsiiProxy_IEksJobDefinition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobDefinition)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IFairshareSchedulingPolicy",
		reflect.TypeOf((*IFairshareSchedulingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeReservation", GoGetter: "ComputeReservation"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyArn", GoGetter: "SchedulingPolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyName", GoGetter: "SchedulingPolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "shareDecay", GoGetter: "ShareDecay"},
			_jsii_.MemberProperty{JsiiProperty: "shares", GoGetter: "Shares"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IFairshareSchedulingPolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISchedulingPolicy)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IFargateComputeEnvironment",
		reflect.TypeOf((*IFargateComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "maxvCpus", GoGetter: "MaxvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "replaceComputeEnvironment", GoGetter: "ReplaceComputeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "terminateOnUpdate", GoGetter: "TerminateOnUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "updateTimeout", GoGetter: "UpdateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "updateToLatestImageVersion", GoGetter: "UpdateToLatestImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_IFargateComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManagedComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IJobDefinition",
		reflect.TypeOf((*IJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addRetryStrategy", GoMethod: "AddRetryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "retryAttempts", GoGetter: "RetryAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "retryStrategies", GoGetter: "RetryStrategies"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPriority", GoGetter: "SchedulingPriority"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
		},
		func() interface{} {
			j := jsiiProxy_IJobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IJobQueue",
		reflect.TypeOf((*IJobQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addComputeEnvironment", GoMethod: "AddComputeEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironments", GoGetter: "ComputeEnvironments"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueArn", GoGetter: "JobQueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueName", GoGetter: "JobQueueName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicy", GoGetter: "SchedulingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IJobQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IManagedComputeEnvironment",
		reflect.TypeOf((*IManagedComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "maxvCpus", GoGetter: "MaxvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "replaceComputeEnvironment", GoGetter: "ReplaceComputeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "terminateOnUpdate", GoGetter: "TerminateOnUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "updateTimeout", GoGetter: "UpdateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "updateToLatestImageVersion", GoGetter: "UpdateToLatestImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_IManagedComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IManagedEc2EcsComputeEnvironment",
		reflect.TypeOf((*IManagedEc2EcsComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInstanceClass", GoMethod: "AddInstanceClass"},
			_jsii_.MemberMethod{JsiiMethod: "addInstanceType", GoMethod: "AddInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "allocationStrategy", GoGetter: "AllocationStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "images", GoGetter: "Images"},
			_jsii_.MemberProperty{JsiiProperty: "instanceClasses", GoGetter: "InstanceClasses"},
			_jsii_.MemberProperty{JsiiProperty: "instanceRole", GoGetter: "InstanceRole"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "maxvCpus", GoGetter: "MaxvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "minvCpus", GoGetter: "MinvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroup", GoGetter: "PlacementGroup"},
			_jsii_.MemberProperty{JsiiProperty: "replaceComputeEnvironment", GoGetter: "ReplaceComputeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "spotBidPercentage", GoGetter: "SpotBidPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "spotFleetRole", GoGetter: "SpotFleetRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "terminateOnUpdate", GoGetter: "TerminateOnUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "updateTimeout", GoGetter: "UpdateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "updateToLatestImageVersion", GoGetter: "UpdateToLatestImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "useOptimalInstanceClasses", GoGetter: "UseOptimalInstanceClasses"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_IManagedEc2EcsComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManagedComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.ISchedulingPolicy",
		reflect.TypeOf((*ISchedulingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyArn", GoGetter: "SchedulingPolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicyName", GoGetter: "SchedulingPolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ISchedulingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IUnmanagedComputeEnvironment",
		reflect.TypeOf((*IUnmanagedComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "unmanagedvCPUs", GoGetter: "UnmanagedvCPUs"},
		},
		func() interface{} {
			j := jsiiProxy_IUnmanagedComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.ImagePullPolicy",
		reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ImagePullPolicy_ALWAYS,
			"IF_NOT_PRESENT": ImagePullPolicy_IF_NOT_PRESENT,
			"NEVER": ImagePullPolicy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.JobDefinitionProps",
		reflect.TypeOf((*JobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.JobQueue",
		reflect.TypeOf((*JobQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addComputeEnvironment", GoMethod: "AddComputeEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironments", GoGetter: "ComputeEnvironments"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueArn", GoGetter: "JobQueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueName", GoGetter: "JobQueueName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPolicy", GoGetter: "SchedulingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_JobQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobQueue)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.JobQueueProps",
		reflect.TypeOf((*JobQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.LinuxParameters",
		reflect.TypeOf((*LinuxParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDevices", GoMethod: "AddDevices"},
			_jsii_.MemberMethod{JsiiMethod: "addTmpfs", GoMethod: "AddTmpfs"},
			_jsii_.MemberProperty{JsiiProperty: "devices", GoGetter: "Devices"},
			_jsii_.MemberProperty{JsiiProperty: "initProcessEnabled", GoGetter: "InitProcessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "maxSwap", GoGetter: "MaxSwap"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "renderLinuxParameters", GoMethod: "RenderLinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "sharedMemorySize", GoGetter: "SharedMemorySize"},
			_jsii_.MemberProperty{JsiiProperty: "swappiness", GoGetter: "Swappiness"},
			_jsii_.MemberProperty{JsiiProperty: "tmpfs", GoGetter: "Tmpfs"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LinuxParameters{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.LinuxParametersProps",
		reflect.TypeOf((*LinuxParametersProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.ManagedComputeEnvironmentProps",
		reflect.TypeOf((*ManagedComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.ManagedEc2EcsComputeEnvironment",
		reflect.TypeOf((*ManagedEc2EcsComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInstanceClass", GoMethod: "AddInstanceClass"},
			_jsii_.MemberMethod{JsiiMethod: "addInstanceType", GoMethod: "AddInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "allocationStrategy", GoGetter: "AllocationStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "images", GoGetter: "Images"},
			_jsii_.MemberProperty{JsiiProperty: "instanceClasses", GoGetter: "InstanceClasses"},
			_jsii_.MemberProperty{JsiiProperty: "instanceRole", GoGetter: "InstanceRole"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "maxvCpus", GoGetter: "MaxvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "minvCpus", GoGetter: "MinvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroup", GoGetter: "PlacementGroup"},
			_jsii_.MemberProperty{JsiiProperty: "replaceComputeEnvironment", GoGetter: "ReplaceComputeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "spotBidPercentage", GoGetter: "SpotBidPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "spotFleetRole", GoGetter: "SpotFleetRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "terminateOnUpdate", GoGetter: "TerminateOnUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateTimeout", GoGetter: "UpdateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "updateToLatestImageVersion", GoGetter: "UpdateToLatestImageVersion"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedEc2EcsComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManagedComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManagedEc2EcsComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.ManagedEc2EcsComputeEnvironmentProps",
		reflect.TypeOf((*ManagedEc2EcsComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.ManagedEc2EksComputeEnvironment",
		reflect.TypeOf((*ManagedEc2EksComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInstanceClass", GoMethod: "AddInstanceClass"},
			_jsii_.MemberMethod{JsiiMethod: "addInstanceType", GoMethod: "AddInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "allocationStrategy", GoGetter: "AllocationStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "eksCluster", GoGetter: "EksCluster"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "images", GoGetter: "Images"},
			_jsii_.MemberProperty{JsiiProperty: "instanceClasses", GoGetter: "InstanceClasses"},
			_jsii_.MemberProperty{JsiiProperty: "instanceRole", GoGetter: "InstanceRole"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesNamespace", GoGetter: "KubernetesNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "maxvCpus", GoGetter: "MaxvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "minvCpus", GoGetter: "MinvCpus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "placementGroup", GoGetter: "PlacementGroup"},
			_jsii_.MemberProperty{JsiiProperty: "replaceComputeEnvironment", GoGetter: "ReplaceComputeEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "spot", GoGetter: "Spot"},
			_jsii_.MemberProperty{JsiiProperty: "spotBidPercentage", GoGetter: "SpotBidPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "terminateOnUpdate", GoGetter: "TerminateOnUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateTimeout", GoGetter: "UpdateTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "updateToLatestImageVersion", GoGetter: "UpdateToLatestImageVersion"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedEc2EksComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManagedComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.ManagedEc2EksComputeEnvironmentProps",
		reflect.TypeOf((*ManagedEc2EksComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.MultiNodeContainer",
		reflect.TypeOf((*MultiNodeContainer)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.MultiNodeJobDefinition",
		reflect.TypeOf((*MultiNodeJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addRetryStrategy", GoMethod: "AddRetryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "containers", GoGetter: "Containers"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "mainNode", GoGetter: "MainNode"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "propagateTags", GoGetter: "PropagateTags"},
			_jsii_.MemberProperty{JsiiProperty: "retryAttempts", GoGetter: "RetryAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "retryStrategies", GoGetter: "RetryStrategies"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingPriority", GoGetter: "SchedulingPriority"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MultiNodeJobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.MultiNodeJobDefinitionProps",
		reflect.TypeOf((*MultiNodeJobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.OrderedComputeEnvironment",
		reflect.TypeOf((*OrderedComputeEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.Reason",
		reflect.TypeOf((*Reason)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Reason{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.RetryStrategy",
		reflect.TypeOf((*RetryStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "on", GoGetter: "On"},
		},
		func() interface{} {
			return &jsiiProxy_RetryStrategy{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.SecretPathVolume",
		reflect.TypeOf((*SecretPathVolume)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "optional", GoGetter: "Optional"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
		},
		func() interface{} {
			j := jsiiProxy_SecretPathVolume{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EksVolume)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.SecretPathVolumeOptions",
		reflect.TypeOf((*SecretPathVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.Share",
		reflect.TypeOf((*Share)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.Tmpfs",
		reflect.TypeOf((*Tmpfs)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.TmpfsMountOption",
		reflect.TypeOf((*TmpfsMountOption)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULTS": TmpfsMountOption_DEFAULTS,
			"RO": TmpfsMountOption_RO,
			"RW": TmpfsMountOption_RW,
			"SUID": TmpfsMountOption_SUID,
			"NOSUID": TmpfsMountOption_NOSUID,
			"DEV": TmpfsMountOption_DEV,
			"NODEV": TmpfsMountOption_NODEV,
			"EXEC": TmpfsMountOption_EXEC,
			"NOEXEC": TmpfsMountOption_NOEXEC,
			"SYNC": TmpfsMountOption_SYNC,
			"ASYNC": TmpfsMountOption_ASYNC,
			"DIRSYNC": TmpfsMountOption_DIRSYNC,
			"REMOUNT": TmpfsMountOption_REMOUNT,
			"MAND": TmpfsMountOption_MAND,
			"NOMAND": TmpfsMountOption_NOMAND,
			"ATIME": TmpfsMountOption_ATIME,
			"NOATIME": TmpfsMountOption_NOATIME,
			"DIRATIME": TmpfsMountOption_DIRATIME,
			"NODIRATIME": TmpfsMountOption_NODIRATIME,
			"BIND": TmpfsMountOption_BIND,
			"RBIND": TmpfsMountOption_RBIND,
			"UNBINDABLE": TmpfsMountOption_UNBINDABLE,
			"RUNBINDABLE": TmpfsMountOption_RUNBINDABLE,
			"PRIVATE": TmpfsMountOption_PRIVATE,
			"RPRIVATE": TmpfsMountOption_RPRIVATE,
			"SHARED": TmpfsMountOption_SHARED,
			"RSHARED": TmpfsMountOption_RSHARED,
			"SLAVE": TmpfsMountOption_SLAVE,
			"RSLAVE": TmpfsMountOption_RSLAVE,
			"RELATIME": TmpfsMountOption_RELATIME,
			"NORELATIME": TmpfsMountOption_NORELATIME,
			"STRICTATIME": TmpfsMountOption_STRICTATIME,
			"NOSTRICTATIME": TmpfsMountOption_NOSTRICTATIME,
			"MODE": TmpfsMountOption_MODE,
			"UID": TmpfsMountOption_UID,
			"GID": TmpfsMountOption_GID,
			"NR_INODES": TmpfsMountOption_NR_INODES,
			"NR_BLOCKS": TmpfsMountOption_NR_BLOCKS,
			"MPOL": TmpfsMountOption_MPOL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.Ulimit",
		reflect.TypeOf((*Ulimit)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.UlimitName",
		reflect.TypeOf((*UlimitName)(nil)).Elem(),
		map[string]interface{}{
			"CORE": UlimitName_CORE,
			"CPU": UlimitName_CPU,
			"DATA": UlimitName_DATA,
			"FSIZE": UlimitName_FSIZE,
			"LOCKS": UlimitName_LOCKS,
			"MEMLOCK": UlimitName_MEMLOCK,
			"MSGQUEUE": UlimitName_MSGQUEUE,
			"NICE": UlimitName_NICE,
			"NOFILE": UlimitName_NOFILE,
			"NPROC": UlimitName_NPROC,
			"RSS": UlimitName_RSS,
			"RTPRIO": UlimitName_RTPRIO,
			"RTTIME": UlimitName_RTTIME,
			"SIGPENDING": UlimitName_SIGPENDING,
			"STACK": UlimitName_STACK,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.UnmanagedComputeEnvironment",
		reflect.TypeOf((*UnmanagedComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unmanagedvCPUs", GoGetter: "UnmanagedvCPUs"},
		},
		func() interface{} {
			j := jsiiProxy_UnmanagedComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUnmanagedComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.UnmanagedComputeEnvironmentProps",
		reflect.TypeOf((*UnmanagedComputeEnvironmentProps)(nil)).Elem(),
	)
}
