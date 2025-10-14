package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceLogsConfigurationProperty := &SourceLogsConfigurationProperty{
//   	EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   	LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration.html
//
type CfnOrganizationCentralizationRule_SourceLogsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration-encryptedloggroupstrategy
	//
	EncryptedLogGroupStrategy *string `field:"required" json:"encryptedLogGroupStrategy" yaml:"encryptedLogGroupStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration-loggroupselectioncriteria
	//
	LogGroupSelectionCriteria *string `field:"required" json:"logGroupSelectionCriteria" yaml:"logGroupSelectionCriteria"`
}

