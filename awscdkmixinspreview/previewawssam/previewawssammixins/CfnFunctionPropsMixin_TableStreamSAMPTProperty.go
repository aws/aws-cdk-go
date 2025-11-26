package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableStreamSAMPTProperty := &TableStreamSAMPTProperty{
//   	StreamName: jsii.String("streamName"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-tablestreamsampt.html
//
type CfnFunctionPropsMixin_TableStreamSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-tablestreamsampt.html#cfn-serverless-function-tablestreamsampt-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-tablestreamsampt.html#cfn-serverless-function-tablestreamsampt-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

