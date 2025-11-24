package mixinsawsomics


// The store's file parsing options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schema interface{}
//
//   storeOptionsProperty := &StoreOptionsProperty{
//   	TsvStoreOptions: &TsvStoreOptionsProperty{
//   		AnnotationType: jsii.String("annotationType"),
//   		FormatToHeader: map[string]*string{
//   			"formatToHeaderKey": jsii.String("formatToHeader"),
//   		},
//   		Schema: schema,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-storeoptions.html
//
type CfnAnnotationStorePropsMixin_StoreOptionsProperty struct {
	// Formatting options for a TSV file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-storeoptions.html#cfn-omics-annotationstore-storeoptions-tsvstoreoptions
	//
	TsvStoreOptions interface{} `field:"optional" json:"tsvStoreOptions" yaml:"tsvStoreOptions"`
}

