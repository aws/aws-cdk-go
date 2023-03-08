package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for API Key authorization in AppSync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var expiration expiration
//
//   apiKeyConfig := &ApiKeyConfig{
//   	Description: jsii.String("description"),
//   	Expires: expiration,
//   	Name: jsii.String("name"),
//   }
//
type ApiKeyConfig struct {
	// Description of API key.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time from creation time after which the API key expires.
	//
	// It must be a minimum of 1 day and a maximum of 365 days from date of creation.
	// Rounded down to the nearest hour.
	Expires awscdk.Expiration `field:"optional" json:"expires" yaml:"expires"`
	// Unique name of the API Key.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

