// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Properties for defining a `CfnPublisher`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPublisherProps := &cfnPublisherProps{
//   	acceptTermsAndConditions: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	connectionArn: jsii.String("connectionArn"),
//   }
//
type CfnPublisherProps struct {
	// Whether you accept the [Terms and Conditions](https://docs.aws.amazon.com/https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf) for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to register to publish public extensions to the CloudFormation registry.
	//
	// The default is `false` .
	AcceptTermsAndConditions interface{} `field:"required" json:"acceptTermsAndConditions" yaml:"acceptTermsAndConditions"`
	// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
	//
	// For more information, see [Registering your account to publish CloudFormation extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs) in the *CloudFormation CLI User Guide* .
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

