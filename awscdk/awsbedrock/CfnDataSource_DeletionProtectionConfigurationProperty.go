package awsbedrock


// Configuration for deletion protection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deletionProtectionConfigurationProperty := &DeletionProtectionConfigurationProperty{
//   	DeletionProtectionStatus: jsii.String("deletionProtectionStatus"),
//
//   	// the properties below are optional
//   	DeletionProtectionThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-deletionprotectionconfiguration.html
//
type CfnDataSource_DeletionProtectionConfigurationProperty struct {
	// Indicates whether a feature is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-deletionprotectionconfiguration.html#cfn-bedrock-datasource-deletionprotectionconfiguration-deletionprotectionstatus
	//
	DeletionProtectionStatus *string `field:"required" json:"deletionProtectionStatus" yaml:"deletionProtectionStatus"`
	// Threshold for deletion protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-deletionprotectionconfiguration.html#cfn-bedrock-datasource-deletionprotectionconfiguration-deletionprotectionthreshold
	//
	// Default: - 15.
	//
	DeletionProtectionThreshold *float64 `field:"optional" json:"deletionProtectionThreshold" yaml:"deletionProtectionThreshold"`
}

