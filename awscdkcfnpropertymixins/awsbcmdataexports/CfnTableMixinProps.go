package awsbcmdataexports


// Properties for CfnTablePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTableMixinProps := &CfnTableMixinProps{
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-table.html
//
type CfnTableMixinProps struct {
	// The name of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-table.html#cfn-bcmdataexports-table-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

