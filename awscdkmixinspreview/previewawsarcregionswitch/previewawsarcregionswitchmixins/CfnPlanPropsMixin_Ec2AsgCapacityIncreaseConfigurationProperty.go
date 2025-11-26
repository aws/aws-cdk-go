package previewawsarcregionswitchmixins


// Configuration for increasing the capacity of Amazon EC2 Auto Scaling groups during a Region switch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ec2AsgCapacityIncreaseConfigurationProperty := &Ec2AsgCapacityIncreaseConfigurationProperty{
//   	Asgs: []interface{}{
//   		&AsgProperty{
//   			Arn: jsii.String("arn"),
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   		},
//   	},
//   	CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   	TargetPercent: jsii.Number(123),
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &Ec2UngracefulProperty{
//   		MinimumSuccessPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.html
//
type CfnPlanPropsMixin_Ec2AsgCapacityIncreaseConfigurationProperty struct {
	// The EC2 Auto Scaling groups for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration-asgs
	//
	Asgs interface{} `field:"optional" json:"asgs" yaml:"asgs"`
	// The monitoring approach that you specify EC2 Auto Scaling groups for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration-capacitymonitoringapproach
	//
	CapacityMonitoringApproach *string `field:"optional" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// The target percentage that you specify for EC2 Auto Scaling groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration-targetpercent
	//
	// Default: - 100.
	//
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.html#cfn-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

