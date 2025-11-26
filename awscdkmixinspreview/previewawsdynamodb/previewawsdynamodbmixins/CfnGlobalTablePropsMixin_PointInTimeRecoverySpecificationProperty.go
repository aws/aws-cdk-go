package previewawsdynamodbmixins


// Represents the settings used to enable point in time recovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pointInTimeRecoverySpecificationProperty := &PointInTimeRecoverySpecificationProperty{
//   	PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	RecoveryPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-pointintimerecoveryspecification.html
//
type CfnGlobalTablePropsMixin_PointInTimeRecoverySpecificationProperty struct {
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-pointintimerecoveryspecification.html#cfn-dynamodb-globaltable-pointintimerecoveryspecification-pointintimerecoveryenabled
	//
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// The number of preceding days for which continuous backups are taken and maintained.
	//
	// Your table data is only recoverable to any point-in-time from within the configured recovery period. This parameter is optional. If no value is provided, the value will default to 35.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-pointintimerecoveryspecification.html#cfn-dynamodb-globaltable-pointintimerecoveryspecification-recoveryperiodindays
	//
	RecoveryPeriodInDays *float64 `field:"optional" json:"recoveryPeriodInDays" yaml:"recoveryPeriodInDays"`
}

