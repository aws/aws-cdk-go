package awsiam


// Properties for a new PolicyDocument.
//
// Example:
//   myTrustedAdminRole := iam.role.fromRoleArn(this, jsii.String("TrustedRole"), jsii.String("arn:aws:iam:...."))
//   // Creates a limited admin policy and assigns to the account root.
//   myCustomPolicy := iam.NewPolicyDocument(&policyDocumentProps{
//   	statements: []policyStatement{
//   		iam.NewPolicyStatement(&policyStatementProps{
//   			actions: []*string{
//   				jsii.String("kms:Create*"),
//   				jsii.String("kms:Describe*"),
//   				jsii.String("kms:Enable*"),
//   				jsii.String("kms:List*"),
//   				jsii.String("kms:Put*"),
//   			},
//   			principals: []iPrincipal{
//   				iam.NewAccountRootPrincipal(),
//   			},
//   			resources: []*string{
//   				jsii.String("*"),
//   			},
//   		}),
//   	},
//   })
//   key := kms.NewKey(this, jsii.String("MyKey"), &keyProps{
//   	policy: myCustomPolicy,
//   })
//
type PolicyDocumentProps struct {
	// Automatically assign Statement Ids to all statements.
	AssignSids *bool `field:"optional" json:"assignSids" yaml:"assignSids"`
	// Try to minimize the policy by merging statements.
	//
	// To avoid overrunning the maximum policy size, combine statements if they produce
	// the same result. Merging happens according to the following rules:
	//
	// - The Effect of both statements is the same
	// - Neither of the statements have a 'Sid'
	// - Combine Principals if the rest of the statement is exactly the same.
	// - Combine Resources if the rest of the statement is exactly the same.
	// - Combine Actions if the rest of the statement is exactly the same.
	// - We will never combine NotPrincipals, NotResources or NotActions, because doing
	//    so would change the meaning of the policy document.
	Minimize *bool `field:"optional" json:"minimize" yaml:"minimize"`
	// Initial statements to add to the policy document.
	Statements *[]PolicyStatement `field:"optional" json:"statements" yaml:"statements"`
}

