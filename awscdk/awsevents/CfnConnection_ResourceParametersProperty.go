package awsevents


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-resourceparameters.html#cfn-events-connection-resourceparameters-resourceconfigurationarn
	//
	ResourceConfigurationArn *string `field:"required" json:"resourceConfigurationArn" yaml:"resourceConfigurationArn"`
	// For connections to private APIs, the Amazon Resource Name (ARN) of the resource association EventBridge created between the connection and the private API's resource configuration.
	//
	// For more information, see [Managing service network resource associations for connections](https://docs.aws.amazon.com/eventbridge/latest/userguide/connection-private.html#connection-private-snra) in the **Amazon EventBridge User Guide** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-resourceparameters.html#cfn-events-connection-resourceparameters-resourceassociationarn
	//
	ResourceAssociationArn *string `field:"optional" json:"resourceAssociationArn" yaml:"resourceAssociationArn"`
}

