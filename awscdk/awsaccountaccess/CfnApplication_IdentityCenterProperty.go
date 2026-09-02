package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityCenterProperty := &IdentityCenterProperty{
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	ApplicationArn: jsii.String("applicationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitycenter.html
//
type CfnApplication_IdentityCenterProperty struct {
	// The ARN of the Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitycenter.html#cfn-accountaccess-application-identitycenter-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The ARN of the associated Identity Center application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-application-identitycenter.html#cfn-accountaccess-application-identitycenter-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
}

