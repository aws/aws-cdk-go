// The CDK Construct Library for AWS::Location
package awscdklocationalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.AllowMapsAction",
		reflect.TypeOf((*AllowMapsAction)(nil)).Elem(),
		map[string]interface{}{
			"GET_STATIC_MAP": AllowMapsAction_GET_STATIC_MAP,
			"GET_TILE": AllowMapsAction_GET_TILE,
			"ANY": AllowMapsAction_ANY,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.AllowPlacesAction",
		reflect.TypeOf((*AllowPlacesAction)(nil)).Elem(),
		map[string]interface{}{
			"AUTOCOMPLETE": AllowPlacesAction_AUTOCOMPLETE,
			"GEOCODE": AllowPlacesAction_GEOCODE,
			"GET_PLACE": AllowPlacesAction_GET_PLACE,
			"REVERSE_GEOCODE": AllowPlacesAction_REVERSE_GEOCODE,
			"SEARCH_NEARBY": AllowPlacesAction_SEARCH_NEARBY,
			"SEARCH_TEXT": AllowPlacesAction_SEARCH_TEXT,
			"SUGGEST": AllowPlacesAction_SUGGEST,
			"ANY": AllowPlacesAction_ANY,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.AllowRoutesAction",
		reflect.TypeOf((*AllowRoutesAction)(nil)).Elem(),
		map[string]interface{}{
			"CALCULATE_ISOLINES": AllowRoutesAction_CALCULATE_ISOLINES,
			"CALCULATE_ROUTES": AllowRoutesAction_CALCULATE_ROUTES,
			"CALCULATE_ROUTE_MATRIX": AllowRoutesAction_CALCULATE_ROUTE_MATRIX,
			"OPTIMIZE_WAYPOINTS": AllowRoutesAction_OPTIMIZE_WAYPOINTS,
			"SNAP_TO_ROADS": AllowRoutesAction_SNAP_TO_ROADS,
			"ANY": AllowRoutesAction_ANY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.ApiKey",
		reflect.TypeOf((*ApiKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiKeyArn", GoGetter: "ApiKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyCreateTime", GoGetter: "ApiKeyCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyName", GoGetter: "ApiKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyUpdateTime", GoGetter: "ApiKeyUpdateTime"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ApiKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiKey)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.ApiKeyProps",
		reflect.TypeOf((*ApiKeyProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.CustomLayer",
		reflect.TypeOf((*CustomLayer)(nil)).Elem(),
		map[string]interface{}{
			"POI": CustomLayer_POI,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
		map[string]interface{}{
			"ESRI": DataSource_ESRI,
			"GRAB": DataSource_GRAB,
			"HERE": DataSource_HERE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.GeofenceCollection",
		reflect.TypeOf((*GeofenceCollection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionArn", GoGetter: "GeofenceCollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionCreateTime", GoGetter: "GeofenceCollectionCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionName", GoGetter: "GeofenceCollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionUpdateTime", GoGetter: "GeofenceCollectionUpdateTime"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_GeofenceCollection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGeofenceCollection)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.GeofenceCollectionProps",
		reflect.TypeOf((*GeofenceCollectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IApiKey",
		reflect.TypeOf((*IApiKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiKeyArn", GoGetter: "ApiKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyName", GoGetter: "ApiKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IApiKey{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IGeofenceCollection",
		reflect.TypeOf((*IGeofenceCollection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionArn", GoGetter: "GeofenceCollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionName", GoGetter: "GeofenceCollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IGeofenceCollection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IMap",
		reflect.TypeOf((*IMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "mapArn", GoGetter: "MapArn"},
			_jsii_.MemberProperty{JsiiProperty: "mapName", GoGetter: "MapName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IMap{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IPlaceIndex",
		reflect.TypeOf((*IPlaceIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexArn", GoGetter: "PlaceIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexName", GoGetter: "PlaceIndexName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IPlaceIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IRouteCalculator",
		reflect.TypeOf((*IRouteCalculator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "routeCalculatorArn", GoGetter: "RouteCalculatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "routeCalculatorName", GoGetter: "RouteCalculatorName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRouteCalculator{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.ITracker",
		reflect.TypeOf((*ITracker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "trackerArn", GoGetter: "TrackerArn"},
			_jsii_.MemberProperty{JsiiProperty: "trackerName", GoGetter: "TrackerName"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ITracker{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.IntendedUse",
		reflect.TypeOf((*IntendedUse)(nil)).Elem(),
		map[string]interface{}{
			"SINGLE_USE": IntendedUse_SINGLE_USE,
			"STORAGE": IntendedUse_STORAGE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.Map",
		reflect.TypeOf((*Map)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRendering", GoMethod: "GrantRendering"},
			_jsii_.MemberProperty{JsiiProperty: "mapArn", GoGetter: "MapArn"},
			_jsii_.MemberProperty{JsiiProperty: "mapCreateTime", GoGetter: "MapCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "mapName", GoGetter: "MapName"},
			_jsii_.MemberProperty{JsiiProperty: "mapUpdateTime", GoGetter: "MapUpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Map{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMap)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.MapProps",
		reflect.TypeOf((*MapProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.PlaceIndex",
		reflect.TypeOf((*PlaceIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantSearch", GoMethod: "GrantSearch"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexArn", GoGetter: "PlaceIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexCreateTime", GoGetter: "PlaceIndexCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexName", GoGetter: "PlaceIndexName"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexUpdateTime", GoGetter: "PlaceIndexUpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PlaceIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPlaceIndex)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.PlaceIndexProps",
		reflect.TypeOf((*PlaceIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.PoliticalView",
		reflect.TypeOf((*PoliticalView)(nil)).Elem(),
		map[string]interface{}{
			"INDIA": PoliticalView_INDIA,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.PositionFiltering",
		reflect.TypeOf((*PositionFiltering)(nil)).Elem(),
		map[string]interface{}{
			"TIME_BASED": PositionFiltering_TIME_BASED,
			"DISTANCE_BASED": PositionFiltering_DISTANCE_BASED,
			"ACCURACY_BASED": PositionFiltering_ACCURACY_BASED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.RouteCalculator",
		reflect.TypeOf((*RouteCalculator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "routeCalculatorArn", GoGetter: "RouteCalculatorArn"},
			_jsii_.MemberProperty{JsiiProperty: "routeCalculatorCreateTime", GoGetter: "RouteCalculatorCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "routeCalculatorName", GoGetter: "RouteCalculatorName"},
			_jsii_.MemberProperty{JsiiProperty: "routeCalculatorUpdateTime", GoGetter: "RouteCalculatorUpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RouteCalculator{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRouteCalculator)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.RouteCalculatorProps",
		reflect.TypeOf((*RouteCalculatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.Style",
		reflect.TypeOf((*Style)(nil)).Elem(),
		map[string]interface{}{
			"VECTOR_ESRI_NAVIGATION": Style_VECTOR_ESRI_NAVIGATION,
			"RASTER_ESRI_IMAGERY": Style_RASTER_ESRI_IMAGERY,
			"VECTOR_ESRI_LIGHT_GRAY_CANVAS": Style_VECTOR_ESRI_LIGHT_GRAY_CANVAS,
			"VECTOR_ESRI_TOPOGRAPHIC": Style_VECTOR_ESRI_TOPOGRAPHIC,
			"VECTOR_ESRI_STREETS": Style_VECTOR_ESRI_STREETS,
			"VECTOR_ESRI_DARK_GRAY_CANVAS": Style_VECTOR_ESRI_DARK_GRAY_CANVAS,
			"VECTOR_HERE_EXPLORE": Style_VECTOR_HERE_EXPLORE,
			"RASTER_HERE_EXPLORE_SATELLITE": Style_RASTER_HERE_EXPLORE_SATELLITE,
			"HYBRID_HERE_EXPLORE_SATELLITE": Style_HYBRID_HERE_EXPLORE_SATELLITE,
			"VECTOR_HERE_CONTRAST": Style_VECTOR_HERE_CONTRAST,
			"VECTOR_HERE_EXPLORE_TRUCK": Style_VECTOR_HERE_EXPLORE_TRUCK,
			"VECTOR_GRAB_STANDARD_LIGHT": Style_VECTOR_GRAB_STANDARD_LIGHT,
			"VECTOR_GRAB_STANDARD_DARK": Style_VECTOR_GRAB_STANDARD_DARK,
			"VECTOR_OPEN_DATA_STANDARD_LIGHT": Style_VECTOR_OPEN_DATA_STANDARD_LIGHT,
			"VECTOR_OPEN_DATA_STANDARD_DARK": Style_VECTOR_OPEN_DATA_STANDARD_DARK,
			"VECTOR_OPEN_DATA_VISUALIZATION_LIGHT": Style_VECTOR_OPEN_DATA_VISUALIZATION_LIGHT,
			"VECTOR_OPEN_DATA_VISUALIZATION_DARK": Style_VECTOR_OPEN_DATA_VISUALIZATION_DARK,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.Tracker",
		reflect.TypeOf((*Tracker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addGeofenceCollections", GoMethod: "AddGeofenceCollections"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantUpdateDevicePositions", GoMethod: "GrantUpdateDevicePositions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trackerArn", GoGetter: "TrackerArn"},
			_jsii_.MemberProperty{JsiiProperty: "trackerCreateTime", GoGetter: "TrackerCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "trackerName", GoGetter: "TrackerName"},
			_jsii_.MemberProperty{JsiiProperty: "trackerUpdateTime", GoGetter: "TrackerUpdateTime"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Tracker{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITracker)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.TrackerProps",
		reflect.TypeOf((*TrackerProps)(nil)).Elem(),
	)
}
