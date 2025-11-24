package mixinsawsgroundstation


// Provides information to AWS Ground Station about which IP endpoints to use during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataflowEndpointConfigProperty := &DataflowEndpointConfigProperty{
//   	DataflowEndpointName: jsii.String("dataflowEndpointName"),
//   	DataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-dataflowendpointconfig.html
//
type CfnConfigPropsMixin_DataflowEndpointConfigProperty struct {
	// The name of the dataflow endpoint to use during contacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-dataflowendpointconfig.html#cfn-groundstation-config-dataflowendpointconfig-dataflowendpointname
	//
	DataflowEndpointName *string `field:"optional" json:"dataflowEndpointName" yaml:"dataflowEndpointName"`
	// The region of the dataflow endpoint to use during contacts.
	//
	// When omitted, Ground Station will use the region of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-dataflowendpointconfig.html#cfn-groundstation-config-dataflowendpointconfig-dataflowendpointregion
	//
	DataflowEndpointRegion *string `field:"optional" json:"dataflowEndpointRegion" yaml:"dataflowEndpointRegion"`
}

