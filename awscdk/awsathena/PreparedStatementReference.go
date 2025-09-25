package awsathena


// A reference to a PreparedStatement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preparedStatementReference := &PreparedStatementReference{
//   	StatementName: jsii.String("statementName"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
type PreparedStatementReference struct {
	// The StatementName of the PreparedStatement resource.
	StatementName *string `field:"required" json:"statementName" yaml:"statementName"`
	// The WorkGroup of the PreparedStatement resource.
	WorkGroup *string `field:"required" json:"workGroup" yaml:"workGroup"`
}

