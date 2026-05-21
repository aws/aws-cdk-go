package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   lambdaEventSourceMappingConfigurationProperty := &LambdaEventSourceMappingConfigurationProperty{
//   	Action: jsii.String("action"),
//   	RegionEventSourceMappings: map[string]interface{}{
//   		"regionEventSourceMappingsKey": &EventSourceMappingProperty{
//   			"arn": jsii.String("arn"),
//   			"crossAccountRole": jsii.String("crossAccountRole"),
//   			"externalId": jsii.String("externalId"),
//   		},
//   	},
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &LambdaEventSourceMappingUngracefulProperty{
//   		Behavior: jsii.String("behavior"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingconfiguration.html
//
type CfnPlanPropsMixin_LambdaEventSourceMappingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingconfiguration.html#cfn-arcregionswitch-plan-lambdaeventsourcemappingconfiguration-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingconfiguration.html#cfn-arcregionswitch-plan-lambdaeventsourcemappingconfiguration-regioneventsourcemappings
	//
	RegionEventSourceMappings interface{} `field:"optional" json:"regionEventSourceMappings" yaml:"regionEventSourceMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingconfiguration.html#cfn-arcregionswitch-plan-lambdaeventsourcemappingconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingconfiguration.html#cfn-arcregionswitch-plan-lambdaeventsourcemappingconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

