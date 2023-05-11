package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathFilterProperty := &PathFilterProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationPortRange: &FilterPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	SourceAddress: jsii.String("sourceAddress"),
//   	SourcePortRange: &FilterPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   }
//
type CfnNetworkInsightsPath_PathFilterProperty struct {
	// `CfnNetworkInsightsPath.PathFilterProperty.DestinationAddress`.
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// `CfnNetworkInsightsPath.PathFilterProperty.DestinationPortRange`.
	DestinationPortRange interface{} `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// `CfnNetworkInsightsPath.PathFilterProperty.SourceAddress`.
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// `CfnNetworkInsightsPath.PathFilterProperty.SourcePortRange`.
	SourcePortRange interface{} `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

