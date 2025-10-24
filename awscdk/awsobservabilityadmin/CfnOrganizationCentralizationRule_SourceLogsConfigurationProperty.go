package awsobservabilityadmin


// Configuration for selecting and handling source log groups for centralization.
//
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
	// A strategy determining whether to centralize source log groups that are encrypted with customer managed KMS keys (CMK).
	//
	// ALLOW will consider CMK encrypted source log groups for centralization while SKIP will skip CMK encrypted source log groups from centralization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration-encryptedloggroupstrategy
	//
	EncryptedLogGroupStrategy *string `field:"required" json:"encryptedLogGroupStrategy" yaml:"encryptedLogGroupStrategy"`
	// The selection criteria that specifies which source log groups to centralize.
	//
	// The selection criteria uses the same format as OAM link filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-sourcelogsconfiguration-loggroupselectioncriteria
	//
	LogGroupSelectionCriteria *string `field:"required" json:"logGroupSelectionCriteria" yaml:"logGroupSelectionCriteria"`
}

