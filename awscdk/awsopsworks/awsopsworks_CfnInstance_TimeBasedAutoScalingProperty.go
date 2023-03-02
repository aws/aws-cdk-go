package awsopsworks


// Describes an instance's time-based auto scaling configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedAutoScalingProperty := &timeBasedAutoScalingProperty{
//   	friday: map[string]*string{
//   		"fridayKey": jsii.String("friday"),
//   	},
//   	monday: map[string]*string{
//   		"mondayKey": jsii.String("monday"),
//   	},
//   	saturday: map[string]*string{
//   		"saturdayKey": jsii.String("saturday"),
//   	},
//   	sunday: map[string]*string{
//   		"sundayKey": jsii.String("sunday"),
//   	},
//   	thursday: map[string]*string{
//   		"thursdayKey": jsii.String("thursday"),
//   	},
//   	tuesday: map[string]*string{
//   		"tuesdayKey": jsii.String("tuesday"),
//   	},
//   	wednesday: map[string]*string{
//   		"wednesdayKey": jsii.String("wednesday"),
//   	},
//   }
//
type CfnInstance_TimeBasedAutoScalingProperty struct {
	// The schedule for Friday.
	Friday interface{} `field:"optional" json:"friday" yaml:"friday"`
	// The schedule for Monday.
	Monday interface{} `field:"optional" json:"monday" yaml:"monday"`
	// The schedule for Saturday.
	Saturday interface{} `field:"optional" json:"saturday" yaml:"saturday"`
	// The schedule for Sunday.
	Sunday interface{} `field:"optional" json:"sunday" yaml:"sunday"`
	// The schedule for Thursday.
	Thursday interface{} `field:"optional" json:"thursday" yaml:"thursday"`
	// The schedule for Tuesday.
	Tuesday interface{} `field:"optional" json:"tuesday" yaml:"tuesday"`
	// The schedule for Wednesday.
	Wednesday interface{} `field:"optional" json:"wednesday" yaml:"wednesday"`
}

