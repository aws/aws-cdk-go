package awsevents


// The parameters for EventBridge to use when invoking the resource endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceParametersProperty := &ResourceParametersProperty{
//   	ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   	// the properties below are optional
//   	ResourceAssociationArn: jsii.String("resourceAssociationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-resourceparameters.html
//
type CfnConnection_ResourceParametersProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon VPC Lattice resource configuration for the resource endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-resourceparameters.html#cfn-events-connection-resourceparameters-resourceconfigurationarn
	//
	ResourceConfigurationArn *string `field:"required" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
	// For connections to private APIs, the Amazon Resource Name (ARN) of the resource association EventBridge created between the connection and the private API's resource configuration.
	//
	// > The value of this property is set by EventBridge . Any value you specify in your template is ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-resourceparameters.html#cfn-events-connection-resourceparameters-resourceassociationarn
	//
	ResourceAssociationArn *string `field:"optional" json:"resourceAssociationArn" yaml:"resourceAssociationArn"`
}

