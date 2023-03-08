package awsauditmanager


// The `AWSAccount` property type specifies the wrapper of the AWS account details, such as account ID, email address, and so on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSAccountProperty := &aWSAccountProperty{
//   	emailAddress: jsii.String("emailAddress"),
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   }
//
type CfnAssessment_AWSAccountProperty struct {
	// The email address that's associated with the AWS account .
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The identifier for the AWS account .
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the AWS account .
	Name *string `field:"optional" json:"name" yaml:"name"`
}

