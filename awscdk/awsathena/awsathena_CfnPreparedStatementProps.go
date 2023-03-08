package awsathena


// Properties for defining a `CfnPreparedStatement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPreparedStatementProps := &cfnPreparedStatementProps{
//   	queryStatement: jsii.String("queryStatement"),
//   	statementName: jsii.String("statementName"),
//   	workGroup: jsii.String("workGroup"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnPreparedStatementProps struct {
	// The query string for the prepared statement.
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// The name of the prepared statement.
	StatementName *string `field:"required" json:"statementName" yaml:"statementName"`
	// The workgroup to which the prepared statement belongs.
	WorkGroup *string `field:"required" json:"workGroup" yaml:"workGroup"`
	// The description of the prepared statement.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

