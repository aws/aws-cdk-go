package awscdklocationalpha


// Actions for Maps that an API key resource grants permissions to perform.
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
// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonlocationservicemaps.html
//
// Experimental.
type AllowMapsAction string

const (
	// Allows getting static map images.
	// Experimental.
	AllowMapsAction_GET_STATIC_MAP AllowMapsAction = "GET_STATIC_MAP"
	// Allows getting map tiles for rendering.
	// Experimental.
	AllowMapsAction_GET_TILE AllowMapsAction = "GET_TILE"
	// Allows any maps actions.
	// Experimental.
	AllowMapsAction_ANY AllowMapsAction = "ANY"
)

