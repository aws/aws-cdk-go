package awsiam


// Properties for looking up an existing Role.
//
// Example:
//   role := iam.Role_FromLookup(this, jsii.String("Role"), &RoleLookupOptions{
//   	RoleName: jsii.String("MyExistingRole"),
//   })
//
type RoleLookupOptions struct {
	// For immutable roles: add grants to resources instead of dropping them.
	//
	// If this is `false` or not specified, grant permissions added to this role are ignored.
	// It is your own responsibility to make sure the role has the required permissions.
	//
	// If this is `true`, any grant permissions will be added to the resource instead.
	// Default: false.
	//
	AddGrantsToResources *bool `field:"optional" json:"addGrantsToResources" yaml:"addGrantsToResources"`
	// Any policies created by this role will use this value as their ID, if specified.
	//
	// Specify this if importing the same role in multiple stacks, and granting it
	// different permissions in at least two stacks. If this is not specified
	// (or if the same name is specified in more than one stack),
	// a CloudFormation issue will result in the policy created in whichever stack
	// is deployed last overwriting the policies created by the others.
	// Default: 'Policy'.
	//
	DefaultPolicyName *string `field:"optional" json:"defaultPolicyName" yaml:"defaultPolicyName"`
	// Whether the imported role can be modified by attaching policy resources to it.
	// Default: true.
	//
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
	// The name of the role to lookup.
	//
	// If the role you want to lookup is a service role, you need to specify
	// the role name without the 'service-role' prefix. For example, if the role arn is
	// 'arn:aws:iam::123456789012:role/service-role/ExampleServiceExecutionRole',
	// you need to specify the role name as 'ExampleServiceExecutionRole'.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

