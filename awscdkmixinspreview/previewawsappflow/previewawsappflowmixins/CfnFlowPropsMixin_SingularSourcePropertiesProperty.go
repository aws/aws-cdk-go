package previewawsappflowmixins


// The properties that are applied when Singular is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   singularSourcePropertiesProperty := &SingularSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-singularsourceproperties.html
//
type CfnFlowPropsMixin_SingularSourcePropertiesProperty struct {
	// The object specified in the Singular flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-singularsourceproperties.html#cfn-appflow-flow-singularsourceproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

