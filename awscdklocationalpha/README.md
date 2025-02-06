# AWS::Location Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

Amazon Location Service lets you add location data and functionality to applications, which
includes capabilities such as maps, points of interest, geocoding, routing, geofences, and
tracking. Amazon Location provides location-based services (LBS) using high-quality data from
global, trusted providers Esri and HERE. With affordable data, tracking and geofencing
capabilities, and built-in metrics for health monitoring, you can build sophisticated
location-enabled applications.

## Geofence Collection

Geofence collection resources allow you to store and manage geofences—virtual boundaries on a map.
You can evaluate locations against a geofence collection resource and get notifications when the location
update crosses the boundary of any of the geofences in the geofence collection.

```go
var key key


location.NewGeofenceCollection(this, jsii.String("GeofenceCollection"), &GeofenceCollectionProps{
	GeofenceCollectionName: jsii.String("MyGeofenceCollection"),
	 // optional, defaults to a generated name
	KmsKey: key,
})
```

Use the `grant()` or `grantRead()` method to grant the given identity permissions to perform actions
on the geofence collection:

```go
var role role


geofenceCollection := location.NewGeofenceCollection(this, jsii.String("GeofenceCollection"), &GeofenceCollectionProps{
	GeofenceCollectionName: jsii.String("MyGeofenceCollection"),
})

geofenceCollection.GrantRead(role)
```

## Tracker

A tracker stores position updates for a collection of devices. The tracker can be used to query the devices' current location or location history. It stores the updates, but reduces storage space and visual noise by filtering the locations before storing them.

For more information, see [Trackers](https://docs.aws.amazon.com/location/latest/developerguide/geofence-tracker-concepts.html#tracking-overview).

To create a tracker, define a `Tracker`:

```go
var key key


location.NewTracker(this, jsii.String("Tracker"), &TrackerProps{
	TrackerName: jsii.String("MyTracker"),
	 // optional, defaults to a generated name
	KmsKey: key,
})
```

Use the `grant()`, `grantUpdateDevicePositions()` or `grantRead()` method to grant the given identity permissions to perform actions
on the geofence collection:

```go
var role role


tracker := location.NewTracker(this, jsii.String("Tracker"), &TrackerProps{
	TrackerName: jsii.String("MyTracker"),
})

tracker.GrantRead(role)
```

If you want to associate a tracker with geofence collections, define a `geofenceCollections` property or use the `addGeofenceCollections()` method.

```go
var geofenceCollection geofenceCollection
var geofenceCollectionForAdd geofenceCollection
var tracker tracker


tracker := location.NewTracker(this, jsii.String("Tracker"), &TrackerProps{
	TrackerName: jsii.String("MyTracker"),
	GeofenceCollections: []iGeofenceCollection{
		geofenceCollection,
	},
})

tracker.AddGeofenceCollections(geofenceCollectionForAdd)
```

## Legacy Resources

AWS has released new [Enhanced Places, Routes, and Maps](https://aws.amazon.com/about-aws/whats-new/2024/11/amazon-location-service-enhanced-places-routes-maps/?nc1=h_ls). Since these use AWS-managed resources, users no longer need to create Maps, Places, and Routes resources themselves.

As a result, the following constructs are now considered legacy.

For more information, see [developer guide](https://docs.aws.amazon.com/location/latest/developerguide/what-is.html).

### Map

The Amazon Location Service Map resource gives you access to the underlying basemap data for a map.
You use the Map resource with a map rendering library to add an interactive map to your application.
You can add other functionality to your map, such as markers (or pins), routes, and polygon areas, as needed for your application.

For information about how to use map resources in practice, see [Using Amazon Location Maps in your application](https://docs.aws.amazon.com/location/latest/developerguide/using-maps.html).

To create a map, define a `Map`:

```go
location.NewMap(this, jsii.String("Map"), &MapProps{
	MapName: jsii.String("my-map"),
	Style: location.Style_VECTOR_ESRI_NAVIGATION,
	CustomLayers: []pOI{
		location.CustomLayer_*pOI,
	},
})
```

Use the `grant()` or `grantRendering()` method to grant the given identity permissions to perform actions
on the map:

```go
var role role


map := location.NewMap(this, jsii.String("Map"), &MapProps{
	Style: location.Style_VECTOR_ESRI_NAVIGATION,
})
map.GrantRendering(role)
```

### Place Index

A key function of Amazon Location Service is the ability to search the geolocation information.
Amazon Location provides this functionality via the Place index resource. The place index includes
which [data provider](https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html)
to use for the search.

To create a place index, define a `PlaceIndex`:

```go
location.NewPlaceIndex(this, jsii.String("PlaceIndex"), &PlaceIndexProps{
	PlaceIndexName: jsii.String("MyPlaceIndex"),
	 // optional, defaults to a generated name
	DataSource: location.DataSource_HERE,
})
```

Use the `grant()` or `grantSearch()` method to grant the given identity permissions to perform actions
on the place index:

```go
var role role


placeIndex := location.NewPlaceIndex(this, jsii.String("PlaceIndex"))
placeIndex.GrantSearch(role)
```

### Route Calculator

Route calculator resources allow you to find routes and estimate travel time based on up-to-date road network and live traffic information from your chosen data provider.

For more information, see [Routes](https://docs.aws.amazon.com/location/latest/developerguide/route-concepts.html).

To create a route calculator, define a `RouteCalculator`:

```go
location.NewRouteCalculator(this, jsii.String("RouteCalculator"), &RouteCalculatorProps{
	RouteCalculatorName: jsii.String("MyRouteCalculator"),
	 // optional, defaults to a generated name
	DataSource: location.DataSource_ESRI,
})
```

Use the `grant()` or `grantRead()` method to grant the given identity permissions to perform actions
on the route calculator:

```go
var role role


routeCalculator := location.NewRouteCalculator(this, jsii.String("RouteCalculator"), &RouteCalculatorProps{
	DataSource: location.DataSource_ESRI,
})
routeCalculator.GrantRead(role)
```
