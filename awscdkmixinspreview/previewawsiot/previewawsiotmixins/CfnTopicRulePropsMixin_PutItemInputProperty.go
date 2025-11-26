package previewawsiotmixins


// The input for the DynamoActionVS action that specifies the DynamoDB table to which the message data will be written.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   putItemInputProperty := &PutItemInputProperty{
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putiteminput.html
//
type CfnTopicRulePropsMixin_PutItemInputProperty struct {
	// The table where the message data will be written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putiteminput.html#cfn-iot-topicrule-putiteminput-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

