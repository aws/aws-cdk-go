package awsmediapackagev2


// <p>The failover settings for the endpoint.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forceEndpointErrorConfigurationProperty := &ForceEndpointErrorConfigurationProperty{
//   	EndpointErrorConditions: []*string{
//   		jsii.String("endpointErrorConditions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-forceendpointerrorconfiguration.html
//
type CfnOriginEndpoint_ForceEndpointErrorConfigurationProperty struct {
	// <p>The failover settings for the endpoint.
	//
	// The options are:</p>
	//          <ul>
	//             <li>
	//                <p>
	//                   <code>STALE_MANIFEST</code> - The manifest stalled and there a no new segments or parts.</p>
	//             </li>
	//             <li>
	//                <p>
	//                   <code>INCOMPLETE_MANIFEST</code> - There is a gap in the manifest.</p>
	//             </li>
	//             <li>
	//                <p>
	//                   <code>MISSING_DRM_KEY</code> - Key rotation is enabled but we're unable to fetch the key for the current key period.</p>
	//             </li>
	// </ul>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-forceendpointerrorconfiguration.html#cfn-mediapackagev2-originendpoint-forceendpointerrorconfiguration-endpointerrorconditions
	//
	EndpointErrorConditions *[]*string `field:"optional" json:"endpointErrorConditions" yaml:"endpointErrorConditions"`
}

