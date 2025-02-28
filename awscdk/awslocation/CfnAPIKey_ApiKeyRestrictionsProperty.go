package awslocation


// API Restrictions on the allowed actions, resources, and referers for an API key resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyRestrictionsProperty := &ApiKeyRestrictionsProperty{
//   	AllowActions: []*string{
//   		jsii.String("allowActions"),
//   	},
//   	AllowResources: []*string{
//   		jsii.String("allowResources"),
//   	},
//
//   	// the properties below are optional
//   	AllowReferers: []*string{
//   		jsii.String("allowReferers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-apikeyrestrictions.html
//
type CfnAPIKey_ApiKeyRestrictionsProperty struct {
	// A list of allowed actions that an API key resource grants permissions to perform.
	//
	// You must have at least one action for each type of resource. For example, if you have a place resource, you must include at least one place action.
	//
	// The following are valid values for the actions.
	//
	// - *Map actions*
	//
	// - `geo:GetMap*` - Allows all actions needed for map rendering.
	// - *Place actions*
	//
	// - `geo:SearchPlaceIndexForText` - Allows geocoding.
	// - `geo:SearchPlaceIndexForPosition` - Allows reverse geocoding.
	// - `geo:SearchPlaceIndexForSuggestions` - Allows generating suggestions from text.
	// - `geo:GetPlace` - Allows finding a place by place ID.
	// - *Route actions*
	//
	// - `geo:CalculateRoute` - Allows point to point routing.
	// - `geo:CalculateRouteMatrix` - Allows calculating a matrix of routes.
	//
	// > You must use these strings exactly. For example, to provide access to map rendering, the only valid action is `geo:GetMap*` as an input to the list. `["geo:GetMap*"]` is valid but `["geo:GetMapTile"]` is not. Similarly, you cannot use `["geo:SearchPlaceIndexFor*"]` - you must list each of the Place actions separately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-apikeyrestrictions.html#cfn-location-apikey-apikeyrestrictions-allowactions
	//
	AllowActions *[]*string `field:"required" json:"allowActions" yaml:"allowActions"`
	// A list of allowed resource ARNs that a API key bearer can perform actions on.
	//
	// - The ARN must be the correct ARN for a map, place, or route ARN. You may include wildcards in the resource-id to match multiple resources of the same type.
	// - The resources must be in the same `partition` , `region` , and `account-id` as the key that is being created.
	// - Other than wildcards, you must include the full ARN, including the `arn` , `partition` , `service` , `region` , `account-id` and `resource-id` delimited by colons (:).
	// - No spaces allowed, even with wildcards. For example, `arn:aws:geo:region: *account-id* :map/ExampleMap*` .
	//
	// For more information about ARN format, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-apikeyrestrictions.html#cfn-location-apikey-apikeyrestrictions-allowresources
	//
	AllowResources *[]*string `field:"required" json:"allowResources" yaml:"allowResources"`
	// An optional list of allowed HTTP referers for which requests must originate from.
	//
	// Requests using this API key from other domains will not be allowed.
	//
	// Requirements:
	//
	// - Contain only alphanumeric characters (A–Z, a–z, 0–9) or any symbols in this list `$\-._+!*`(),;/?:@=&`
	// - May contain a percent (%) if followed by 2 hexadecimal digits (A-F, a-f, 0-9); this is used for URL encoding purposes.
	// - May contain wildcard characters question mark (?) and asterisk (*).
	//
	// Question mark (?) will replace any single character (including hexadecimal digits).
	//
	// Asterisk (*) will replace any multiple characters (including multiple hexadecimal digits).
	// - No spaces allowed. For example, `https://example.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-location-apikey-apikeyrestrictions.html#cfn-location-apikey-apikeyrestrictions-allowreferers
	//
	AllowReferers *[]*string `field:"optional" json:"allowReferers" yaml:"allowReferers"`
}

