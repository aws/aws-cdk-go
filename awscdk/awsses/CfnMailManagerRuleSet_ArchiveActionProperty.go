package awsses


// The action to archive the email by delivering the email to an Amazon SES archive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveActionProperty := &ArchiveActionProperty{
//   	TargetArchive: jsii.String("targetArchive"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-archiveaction.html
//
type CfnMailManagerRuleSet_ArchiveActionProperty struct {
	// The identifier of the archive to send the email to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-archiveaction.html#cfn-ses-mailmanagerruleset-archiveaction-targetarchive
	//
	TargetArchive *string `field:"required" json:"targetArchive" yaml:"targetArchive"`
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, the specified archive has been deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-archiveaction.html#cfn-ses-mailmanagerruleset-archiveaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
}

