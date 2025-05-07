package awscdklocationalpha


// Actions for Places that an API key resource grants permissions to perform.
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
// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonlocationserviceplaces.html
//
// Experimental.
type AllowPlacesAction string

const (
	// Allows auto-completion of search text.
	// Experimental.
	AllowPlacesAction_AUTOCOMPLETE AllowPlacesAction = "AUTOCOMPLETE"
	// Allows finding geo coordinates of a known place.
	// Experimental.
	AllowPlacesAction_GEOCODE AllowPlacesAction = "GEOCODE"
	// Allows getting details of a place.
	// Experimental.
	AllowPlacesAction_GET_PLACE AllowPlacesAction = "GET_PLACE"
	// Allows getting nearest address to geo coordinates.
	// Experimental.
	AllowPlacesAction_REVERSE_GEOCODE AllowPlacesAction = "REVERSE_GEOCODE"
	// Allows category based places search around geo coordinates.
	// Experimental.
	AllowPlacesAction_SEARCH_NEARBY AllowPlacesAction = "SEARCH_NEARBY"
	// Allows place or address search based on free-form text.
	// Experimental.
	AllowPlacesAction_SEARCH_TEXT AllowPlacesAction = "SEARCH_TEXT"
	// Allows suggestions based on an incomplete or misspelled query.
	// Experimental.
	AllowPlacesAction_SUGGEST AllowPlacesAction = "SUGGEST"
	// Allows any places actions.
	// Experimental.
	AllowPlacesAction_ANY AllowPlacesAction = "ANY"
)

