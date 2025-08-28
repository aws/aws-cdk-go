package awsssmquicksetup


// A summarized description of the status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statusSummaryProperty := &StatusSummaryProperty{
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	StatusType: jsii.String("statusType"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   	StatusDetails: map[string]*string{
//   		"statusDetailsKey": jsii.String("statusDetails"),
//   	},
//   	StatusMessage: jsii.String("statusMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-statussummary.html
//
type CfnConfigurationManager_StatusSummaryProperty struct {
	// The datetime stamp when the status was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-statussummary.html#cfn-ssmquicksetup-configurationmanager-statussummary-lastupdatedat
	//
	LastUpdatedAt *string `field:"required" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// The type of a status summary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-statussummary.html#cfn-ssmquicksetup-configurationmanager-statussummary-statustype
	//
	StatusType *string `field:"required" json:"statusType" yaml:"statusType"`
	// The current status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-statussummary.html#cfn-ssmquicksetup-configurationmanager-statussummary-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Details about the status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-statussummary.html#cfn-ssmquicksetup-configurationmanager-statussummary-statusdetails
	//
	StatusDetails interface{} `field:"optional" json:"statusDetails" yaml:"statusDetails"`
	// When applicable, returns an informational message relevant to the current status and status type of the status summary object.
	//
	// We don't recommend implementing parsing logic around this value since the messages returned can vary in format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmquicksetup-configurationmanager-statussummary.html#cfn-ssmquicksetup-configurationmanager-statussummary-statusmessage
	//
	StatusMessage *string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}

