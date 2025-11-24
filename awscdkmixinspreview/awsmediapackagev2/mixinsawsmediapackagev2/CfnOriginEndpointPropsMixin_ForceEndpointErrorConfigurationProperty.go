package mixinsawsmediapackagev2


// The failover settings for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   forceEndpointErrorConfigurationProperty := &ForceEndpointErrorConfigurationProperty{
//   	EndpointErrorConditions: []*string{
//   		jsii.String("endpointErrorConditions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-forceendpointerrorconfiguration.html
//
type CfnOriginEndpointPropsMixin_ForceEndpointErrorConfigurationProperty struct {
	// The failover conditions for the endpoint. The options are:.
	//
	// - `STALE_MANIFEST` - The manifest stalled and there are no new segments or parts.
	// - `INCOMPLETE_MANIFEST` - There is a gap in the manifest.
	// - `MISSING_DRM_KEY` - Key rotation is enabled but we're unable to fetch the key for the current key period.
	// - `SLATE_INPUT` - The segments which contain slate content are considered to be missing content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-forceendpointerrorconfiguration.html#cfn-mediapackagev2-originendpoint-forceendpointerrorconfiguration-endpointerrorconditions
	//
	EndpointErrorConditions *[]*string `field:"optional" json:"endpointErrorConditions" yaml:"endpointErrorConditions"`
}

