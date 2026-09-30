package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postCallAnalyticsSettingsProperty := &PostCallAnalyticsSettingsProperty{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	OutputLocation: jsii.String("outputLocation"),
//
//   	// the properties below are optional
//   	ContentRedactionOutput: jsii.String("contentRedactionOutput"),
//   	OutputEncryptionKmsKeyId: jsii.String("outputEncryptionKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html
//
type CfnMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty struct {
	// The ARN of the role used by Transcribe to upload post-call analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The URL of the Amazon S3 bucket for post-call data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-outputlocation
	//
	OutputLocation *string `field:"required" json:"outputLocation" yaml:"outputLocation"`
	// The content redaction output settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-contentredactionoutput
	//
	ContentRedactionOutput *string `field:"optional" json:"contentRedactionOutput" yaml:"contentRedactionOutput"`
	// The ID of the KMS key used to encrypt the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-outputencryptionkmskeyid
	//
	OutputEncryptionKmsKeyId *string `field:"optional" json:"outputEncryptionKmsKeyId" yaml:"outputEncryptionKmsKeyId"`
}

