package previewawsappflowmixins


// The properties that are applied when Marketo is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   marketoSourcePropertiesProperty := &MarketoSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-marketosourceproperties.html
//
type CfnFlowPropsMixin_MarketoSourcePropertiesProperty struct {
	// The object specified in the Marketo flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-marketosourceproperties.html#cfn-appflow-flow-marketosourceproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

