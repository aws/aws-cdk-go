package awsappflow


// The properties that are applied when using Trend Micro as a flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trendmicroSourcePropertiesProperty := &TrendmicroSourcePropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-trendmicrosourceproperties.html
//
type CfnFlow_TrendmicroSourcePropertiesProperty struct {
	// The object specified in the Trend Micro flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-trendmicrosourceproperties.html#cfn-appflow-flow-trendmicrosourceproperties-object
	//
	Object *string `field:"required" json:"object" yaml:"object"`
}

