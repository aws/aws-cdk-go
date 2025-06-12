package awsapprunner


// Describes resources needed to authenticate access to some source repositories.
//
// The specific resource depends on the repository provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationConfigurationProperty := &AuthenticationConfigurationProperty{
//   	AccessRoleArn: jsii.String("accessRoleArn"),
//   	ConnectionArn: jsii.String("connectionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-authenticationconfiguration.html
//
type CfnService_AuthenticationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants the App Runner service access to a source repository.
	//
	// It's required for ECR image repositories (but not for ECR Public repositories).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-authenticationconfiguration.html#cfn-apprunner-service-authenticationconfiguration-accessrolearn
	//
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// The Amazon Resource Name (ARN) of the App Runner connection that enables the App Runner service to connect to a source repository.
	//
	// It's required for GitHub code repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-authenticationconfiguration.html#cfn-apprunner-service-authenticationconfiguration-connectionarn
	//
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

