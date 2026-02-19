package interfaces


// Represents the environment a given resource lives in.
//
// Used as the return value for the `IEnvironmentAware.env` property.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//
//   var scope Construct
//   type myFactory struct {
//   }
//
//   func (this *myFactory) forResource(resource CfnResource) IResourceWithPolicyV2 {
//   	return map[string]ResourceEnvironment{
//   		"env": *resource.env,
//   		(MethodDeclaration addToResourcePolicy(statement: PolicyStatement) {
//   		        // custom implementation to add the statement to the resource policy
//   		        return { statementAdded: true, policyDependable: resource };
//   		      }
//   				addToResourcePolicy
//   				statement PolicyStatement
//   				{
//   					// custom implementation to add the statement to the resource policy
//   					return &AddToResourcePolicyResult{
//   						"statementAdded": jsii.Boolean(true),
//   						"policyDependable": resource,
//   					}
//   				}),
//   	}
//   }
//
//   awscdk.ResourceWithPolicies_Register(scope, jsii.String("AWS::KMS::Key"), NewMyFactory())
//
type ResourceEnvironment struct {
	// The AWS Account ID that this resource belongs to.
	//
	// Since this can be a Token (for example, when the account is
	// CloudFormation's `AWS::AccountId` intrinsic), make sure to use
	// `Token.compareStrings()` instead of comparing the values with direct
	// string equality.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The AWS Region that this resource belongs to.
	//
	// Since this can be a Token (for example, when the region is CloudFormation's
	// `AWS::Region` intrinsic), make sure to use `Token.compareStrings()` instead
	// of comparing the values with direct string equality.
	Region *string `field:"required" json:"region" yaml:"region"`
}

