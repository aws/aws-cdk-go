package previewawscodebuildevents


// Type definition for Artifacts_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   artifacts1 := &Artifacts1{
//   	EncryptionDisabled: []*string{
//   		jsii.String("encryptionDisabled"),
//   	},
//   	Location: []*string{
//   		jsii.String("location"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	NamespaceType: []*string{
//   		jsii.String("namespaceType"),
//   	},
//   	Packaging: []*string{
//   		jsii.String("packaging"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Artifacts1 struct {
	// encryptionDisabled property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionDisabled is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionDisabled *[]*string `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// location property.
	//
	// Specify an array of string values to match this event if the actual value of location is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Location *[]*string `field:"optional" json:"location" yaml:"location"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// namespaceType property.
	//
	// Specify an array of string values to match this event if the actual value of namespaceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NamespaceType *[]*string `field:"optional" json:"namespaceType" yaml:"namespaceType"`
	// packaging property.
	//
	// Specify an array of string values to match this event if the actual value of packaging is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Packaging *[]*string `field:"optional" json:"packaging" yaml:"packaging"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

