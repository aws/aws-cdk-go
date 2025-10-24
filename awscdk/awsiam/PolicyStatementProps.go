package awsiam


// Interface for creating a policy statement.
//
// Example:
//   accessLogsBucket := s3.NewBucket(this, jsii.String("AccessLogsBucket"), &BucketProps{
//   	ObjectOwnership: s3.ObjectOwnership_BUCKET_OWNER_ENFORCED,
//   })
//
//   accessLogsBucket.AddToResourcePolicy(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:*"),
//   	},
//   	Resources: []*string{
//   		accessLogsBucket.BucketArn,
//   		accessLogsBucket.ArnForObjects(jsii.String("*")),
//   	},
//   	Principals: []IPrincipal{
//   		iam.NewAnyPrincipal(),
//   	},
//   }))
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	ServerAccessLogsBucket: accessLogsBucket,
//   	ServerAccessLogsPrefix: jsii.String("logs"),
//   })
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

