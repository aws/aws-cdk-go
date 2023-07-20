package awscustomerprofiles


// The properties that are applied when using Zendesk as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zendeskSourcePropertiesProperty := &ZendeskSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-zendesksourceproperties.html
//
type CfnIntegration_ZendeskSourcePropertiesProperty struct {
	// The object specified in the Zendesk flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-zendesksourceproperties.html#cfn-customerprofiles-integration-zendesksourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

