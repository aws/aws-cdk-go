package awsappflow


// The properties that are applied when using Veeva as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   veevaSourcePropertiesProperty := &VeevaSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//
//   	// the properties below are optional
//   	DocumentType: jsii.String("documentType"),
//   	IncludeAllVersions: jsii.Boolean(false),
//   	IncludeRenditions: jsii.Boolean(false),
//   	IncludeSourceFiles: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-veevasourceproperties.html
//
type CfnFlow_VeevaSourcePropertiesProperty struct {
	// The object specified in the Veeva flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-veevasourceproperties.html#cfn-appflow-flow-veevasourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
	// The document type specified in the Veeva document extract flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-veevasourceproperties.html#cfn-appflow-flow-veevasourceproperties-documenttype
	//
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Boolean value to include All Versions of files in Veeva document extract flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-veevasourceproperties.html#cfn-appflow-flow-veevasourceproperties-includeallversions
	//
	IncludeAllVersions interface{} `field:"optional" json:"includeAllVersions" yaml:"includeAllVersions"`
	// Boolean value to include file renditions in Veeva document extract flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-veevasourceproperties.html#cfn-appflow-flow-veevasourceproperties-includerenditions
	//
	IncludeRenditions interface{} `field:"optional" json:"includeRenditions" yaml:"includeRenditions"`
	// Boolean value to include source files in Veeva document extract flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-veevasourceproperties.html#cfn-appflow-flow-veevasourceproperties-includesourcefiles
	//
	IncludeSourceFiles interface{} `field:"optional" json:"includeSourceFiles" yaml:"includeSourceFiles"`
}

