package awsdynamodb


// Represents the settings used to enable point in time recovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pointInTimeRecoverySpecificationProperty := &pointInTimeRecoverySpecificationProperty{
//   	pointInTimeRecoveryEnabled: jsii.Boolean(false),
//   }
//
type CfnGlobalTable_PointInTimeRecoverySpecificationProperty struct {
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
}

