package previewawsqbusinessmixins


// Provides information about index capacity configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   indexCapacityConfigurationProperty := &IndexCapacityConfigurationProperty{
//   	Units: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-indexcapacityconfiguration.html
//
type CfnIndexPropsMixin_IndexCapacityConfigurationProperty struct {
	// The number of storage units configured for an Amazon Q Business index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-indexcapacityconfiguration.html#cfn-qbusiness-index-indexcapacityconfiguration-units
	//
	Units *float64 `field:"optional" json:"units" yaml:"units"`
}

