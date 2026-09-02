package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsSwitchoverReadReplicaConfigurationProperty := &RdsSwitchoverReadReplicaConfigurationProperty{
//   	DbInstanceArnMap: map[string]*string{
//   		"dbInstanceArnMapKey": jsii.String("dbInstanceArnMap"),
//   	},
//
//   	// the properties below are optional
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &RdsUngracefulProperty{
//   		Ungraceful: jsii.String("ungraceful"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration.html
//
type CfnPlan_RdsSwitchoverReadReplicaConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration-dbinstancearnmap
	//
	DbInstanceArnMap interface{} `field:"required" json:"dbInstanceArnMap" yaml:"dbInstanceArnMap"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdsswitchoverreadreplicaconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

