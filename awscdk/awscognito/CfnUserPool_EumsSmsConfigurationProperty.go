package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eumsSmsConfigurationProperty := &EumsSmsConfigurationProperty{
//   	CallerArn: jsii.String("callerArn"),
//
//   	// the properties below are optional
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	ExternalId: jsii.String("externalId"),
//   	InEntityId: jsii.String("inEntityId"),
//   	InTemplateId: jsii.String("inTemplateId"),
//   	OriginationIdentity: jsii.String("originationIdentity"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html
//
type CfnUserPool_EumsSmsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-callerarn
	//
	CallerArn *string `field:"required" json:"callerArn" yaml:"callerArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-configurationsetname
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-inentityid
	//
	InEntityId *string `field:"optional" json:"inEntityId" yaml:"inEntityId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-intemplateid
	//
	InTemplateId *string `field:"optional" json:"inTemplateId" yaml:"inTemplateId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-originationidentity
	//
	OriginationIdentity *string `field:"optional" json:"originationIdentity" yaml:"originationIdentity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-eumssmsconfiguration.html#cfn-cognito-userpool-eumssmsconfiguration-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

