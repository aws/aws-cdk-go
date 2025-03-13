package awssso


// A structure that describes the options for the portal associated with an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalOptionsConfigurationProperty := &PortalOptionsConfigurationProperty{
//   	SignInOptions: &SignInOptionsProperty{
//   		Origin: jsii.String("origin"),
//
//   		// the properties below are optional
//   		ApplicationUrl: jsii.String("applicationUrl"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-application-portaloptionsconfiguration.html
//
type CfnApplication_PortalOptionsConfigurationProperty struct {
	// A structure that describes the sign-in options for the access portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-application-portaloptionsconfiguration.html#cfn-sso-application-portaloptionsconfiguration-signinoptions
	//
	SignInOptions interface{} `field:"optional" json:"signInOptions" yaml:"signInOptions"`
	// Indicates whether this application is visible in the access portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-application-portaloptionsconfiguration.html#cfn-sso-application-portaloptionsconfiguration-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

