package awssagemaker


// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userContextProperty := &UserContextProperty{
//   	DomainId: jsii.String("domainId"),
//   	UserProfileArn: jsii.String("userProfileArn"),
//   	UserProfileName: jsii.String("userProfileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-usercontext.html
//
type CfnModelCard_UserContextProperty struct {
	// The domain associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-usercontext.html#cfn-sagemaker-modelcard-usercontext-domainid
	//
	// Default: - "UnsetValue".
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The Amazon Resource Name (ARN) of the user's profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-usercontext.html#cfn-sagemaker-modelcard-usercontext-userprofilearn
	//
	// Default: - "UnsetValue".
	//
	UserProfileArn *string `field:"optional" json:"userProfileArn" yaml:"userProfileArn"`
	// The name of the user's profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-usercontext.html#cfn-sagemaker-modelcard-usercontext-userprofilename
	//
	// Default: - "UnsetValue".
	//
	UserProfileName *string `field:"optional" json:"userProfileName" yaml:"userProfileName"`
}

