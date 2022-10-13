// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Properties of a DataFormat instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var classificationString classificationString
//   var inputFormat inputFormat
//   var outputFormat outputFormat
//   var serializationLibrary serializationLibrary
//
//   dataFormatProps := &dataFormatProps{
//   	inputFormat: inputFormat,
//   	outputFormat: outputFormat,
//   	serializationLibrary: serializationLibrary,
//
//   	// the properties below are optional
//   	classificationString: classificationString,
//   }
//
// Experimental.
type DataFormatProps struct {
	// `InputFormat` for this data format.
	// Experimental.
	InputFormat InputFormat `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// `OutputFormat` for this data format.
	// Experimental.
	OutputFormat OutputFormat `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// Serialization library for this data format.
	// Experimental.
	SerializationLibrary SerializationLibrary `field:"required" json:"serializationLibrary" yaml:"serializationLibrary"`
	// Classification string given to tables with this data format.
	// Experimental.
	ClassificationString ClassificationString `field:"optional" json:"classificationString" yaml:"classificationString"`
}

