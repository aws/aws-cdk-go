package awscustomerprofiles


// The properties that are applied when ServiceNow is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowSourcePropertiesProperty := &ServiceNowSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-servicenowsourceproperties.html
//
type CfnIntegration_ServiceNowSourcePropertiesProperty struct {
	// The object specified in the ServiceNow flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-servicenowsourceproperties.html#cfn-customerprofiles-integration-servicenowsourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

