package awsappflow


// The properties that are applied when Infor Nexus is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inforNexusSourcePropertiesProperty := &InforNexusSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-infornexussourceproperties.html
//
type CfnFlow_InforNexusSourcePropertiesProperty struct {
	// The object specified in the Infor Nexus flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-infornexussourceproperties.html#cfn-appflow-flow-infornexussourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

