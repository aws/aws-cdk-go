package awsauditmanager


// The `AWSAccount` property type specifies the wrapper of the AWS account details, such as account ID, email address, and so on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSAccountProperty := &AWSAccountProperty{
//   	EmailAddress: jsii.String("emailAddress"),
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsaccount.html
//
type CfnAssessment_AWSAccountProperty struct {
	// The email address that's associated with the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsaccount.html#cfn-auditmanager-assessment-awsaccount-emailaddress
	//
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The identifier for the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsaccount.html#cfn-auditmanager-assessment-awsaccount-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsaccount.html#cfn-auditmanager-assessment-awsaccount-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

