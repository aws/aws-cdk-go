package awsentityresolution


// Configuration for matching behavior within rule condition properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchingConfigProperty := &MatchingConfigProperty{
//   	EnableTransitiveMatching: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-matchingconfig.html
//
type CfnMatchingWorkflow_MatchingConfigProperty struct {
	// Enables transitive matching to process records across all rule levels and connect unmatched records to existing match groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-matchingconfig.html#cfn-entityresolution-matchingworkflow-matchingconfig-enabletransitivematching
	//
	EnableTransitiveMatching interface{} `field:"optional" json:"enableTransitiveMatching" yaml:"enableTransitiveMatching"`
}

