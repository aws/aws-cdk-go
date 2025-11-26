package previewawsdmsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceDataSettingsProperty := &SourceDataSettingsProperty{
//   	CdcStartPosition: jsii.String("cdcStartPosition"),
//   	CdcStartTime: jsii.String("cdcStartTime"),
//   	CdcStopTime: jsii.String("cdcStopTime"),
//   	SlotName: jsii.String("slotName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-sourcedatasettings.html
//
type CfnDataMigrationPropsMixin_SourceDataSettingsProperty struct {
	// The property is a point in the database engine's log that defines a time where you can begin CDC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-sourcedatasettings.html#cfn-dms-datamigration-sourcedatasettings-cdcstartposition
	//
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// The property indicates the start time for a change data capture (CDC) operation.
	//
	// The value is server time in UTC format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-sourcedatasettings.html#cfn-dms-datamigration-sourcedatasettings-cdcstarttime
	//
	CdcStartTime *string `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// The property indicates the stop time for a change data capture (CDC) operation.
	//
	// The value is server time in UTC format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-sourcedatasettings.html#cfn-dms-datamigration-sourcedatasettings-cdcstoptime
	//
	CdcStopTime *string `field:"optional" json:"cdcStopTime" yaml:"cdcStopTime"`
	// The property sets the name of a previously created logical replication slot for a change data capture (CDC) load of the source instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-datamigration-sourcedatasettings.html#cfn-dms-datamigration-sourcedatasettings-slotname
	//
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

