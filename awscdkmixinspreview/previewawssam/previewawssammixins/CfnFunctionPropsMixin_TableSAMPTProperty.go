package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableSAMPTProperty := &TableSAMPTProperty{
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-tablesampt.html
//
type CfnFunctionPropsMixin_TableSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-tablesampt.html#cfn-serverless-function-tablesampt-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

