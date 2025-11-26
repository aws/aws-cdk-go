package previewawsappflowmixins


// The properties that are applied when ServiceNow is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceNowSourcePropertiesProperty := &ServiceNowSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-servicenowsourceproperties.html
//
type CfnFlowPropsMixin_ServiceNowSourcePropertiesProperty struct {
	// The object specified in the ServiceNow flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-servicenowsourceproperties.html#cfn-appflow-flow-servicenowsourceproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

