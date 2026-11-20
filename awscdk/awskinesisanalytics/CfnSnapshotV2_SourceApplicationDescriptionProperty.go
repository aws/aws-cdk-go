package awskinesisanalytics


// Details about the application the snapshot was taken from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceApplicationDescriptionProperty := &SourceApplicationDescriptionProperty{
//   	ApplicationCreationTimestamp: jsii.String("applicationCreationTimestamp"),
//   	ApplicationName: jsii.String("applicationName"),
//   	ApplicationVersionId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-sourceapplicationdescription.html
//
type CfnSnapshotV2_SourceApplicationDescriptionProperty struct {
	// The time at which the source application was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-sourceapplicationdescription.html#cfn-kinesisanalyticsv2-snapshot-sourceapplicationdescription-applicationcreationtimestamp
	//
	ApplicationCreationTimestamp *string `field:"required" json:"applicationCreationTimestamp" yaml:"applicationCreationTimestamp"`
	// The name of the source application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-sourceapplicationdescription.html#cfn-kinesisanalyticsv2-snapshot-sourceapplicationdescription-applicationname
	//
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The version of the source application captured in the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-sourceapplicationdescription.html#cfn-kinesisanalyticsv2-snapshot-sourceapplicationdescription-applicationversionid
	//
	ApplicationVersionId *float64 `field:"required" json:"applicationVersionId" yaml:"applicationVersionId"`
}

