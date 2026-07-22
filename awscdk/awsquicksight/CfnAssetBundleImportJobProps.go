package awsquicksight


// Properties for defining a `CfnAssetBundleImportJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetBundleImportJobProps := &CfnAssetBundleImportJobProps{
//   	AssetBundleImportJobId: jsii.String("assetBundleImportJobId"),
//
//   	// the properties below are optional
//   	AssetBundleImportSource: &AssetBundleImportSourceDescriptionProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	FailureAction: jsii.String("failureAction"),
//   	OverrideValidationStrategy: &AssetBundleImportJobOverrideValidationStrategyProperty{
//   		StrictModeForAllResources: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html
//
type CfnAssetBundleImportJobProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html#cfn-quicksight-assetbundleimportjob-assetbundleimportjobid
	//
	AssetBundleImportJobId *string `field:"required" json:"assetBundleImportJobId" yaml:"assetBundleImportJobId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html#cfn-quicksight-assetbundleimportjob-assetbundleimportsource
	//
	AssetBundleImportSource interface{} `field:"optional" json:"assetBundleImportSource" yaml:"assetBundleImportSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html#cfn-quicksight-assetbundleimportjob-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html#cfn-quicksight-assetbundleimportjob-failureaction
	//
	FailureAction *string `field:"optional" json:"failureAction" yaml:"failureAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleimportjob.html#cfn-quicksight-assetbundleimportjob-overridevalidationstrategy
	//
	OverrideValidationStrategy interface{} `field:"optional" json:"overrideValidationStrategy" yaml:"overrideValidationStrategy"`
}

