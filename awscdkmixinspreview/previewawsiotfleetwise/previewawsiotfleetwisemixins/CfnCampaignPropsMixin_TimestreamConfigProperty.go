package previewawsiotfleetwisemixins


// The Amazon Timestream table where the AWS IoT FleetWise campaign sends data.
//
// Timestream stores and organizes data to optimize query processing time and to reduce storage costs. For more information, see [Data modeling](https://docs.aws.amazon.com/timestream/latest/developerguide/data-modeling.html) in the *Amazon Timestream Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timestreamConfigProperty := &TimestreamConfigProperty{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	TimestreamTableArn: jsii.String("timestreamTableArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timestreamconfig.html
//
type CfnCampaignPropsMixin_TimestreamConfigProperty struct {
	// The Amazon Resource Name (ARN) of the task execution role that grants AWS IoT FleetWise permission to deliver data to the Amazon Timestream table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timestreamconfig.html#cfn-iotfleetwise-campaign-timestreamconfig-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The Amazon Resource Name (ARN) of the Amazon Timestream table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timestreamconfig.html#cfn-iotfleetwise-campaign-timestreamconfig-timestreamtablearn
	//
	TimestreamTableArn *string `field:"optional" json:"timestreamTableArn" yaml:"timestreamTableArn"`
}

