package awscdklocationalpha


// The position filtering for the tracker resource.
// Experimental.
type PositionFiltering string

const (
	// Location updates are evaluated against linked geofence collections, but not every location update is stored.
	//
	// If your update frequency is more often than 30 seconds, only one update per 30 seconds is stored for each unique device ID.
	// Experimental.
	PositionFiltering_TIME_BASED PositionFiltering = "TIME_BASED"
	// If the device has moved less than 30 m (98.4 ft), location updates are ignored. Location updates within this area are neither evaluated against linked geofence collections, nor stored. This helps control costs by reducing the number of geofence evaluations and historical device positions to paginate through. Distance-based filtering can also reduce the effects of GPS noise when displaying device trajectories on a map.
	// Experimental.
	PositionFiltering_DISTANCE_BASED PositionFiltering = "DISTANCE_BASED"
	// If the device has moved less than the measured accuracy, location updates are ignored.
	//
	// For example, if two consecutive updates from a device have a horizontal accuracy of 5 m and 10 m,
	// the second update is ignored if the device has moved less than 15 m.
	// Ignored location updates are neither evaluated against linked geofence collections, nor stored.
	// This can reduce the effects of GPS noise when displaying device trajectories on a map,
	// and can help control your costs by reducing the number of geofence evaluations.
	// Experimental.
	PositionFiltering_ACCURACY_BASED PositionFiltering = "ACCURACY_BASED"
)

