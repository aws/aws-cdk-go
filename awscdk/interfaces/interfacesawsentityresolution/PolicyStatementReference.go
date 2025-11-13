package interfacesawsentityresolution


// A reference to a PolicyStatement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStatementReference := &PolicyStatementReference{
//   	PolicyStatementArn: jsii.String("policyStatementArn"),
//   	StatementId: jsii.String("statementId"),
//   }
//
type PolicyStatementReference struct {
	// The Arn of the PolicyStatement resource.
	PolicyStatementArn *string `field:"required" json:"policyStatementArn" yaml:"policyStatementArn"`
	// The StatementId of the PolicyStatement resource.
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
}

