package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   postCallAnalyticsSettingsProperty := &PostCallAnalyticsSettingsProperty{
//   	ContentRedactionOutput: jsii.String("contentRedactionOutput"),
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	OutputEncryptionKmsKeyId: jsii.String("outputEncryptionKmsKeyId"),
//   	OutputLocation: jsii.String("outputLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_PostCallAnalyticsSettingsProperty struct {
	// The content redaction output settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-contentredactionoutput
	//
	ContentRedactionOutput *string `field:"optional" json:"contentRedactionOutput" yaml:"contentRedactionOutput"`
	// The ARN of the role used by Transcribe to upload post-call analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The ID of the KMS key used to encrypt the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-outputencryptionkmskeyid
	//
	OutputEncryptionKmsKeyId *string `field:"optional" json:"outputEncryptionKmsKeyId" yaml:"outputEncryptionKmsKeyId"`
	// The URL of the Amazon S3 bucket for post-call data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-postcallanalyticssettings.html#cfn-chime-mediainsightspipelineconfiguration-postcallanalyticssettings-outputlocation
	//
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

