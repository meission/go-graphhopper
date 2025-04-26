# RouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The routing profile. It determines the network, speed and other physical attributes used when computing the route. See the section about [routing profiles](#tag/Map-Data-and-Routing-Profiles) for more details and valid profile values. | [optional] [default to "car"]
**Points** | Pointer to **[][]float64** | The points for the route in an array of &#x60;[longitude,latitude]&#x60;. For instance, if you want to calculate a route from point A to B to C then you specify &#x60;points: [ [A_longitude, A_latitude], [B_longitude, B_latitude], [C_longitude, C_latitude]]  | [optional] 
**PointHints** | Pointer to **[]string** | Optional parameter. Specifies a hint for each point in the &#x60;points&#x60; array to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. Make sure you do not include the house number of city name and only the street name to improve the quality of the matching. | [optional] 
**SnapPreventions** | Pointer to **[]string** | Optional parameter. &#39;Snapping&#39; is the process of finding the closest road location for GPS coordinates provided in the &#x60;points&#x60; array. The &#x60;snap_preventions&#x60; array allows you to prevent snapping to specific types of roads. For example, if the array includes &#x60;bridge&#x60;, then the routing engine will avoid snapping to a bridge, even if it is the closest road for the given point. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the &#x60;custom_model&#x60;. | [optional] 
**Curbsides** | Pointer to **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicle profiles and OpenStreetMap. | [optional] 
**Locale** | Pointer to **string** | The locale of the resulting turn instructions. E.g. &#x60;pt_PT&#x60; for Portuguese or &#x60;de&#x60; for German.  | [optional] [default to "en"]
**Elevation** | Pointer to **bool** | If &#x60;true&#x60;, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the &#x60;points&#x60; and &#x60;snapped_waypoints&#x60; fields of the response, in both their encodings. Unless you switch off the &#x60;points_encoded&#x60; parameter, you need special code on the client side that can handle three-dimensional coordinates.  | [optional] [default to false]
**Details** | Pointer to **[]string** | Optional parameter. The following path details are available: &#x60;street_name&#x60;, &#x60;street_ref&#x60;, &#x60;street_destination&#x60;, &#x60;leg_time&#x60;, &#x60;leg_distance&#x60;, &#x60;roundabout&#x60;, &#x60;country&#x60;, &#x60;time&#x60;, &#x60;distance&#x60;, &#x60;max_speed&#x60;, &#x60;max_weight&#x60;, &#x60;max_width&#x60;, &#x60;toll&#x60;, &#x60;road_class&#x60;, &#x60;road_class_link&#x60;, &#x60;road_access&#x60;, &#x60;road_environment&#x60;, &#x60;hazmat&#x60;, &#x60;hazmat_tunnel&#x60;, &#x60;hazmat_water&#x60;,  &#x60;lanes&#x60;, &#x60;surface&#x60;, &#x60;smoothness&#x60;, &#x60;hike_rating&#x60;, &#x60;mtb_rating&#x60;, &#x60;foot_network&#x60;, &#x60;bike_network&#x60;. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  | [optional] 
**Optimize** | Pointer to **string** | Normally, the calculated route will visit the points in the order you specified them. If you have more than two points, you can set this parameter to &#x60;\&quot;true\&quot;&#x60; and the points may be re-ordered to minimize the total travel time. Keep in mind that the limits on the number of locations of the Route Optimization API applies, and the request costs more credits.  | [optional] [default to "false"]
**Instructions** | Pointer to **bool** | If instructions should be calculated and returned  | [optional] [default to true]
**CalcPoints** | Pointer to **bool** | If the points for the route should be calculated at all.  | [optional] [default to true]
**Debug** | Pointer to **bool** | If &#x60;true&#x60;, the output will be formatted.  | [optional] [default to false]
**PointsEncoded** | Pointer to **bool** | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to &#x60;false&#x60; to switch the encoding to simple coordinate pairs like &#x60;[lon,lat]&#x60;, or &#x60;[lon,lat,elevation]&#x60;. See the description of the response format for more information.  | [optional] [default to true]
**ChDisable** | Pointer to **bool** | Use this parameter in combination with one or more parameters from below.  | [optional] [default to false]
**CustomModel** | Pointer to [**CustomModel**](CustomModel.md) |  | [optional] 
**Headings** | Pointer to **[]int32** | Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with &#x60;algorithm&#x3D;round_trip&#x60; and forces the initial direction.  Requires &#x60;ch.disable&#x3D;true&#x60;.  | [optional] 
**HeadingPenalty** | Pointer to **int32** | Time penalty in seconds for not obeying a specified heading. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [optional] [default to 300]
**PassThrough** | Pointer to **bool** | If &#x60;true&#x60;, u-turns are avoided at via-points with regard to the &#x60;heading_penalty&#x60;. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [optional] [default to false]
**Algorithm** | Pointer to **string** | Rather than looking for the shortest or fastest path, this parameter lets you solve two different problems related to routing: With &#x60;alternative_route&#x60;, we give you not one but several routes that are close to optimal, but not too similar to each other.  With &#x60;round_trip&#x60;, the route will get you back to where you started. This is meant for fun (think of a bike trip), so we will add some randomness. The &#x60;round_trip&#x60; option requires &#x60;ch.disable&#x3D;true&#x60;. You can control both of these features with additional parameters, see below.   | [optional] 
**RoundTripDistance** | Pointer to **int32** | If &#x60;algorithm&#x3D;round_trip&#x60;, this parameter configures approximative length of the resulting round trip. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [optional] [default to 10000]
**RoundTripSeed** | Pointer to **int64** | If &#x60;algorithm&#x3D;round_trip&#x60;, this sets the random seed. Change this to get a different tour for each value.  | [optional] 
**AlternativeRouteMaxPaths** | Pointer to **int32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives.  | [optional] [default to 2]
**AlternativeRouteMaxWeightFactor** | Pointer to **float32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives.  | [optional] [default to 1.4]
**AlternativeRouteMaxShareFactor** | Pointer to **float32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter specifies how similar an alternative route can be to the optimal route. Increasing can lead to worse alternatives.  | [optional] [default to 0.6]

