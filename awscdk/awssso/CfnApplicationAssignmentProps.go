package awssso


// Properties for defining a `CfnApplicationAssignment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationAssignmentProps := &CfnApplicationAssignmentProps{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	PrincipalId: jsii.String("principalId"),
//   	PrincipalType: jsii.String("principalType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-applicationassignment.html
//
type CfnApplicationAssignmentProps struct {
	// The ARN of the application that has principals assigned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-applicationassignment.html#cfn-sso-applicationassignment-applicationarn
	//
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The unique identifier of the principal assigned to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-applicationassignment.html#cfn-sso-applicationassignment-principalid
	//
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The type of the principal assigned to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-applicationassignment.html#cfn-sso-applicationassignment-principaltype
	//
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
}

