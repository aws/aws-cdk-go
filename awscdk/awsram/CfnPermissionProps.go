package awsram

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyTemplate interface{}
//
//   cfnPermissionProps := &CfnPermissionProps{
//   	Name: jsii.String("name"),
//   	PolicyTemplate: policyTemplate,
//   	ResourceType: jsii.String("resourceType"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPermissionProps struct {
	// Specifies the name of the customer managed permission.
	//
	// The name must be unique within the AWS Region .
	Name *string `field:"required" json:"name" yaml:"name"`
	// A string in JSON format string that contains the following elements of a resource-based policy:.
	//
	// - *Effect* : must be set to `ALLOW` .
	// - *Action* : specifies the actions that are allowed by this customer managed permission. The list must contain only actions that are supported by the specified resource type. For a list of all actions supported by each resource type, see [Actions, resources, and condition keys for AWS services](https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html) in the *AWS Identity and Access Management User Guide* .
	// - *Condition* : (optional) specifies conditional parameters that must evaluate to true when a user attempts an action for that action to be allowed. For more information about the Condition element, see [IAM policies: Condition element](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html) in the *AWS Identity and Access Management User Guide* .
	//
	// This template can't include either the `Resource` or `Principal` elements. Those are both filled in by AWS RAM when it instantiates the resource-based policy on each resource shared using this managed permission. The `Resource` comes from the ARN of the specific resource that you are sharing. The `Principal` comes from the list of identities added to the resource share.
	PolicyTemplate interface{} `field:"required" json:"policyTemplate" yaml:"policyTemplate"`
	// Specifies the name of the resource type that this customer managed permission applies to.
	//
	// The format is `*<service-code>* : *<resource-type>*` and is not case sensitive. For example, to specify an Amazon EC2 Subnet, you can use the string `ec2:subnet` . To see the list of valid values for this parameter, query the [ListResourceTypes](https://docs.aws.amazon.com/ram/latest/APIReference/API_ListResourceTypes.html) operation.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Specifies a list of one or more tag key and value pairs to attach to the permission.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

