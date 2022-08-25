package awsiam


// Properties for defining an IAM group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var managedPolicy managedPolicy
//
//   groupProps := &groupProps{
//   	groupName: jsii.String("groupName"),
//   	managedPolicies: []iManagedPolicy{
//   		managedPolicy,
//   	},
//   	path: jsii.String("path"),
//   }
//
type GroupProps struct {
	// A name for the IAM group.
	//
	// For valid values, see the GroupName parameter
	// for the CreateGroup action in the IAM API Reference. If you don't specify
	// a name, AWS CloudFormation generates a unique physical ID and uses that
	// ID for the group name.
	//
	// If you specify a name, you must specify the CAPABILITY_NAMED_IAM value to
	// acknowledge your template's capabilities. For more information, see
	// Acknowledging IAM Resources in AWS CloudFormation Templates.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// A list of managed policies associated with this role.
	//
	// You can add managed policies later using
	// `addManagedPolicy(ManagedPolicy.fromAwsManagedPolicyName(policyName))`.
	ManagedPolicies *[]IManagedPolicy `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// The path to the group.
	//
	// For more information about paths, see [IAM
	// Identifiers](http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html)
	// in the IAM User Guide.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

