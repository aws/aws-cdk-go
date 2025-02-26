package awsiot


// Describes an action that writes records into an Amazon Timestream table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timestreamActionProperty := &TimestreamActionProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	Dimensions: []interface{}{
//   		&TimestreamDimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	Timestamp: &TimestreamTimestampProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamaction.html
//
type CfnTopicRule_TimestreamActionProperty struct {
	// The name of an Amazon Timestream database that has the table to write records into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamaction.html#cfn-iot-topicrule-timestreamaction-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Metadata attributes of the time series that are written in each measure record.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamaction.html#cfn-iot-topicrule-timestreamaction-dimensions
	//
	Dimensions interface{} `field:"required" json:"dimensions" yaml:"dimensions"`
	// The Amazon Resource Name (ARN) of the role that grants AWS IoT permission to write to the Timestream database table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamaction.html#cfn-iot-topicrule-timestreamaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The table where the message data will be written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamaction.html#cfn-iot-topicrule-timestreamaction-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The value to use for the entry's timestamp.
	//
	// If blank, the time that the entry was processed is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-timestreamaction.html#cfn-iot-topicrule-timestreamaction-timestamp
	//
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}

