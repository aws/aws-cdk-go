package mixinsawsfsx


// Specifies the retention period of an FSx for ONTAP SnapLock volume.
//
// After it is set, it can't be changed. Files can't be deleted or modified during the retention period.
//
// For more information, see [Working with the retention period in SnapLock](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/snaplock-retention.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionPeriodProperty := &RetentionPeriodProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-retentionperiod.html
//
type CfnVolumePropsMixin_RetentionPeriodProperty struct {
	// Defines the type of time for the retention period of an FSx for ONTAP SnapLock volume.
	//
	// Set it to one of the valid types. If you set it to `INFINITE` , the files are retained forever. If you set it to `UNSPECIFIED` , the files are retained until you set an explicit retention period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-retentionperiod.html#cfn-fsx-volume-retentionperiod-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Defines the amount of time for the retention period of an FSx for ONTAP SnapLock volume.
	//
	// You can't set a value for `INFINITE` or `UNSPECIFIED` . For all other options, the following ranges are valid:
	//
	// - `Seconds` : 0 - 65,535
	// - `Minutes` : 0 - 65,535
	// - `Hours` : 0 - 24
	// - `Days` : 0 - 365
	// - `Months` : 0 - 12
	// - `Years` : 0 - 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-retentionperiod.html#cfn-fsx-volume-retentionperiod-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

