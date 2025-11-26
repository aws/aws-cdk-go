package previewawsapsmixins


// The role configuration in an Amazon Managed Service for Prometheus scraper.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   roleConfigurationProperty := &RoleConfigurationProperty{
//   	SourceRoleArn: jsii.String("sourceRoleArn"),
//   	TargetRoleArn: jsii.String("targetRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-roleconfiguration.html
//
type CfnScraperPropsMixin_RoleConfigurationProperty struct {
	// The ARN of the source role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-roleconfiguration.html#cfn-aps-scraper-roleconfiguration-sourcerolearn
	//
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// The ARN of the target role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-roleconfiguration.html#cfn-aps-scraper-roleconfiguration-targetrolearn
	//
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

