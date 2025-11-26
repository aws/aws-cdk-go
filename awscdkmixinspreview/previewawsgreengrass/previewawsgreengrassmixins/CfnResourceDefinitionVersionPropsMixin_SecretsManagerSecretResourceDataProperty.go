package previewawsgreengrassmixins


// Settings for a secret resource, which references a secret from AWS Secrets Manager .
//
// AWS IoT Greengrass stores a local, encrypted copy of the secret on the Greengrass core, where it can be securely accessed by connectors and Lambda functions. For more information, see [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/v1/developerguide/secrets.html) in the *Developer Guide* .
//
// In an CloudFormation template, `SecretsManagerSecretResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   secretsManagerSecretResourceDataProperty := &SecretsManagerSecretResourceDataProperty{
//   	AdditionalStagingLabelsToDownload: []*string{
//   		jsii.String("additionalStagingLabelsToDownload"),
//   	},
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata.html
//
type CfnResourceDefinitionVersionPropsMixin_SecretsManagerSecretResourceDataProperty struct {
	// The staging labels whose values you want to make available on the core, in addition to `AWSCURRENT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata.html#cfn-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata-additionalstaginglabelstodownload
	//
	AdditionalStagingLabelsToDownload *[]*string `field:"optional" json:"additionalStagingLabelsToDownload" yaml:"additionalStagingLabelsToDownload"`
	// The Amazon Resource Name (ARN) of the Secrets Manager secret to make available on the core.
	//
	// The value of the secret's latest version (represented by the `AWSCURRENT` staging label) is included by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata.html#cfn-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