## Methods

### NewRouteRequest

`func NewRouteRequest() *RouteRequest`

NewRouteRequest instantiates a new RouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteRequestWithDefaults

`func NewRouteRequestWithDefaults() *RouteRequest`

NewRouteRequestWithDefaults instantiates a new RouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *RouteRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *RouteRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *RouteRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *RouteRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetPoints

`func (o *RouteRequest) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *RouteRequest) GetPointsOk() (*[][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *RouteRequest) SetPoints(v [][]float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *RouteRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPointHints

`func (o *RouteRequest) GetPointHints() []string`

GetPointHints returns the PointHints field if non-nil, zero value otherwise.

### GetPointHintsOk

`func (o *RouteRequest) GetPointHintsOk() (*[]string, bool)`

GetPointHintsOk returns a tuple with the PointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointHints

`func (o *RouteRequest) SetPointHints(v []string)`

SetPointHints sets PointHints field to given value.

### HasPointHints

`func (o *RouteRequest) HasPointHints() bool`

HasPointHints returns a boolean if a field has been set.

### GetSnapPreventions

`func (o *RouteRequest) GetSnapPreventions() []string`

GetSnapPreventions returns the SnapPreventions field if non-nil, zero value otherwise.

### GetSnapPreventionsOk

`func (o *RouteRequest) GetSnapPreventionsOk() (*[]string, bool)`

GetSnapPreventionsOk returns a tuple with the SnapPreventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPreventions

`func (o *RouteRequest) SetSnapPreventions(v []string)`

SetSnapPreventions sets SnapPreventions field to given value.

### HasSnapPreventions

`func (o *RouteRequest) HasSnapPreventions() bool`

HasSnapPreventions returns a boolean if a field has been set.

### GetCurbsides

`func (o *RouteRequest) GetCurbsides() []string`

GetCurbsides returns the Curbsides field if non-nil, zero value otherwise.

### GetCurbsidesOk

`func (o *RouteRequest) GetCurbsidesOk() (*[]string, bool)`

GetCurbsidesOk returns a tuple with the Curbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurbsides

`func (o *RouteRequest) SetCurbsides(v []string)`

SetCurbsides sets Curbsides field to given value.

### HasCurbsides

`func (o *RouteRequest) HasCurbsides() bool`

HasCurbsides returns a boolean if a field has been set.

### GetLocale

`func (o *RouteRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *RouteRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *RouteRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *RouteRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetElevation

`func (o *RouteRequest) GetElevation() bool`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *RouteRequest) GetElevationOk() (*bool, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *RouteRequest) SetElevation(v bool)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *RouteRequest) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetDetails

