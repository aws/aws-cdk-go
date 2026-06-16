package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraServerlessScalingConfigurationProperty := &AuroraServerlessScalingConfigurationProperty{
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	RegionDatabaseClusterArns: map[string]*string{
//   		"regionDatabaseClusterArnsKey": jsii.String("regionDatabaseClusterArns"),
//   	},
//
//   	// the properties below are optional
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	TargetPercent: jsii.Number(123),
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html
//
type CfnPlan_AuroraServerlessScalingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html#cfn-arcregionswitch-plan-auroraserverlessscalingconfiguration-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"required" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html#cfn-arcregionswitch-plan-auroraserverlessscalingconfiguration-regiondatabaseclusterarns
	//
	RegionDatabaseClusterArns interface{} `field:"required" json:"regionDatabaseClusterArns" yaml:"regionDatabaseClusterArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html#cfn-arcregionswitch-plan-auroraserverlessscalingconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html#cfn-arcregionswitch-plan-auroraserverlessscalingconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html#cfn-arcregionswitch-plan-auroraserverlessscalingconfiguration-targetpercent
	//
	// Default: - 100.
	//
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraserverlessscalingconfiguration.html#cfn-arcregionswitch-plan-auroraserverlessscalingconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

