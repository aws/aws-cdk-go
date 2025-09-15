package awsdynamodb


// Reference to PointInTimeRecovey Specification for continuous backups.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ContributorInsightsSpecification: &ContributorInsightsSpecification{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecification{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(true),
//   	},
//   })
//
type PointInTimeRecoverySpecification struct {
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	// Default: false.
	//
	PointInTimeRecoveryEnabled *bool `field:"required" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// The number of preceding days for which continuous backups are taken and maintained.
	//
	// Your table data is only recoverable to any point-in-time from within the configured recovery period.
	// If no value is provided, the value will default to 35.
	// Default: 35.
	//
	RecoveryPeriodInDays *float64 `field:"optional" json:"recoveryPeriodInDays" yaml:"recoveryPeriodInDays"`
}

