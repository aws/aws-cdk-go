package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnInfluxDBBackupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnInfluxDBBackupMixinProps := &CfnInfluxDBBackupMixinProps{
//   	DbResourceId: jsii.String("dbResourceId"),
//   	Name: jsii.String("name"),
//   	RetentionDays: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html
//
type CfnInfluxDBBackupMixinProps struct {
	// The id of the DB instance or DB cluster to back up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html#cfn-timestream-influxdbbackup-dbresourceid
	//
	DbResourceId *string `field:"optional" json:"dbResourceId" yaml:"dbResourceId"`
	// The name of the backup.
	//
	// Must be unique within the account and region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html#cfn-timestream-influxdbbackup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of days to retain the backup.
	//
	// Valid values are 1 to 3650.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html#cfn-timestream-influxdbbackup-retentiondays
	//
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// A list of key-value pairs to associate with the backup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html#cfn-timestream-influxdbbackup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

