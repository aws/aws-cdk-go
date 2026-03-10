package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsCreateCrossRegionReplicaConfigurationProperty := &RdsCreateCrossRegionReplicaConfigurationProperty{
//   	DbInstanceArnMap: map[string]*string{
//   		"dbInstanceArnMapKey": jsii.String("dbInstanceArnMap"),
//   	},
//
//   	// the properties below are optional
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration.html
//
type CfnPlan_RdsCreateCrossRegionReplicaConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration.html#cfn-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration-dbinstancearnmap
	//
	DbInstanceArnMap interface{} `field:"required" json:"dbInstanceArnMap" yaml:"dbInstanceArnMap"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration.html#cfn-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration.html#cfn-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration.html#cfn-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

