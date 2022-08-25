// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for API Key authorization in AppSync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var expiration expiration
//
//   apiKeyConfig := &apiKeyConfig{
//   	description: jsii.String("description"),
//   	expires: expiration,
//   	name: jsii.String("name"),
//   }
//
// Experimental.
type ApiKeyConfig struct {
	// Description of API key.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time from creation time after which the API key expires.
	//
	// It must be a minimum of 1 day and a maximum of 365 days from date of creation.
	// Rounded down to the nearest hour.
	// Experimental.
	Expires awscdk.Expiration `field:"optional" json:"expires" yaml:"expires"`
	// Unique name of the API Key.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

