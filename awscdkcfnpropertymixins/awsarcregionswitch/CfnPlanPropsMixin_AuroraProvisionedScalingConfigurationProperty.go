package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   auroraProvisionedScalingConfigurationProperty := &AuroraProvisionedScalingConfigurationProperty{
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	InstanceArns: map[string]*string{
//   		"instanceArnsKey": jsii.String("instanceArns"),
//   	},
//   	RegionDatabaseClusterArns: map[string]*string{
//   		"regionDatabaseClusterArnsKey": jsii.String("regionDatabaseClusterArns"),
//   	},
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html
//
type CfnPlanPropsMixin_AuroraProvisionedScalingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html#cfn-arcregionswitch-plan-auroraprovisionedscalingconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html#cfn-arcregionswitch-plan-auroraprovisionedscalingconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html#cfn-arcregionswitch-plan-auroraprovisionedscalingconfiguration-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html#cfn-arcregionswitch-plan-auroraprovisionedscalingconfiguration-instancearns
	//
	InstanceArns interface{} `field:"optional" json:"instanceArns" yaml:"instanceArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html#cfn-arcregionswitch-plan-auroraprovisionedscalingconfiguration-regiondatabaseclusterarns
	//
	RegionDatabaseClusterArns interface{} `field:"optional" json:"regionDatabaseClusterArns" yaml:"regionDatabaseClusterArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-auroraprovisionedscalingconfiguration.html#cfn-arcregionswitch-plan-auroraprovisionedscalingconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

