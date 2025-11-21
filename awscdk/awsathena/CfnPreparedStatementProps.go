package awsathena


// Properties for defining a `CfnPreparedStatement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPreparedStatementProps := &CfnPreparedStatementProps{
//   	QueryStatement: jsii.String("queryStatement"),
//   	StatementName: jsii.String("statementName"),
//   	WorkGroup: jsii.String("workGroup"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html
//
type CfnPreparedStatementProps struct {
	// The query string for the prepared statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-querystatement
	//
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// The name of the prepared statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-statementname
	//
	StatementName *string `field:"required" json:"statementName" yaml:"statementName"`
	// The workgroup to which the prepared statement belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-workgroup
	//
	WorkGroup interface{} `field:"required" json:"workGroup" yaml:"workGroup"`
	// The description of the prepared statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html#cfn-athena-preparedstatement-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

