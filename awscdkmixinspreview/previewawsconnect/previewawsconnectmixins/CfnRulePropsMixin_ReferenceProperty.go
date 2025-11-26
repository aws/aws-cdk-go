package previewawsconnectmixins


// Information about the reference when the `referenceType` is `URL` .
//
// Otherwise, null. (Supports variable injection in the `Value` field.)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceProperty := &ReferenceProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-reference.html
//
type CfnRulePropsMixin_ReferenceProperty struct {
	// The type of the reference. `DATE` must be of type Epoch timestamp.
	//
	// *Allowed values* : `URL` | `ATTACHMENT` | `NUMBER` | `STRING` | `DATE` | `EMAIL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-reference.html#cfn-connect-rule-reference-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// A valid value for the reference.
	//
	// For example, for a URL reference, a formatted URL that is displayed to an agent in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-reference.html#cfn-connect-rule-reference-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

