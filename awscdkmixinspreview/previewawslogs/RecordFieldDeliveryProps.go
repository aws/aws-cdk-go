package previewawslogs


// Props for RecordFields involved in log delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordFieldDeliveryProps := &RecordFieldDeliveryProps{
//   	MandatoryFields: []*string{
//   		jsii.String("mandatoryFields"),
//   	},
//   	ProvidedFields: []*string{
//   		jsii.String("providedFields"),
//   	},
//   }
//
// Experimental.
type RecordFieldDeliveryProps struct {
	// Any recordFields that a mandatory to be included in a log delivery of a certain log type.
	// Default: - log type has no mandatory fields.
	//
	// Experimental.
	MandatoryFields *[]*string `field:"optional" json:"mandatoryFields" yaml:"mandatoryFields"`
	// RecordFields the user has defined to be used in log delivery.
	// Experimental.
	ProvidedFields *[]*string `field:"optional" json:"providedFields" yaml:"providedFields"`
}

