package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rdsPromoteReadReplicaConfigurationProperty := &RdsPromoteReadReplicaConfigurationProperty{
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	DbInstanceArnMap: map[string]*string{
//   		"dbInstanceArnMapKey": jsii.String("dbInstanceArnMap"),
//   	},
//   	ExternalId: jsii.String("externalId"),
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdspromotereadreplicaconfiguration.html
//
type CfnPlanPropsMixin_RdsPromoteReadReplicaConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdspromotereadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdspromotereadreplicaconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdspromotereadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdspromotereadreplicaconfiguration-dbinstancearnmap
	//
	DbInstanceArnMap interface{} `field:"optional" json:"dbInstanceArnMap" yaml:"dbInstanceArnMap"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdspromotereadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdspromotereadreplicaconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-rdspromotereadreplicaconfiguration.html#cfn-arcregionswitch-plan-rdspromotereadreplicaconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

