package previewawsfsxmixins


// Sets the autocommit period of files in an FSx for ONTAP SnapLock volume, which determines how long the files must remain unmodified before they're automatically transitioned to the write once, read many (WORM) state.
//
// For more information, see [Autocommit](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/worm-state.html#worm-state-autocommit) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autocommitPeriodProperty := &AutocommitPeriodProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-autocommitperiod.html
//
type CfnVolumePropsMixin_AutocommitPeriodProperty struct {
	// Defines the type of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
	//
	// Setting this value to `NONE` disables autocommit. The default value is `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-autocommitperiod.html#cfn-fsx-volume-autocommitperiod-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Defines the amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
	//
	// The following ranges are valid:
	//
	// - `Minutes` : 5 - 65,535
	// - `Hours` : 1 - 65,535
	// - `Days` : 1 - 3,650
	// - `Months` : 1 - 120
	// - `Years` : 1 - 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-autocommitperiod.html#cfn-fsx-volume-autocommitperiod-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

