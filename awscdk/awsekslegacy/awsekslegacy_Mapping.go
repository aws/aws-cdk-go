package awsekslegacy


// Example:
//   var cluster cluster
//
//   adminUser := iam.NewUser(this, jsii.String("Admin"))
//   cluster.awsAuth.addUserMapping(adminUser, &mapping{
//   	groups: []*string{
//   		jsii.String("system:masters"),
//   	},
//   })
//
// Experimental.
type Mapping struct {
	// A list of groups within Kubernetes to which the role is mapped.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	// Experimental.
	Groups *[]*string `field:"required" json:"groups" yaml:"groups"`
	// The user name within Kubernetes to map to the IAM role.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

