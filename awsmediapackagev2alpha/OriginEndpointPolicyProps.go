package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for Origin Endpoint policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var originEndpoint OriginEndpoint
//   var policyDocument PolicyDocument
//   var role Role
//   var secret Secret
//
//   originEndpointPolicyProps := &OriginEndpointPolicyProps{
//   	OriginEndpoint: originEndpoint,
//
//   	// the properties below are optional
//   	CdnAuth: &CdnAuthConfiguration{
//   		Secrets: []ISecret{
//   			secret,
//   		},
//
//   		// the properties below are optional
//   		Role: role,
//   	},
//   	PolicyDocument: policyDocument,
//   }
//
// Experimental.
type OriginEndpointPolicyProps struct {
	// `OriginEndpoint` to apply the Origin Endpoint Policy to.
	// Experimental.
	OriginEndpoint IOriginEndpoint `field:"required" json:"originEndpoint" yaml:"originEndpoint"`
	// Optional CDN Authorization configuration.
	// Default: - No header based CDN authorization.
	//
	// Experimental.
	CdnAuth *CdnAuthConfiguration `field:"optional" json:"cdnAuth" yaml:"cdnAuth"`
	// Initial policy document to apply.
	// Default: - empty policy document.
	//
	// Experimental.
	PolicyDocument awsiam.PolicyDocument `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

