package awsauditmanager


// The `Scope` property type specifies the wrapper that contains the AWS accounts and services in scope for the assessment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeProperty := &scopeProperty{
//   	awsAccounts: []interface{}{
//   		&aWSAccountProperty{
//   			emailAddress: jsii.String("emailAddress"),
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	awsServices: []interface{}{
//   		&aWSServiceProperty{
//   			serviceName: jsii.String("serviceName"),
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

