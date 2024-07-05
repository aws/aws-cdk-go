package awskinesisanalyticsv2


// Describes whether system initiated rollbacks are enabled for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSystemRollbackConfigurationProperty := &ApplicationSystemRollbackConfigurationProperty{
//   	RollbackEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsystemrollbackconfiguration.html
//
type CfnApplication_ApplicationSystemRollbackConfigurationProperty struct {
	// Describes whether system initiated rollbacks are enabled for a Flink-based Kinesis Data Analytics application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationsystemrollbackconfiguration.html#cfn-kinesisanalyticsv2-application-applicationsystemrollbackconfiguration-rollbackenabled
	//
	RollbackEnabled interface{} `field:"required" json:"rollbackEnabled" yaml:"rollbackEnabled"`
}