`func (o *RouteRequest) GetDetails() []string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RouteRequest) GetDetailsOk() (*[]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RouteRequest) SetDetails(v []string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RouteRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetOptimize

`func (o *RouteRequest) GetOptimize() string`

GetOptimize returns the Optimize field if non-nil, zero value otherwise.

### GetOptimizeOk

`func (o *RouteRequest) GetOptimizeOk() (*string, bool)`

GetOptimizeOk returns a tuple with the Optimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimize

`func (o *RouteRequest) SetOptimize(v string)`

SetOptimize sets Optimize field to given value.

### HasOptimize

`func (o *RouteRequest) HasOptimize() bool`

HasOptimize returns a boolean if a field has been set.

### GetInstructions

`func (o *RouteRequest) GetInstructions() bool`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *RouteRequest) GetInstructionsOk() (*bool, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *RouteRequest) SetInstructions(v bool)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *RouteRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetCalcPoints

`func (o *RouteRequest) GetCalcPoints() bool`

GetCalcPoints returns the CalcPoints field if non-nil, zero value otherwise.

### GetCalcPointsOk

`func (o *RouteRequest) GetCalcPointsOk() (*bool, bool)`

GetCalcPointsOk returns a tuple with the CalcPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalcPoints

`func (o *RouteRequest) SetCalcPoints(v bool)`

SetCalcPoints sets CalcPoints field to given value.

### HasCalcPoints

`func (o *RouteRequest) HasCalcPoints() bool`

HasCalcPoints returns a boolean if a field has been set.

### GetDebug

`func (o *RouteRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *RouteRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *RouteRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *RouteRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetPointsEncoded

`func (o *RouteRequest) GetPointsEncoded() bool`

GetPointsEncoded returns the PointsEncoded field if non-nil, zero value otherwise.

### GetPointsEncodedOk

`func (o *RouteRequest) GetPointsEncodedOk() (*bool, bool)`

GetPointsEncodedOk returns a tuple with the PointsEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsEncoded

`func (o *RouteRequest) SetPointsEncoded(v bool)`

SetPointsEncoded sets PointsEncoded field to given value.

### HasPointsEncoded

`func (o *RouteRequest) HasPointsEncoded() bool`

HasPointsEncoded returns a boolean if a field has been set.

### GetChDisable

`func (o *RouteRequest) GetChDisable() bool`

GetChDisable returns the ChDisable field if non-nil, zero value otherwise.

### GetChDisableOk

`func (o *RouteRequest) GetChDisableOk() (*bool, bool)`

GetChDisableOk returns a tuple with the ChDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChDisable

`func (o *RouteRequest) SetChDisable(v bool)`

SetChDisable sets ChDisable field to given value.

### HasChDisable

`func (o *RouteRequest) HasChDisable() bool`

HasChDisable returns a boolean if a field has been set.

### GetCustomModel

`func (o *RouteRequest) GetCustomModel() CustomModel`

GetCustomModel returns the CustomModel field if non-nil, zero value otherwise.

### GetCustomModelOk

`func (o *RouteRequest) GetCustomModelOk() (*CustomModel, bool)`

GetCustomModelOk returns a tuple with the CustomModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomModel

`func (o *RouteRequest) SetCustomModel(v CustomModel)`

SetCustomModel sets CustomModel field to given value.

### HasCustomModel

`func (o *RouteRequest) HasCustomModel() bool`

HasCustomModel returns a boolean if a field has been set.

### GetHeadings

`func (o *RouteRequest) GetHeadings() []int32`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *RouteRequest) GetHeadingsOk() (*[]int32, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *RouteRequest) SetHeadings(v []int32)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *RouteRequest) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### GetHeadingPenalty

`func (o *RouteRequest) GetHeadingPenalty() int32`

GetHeadingPenalty returns the HeadingPenalty field if non-nil, zero value otherwise.

### GetHeadingPenaltyOk

`func (o *RouteRequest) GetHeadingPenaltyOk() (*int32, bool)`

GetHeadingPenaltyOk returns a tuple with the HeadingPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadingPenalty

`func (o *RouteRequest) SetHeadingPenalty(v int32)`

SetHeadingPenalty sets HeadingPenalty field to given value.

### HasHeadingPenalty

`func (o *RouteRequest) HasHeadingPenalty() bool`

HasHeadingPenalty returns a boolean if a field has been set.

### GetPassThrough

`func (o *RouteRequest) GetPassThrough() bool`

GetPassThrough returns the PassThrough field if non-nil, zero value otherwise.

### GetPassThroughOk

`func (o *RouteRequest) GetPassThroughOk() (*bool, bool)`

GetPassThroughOk returns a tuple with the PassThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThrough

`func (o *RouteRequest) SetPassThrough(v bool)`

SetPassThrough sets PassThrough field to given value.

