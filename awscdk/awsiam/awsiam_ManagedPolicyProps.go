package awsiam


// Properties for defining an IAM managed policy.
//
// Example:
//   policyDocument := map[string]interface{}{
//   	"Version": jsii.String("2012-10-17"),
//   	"Statement": []interface{}{
//   		map[string]interface{}{
//   			"Sid": jsii.String("FirstStatement"),
//   			"Effect": jsii.String("Allow"),
//   			"Action": []*string{
//   				jsii.String("iam:ChangePassword"),
//   			},
//   			"Resource": jsii.String("*"),
//   		},
//   		map[string]*string{
//   			"Sid": jsii.String("SecondStatement"),
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("s3:ListAllMyBuckets"),
//   			"Resource": jsii.String("*"),
//   		},
//   		map[string]interface{}{
//   			"Sid": jsii.String("ThirdStatement"),
//   			"Effect": jsii.String("Allow"),
//   			"Action": []*string{
//   				jsii.String("s3:List*"),
//   				jsii.String("s3:Get*"),
//   			},
//   			"Resource": []*string{
//   				jsii.String("arn:aws:s3:::confidential-data"),
//   				jsii.String("arn:aws:s3:::confidential-data/*"),
//   			},
//   			"Condition": map[string]map[string]*string{
//   				"Bool": map[string]*string{
//   					"aws:MultiFactorAuthPresent": jsii.String("true"),
//   				},
//   			},
//   		},
//   	},
//   }
//
//   customPolicyDocument := iam.policyDocument.fromJson(policyDocument)
//
//   // You can pass this document as an initial document to a ManagedPolicy
//   // or inline Policy.
//   newManagedPolicy := iam.NewManagedPolicy(this, jsii.String("MyNewManagedPolicy"), &managedPolicyProps{
//   	document: customPolicyDocument,
//   })
//   newPolicy := iam.NewPolicy(this, jsii.String("MyNewPolicy"), &policyProps{
//   	document: customPolicyDocument,
//   })
//
type ManagedPolicyProps struct {
	// A description of the managed policy.
	//
	// Typically used to store information about the
	// permissions defined in the policy. For example, "Grants access to production DynamoDB tables."
	// The policy description is immutable. After a value is assigned, it cannot be changed.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Initial PolicyDocument to use for this ManagedPolicy.
	//
	// If omited, any
	// `PolicyStatement` provided in the `statements` property will be applied
	// against the empty default `PolicyDocument`.
	Document PolicyDocument `field:"optional" json:"document" yaml:"document"`
	// Groups to attach this policy to.
	//
	// You can also use `attachToGroup(group)` to attach this policy to a group.
	Groups *[]IGroup `field:"optional" json:"groups" yaml:"groups"`
	// The name of the managed policy.
	//
	// If you specify multiple policies for an entity,
	// specify unique names. For example, if you specify a list of policies for
	// an IAM role, each policy must have a unique name.
	ManagedPolicyName *string `field:"optional" json:"managedPolicyName" yaml:"managedPolicyName"`
	// The path for the policy.
	//
	// This parameter allows (through its regex pattern) a string of characters
	// consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes.
	// In addition, it can contain any ASCII character from the ! (\u0021) through the DEL character (\u007F),
	// including most punctuation characters, digits, and upper and lowercased letters.
	//
	// For more information about paths, see IAM Identifiers in the IAM User Guide.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Roles to attach this policy to.
	//
	// You can also use `attachToRole(role)` to attach this policy to a role.
	Roles *[]IRole `field:"optional" json:"roles" yaml:"roles"`
	// Initial set of permissions to add to this policy document.
	//
	// You can also use `addPermission(statement)` to add permissions later.
	Statements *[]PolicyStatement `field:"optional" json:"statements" yaml:"statements"`
	// Users to attach this policy to.
	//
	// You can also use `attachToUser(user)` to attach this policy to a user.
	Users *[]IUser `field:"optional" json:"users" yaml:"users"`
}

