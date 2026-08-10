package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eumsSmsConfigurationProperty := &EumsSmsConfigurationProperty{
//   	CallerArn: jsii.String("callerArn"),
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	ExternalId: jsii.String("externalId"),
//   	InEntityId: jsii.String("inEntityId"),
//   	InTemplateId: jsii.String("inTemplateId"),
//   	OriginationIdentity: jsii.String("originationIdentity"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html
//
type CfnUserPoolRegionalConfigurationAttachmentPropsMixin_EumsSmsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-callerarn
	//
	CallerArn *string `field:"optional" json:"callerArn" yaml:"callerArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-configurationsetname
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-inentityid
	//
	InEntityId *string `field:"optional" json:"inEntityId" yaml:"inEntityId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-intemplateid
	//
	InTemplateId *string `field:"optional" json:"inTemplateId" yaml:"inTemplateId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-originationidentity
	//
	OriginationIdentity *string `field:"optional" json:"originationIdentity" yaml:"originationIdentity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-eumssmsconfiguration-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

