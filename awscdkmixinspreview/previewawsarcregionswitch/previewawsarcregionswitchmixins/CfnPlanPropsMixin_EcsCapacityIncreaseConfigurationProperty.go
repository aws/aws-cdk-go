package previewawsarcregionswitchmixins


// The configuration for an AWS ECS capacity increase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ecsCapacityIncreaseConfigurationProperty := &EcsCapacityIncreaseConfigurationProperty{
//   	CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   	Services: []interface{}{
//   		&ServiceProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   			ServiceArn: jsii.String("serviceArn"),
//   		},
//   	},
//   	TargetPercent: jsii.Number(123),
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &EcsUngracefulProperty{
//   		MinimumSuccessPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.html
//
type CfnPlanPropsMixin_EcsCapacityIncreaseConfigurationProperty struct {
	// The monitoring approach specified for the configuration, for example, `Most_Recent` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ecscapacityincreaseconfiguration-capacitymonitoringapproach
	//
	CapacityMonitoringApproach *string `field:"optional" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// The services specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ecscapacityincreaseconfiguration-services
	//
	Services interface{} `field:"optional" json:"services" yaml:"services"`
	// The target percentage specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ecscapacityincreaseconfiguration-targetpercent
	//
	// Default: - 100.
	//
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ecscapacityincreaseconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ecscapacityincreaseconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

