package previewawsopsworksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeBasedAutoScalingProperty := &TimeBasedAutoScalingProperty{
//   	Friday: map[string]*string{
//   		"fridayKey": jsii.String("friday"),
//   	},
//   	Monday: map[string]*string{
//   		"mondayKey": jsii.String("monday"),
//   	},
//   	Saturday: map[string]*string{
//   		"saturdayKey": jsii.String("saturday"),
//   	},
//   	Sunday: map[string]*string{
//   		"sundayKey": jsii.String("sunday"),
//   	},
//   	Thursday: map[string]*string{
//   		"thursdayKey": jsii.String("thursday"),
//   	},
//   	Tuesday: map[string]*string{
//   		"tuesdayKey": jsii.String("tuesday"),
//   	},
//   	Wednesday: map[string]*string{
//   		"wednesdayKey": jsii.String("wednesday"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html
//
type CfnInstancePropsMixin_TimeBasedAutoScalingProperty struct {
	// The schedule for Friday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-friday
	//
	Friday interface{} `field:"optional" json:"friday" yaml:"friday"`
	// The schedule for Monday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-monday
	//
	Monday interface{} `field:"optional" json:"monday" yaml:"monday"`
	// The schedule for Saturday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-saturday
	//
	Saturday interface{} `field:"optional" json:"saturday" yaml:"saturday"`
	// The schedule for Sunday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-sunday
	//
	Sunday interface{} `field:"optional" json:"sunday" yaml:"sunday"`
	// The schedule for Thursday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-thursday
	//
	Thursday interface{} `field:"optional" json:"thursday" yaml:"thursday"`
	// The schedule for Tuesday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-tuesday
	//
	Tuesday interface{} `field:"optional" json:"tuesday" yaml:"tuesday"`
	// The schedule for Wednesday.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-wednesday
	//
	Wednesday interface{} `field:"optional" json:"wednesday" yaml:"wednesday"`
}

