package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayConnectorFiltersProperty := &GatewayConnectorFiltersProperty{
//   	AccessTypes: []*string{
//   		jsii.String("accessTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorfilters.html
//
type CfnServicePropsMixin_GatewayConnectorFiltersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorfilters.html#cfn-ecs-service-gatewayconnectorfilters-accesstypes
	//
	AccessTypes *[]*string `field:"optional" json:"accessTypes" yaml:"accessTypes"`
}

