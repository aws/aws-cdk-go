package awsdatazone


// The group principal to whom the policy is granted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupPolicyGrantPrincipalProperty := &GroupPolicyGrantPrincipalProperty{
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-grouppolicygrantprincipal.html
//
type CfnPolicyGrant_GroupPolicyGrantPrincipalProperty struct {
	// The ID Of the group of the group principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-grouppolicygrantprincipal.html#cfn-datazone-policygrant-grouppolicygrantprincipal-groupidentifier
	//
	GroupIdentifier *string `field:"required" json:"groupIdentifier" yaml:"groupIdentifier"`
}

