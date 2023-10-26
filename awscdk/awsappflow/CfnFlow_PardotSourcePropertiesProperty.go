package awsappflow


// The properties that are applied when Salesforce Pardot is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pardotSourcePropertiesProperty := &PardotSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-pardotsourceproperties.html
//
type CfnFlow_PardotSourcePropertiesProperty struct {
	// The object specified in the Salesforce Pardot flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-pardotsourceproperties.html#cfn-appflow-flow-pardotsourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

