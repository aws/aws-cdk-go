package previewawsauditmanagermixins


// The `Scope` property type specifies the wrapper that contains the AWS accounts and services that are in scope for the assessment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-scope.html
//
type CfnAssessmentPropsMixin_ScopeProperty struct {
	// The AWS accounts that are included in the scope of the assessment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-scope.html#cfn-auditmanager-assessment-scope-awsaccounts
	//
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// The AWS services that are included in the scope of the assessment.
	//
	// > This API parameter is no longer supported. If you use this parameter to specify one or more AWS services , Audit Manager ignores this input. Instead, the value for `awsServices` will show as empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-scope.html#cfn-auditmanager-assessment-scope-awsservices
	//
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}

