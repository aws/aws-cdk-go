package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customizationConfigurationProperty := &CustomizationConfigurationProperty{
//   	CustomCssUrl: jsii.String("customCssUrl"),
//   	FaviconUrl: jsii.String("faviconUrl"),
//   	FontUrl: jsii.String("fontUrl"),
//   	LogoUrl: jsii.String("logoUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-customizationconfiguration.html
//
type CfnWebExperience_CustomizationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-customizationconfiguration.html#cfn-qbusiness-webexperience-customizationconfiguration-customcssurl
	//
	CustomCssUrl *string `field:"optional" json:"customCssUrl" yaml:"customCssUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-customizationconfiguration.html#cfn-qbusiness-webexperience-customizationconfiguration-faviconurl
	//
	FaviconUrl *string `field:"optional" json:"faviconUrl" yaml:"faviconUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-customizationconfiguration.html#cfn-qbusiness-webexperience-customizationconfiguration-fonturl
	//
	FontUrl *string `field:"optional" json:"fontUrl" yaml:"fontUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-customizationconfiguration.html#cfn-qbusiness-webexperience-customizationconfiguration-logourl
	//
	LogoUrl *string `field:"optional" json:"logoUrl" yaml:"logoUrl"`
}

