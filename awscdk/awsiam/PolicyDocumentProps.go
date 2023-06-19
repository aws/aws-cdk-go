package awsiam


// Properties for a new PolicyDocument.
//
// Example:
//   myTrustedAdminRole := iam.Role_FromRoleArn(this, jsii.String("TrustedRole"), jsii.String("arn:aws:iam:...."))
//   // Creates a limited admin policy and assigns to the account root.
//   myCustomPolicy := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("kms:Create*"),
//   				jsii.String("kms:Describe*"),
//   				jsii.String("kms:Enable*"),
//   				jsii.String("kms:List*"),
//   				jsii.String("kms:Put*"),
//   			},
//   			Principals: []iPrincipal{
//   				iam.NewAccountRootPrincipal(),
//   			},
//   			Resources: []*string{
//   				jsii.String("*"),
//   			},
//   		}),
//   	},
//   })
//   key := kms.NewKey(this, jsii.String("MyKey"), &KeyProps{
//   	Policy: myCustomPolicy,
//   })
//
// Experimental.
type PolicyDocumentProps struct {
	// Automatically assign Statement Ids to all statements.
	// Experimental.
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
	// Experimental.
	Minimize *bool `field:"optional" json:"minimize" yaml:"minimize"`
	// Initial statements to add to the policy document.
	// Experimental.
	Statements *[]PolicyStatement `field:"optional" json:"statements" yaml:"statements"`
}

