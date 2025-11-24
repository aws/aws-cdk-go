package mixinsawsses


// The retention policy for an email archive that specifies how long emails are kept before being automatically deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archiveRetentionProperty := &ArchiveRetentionProperty{
//   	RetentionPeriod: jsii.String("retentionPeriod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerarchive-archiveretention.html
//
type CfnMailManagerArchivePropsMixin_ArchiveRetentionProperty struct {
	// The enum value sets the period for retaining emails in an archive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerarchive-archiveretention.html#cfn-ses-mailmanagerarchive-archiveretention-retentionperiod
	//
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

