package awscdklocationalpha


// Actions for Routes that an API key resource grants permissions to perform.
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
// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonlocationserviceroutes.html
//
// Experimental.
type AllowRoutesAction string

const (
	// Allows isoline calculation.
	// Experimental.
	AllowRoutesAction_CALCULATE_ISOLINES AllowRoutesAction = "CALCULATE_ISOLINES"
	// Allows point to point routing.
	// Experimental.
	AllowRoutesAction_CALCULATE_ROUTES AllowRoutesAction = "CALCULATE_ROUTES"
	// Allows matrix routing.
	// Experimental.
	AllowRoutesAction_CALCULATE_ROUTE_MATRIX AllowRoutesAction = "CALCULATE_ROUTE_MATRIX"
	// Allows computing the best sequence of waypoints.
	// Experimental.
	AllowRoutesAction_OPTIMIZE_WAYPOINTS AllowRoutesAction = "OPTIMIZE_WAYPOINTS"
	// Allows snapping GPS points to a likely route.
	// Experimental.
	AllowRoutesAction_SNAP_TO_ROADS AllowRoutesAction = "SNAP_TO_ROADS"
	// Allows any routes actions.
	// Experimental.
	AllowRoutesAction_ANY AllowRoutesAction = "ANY"
)

