package awsappflow


// The properties that are applied when Amazon Lookout for Metrics is used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lookoutMetricsDestinationPropertiesProperty := &LookoutMetricsDestinationPropertiesProperty{
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-lookoutmetricsdestinationproperties.html
//
type CfnFlow_LookoutMetricsDestinationPropertiesProperty struct {
	// The object specified in the Amazon Lookout for Metrics flow destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-lookoutmetricsdestinationproperties.html#cfn-appflow-flow-lookoutmetricsdestinationproperties-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

