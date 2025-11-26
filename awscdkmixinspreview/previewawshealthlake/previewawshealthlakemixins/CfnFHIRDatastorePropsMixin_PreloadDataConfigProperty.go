package previewawshealthlakemixins


// An optional parameter to preload (import) open source Synthea FHIR data upon creation of the data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   preloadDataConfigProperty := &PreloadDataConfigProperty{
//   	PreloadDataType: jsii.String("preloadDataType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-preloaddataconfig.html
//
type CfnFHIRDatastorePropsMixin_PreloadDataConfigProperty struct {
	// The type of preloaded data.
	//
	// Only Synthea preloaded data is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-preloaddataconfig.html#cfn-healthlake-fhirdatastore-preloaddataconfig-preloaddatatype
	//
	PreloadDataType *string `field:"optional" json:"preloadDataType" yaml:"preloadDataType"`
}

