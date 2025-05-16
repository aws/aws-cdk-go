package awscdklocationalpha

import (
	"time"
)

// Properties for an API Key.
//
// Example:
//   location.NewApiKey(this, jsii.String("APIKeyAny"), &ApiKeyProps{
//   	// specify allowed actions
//   	AllowMapsActions: []allowMapsAction{
//   		location.*allowMapsAction_GET_STATIC_MAP,
//   	},
//   	AllowPlacesActions: []allowPlacesAction{
//   		location.*allowPlacesAction_GET_PLACE,
//   	},
//   	AllowRoutesActions: []allowRoutesAction{
//   		location.*allowRoutesAction_CALCULATE_ISOLINES,
//   	},
//   })
//
// Experimental.
type ApiKeyProps struct {
	// A list of allowed actions for Maps that an API key resource grants permissions to perform.
	// Default: - no actions for Maps are permitted.
	//
	// Experimental.
	AllowMapsActions *[]AllowMapsAction `field:"optional" json:"allowMapsActions" yaml:"allowMapsActions"`
	// A list of allowed actions for Places that an API key resource grants permissions to perform.
	// Default: - no actions for Places are permitted.
	//
	// Experimental.
	AllowPlacesActions *[]AllowPlacesAction `field:"optional" json:"allowPlacesActions" yaml:"allowPlacesActions"`
	// An optional list of allowed HTTP referers for which requests must originate from.
	//
	// Requests using this API key from other domains will not be allowed.
	// See: https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-apikeyrestrictions.html#cfn-location-apikey-apikeyrestrictions-allowreferers
	//
	// Default: - no Referer.
	//
	// Experimental.
	AllowReferers *[]*string `field:"optional" json:"allowReferers" yaml:"allowReferers"`
	// A list of allowed actions for Routes that an API key resource grants permissions to perform.
	// Default: - no actions for Routes are permitted.
	//
	// Experimental.
	AllowRoutesActions *[]AllowRoutesAction `field:"optional" json:"allowRoutesActions" yaml:"allowRoutesActions"`
	// A name for the api key.
	//
	// Must be between 1 and 100 characters and contain only alphanumeric characters,
	// hyphens, periods and underscores.
	//
	// Must be a unique API key name.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description for the api key.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The optional timestamp for when the API key resource will expire.
	//
	// `expireTime` must be set when `noExpiry` is false or undefined.
	// When `expireTime` is not set, `noExpiry` must be `true`.
	// Default: undefined - The API Key never expires.
	//
	// Experimental.
	ExpireTime *time.Time `field:"optional" json:"expireTime" yaml:"expireTime"`
	// `forceDelete` bypasses an API key's expiry conditions and deletes the key.
	//
	// Set the parameter true to delete the key or to false to not preemptively delete the API key.
	// Default: undefined - not force delete.
	//
	// Experimental.
	ForceDelete *bool `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// The boolean flag to be included for updating ExpireTime or Restrictions details.
	//
	// Must be set to true to update an API key resource that has been used in the past 7 days.
	// False if force update is not preferred.
	// Default: undefined - not force update.
	//
	// Experimental.
	ForceUpdate *bool `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// Whether the API key should expire.
	//
	// Set to `true` when `expireTime` is not set.
	// When you set `expireTime`, `noExpiry` must be `false` or `undefined`.
	// Default: undefined - The API Key expires at `expireTime`.
	//
	// Experimental.
	NoExpiry *bool `field:"optional" json:"noExpiry" yaml:"noExpiry"`
}

