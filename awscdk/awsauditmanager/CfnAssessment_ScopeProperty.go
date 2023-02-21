package awsauditmanager


// The `Scope` property type specifies the wrapper that contains the AWS accounts and services that are in scope for the assessment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeProperty := &ScopeProperty{
//   	AwsAccounts: []interface{}{
//   		&AWSAccountProperty{
//   			EmailAddress: jsii.String("emailAddress"),
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	AwsServices: []interface{}{
//   		&AWSServiceProperty{
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   	},
//   }
//
type CfnAssessment_ScopeProperty struct {
	// The AWS accounts that are included in the scope of the assessment.
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// The AWS services that are included in the scope of the assessment.
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}

