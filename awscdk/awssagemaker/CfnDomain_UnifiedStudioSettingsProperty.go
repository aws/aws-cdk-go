package awssagemaker


// A collection of settings that apply to an Amazon SageMaker AI domain when you use it in Amazon SageMaker Unified Studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   unifiedStudioSettingsProperty := &UnifiedStudioSettingsProperty{
//   	DomainAccountId: jsii.String("domainAccountId"),
//   	DomainId: jsii.String("domainId"),
//   	DomainRegion: jsii.String("domainRegion"),
//   	EnvironmentId: jsii.String("environmentId"),
//   	ProjectId: jsii.String("projectId"),
//   	ProjectS3Path: jsii.String("projectS3Path"),
//   	StudioWebPortalAccess: jsii.String("studioWebPortalAccess"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html
//
type CfnDomain_UnifiedStudioSettingsProperty struct {
	// The ID of the AWS account that has the Amazon SageMaker Unified Studio domain.
	//
	// The default value, if you don't specify an ID, is the ID of the account that has the Amazon SageMaker AI domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-domainaccountid
	//
	DomainAccountId *string `field:"optional" json:"domainAccountId" yaml:"domainAccountId"`
	// The ID of the Amazon SageMaker Unified Studio domain associated with this domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The AWS Region where the domain is located in Amazon SageMaker Unified Studio.
	//
	// The default value, if you don't specify a Region, is the Region where the Amazon SageMaker AI domain is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-domainregion
	//
	DomainRegion *string `field:"optional" json:"domainRegion" yaml:"domainRegion"`
	// The ID of the environment that Amazon SageMaker Unified Studio associates with the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-environmentid
	//
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// The ID of the Amazon SageMaker Unified Studio project that corresponds to the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-projectid
	//
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The location where Amazon S3 stores temporary execution data and other artifacts for the project that corresponds to the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-projects3path
	//
	ProjectS3Path *string `field:"optional" json:"projectS3Path" yaml:"projectS3Path"`
	// Sets whether you can access the domain in Amazon SageMaker Studio:.
	//
	// ENABLED
	// You can access the domain in Amazon SageMaker Studio. If you migrate the domain to Amazon SageMaker Unified Studio, you can access it in both studio interfaces.
	// DISABLED
	// You can't access the domain in Amazon SageMaker Studio. If you migrate the domain to Amazon SageMaker Unified Studio, you can access it only in that studio interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-unifiedstudiosettings.html#cfn-sagemaker-domain-unifiedstudiosettings-studiowebportalaccess
	//
	StudioWebPortalAccess *string `field:"optional" json:"studioWebPortalAccess" yaml:"studioWebPortalAccess"`
}

