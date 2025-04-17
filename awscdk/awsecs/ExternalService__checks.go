//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_ExternalService) validateAddVolumeParameters(volume ServiceManagedVolume) error {
	if volume == nil {
		return fmt.Errorf("parameter volume is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateAssociateCloudMapServiceParameters(_options *AssociateCloudMapServiceOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateAttachToApplicationTargetGroupParameters(_targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	if _targetGroup == nil {
		return fmt.Errorf("parameter _targetGroup is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateAutoScaleTaskCountParameters(_props *awsapplicationautoscaling.EnableScalingProps) error {
	if _props == nil {
		return fmt.Errorf("parameter _props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_props, func() string { return "parameter _props" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(_vpc awsec2.IVpc, _vpcSubnets *awsec2.SubnetSelection) error {
	if _vpc == nil {
		return fmt.Errorf("parameter _vpc is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(_vpcSubnets, func() string { return "parameter _vpcSubnets" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateEnableCloudMapParameters(_options *CloudMapOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateEnableDeploymentAlarmsParameters(alarmNames *[]*string, options *DeploymentAlarmOptions) error {
	if alarmNames == nil {
		return fmt.Errorf("parameter alarmNames is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateEnableServiceConnectParameters(config *ServiceConnectProps) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateLoadBalancerTargetParameters(_options *LoadBalancerTargetOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ExternalService) validateRegisterLoadBalancerTargetsParameters(_targets *[]*EcsTarget) error {
	for idx_d55176, v := range *_targets {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter _targets[%#v]", idx_d55176) }); err != nil {
			return err
		}
	}

	return nil
}

func validateExternalService_FromExternalServiceArnParameters(scope constructs.Construct, id *string, externalServiceArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if externalServiceArn == nil {
		return fmt.Errorf("parameter externalServiceArn is required, but nil was provided")
	}

	return nil
}

func validateExternalService_FromExternalServiceAttributesParameters(scope constructs.Construct, id *string, attrs *ExternalServiceAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateExternalService_FromServiceArnWithClusterParameters(scope constructs.Construct, id *string, serviceArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if serviceArn == nil {
		return fmt.Errorf("parameter serviceArn is required, but nil was provided")
	}

	return nil
}

func validateExternalService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateExternalService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateExternalService_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ExternalService) validateSetDeploymentAlarmsParameters(val *CfnService_DeploymentAlarmsProperty) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_ExternalService) validateSetLoadBalancersParameters(val *[]*CfnService_LoadBalancerProperty) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func (j *jsiiProxy_ExternalService) validateSetNetworkConfigurationParameters(val *CfnService_NetworkConfigurationProperty) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_ExternalService) validateSetServiceRegistriesParameters(val *[]*CfnService_ServiceRegistryProperty) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNewExternalServiceParameters(scope constructs.Construct, id *string, props *ExternalServiceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

