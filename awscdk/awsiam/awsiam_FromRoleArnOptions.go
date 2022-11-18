package awsiam


// Options allowing customizing the behavior of {@link Role.fromRoleArn}.
//
// Example:
//   role := iam.role.fromRoleArn(this, jsii.String("Role"), jsii.String("arn:aws:iam::123456789012:role/MyExistingRole"), &fromRoleArnOptions{
//   	// Set 'mutable' to 'false' to use the role as-is and prevent adding new
//   	// policies to it. The default is 'true', which means the role may be
//   	// modified as part of the deployment.
//   	mutable: jsii.Boolean(false),
//   })
//
type FromRoleArnOptions struct {
	// For immutable roles: add grants to resources instead of dropping them.
	//
	// If this is `false` or not specified, grant permissions added to this role are ignored.
	// It is your own responsibility to make sure the role has the required permissions.
	//
	// If this is `true`, any grant permissions will be added to the resource instead.
	AddGrantsToResources *bool `field:"optional" json:"addGrantsToResources" yaml:"addGrantsToResources"`
	// Any policies created by this role will use this value as their ID, if specified.
	//
	// Specify this if importing the same role in multiple stacks, and granting it
	// different permissions in at least two stacks. If this is not specified
	// (or if the same name is specified in more than one stack),
	// a CloudFormation issue will result in the policy created in whichever stack
	// is deployed last overwriting the policies created by the others.
	DefaultPolicyName *string `field:"optional" json:"defaultPolicyName" yaml:"defaultPolicyName"`
	// Whether the imported role can be modified by attaching policy resources to it.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

