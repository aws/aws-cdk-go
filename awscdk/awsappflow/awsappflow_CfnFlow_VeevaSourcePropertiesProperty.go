package awsappflow


// The properties that are applied when using Veeva as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   veevaSourcePropertiesProperty := &veevaSourcePropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	documentType: jsii.String("documentType"),
//   	includeAllVersions: jsii.Boolean(false),
//   	includeRenditions: jsii.Boolean(false),
//   	includeSourceFiles: jsii.Boolean(false),
//   }
//
type CfnFlow_VeevaSourcePropertiesProperty struct {
	// The object specified in the Veeva flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
	// The document type specified in the Veeva document extract flow.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Boolean value to include All Versions of files in Veeva document extract flow.
	IncludeAllVersions interface{} `field:"optional" json:"includeAllVersions" yaml:"includeAllVersions"`
	// Boolean value to include file renditions in Veeva document extract flow.
	IncludeRenditions interface{} `field:"optional" json:"includeRenditions" yaml:"includeRenditions"`
	// Boolean value to include source files in Veeva document extract flow.
	IncludeSourceFiles interface{} `field:"optional" json:"includeSourceFiles" yaml:"includeSourceFiles"`
}

