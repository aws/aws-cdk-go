package previewawsopsworksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceProperty := &DataSourceProperty{
//   	Arn: jsii.String("arn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html
//
type CfnAppPropsMixin_DataSourceProperty struct {
	// The data source's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The database name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The data source's type, `AutoSelectOpsworksMysqlInstance` , `OpsworksMysqlInstance` , `RdsDbInstance` , or `None` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

