package awsiam


// Options allowing customizing the behavior of {@link Role.fromRoleArn}.
//
// Example:
//   role := iam.Role_FromRoleArn(this, jsii.String("Role"), jsii.String("arn:aws:iam::123456789012:role/MyExistingRole"), &FromRoleArnOptions{
//   	// Set 'mutable' to 'false' to use the role as-is and prevent adding new
//   	// policies to it. The default is 'true', which means the role may be
//   	// modified as part of the deployment.
//   	Mutable: jsii.Boolean(false),
//   })
//
// Experimental.
type FromRoleArnOptions struct {
	// For immutable roles: add grants to resources instead of dropping them.
	//
	// If this is `false` or not specified, grant permissions added to this role are ignored.
	// It is your own responsibility to make sure the role has the required permissions.
	//
	// If this is `true`, any grant permissions will be added to the resource instead.
	// Experimental.
	AddGrantsToResources *bool `field:"optional" json:"addGrantsToResources" yaml:"addGrantsToResources"`
	// Whether the imported role can be modified by attaching policy resources to it.
	// Experimental.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

