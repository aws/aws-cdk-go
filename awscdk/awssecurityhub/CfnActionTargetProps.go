package awssecurityhub


// Properties for defining a `CfnActionTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnActionTargetProps := &CfnActionTargetProps{
//   	Description: jsii.String("description"),
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-actiontarget.html
//
type CfnActionTargetProps struct {
	// The description for the custom action target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-actiontarget.html#cfn-securityhub-actiontarget-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The ID for the custom action target.
	//
	// Can contain up to 20 alphanumeric characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-actiontarget.html#cfn-securityhub-actiontarget-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The name of the custom action target.
	//
	// Can contain up to 20 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-actiontarget.html#cfn-securityhub-actiontarget-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

