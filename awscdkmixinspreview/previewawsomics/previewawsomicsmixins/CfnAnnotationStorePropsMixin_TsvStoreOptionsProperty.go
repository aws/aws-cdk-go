package previewawsomicsmixins


// The store's parsing options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schema interface{}
//
//   tsvStoreOptionsProperty := &TsvStoreOptionsProperty{
//   	AnnotationType: jsii.String("annotationType"),
//   	FormatToHeader: map[string]*string{
//   		"formatToHeaderKey": jsii.String("formatToHeader"),
//   	},
//   	Schema: schema,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-tsvstoreoptions.html
//
type CfnAnnotationStorePropsMixin_TsvStoreOptionsProperty struct {
	// The store's annotation type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-tsvstoreoptions.html#cfn-omics-annotationstore-tsvstoreoptions-annotationtype
	//
	AnnotationType *string `field:"optional" json:"annotationType" yaml:"annotationType"`
	// The store's header key to column name mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-tsvstoreoptions.html#cfn-omics-annotationstore-tsvstoreoptions-formattoheader
	//
	FormatToHeader interface{} `field:"optional" json:"formatToHeader" yaml:"formatToHeader"`
	// The schema of an annotation store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-tsvstoreoptions.html#cfn-omics-annotationstore-tsvstoreoptions-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

