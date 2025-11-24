package mixinsawsathena


// Properties for CfnPreparedStatementPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPreparedStatementMixinProps := &CfnPreparedStatementMixinProps{
//   	Description: jsii.String("description"),
//   	QueryStatement: jsii.String("queryStatement"),
//   	StatementName: jsii.String("statementName"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html
//
type CfnPreparedStatementMixinProps struct {
	// The description of the prepared statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The query string for the prepared statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-querystatement
	//
	QueryStatement *string `field:"optional" json:"queryStatement" yaml:"queryStatement"`
	// The name of the prepared statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-statementname
	//
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// The workgroup to which the prepared statement belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-workgroup
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

