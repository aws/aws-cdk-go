package awsomics


// The store's parsing options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   tsvStoreOptionsProperty := &tsvStoreOptionsProperty{
//   	annotationType: jsii.String("annotationType"),
//   	formatToHeader: map[string]*string{
//   		"formatToHeaderKey": jsii.String("formatToHeader"),
//   	},
//   	schema: schema,
//   }
//
type CfnAnnotationStore_TsvStoreOptionsProperty struct {
	// The store's annotation type.
	AnnotationType *string `field:"optional" json:"annotationType" yaml:"annotationType"`
	// The store's header key to column name mapping.
	FormatToHeader interface{} `field:"optional" json:"formatToHeader" yaml:"formatToHeader"`
	// The schema of an annotation store.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

