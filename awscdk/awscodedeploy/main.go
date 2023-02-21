package awscodedeploy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.AllAtOnceTrafficRouting",
		reflect.TypeOf((*AllAtOnceTrafficRouting)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AllAtOnceTrafficRouting{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TrafficRouting)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.AutoRollbackConfig",
		reflect.TypeOf((*AutoRollbackConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.BaseDeploymentConfig",
		reflect.TypeOf((*BaseDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
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
			j := jsiiProxy_BaseDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.BaseDeploymentConfigOptions",
		reflect.TypeOf((*BaseDeploymentConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.BaseDeploymentConfigProps",
		reflect.TypeOf((*BaseDeploymentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.BaseTrafficShiftingConfigProps",
		reflect.TypeOf((*BaseTrafficShiftingConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CanaryTrafficRoutingConfig",
		reflect.TypeOf((*CanaryTrafficRoutingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.CfnApplication",
		reflect.TypeOf((*CfnApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "computePlatform", GoGetter: "ComputePlatform"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnApplicationProps",
		reflect.TypeOf((*CfnApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig",
		reflect.TypeOf((*CfnDeploymentConfig)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "computePlatform", GoGetter: "ComputePlatform"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "minimumHealthyHosts", GoGetter: "MinimumHealthyHosts"},
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
			_jsii_.MemberProperty{JsiiProperty: "trafficRoutingConfig", GoGetter: "TrafficRoutingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig.MinimumHealthyHostsProperty",
		reflect.TypeOf((*CfnDeploymentConfig_MinimumHealthyHostsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig.TimeBasedCanaryProperty",
		reflect.TypeOf((*CfnDeploymentConfig_TimeBasedCanaryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig.TimeBasedLinearProperty",
		reflect.TypeOf((*CfnDeploymentConfig_TimeBasedLinearProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfig.TrafficRoutingConfigProperty",
		reflect.TypeOf((*CfnDeploymentConfig_TrafficRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentConfigProps",
		reflect.TypeOf((*CfnDeploymentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup",
		reflect.TypeOf((*CfnDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alarmConfiguration", GoGetter: "AlarmConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoRollbackConfiguration", GoGetter: "AutoRollbackConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroups", GoGetter: "AutoScalingGroups"},
			_jsii_.MemberProperty{JsiiProperty: "blueGreenDeploymentConfiguration", GoGetter: "BlueGreenDeploymentConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deployment", GoGetter: "Deployment"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStyle", GoGetter: "DeploymentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "ec2TagFilters", GoGetter: "Ec2TagFilters"},
			_jsii_.MemberProperty{JsiiProperty: "ec2TagSet", GoGetter: "Ec2TagSet"},
			_jsii_.MemberProperty{JsiiProperty: "ecsServices", GoGetter: "EcsServices"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerInfo", GoGetter: "LoadBalancerInfo"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "onPremisesInstanceTagFilters", GoGetter: "OnPremisesInstanceTagFilters"},
			_jsii_.MemberProperty{JsiiProperty: "onPremisesTagSet", GoGetter: "OnPremisesTagSet"},
			_jsii_.MemberProperty{JsiiProperty: "outdatedInstancesStrategy", GoGetter: "OutdatedInstancesStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfigurations", GoGetter: "TriggerConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.AlarmConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroup_AlarmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.AlarmProperty",
		reflect.TypeOf((*CfnDeploymentGroup_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.AutoRollbackConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroup_AutoRollbackConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.BlueGreenDeploymentConfigurationProperty",
		reflect.TypeOf((*CfnDeploymentGroup_BlueGreenDeploymentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.BlueInstanceTerminationOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroup_BlueInstanceTerminationOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.DeploymentProperty",
		reflect.TypeOf((*CfnDeploymentGroup_DeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.DeploymentReadyOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroup_DeploymentReadyOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.DeploymentStyleProperty",
		reflect.TypeOf((*CfnDeploymentGroup_DeploymentStyleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.EC2TagFilterProperty",
		reflect.TypeOf((*CfnDeploymentGroup_EC2TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.EC2TagSetListObjectProperty",
		reflect.TypeOf((*CfnDeploymentGroup_EC2TagSetListObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.EC2TagSetProperty",
		reflect.TypeOf((*CfnDeploymentGroup_EC2TagSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.ECSServiceProperty",
		reflect.TypeOf((*CfnDeploymentGroup_ECSServiceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.ELBInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroup_ELBInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.GitHubLocationProperty",
		reflect.TypeOf((*CfnDeploymentGroup_GitHubLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.GreenFleetProvisioningOptionProperty",
		reflect.TypeOf((*CfnDeploymentGroup_GreenFleetProvisioningOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.LoadBalancerInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroup_LoadBalancerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.OnPremisesTagSetListObjectProperty",
		reflect.TypeOf((*CfnDeploymentGroup_OnPremisesTagSetListObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.OnPremisesTagSetProperty",
		reflect.TypeOf((*CfnDeploymentGroup_OnPremisesTagSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.RevisionLocationProperty",
		reflect.TypeOf((*CfnDeploymentGroup_RevisionLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.S3LocationProperty",
		reflect.TypeOf((*CfnDeploymentGroup_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.TagFilterProperty",
		reflect.TypeOf((*CfnDeploymentGroup_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.TargetGroupInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroup_TargetGroupInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.TargetGroupPairInfoProperty",
		reflect.TypeOf((*CfnDeploymentGroup_TargetGroupPairInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.TrafficRouteProperty",
		reflect.TypeOf((*CfnDeploymentGroup_TrafficRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroup.TriggerConfigProperty",
		reflect.TypeOf((*CfnDeploymentGroup_TriggerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CfnDeploymentGroupProps",
		reflect.TypeOf((*CfnDeploymentGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codedeploy.ComputePlatform",
		reflect.TypeOf((*ComputePlatform)(nil)).Elem(),
		map[string]interface{}{
			"SERVER": ComputePlatform_SERVER,
			"LAMBDA": ComputePlatform_LAMBDA,
			"ECS": ComputePlatform_ECS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		reflect.TypeOf((*CustomLambdaDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
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
			j := jsiiProxy_CustomLambdaDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILambdaDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfigProps",
		reflect.TypeOf((*CustomLambdaDeploymentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfigType",
		reflect.TypeOf((*CustomLambdaDeploymentConfigType)(nil)).Elem(),
		map[string]interface{}{
			"CANARY": CustomLambdaDeploymentConfigType_CANARY,
			"LINEAR": CustomLambdaDeploymentConfigType_LINEAR,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.EcsApplication",
		reflect.TypeOf((*EcsApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
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
			j := jsiiProxy_EcsApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsApplication)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.EcsApplicationProps",
		reflect.TypeOf((*EcsApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.EcsBlueGreenDeploymentConfig",
		reflect.TypeOf((*EcsBlueGreenDeploymentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentConfig",
		reflect.TypeOf((*EcsDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
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
			j := jsiiProxy_EcsDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDeploymentConfig)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentConfigProps",
		reflect.TypeOf((*EcsDeploymentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentGroup",
		reflect.TypeOf((*EcsDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlarm", GoMethod: "AddAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupArn", GoGetter: "DeploymentGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsDeploymentGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentGroupAttributes",
		reflect.TypeOf((*EcsDeploymentGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.EcsDeploymentGroupProps",
		reflect.TypeOf((*EcsDeploymentGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IBaseDeploymentConfig",
		reflect.TypeOf((*IBaseDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
		},
		func() interface{} {
			return &jsiiProxy_IBaseDeploymentConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IEcsApplication",
		reflect.TypeOf((*IEcsApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IEcsApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IEcsDeploymentConfig",
		reflect.TypeOf((*IEcsDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
		},
		func() interface{} {
			j := jsiiProxy_IEcsDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IEcsDeploymentGroup",
		reflect.TypeOf((*IEcsDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupArn", GoGetter: "DeploymentGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IEcsDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.ILambdaApplication",
		reflect.TypeOf((*ILambdaApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ILambdaApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.ILambdaDeploymentConfig",
		reflect.TypeOf((*ILambdaDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
		},
		func() interface{} {
			j := jsiiProxy_ILambdaDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.ILambdaDeploymentGroup",
		reflect.TypeOf((*ILambdaDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupArn", GoGetter: "DeploymentGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ILambdaDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IServerApplication",
		reflect.TypeOf((*IServerApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IServerApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IServerDeploymentConfig",
		reflect.TypeOf((*IServerDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
		},
		func() interface{} {
			j := jsiiProxy_IServerDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBaseDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codedeploy.IServerDeploymentGroup",
		reflect.TypeOf((*IServerDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroups", GoGetter: "AutoScalingGroups"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupArn", GoGetter: "DeploymentGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IServerDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.InstanceTagSet",
		reflect.TypeOf((*InstanceTagSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "instanceTagGroups", GoGetter: "InstanceTagGroups"},
		},
		func() interface{} {
			return &jsiiProxy_InstanceTagSet{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.LambdaApplication",
		reflect.TypeOf((*LambdaApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
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
			j := jsiiProxy_LambdaApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILambdaApplication)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.LambdaApplicationProps",
		reflect.TypeOf((*LambdaApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfig",
		reflect.TypeOf((*LambdaDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
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
			j := jsiiProxy_LambdaDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDeploymentConfig)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILambdaDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfigImportProps",
		reflect.TypeOf((*LambdaDeploymentConfigImportProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentConfigProps",
		reflect.TypeOf((*LambdaDeploymentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroup",
		reflect.TypeOf((*LambdaDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlarm", GoMethod: "AddAlarm"},
			_jsii_.MemberMethod{JsiiMethod: "addPostHook", GoMethod: "AddPostHook"},
			_jsii_.MemberMethod{JsiiMethod: "addPreHook", GoMethod: "AddPreHook"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupArn", GoGetter: "DeploymentGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutLifecycleEventHookExecutionStatus", GoMethod: "GrantPutLifecycleEventHookExecutionStatus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILambdaDeploymentGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroupAttributes",
		reflect.TypeOf((*LambdaDeploymentGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.LambdaDeploymentGroupProps",
		reflect.TypeOf((*LambdaDeploymentGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.LinearTrafficRoutingConfig",
		reflect.TypeOf((*LinearTrafficRoutingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		reflect.TypeOf((*LoadBalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "generation", GoGetter: "Generation"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_LoadBalancer{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codedeploy.LoadBalancerGeneration",
		reflect.TypeOf((*LoadBalancerGeneration)(nil)).Elem(),
		map[string]interface{}{
			"FIRST": LoadBalancerGeneration_FIRST,
			"SECOND": LoadBalancerGeneration_SECOND,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.MinimumHealthyHosts",
		reflect.TypeOf((*MinimumHealthyHosts)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_MinimumHealthyHosts{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.ServerApplication",
		reflect.TypeOf((*ServerApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationName", GoGetter: "ApplicationName"},
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
			j := jsiiProxy_ServerApplication{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServerApplication)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.ServerApplicationProps",
		reflect.TypeOf((*ServerApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfig",
		reflect.TypeOf((*ServerDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigArn", GoGetter: "DeploymentConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
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
			j := jsiiProxy_ServerDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDeploymentConfig)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServerDeploymentConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentConfigProps",
		reflect.TypeOf((*ServerDeploymentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroup",
		reflect.TypeOf((*ServerDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlarm", GoMethod: "AddAlarm"},
			_jsii_.MemberMethod{JsiiMethod: "addAutoScalingGroup", GoMethod: "AddAutoScalingGroup"},
			_jsii_.MemberProperty{JsiiProperty: "application", GoGetter: "Application"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroups", GoGetter: "AutoScalingGroups"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupArn", GoGetter: "DeploymentGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServerDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IServerDeploymentGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroupAttributes",
		reflect.TypeOf((*ServerDeploymentGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.ServerDeploymentGroupProps",
		reflect.TypeOf((*ServerDeploymentGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRouting",
		reflect.TypeOf((*TimeBasedCanaryTrafficRouting)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "percentage", GoGetter: "Percentage"},
		},
		func() interface{} {
			j := jsiiProxy_TimeBasedCanaryTrafficRouting{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TrafficRouting)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRoutingProps",
		reflect.TypeOf((*TimeBasedCanaryTrafficRoutingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRouting",
		reflect.TypeOf((*TimeBasedLinearTrafficRouting)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "percentage", GoGetter: "Percentage"},
		},
		func() interface{} {
			j := jsiiProxy_TimeBasedLinearTrafficRouting{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TrafficRouting)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRoutingProps",
		reflect.TypeOf((*TimeBasedLinearTrafficRoutingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codedeploy.TrafficRouting",
		reflect.TypeOf((*TrafficRouting)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_TrafficRouting{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codedeploy.TrafficRoutingConfig",
		reflect.TypeOf((*TrafficRoutingConfig)(nil)).Elem(),
	)
}
