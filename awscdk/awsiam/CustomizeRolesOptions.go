package awsiam


// Options for customizing IAM role creation.
//
// Example:
//   var app app
//
//   stack := awscdk.Newstack(app, jsii.String("MyStack"))
//   iam.Role_CustomizeRoles(this, &CustomizeRolesOptions{
//   	UsePrecreatedRoles: map[string]*string{
//   		"MyStack/MyLambda/ServiceRole": jsii.String("my-role-name"),
//   	},
//   })
//
type CustomizeRolesOptions struct {
	// Whether or not to synthesize the resource into the CFN template.
	//
	// Set this to `false` if you still want to create the resources _and_
	// you also want to create the policy report.
	// Default: true.
	//
	PreventSynthesis *bool `field:"optional" json:"preventSynthesis" yaml:"preventSynthesis"`
	// A list of precreated IAM roles to substitute for roles that CDK is creating.
	//
	// The constructPath can be either a relative or absolute path
	// from the scope that `customizeRoles` is used on to the role being created.
	//
	// Example:
	//   var app app
	//
	//
	//   stack := awscdk.Newstack(app, jsii.String("MyStack"))
	//   iam.NewRole(stack, jsii.String("MyRole"), &RoleProps{
	//   	AssumedBy: iam.NewAccountPrincipal(jsii.String("1111111111")),
	//   })
	//
	//   iam.Role_CustomizeRoles(stack, &CustomizeRolesOptions{
	//   	UsePrecreatedRoles: map[string]*string{
	//   		// absolute path
	//   		"MyStack/MyRole": jsii.String("my-precreated-role-name"),
	//   		// or relative path from `stack`
	//   		"MyRole": jsii.String("my-precreated-role"),
	//   	},
	//   })
	//
	// Default: - there are no precreated roles. Synthesis will fail if `preventSynthesis=true`
	//
	UsePrecreatedRoles *map[string]*string `field:"optional" json:"usePrecreatedRoles" yaml:"usePrecreatedRoles"`
}

