package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   neptuneGlobalDatabaseConfigurationProperty := &NeptuneGlobalDatabaseConfigurationProperty{
//   	Behavior: jsii.String("behavior"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	RegionDatabaseClusterArns: map[string]*string{
//   		"regionDatabaseClusterArnsKey": jsii.String("regionDatabaseClusterArns"),
//   	},
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &NeptuneUngracefulProperty{
//   		Ungraceful: jsii.String("ungraceful"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html
//
type CfnPlanPropsMixin_NeptuneGlobalDatabaseConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-regiondatabaseclusterarns
	//
	RegionDatabaseClusterArns interface{} `field:"optional" json:"regionDatabaseClusterArns" yaml:"regionDatabaseClusterArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-neptuneglobaldatabaseconfiguration.html#cfn-arcregionswitch-plan-neptuneglobaldatabaseconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

