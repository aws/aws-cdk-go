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

func (b *jsiiProxy_BaseService) validateAddVolumeParameters(volume ServiceManagedVolume) error {
	if volume == nil {
		return fmt.Errorf("parameter volume is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateAssociateCloudMapServiceParameters(options *AssociateCloudMapServiceOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateAttachToClassicLBParameters(loadBalancer awselasticloadbalancing.LoadBalancer) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateAutoScaleTaskCountParameters(props *awsapplicationautoscaling.EnableScalingProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(vpcSubnets, func() string { return "parameter vpcSubnets" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateEnableCloudMapParameters(options *CloudMapOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateEnableDeploymentAlarmsParameters(alarmNames *[]*string, options *DeploymentAlarmOptions) error {
	if alarmNames == nil {
		return fmt.Errorf("parameter alarmNames is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateEnableServiceConnectParameters(config *ServiceConnectProps) error {
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (b *jsiiProxy_BaseService) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateLoadBalancerTargetParameters(options *LoadBalancerTargetOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateMetricCpuUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateMetricMemoryUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BaseService) validateRegisterLoadBalancerTargetsParameters(targets *[]*EcsTarget) error {
	for idx_26cafb, v := range *targets {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter targets[%#v]", idx_26cafb) }); err != nil {
			return err
		}
	}

	return nil
}

func validateBaseService_FromServiceArnWithClusterParameters(scope constructs.Construct, id *string, serviceArn *string) error {
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

func validateBaseService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateBaseService_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBaseService_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseService) validateSetDeploymentAlarmsParameters(val *CfnService_DeploymentAlarmsProperty) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_BaseService) validateSetLoadBalancersParameters(val *[]*CfnService_LoadBalancerProperty) error {
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

func (j *jsiiProxy_BaseService) validateSetNetworkConfigurationParameters(val *CfnService_NetworkConfigurationProperty) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_BaseService) validateSetServiceRegistriesParameters(val *[]*CfnService_ServiceRegistryProperty) error {
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

func validateNewBaseServiceParameters(scope constructs.Construct, id *string, props *BaseServiceProps, additionalProps interface{}, taskDefinition TaskDefinition) error {
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

	if additionalProps == nil {
		return fmt.Errorf("parameter additionalProps is required, but nil was provided")
	}

	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

