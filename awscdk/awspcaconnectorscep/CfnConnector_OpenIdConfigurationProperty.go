package awspcaconnectorscep


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openIdConfigurationProperty := &OpenIdConfigurationProperty{
//   	Audience: jsii.String("audience"),
//   	Issuer: jsii.String("issuer"),
//   	Subject: jsii.String("subject"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html
//
type CfnConnector_OpenIdConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-audience
	//
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-subject
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

