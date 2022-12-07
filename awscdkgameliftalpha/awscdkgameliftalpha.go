package awscdkgameliftalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		reflect.TypeOf((*Alias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fleet", GoGetter: "Fleet"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Alias{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AliasBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.AliasAttributes",
		reflect.TypeOf((*AliasAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.AliasBase",
		reflect.TypeOf((*AliasBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
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
			j := jsiiProxy_AliasBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlias)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.AliasOptions",
		reflect.TypeOf((*AliasOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.AliasProps",
		reflect.TypeOf((*AliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.AssetContent",
		reflect.TypeOf((*AssetContent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
		},
		func() interface{} {
			j := jsiiProxy_AssetContent{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Content)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.AutoScalingPolicy",
		reflect.TypeOf((*AutoScalingPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-gamelift-alpha.BalancingStrategy",
		reflect.TypeOf((*BalancingStrategy)(nil)).Elem(),
		map[string]interface{}{
			"SPOT_ONLY": BalancingStrategy_SPOT_ONLY,
			"SPOT_PREFERRED": BalancingStrategy_SPOT_PREFERRED,
			"ON_DEMAND_ONLY": BalancingStrategy_ON_DEMAND_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.Build",
		reflect.TypeOf((*Build)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "buildArn", GoGetter: "BuildArn"},
			_jsii_.MemberProperty{JsiiProperty: "buildId", GoGetter: "BuildId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Build{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BuildBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.BuildAttributes",
		reflect.TypeOf((*BuildAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.BuildBase",
		reflect.TypeOf((*BuildBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "buildArn", GoGetter: "BuildArn"},
			_jsii_.MemberProperty{JsiiProperty: "buildId", GoGetter: "BuildId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BuildBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBuild)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.BuildFleet",
		reflect.TypeOf((*BuildFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addIngressRule", GoMethod: "AddIngressRule"},
			_jsii_.MemberMethod{JsiiMethod: "addInternalLocation", GoMethod: "AddInternalLocation"},
			_jsii_.MemberMethod{JsiiMethod: "addLocation", GoMethod: "AddLocation"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fleetArn", GoGetter: "FleetArn"},
			_jsii_.MemberProperty{JsiiProperty: "fleetId", GoGetter: "FleetId"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveInstances", GoMethod: "MetricActiveInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricDesiredInstances", GoMethod: "MetricDesiredInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricIdleInstances", GoMethod: "MetricIdleInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricInstanceInterruptions", GoMethod: "MetricInstanceInterruptions"},
			_jsii_.MemberMethod{JsiiMethod: "metricMaxInstances", GoMethod: "MetricMaxInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricMinInstances", GoMethod: "MetricMinInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricPercentIdleInstances", GoMethod: "MetricPercentIdleInstances"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseLocationCapacity", GoMethod: "ParseLocationCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "parseLocations", GoMethod: "ParseLocations"},
			_jsii_.MemberMethod{JsiiMethod: "parseResourceCreationLimitPolicy", GoMethod: "ParseResourceCreationLimitPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "parseRuntimeConfiguration", GoMethod: "ParseRuntimeConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "warnVpcPeeringAuthorizations", GoMethod: "WarnVpcPeeringAuthorizations"},
		},
		func() interface{} {
			j := jsiiProxy_BuildFleet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_FleetBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IBuildFleet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.BuildFleetProps",
		reflect.TypeOf((*BuildFleetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.BuildProps",
		reflect.TypeOf((*BuildProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.Content",
		reflect.TypeOf((*Content)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Content{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.ContentConfig",
		reflect.TypeOf((*ContentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-gamelift-alpha.DeleteOption",
		reflect.TypeOf((*DeleteOption)(nil)).Elem(),
		map[string]interface{}{
			"SAFE_DELETE": DeleteOption_SAFE_DELETE,
			"FORCE_DELETE": DeleteOption_FORCE_DELETE,
			"RETAIN": DeleteOption_RETAIN,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.FleetAttributes",
		reflect.TypeOf((*FleetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.FleetBase",
		reflect.TypeOf((*FleetBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addInternalLocation", GoMethod: "AddInternalLocation"},
			_jsii_.MemberMethod{JsiiMethod: "addLocation", GoMethod: "AddLocation"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fleetArn", GoGetter: "FleetArn"},
			_jsii_.MemberProperty{JsiiProperty: "fleetId", GoGetter: "FleetId"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveInstances", GoMethod: "MetricActiveInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricDesiredInstances", GoMethod: "MetricDesiredInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricIdleInstances", GoMethod: "MetricIdleInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricInstanceInterruptions", GoMethod: "MetricInstanceInterruptions"},
			_jsii_.MemberMethod{JsiiMethod: "metricMaxInstances", GoMethod: "MetricMaxInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricMinInstances", GoMethod: "MetricMinInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricPercentIdleInstances", GoMethod: "MetricPercentIdleInstances"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseLocationCapacity", GoMethod: "ParseLocationCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "parseLocations", GoMethod: "ParseLocations"},
			_jsii_.MemberMethod{JsiiMethod: "parseResourceCreationLimitPolicy", GoMethod: "ParseResourceCreationLimitPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "parseRuntimeConfiguration", GoMethod: "ParseRuntimeConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "warnVpcPeeringAuthorizations", GoMethod: "WarnVpcPeeringAuthorizations"},
		},
		func() interface{} {
			j := jsiiProxy_FleetBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFleet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.FleetProps",
		reflect.TypeOf((*FleetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		reflect.TypeOf((*GameServerGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupArn", GoGetter: "AutoScalingGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gameServerGroupArn", GoGetter: "GameServerGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "gameServerGroupName", GoGetter: "GameServerGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "parseAutoScalingPolicy", GoMethod: "ParseAutoScalingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "parseInstanceDefinitions", GoMethod: "ParseInstanceDefinitions"},
			_jsii_.MemberMethod{JsiiMethod: "parseLaunchTemplate", GoMethod: "ParseLaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "vpcSubnets", GoGetter: "VpcSubnets"},
		},
		func() interface{} {
			j := jsiiProxy_GameServerGroup{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GameServerGroupBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroupAttributes",
		reflect.TypeOf((*GameServerGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroupBase",
		reflect.TypeOf((*GameServerGroupBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupArn", GoGetter: "AutoScalingGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gameServerGroupArn", GoGetter: "GameServerGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "gameServerGroupName", GoGetter: "GameServerGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GameServerGroupBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGameServerGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroupProps",
		reflect.TypeOf((*GameServerGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IAlias",
		reflect.TypeOf((*IAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IBuild",
		reflect.TypeOf((*IBuild)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "buildArn", GoGetter: "BuildArn"},
			_jsii_.MemberProperty{JsiiProperty: "buildId", GoGetter: "BuildId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBuild{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IBuildFleet",
		reflect.TypeOf((*IBuildFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fleetArn", GoGetter: "FleetArn"},
			_jsii_.MemberProperty{JsiiProperty: "fleetId", GoGetter: "FleetId"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveInstances", GoMethod: "MetricActiveInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricDesiredInstances", GoMethod: "MetricDesiredInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricIdleInstances", GoMethod: "MetricIdleInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricInstanceInterruptions", GoMethod: "MetricInstanceInterruptions"},
			_jsii_.MemberMethod{JsiiMethod: "metricMaxInstances", GoMethod: "MetricMaxInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricMinInstances", GoMethod: "MetricMinInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricPercentIdleInstances", GoMethod: "MetricPercentIdleInstances"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IBuildFleet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFleet)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IFleet",
		reflect.TypeOf((*IFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "fleetArn", GoGetter: "FleetArn"},
			_jsii_.MemberProperty{JsiiProperty: "fleetId", GoGetter: "FleetId"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricActiveInstances", GoMethod: "MetricActiveInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricDesiredInstances", GoMethod: "MetricDesiredInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricIdleInstances", GoMethod: "MetricIdleInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricInstanceInterruptions", GoMethod: "MetricInstanceInterruptions"},
			_jsii_.MemberMethod{JsiiMethod: "metricMaxInstances", GoMethod: "MetricMaxInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricMinInstances", GoMethod: "MetricMinInstances"},
			_jsii_.MemberMethod{JsiiMethod: "metricPercentIdleInstances", GoMethod: "MetricPercentIdleInstances"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IFleet{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IGameServerGroup",
		reflect.TypeOf((*IGameServerGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingGroupArn", GoGetter: "AutoScalingGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "gameServerGroupArn", GoGetter: "GameServerGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "gameServerGroupName", GoGetter: "GameServerGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IGameServerGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IMatchmakingRuleSet",
		reflect.TypeOf((*IMatchmakingRuleSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "matchmakingRuleSetArn", GoGetter: "MatchmakingRuleSetArn"},
			_jsii_.MemberProperty{JsiiProperty: "matchmakingRuleSetName", GoGetter: "MatchmakingRuleSetName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluationsFailed", GoMethod: "MetricRuleEvaluationsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluationsPassed", GoMethod: "MetricRuleEvaluationsPassed"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IMatchmakingRuleSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IPeer",
		reflect.TypeOf((*IPeer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueId", GoGetter: "UniqueId"},
		},
		func() interface{} {
			return &jsiiProxy_IPeer{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IRuleSetBody",
		reflect.TypeOf((*IRuleSetBody)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IRuleSetBody{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IRuleSetContent",
		reflect.TypeOf((*IRuleSetContent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
		},
		func() interface{} {
			return &jsiiProxy_IRuleSetContent{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-gamelift-alpha.IScript",
		reflect.TypeOf((*IScript)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scriptArn", GoGetter: "ScriptArn"},
			_jsii_.MemberProperty{JsiiProperty: "scriptId", GoGetter: "ScriptId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IScript{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.InstanceDefinition",
		reflect.TypeOf((*InstanceDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.Location",
		reflect.TypeOf((*Location)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.LocationCapacity",
		reflect.TypeOf((*LocationCapacity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		reflect.TypeOf((*MatchmakingRuleSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "matchmakingRuleSetArn", GoGetter: "MatchmakingRuleSetArn"},
			_jsii_.MemberProperty{JsiiProperty: "matchmakingRuleSetName", GoGetter: "MatchmakingRuleSetName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluationsFailed", GoMethod: "MetricRuleEvaluationsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluationsPassed", GoMethod: "MetricRuleEvaluationsPassed"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MatchmakingRuleSet{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_MatchmakingRuleSetBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSetAttributes",
		reflect.TypeOf((*MatchmakingRuleSetAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSetBase",
		reflect.TypeOf((*MatchmakingRuleSetBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "matchmakingRuleSetArn", GoGetter: "MatchmakingRuleSetArn"},
			_jsii_.MemberProperty{JsiiProperty: "matchmakingRuleSetName", GoGetter: "MatchmakingRuleSetName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluationsFailed", GoMethod: "MetricRuleEvaluationsFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricRuleEvaluationsPassed", GoMethod: "MetricRuleEvaluationsPassed"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MatchmakingRuleSetBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMatchmakingRuleSet)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSetProps",
		reflect.TypeOf((*MatchmakingRuleSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-gamelift-alpha.OperatingSystem",
		reflect.TypeOf((*OperatingSystem)(nil)).Elem(),
		map[string]interface{}{
			"AMAZON_LINUX": OperatingSystem_AMAZON_LINUX,
			"AMAZON_LINUX_2": OperatingSystem_AMAZON_LINUX_2,
			"WINDOWS_2012": OperatingSystem_WINDOWS_2012,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.Peer",
		reflect.TypeOf((*Peer)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Peer{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.Port",
		reflect.TypeOf((*Port)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
		},
		func() interface{} {
			return &jsiiProxy_Port{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.PortProps",
		reflect.TypeOf((*PortProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-gamelift-alpha.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": Protocol_TCP,
			"UDP": Protocol_UDP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.ResourceCreationLimitPolicy",
		reflect.TypeOf((*ResourceCreationLimitPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.RuleSetBodyConfig",
		reflect.TypeOf((*RuleSetBodyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.RuleSetContent",
		reflect.TypeOf((*RuleSetContent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
		},
		func() interface{} {
			j := jsiiProxy_RuleSetContent{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRuleSetContent)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.RuleSetContentProps",
		reflect.TypeOf((*RuleSetContentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.RuntimeConfiguration",
		reflect.TypeOf((*RuntimeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.S3Content",
		reflect.TypeOf((*S3Content)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Content{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Content)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.Script",
		reflect.TypeOf((*Script)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "scriptArn", GoGetter: "ScriptArn"},
			_jsii_.MemberProperty{JsiiProperty: "scriptId", GoGetter: "ScriptId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Script{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScriptBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.ScriptAttributes",
		reflect.TypeOf((*ScriptAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-gamelift-alpha.ScriptBase",
		reflect.TypeOf((*ScriptBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "scriptArn", GoGetter: "ScriptArn"},
			_jsii_.MemberProperty{JsiiProperty: "scriptId", GoGetter: "ScriptId"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ScriptBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScript)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.ScriptProps",
		reflect.TypeOf((*ScriptProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-gamelift-alpha.ServerProcess",
		reflect.TypeOf((*ServerProcess)(nil)).Elem(),
	)
}
