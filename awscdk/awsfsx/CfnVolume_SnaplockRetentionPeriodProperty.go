package awsfsx


// The configuration to set the retention period of an FSx for ONTAP SnapLock volume.
//
// The retention period includes default, maximum, and minimum settings. For more information, see [Working with the retention period in SnapLock](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snaplock-retention.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snaplockRetentionPeriodProperty := &SnaplockRetentionPeriodProperty{
//   	DefaultRetention: &RetentionPeriodProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Value: jsii.Number(123),
//   	},
//   	MaximumRetention: &RetentionPeriodProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Value: jsii.Number(123),
//   	},
//   	MinimumRetention: &RetentionPeriodProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockretentionperiod.html
//
type CfnVolume_SnaplockRetentionPeriodProperty struct {
	// The retention period assigned to a write once, read many (WORM) file by default if an explicit retention period is not set for an FSx for ONTAP SnapLock volume.
	//
	// The default retention period must be greater than or equal to the minimum retention period and less than or equal to the maximum retention period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockretentionperiod.html#cfn-fsx-volume-snaplockretentionperiod-defaultretention
	//
	DefaultRetention interface{} `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
	// The longest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockretentionperiod.html#cfn-fsx-volume-snaplockretentionperiod-maximumretention
	//
	MaximumRetention interface{} `field:"required" json:"maximumRetention" yaml:"maximumRetention"`
	// The shortest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-snaplockretentionperiod.html#cfn-fsx-volume-snaplockretentionperiod-minimumretention
	//
	MinimumRetention interface{} `field:"required" json:"minimumRetention" yaml:"minimumRetention"`
}

