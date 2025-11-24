package mixinsawsappflow


// The properties that are applied when Dynatrace is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynatraceSourcePropertiesProperty := &DynatraceSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-dynatracesourceproperties.html
//
type CfnFlowPropsMixin_DynatraceSourcePropertiesProperty struct {
	// The object specified in the Dynatrace flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-dynatracesourceproperties.html#cfn-appflow-flow-dynatracesourceproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

