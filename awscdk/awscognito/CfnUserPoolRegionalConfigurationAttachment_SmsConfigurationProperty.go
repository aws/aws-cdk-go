package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smsConfigurationProperty := &SmsConfigurationProperty{
//   	EumsSms: &EumsSmsConfigurationProperty{
//   		CallerArn: jsii.String("callerArn"),
//
//   		// the properties below are optional
//   		ConfigurationSetName: jsii.String("configurationSetName"),
//   		ExternalId: jsii.String("externalId"),
//   		InEntityId: jsii.String("inEntityId"),
//   		InTemplateId: jsii.String("inTemplateId"),
//   		OriginationIdentity: jsii.String("originationIdentity"),
//   		Region: jsii.String("region"),
//   	},
//   	ExternalId: jsii.String("externalId"),
//   	SnsCallerArn: jsii.String("snsCallerArn"),
//   	SnsRegion: jsii.String("snsRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-smsconfiguration.html
//
type CfnUserPoolRegionalConfigurationAttachment_SmsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-smsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-smsconfiguration-eumssms
	//
	EumsSms interface{} `field:"optional" json:"eumsSms" yaml:"eumsSms"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-smsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-smsconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-smsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-smsconfiguration-snscallerarn
	//
	SnsCallerArn *string `field:"optional" json:"snsCallerArn" yaml:"snsCallerArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-smsconfiguration.html#cfn-cognito-userpoolregionalconfigurationattachment-smsconfiguration-snsregion
	//
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
}

