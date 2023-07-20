package awssagemaker


// A list of user groups that exist in your OIDC Identity Provider (IdP).
//
// One to ten groups can be used to create a single private work team. When you add a user group to the list of `Groups` , you can add that user group to one or more private work teams. If you add a user group to a private work team, all workers in that user group are added to the work team.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcMemberDefinitionProperty := &OidcMemberDefinitionProperty{
//   	OidcGroups: []*string{
//   		jsii.String("oidcGroups"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-oidcmemberdefinition.html
//
type CfnWorkteam_OidcMemberDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-oidcmemberdefinition.html#cfn-sagemaker-workteam-oidcmemberdefinition-oidcgroups
	//
	OidcGroups *[]*string `field:"required" json:"oidcGroups" yaml:"oidcGroups"`
}

