package previewawsappflowmixins


// The connector-specific profile credentials required by Dynatrace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynatraceConnectorProfileCredentialsProperty := &DynatraceConnectorProfileCredentialsProperty{
//   	ApiToken: jsii.String("apiToken"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-dynatraceconnectorprofilecredentials.html
//
type CfnConnectorProfilePropsMixin_DynatraceConnectorProfileCredentialsProperty struct {
	// The API tokens used by Dynatrace API to authenticate various API calls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-dynatraceconnectorprofilecredentials.html#cfn-appflow-connectorprofile-dynatraceconnectorprofilecredentials-apitoken
	//
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
}

