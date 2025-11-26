package previewawscassandramixins


// The optional auto scaling capacity settings for a table in provisioned capacity mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoScalingSpecificationProperty := &AutoScalingSpecificationProperty{
//   	ReadCapacityAutoScaling: &AutoScalingSettingProperty{
//   		AutoScalingDisabled: jsii.Boolean(false),
//   		MaximumUnits: jsii.Number(123),
//   		MinimumUnits: jsii.Number(123),
//   		ScalingPolicy: &ScalingPolicyProperty{
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	WriteCapacityAutoScaling: &AutoScalingSettingProperty{
//   		AutoScalingDisabled: jsii.Boolean(false),
//   		MaximumUnits: jsii.Number(123),
//   		MinimumUnits: jsii.Number(123),
//   		ScalingPolicy: &ScalingPolicyProperty{
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingspecification.html
//
type CfnTablePropsMixin_AutoScalingSpecificationProperty struct {
	// The auto scaling settings for the table's read capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingspecification.html#cfn-cassandra-table-autoscalingspecification-readcapacityautoscaling
	//
	ReadCapacityAutoScaling interface{} `field:"optional" json:"readCapacityAutoScaling" yaml:"readCapacityAutoScaling"`
	// The auto scaling settings for the table's write capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-autoscalingspecification.html#cfn-cassandra-table-autoscalingspecification-writecapacityautoscaling
	//
	WriteCapacityAutoScaling interface{} `field:"optional" json:"writeCapacityAutoScaling" yaml:"writeCapacityAutoScaling"`
}

