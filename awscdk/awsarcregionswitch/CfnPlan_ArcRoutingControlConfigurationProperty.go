package awsarcregionswitch


// Configuration for ARC routing controls used in a Region switch plan.
//
// Routing controls are simple on/off switches that you can use to shift traffic away from an impaired Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcRoutingControlConfigurationProperty := &ArcRoutingControlConfigurationProperty{
//   	RegionAndRoutingControls: map[string]interface{}{
//   		"regionAndRoutingControlsKey": []interface{}{
//   			&ArcRoutingControlStateProperty{
//   				"routingControlArn": jsii.String("routingControlArn"),
//   				"state": jsii.String("state"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolconfiguration.html
//
type CfnPlan_ArcRoutingControlConfigurationProperty struct {
	// The Region and ARC routing controls for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolconfiguration.html#cfn-arcregionswitch-plan-arcroutingcontrolconfiguration-regionandroutingcontrols
	//
	RegionAndRoutingControls interface{} `field:"required" json:"regionAndRoutingControls" yaml:"regionAndRoutingControls"`
	// The cross account role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolconfiguration.html#cfn-arcregionswitch-plan-arcroutingcontrolconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The external ID (secret key) for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolconfiguration.html#cfn-arcregionswitch-plan-arcroutingcontrolconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-arcroutingcontrolconfiguration.html#cfn-arcregionswitch-plan-arcroutingcontrolconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

