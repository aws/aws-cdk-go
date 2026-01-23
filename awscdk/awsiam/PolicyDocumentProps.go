package awsiam


// Properties for a new PolicyDocument.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//   policyDocument := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	AssignSids: jsii.Boolean(true),
//   	Statements: []PolicyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("sns:Publish"),
//   			},
//   			Principals: []IPrincipal{
//   				iam.NewServicePrincipal(jsii.String("s3.amazonaws.com")),
//   			},
//   			Resources: []*string{
//   				topic.topicArn,
//   			},
//   		}),
//   	},
//   })
//
//   topicPolicy := sns.NewTopicPolicy(this, jsii.String("Policy"), &TopicPolicyProps{
//   	Topics: []ITopic{
//   		topic,
//   	},
//   	PolicyDocument: PolicyDocument,
//   	EnforceSSL: jsii.Boolean(true),
//   })
//
type PolicyDocumentProps struct {
	// Automatically assign Statement Ids to all statements.
	// Default: false.
	//
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
	//   so would change the meaning of the policy document.
	// Default: - false, unless the feature flag `@aws-cdk/aws-iam:minimizePolicies` is set.
	//
	Minimize *bool `field:"optional" json:"minimize" yaml:"minimize"`
	// Initial statements to add to the policy document.
	// Default: - No statements.
	//
	Statements *[]PolicyStatement `field:"optional" json:"statements" yaml:"statements"`
}

