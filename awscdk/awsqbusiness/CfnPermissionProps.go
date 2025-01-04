package awsqbusiness


// Properties for defining a `CfnPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionProps := &CfnPermissionProps{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ApplicationId: jsii.String("applicationId"),
//   	Principal: jsii.String("principal"),
//   	StatementId: jsii.String("statementId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html
//
type CfnPermissionProps struct {
	// The list of Amazon Q Business actions that the ISV is allowed to perform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The unique identifier of the Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Provides user and group information used for filtering documents to use for generating Amazon Q Business conversation responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// A unique identifier for the policy statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-permission.html#cfn-qbusiness-permission-statementid
	//
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
}