### HasPassThrough

`func (o *RouteRequest) HasPassThrough() bool`

HasPassThrough returns a boolean if a field has been set.

### GetAlgorithm

`func (o *RouteRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RouteRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RouteRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RouteRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetRoundTripDistance

`func (o *RouteRequest) GetRoundTripDistance() int32`

GetRoundTripDistance returns the RoundTripDistance field if non-nil, zero value otherwise.

### GetRoundTripDistanceOk

`func (o *RouteRequest) GetRoundTripDistanceOk() (*int32, bool)`

GetRoundTripDistanceOk returns a tuple with the RoundTripDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTripDistance

`func (o *RouteRequest) SetRoundTripDistance(v int32)`

SetRoundTripDistance sets RoundTripDistance field to given value.

### HasRoundTripDistance

`func (o *RouteRequest) HasRoundTripDistance() bool`

HasRoundTripDistance returns a boolean if a field has been set.

### GetRoundTripSeed

`func (o *RouteRequest) GetRoundTripSeed() int64`

GetRoundTripSeed returns the RoundTripSeed field if non-nil, zero value otherwise.

### GetRoundTripSeedOk

`func (o *RouteRequest) GetRoundTripSeedOk() (*int64, bool)`

GetRoundTripSeedOk returns a tuple with the RoundTripSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTripSeed

`func (o *RouteRequest) SetRoundTripSeed(v int64)`

SetRoundTripSeed sets RoundTripSeed field to given value.

### HasRoundTripSeed

`func (o *RouteRequest) HasRoundTripSeed() bool`

HasRoundTripSeed returns a boolean if a field has been set.

### GetAlternativeRouteMaxPaths

`func (o *RouteRequest) GetAlternativeRouteMaxPaths() int32`

GetAlternativeRouteMaxPaths returns the AlternativeRouteMaxPaths field if non-nil, zero value otherwise.

### GetAlternativeRouteMaxPathsOk

`func (o *RouteRequest) GetAlternativeRouteMaxPathsOk() (*int32, bool)`

GetAlternativeRouteMaxPathsOk returns a tuple with the AlternativeRouteMaxPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeRouteMaxPaths

`func (o *RouteRequest) SetAlternativeRouteMaxPaths(v int32)`

SetAlternativeRouteMaxPaths sets AlternativeRouteMaxPaths field to given value.

### HasAlternativeRouteMaxPaths

`func (o *RouteRequest) HasAlternativeRouteMaxPaths() bool`

HasAlternativeRouteMaxPaths returns a boolean if a field has been set.

### GetAlternativeRouteMaxWeightFactor

`func (o *RouteRequest) GetAlternativeRouteMaxWeightFactor() float32`

GetAlternativeRouteMaxWeightFactor returns the AlternativeRouteMaxWeightFactor field if non-nil, zero value otherwise.

### GetAlternativeRouteMaxWeightFactorOk

`func (o *RouteRequest) GetAlternativeRouteMaxWeightFactorOk() (*float32, bool)`

GetAlternativeRouteMaxWeightFactorOk returns a tuple with the AlternativeRouteMaxWeightFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeRouteMaxWeightFactor

`func (o *RouteRequest) SetAlternativeRouteMaxWeightFactor(v float32)`

SetAlternativeRouteMaxWeightFactor sets AlternativeRouteMaxWeightFactor field to given value.

### HasAlternativeRouteMaxWeightFactor

`func (o *RouteRequest) HasAlternativeRouteMaxWeightFactor() bool`

HasAlternativeRouteMaxWeightFactor returns a boolean if a field has been set.

### GetAlternativeRouteMaxShareFactor

`func (o *RouteRequest) GetAlternativeRouteMaxShareFactor() float32`

GetAlternativeRouteMaxShareFactor returns the AlternativeRouteMaxShareFactor field if non-nil, zero value otherwise.

### GetAlternativeRouteMaxShareFactorOk

`func (o *RouteRequest) GetAlternativeRouteMaxShareFactorOk() (*float32, bool)`

GetAlternativeRouteMaxShareFactorOk returns a tuple with the AlternativeRouteMaxShareFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeRouteMaxShareFactor

`func (o *RouteRequest) SetAlternativeRouteMaxShareFactor(v float32)`

SetAlternativeRouteMaxShareFactor sets AlternativeRouteMaxShareFactor field to given value.

### HasAlternativeRouteMaxShareFactor

`func (o *RouteRequest) HasAlternativeRouteMaxShareFactor() bool`

HasAlternativeRouteMaxShareFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


