package awsiam


// Interface for creating a policy statement.
//
// Example:
//   var destinationBucket bucket
//
//
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployFiles"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Asset(path.join(__dirname, jsii.String("source-files"))),
//   	},
//   	DestinationBucket: DestinationBucket,
//   })
//
//   deployment.HandlerRole.AddToPolicy(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("kms:Decrypt"),
//   		jsii.String("kms:DescribeKey"),
//   	},
//   	Effect: iam.Effect_ALLOW,
//   	Resources: []*string{
//   		jsii.String("<encryption key ARN>"),
//   	},
//   }))
//
type PolicyStatementProps struct {
	// List of actions to add to the statement.
	// Default: - no actions.
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Conditions to add to the statement.
	// Default: - no condition.
	//
	Conditions *map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether to allow or deny the actions in this statement.
	// Default: Effect.ALLOW
	//
	Effect Effect `field:"optional" json:"effect" yaml:"effect"`
	// List of not actions to add to the statement.
	// Default: - no not-actions.
	//
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// List of not principals to add to the statement.
	// Default: - no not principals.
	//
	NotPrincipals *[]IPrincipal `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// NotResource ARNs to add to the statement.
	// Default: - no not-resources.
	//
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// List of principals to add to the statement.
	// Default: - no principals.
	//
	Principals *[]IPrincipal `field:"optional" json:"principals" yaml:"principals"`
	// Resource ARNs to add to the statement.
	// Default: - no resources.
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The Sid (statement ID) is an optional identifier that you provide for the policy statement.
	//
	// You can assign a Sid value to each statement in a
	// statement array. In services that let you specify an ID element, such as
	// SQS and SNS, the Sid value is just a sub-ID of the policy document's ID. In
	// IAM, the Sid value must be unique within a JSON policy.
	// Default: - no sid.
	//
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

