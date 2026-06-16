package awsresiliencehubv2


// An assertion about the service's resilience posture.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   assertionDefinitionProperty := &AssertionDefinitionProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-assertiondefinition.html
//
type CfnServicePropsMixin_AssertionDefinitionProperty struct {
	// The text of the assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-assertiondefinition.html#cfn-resiliencehubv2-service-assertiondefinition-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

