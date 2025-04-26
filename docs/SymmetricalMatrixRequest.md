# SymmetricalMatrixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The routing profile. It determines the network, speed and other physical attributes used when computing the route. See the section about [routing profiles](#tag/Map-Data-and-Routing-Profiles) for more details and valid profile values. | [optional] [default to "car"]
**Points** | Pointer to **[][]float64** | Specify multiple points for which the weight-, route-, time- or distance-matrix should be calculated as follows: &#x60;[longitude,latitude]&#x60;. In this case the origins are identical to the destinations. Thus, if there are N points, NxN entries are calculated. The order of the point parameter is important. Specify at least three points. Cannot be used together with &#x60;from_point&#x60; or &#x60;to_point.&#x60;. | [optional] 
**PointHints** | Pointer to **[]string** | Optional parameter. Specifies a hint for each point in the &#x60;points&#x60; array to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | [optional] 
**SnapPreventions** | Pointer to **[]string** | Optional parameter. &#39;Snapping&#39; is the process of finding the closest road location for GPS coordinates provided in the &#x60;points&#x60; array. The &#x60;snap_preventions&#x60; array allows you to prevent snapping to specific types of roads. For example, if the array includes &#x60;bridge&#x60;, then the routing engine will avoid snapping to a bridge, even if it is the closest road for the given point. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the &#x60;custom_model&#x60;. | [optional] 
**Curbsides** | Pointer to **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap. | [optional] 
**OutArrays** | Pointer to **[]string** | Specifies which matrices should be included in the response. Specify one or more of the following options &#x60;weights&#x60;, &#x60;times&#x60;, &#x60;distances&#x60;. The units of the entries of &#x60;distances&#x60; are meters, of &#x60;times&#x60; are seconds and of &#x60;weights&#x60; is arbitrary and it can differ for different vehicle profiles or versions of this API. | [optional] 
**FailFast** | Pointer to **bool** | Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to &#x60;false&#x60; the time/weight/distance matrix will be calculated for all valid points and contain the &#x60;null&#x60; value for all entries that could not be calculated. The &#x60;hint&#x60; field of the response will also contain additional information about what went wrong (see its documentation). | [optional] [default to true]

## Methods

### NewSymmetricalMatrixRequest

`func NewSymmetricalMatrixRequest() *SymmetricalMatrixRequest`

NewSymmetricalMatrixRequest instantiates a new SymmetricalMatrixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymmetricalMatrixRequestWithDefaults

`func NewSymmetricalMatrixRequestWithDefaults() *SymmetricalMatrixRequest`

NewSymmetricalMatrixRequestWithDefaults instantiates a new SymmetricalMatrixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *SymmetricalMatrixRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SymmetricalMatrixRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SymmetricalMatrixRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SymmetricalMatrixRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetPoints

`func (o *SymmetricalMatrixRequest) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *SymmetricalMatrixRequest) GetPointsOk() (*[][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *SymmetricalMatrixRequest) SetPoints(v [][]float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *SymmetricalMatrixRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPointHints

`func (o *SymmetricalMatrixRequest) GetPointHints() []string`

GetPointHints returns the PointHints field if non-nil, zero value otherwise.

### GetPointHintsOk

`func (o *SymmetricalMatrixRequest) GetPointHintsOk() (*[]string, bool)`

GetPointHintsOk returns a tuple with the PointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointHints

`func (o *SymmetricalMatrixRequest) SetPointHints(v []string)`

SetPointHints sets PointHints field to given value.

### HasPointHints

`func (o *SymmetricalMatrixRequest) HasPointHints() bool`

HasPointHints returns a boolean if a field has been set.

### GetSnapPreventions

`func (o *SymmetricalMatrixRequest) GetSnapPreventions() []string`

GetSnapPreventions returns the SnapPreventions field if non-nil, zero value otherwise.

### GetSnapPreventionsOk

`func (o *SymmetricalMatrixRequest) GetSnapPreventionsOk() (*[]string, bool)`

GetSnapPreventionsOk returns a tuple with the SnapPreventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPreventions

`func (o *SymmetricalMatrixRequest) SetSnapPreventions(v []string)`

SetSnapPreventions sets SnapPreventions field to given value.

### HasSnapPreventions

`func (o *SymmetricalMatrixRequest) HasSnapPreventions() bool`

HasSnapPreventions returns a boolean if a field has been set.

### GetCurbsides

`func (o *SymmetricalMatrixRequest) GetCurbsides() []string`

GetCurbsides returns the Curbsides field if non-nil, zero value otherwise.

### GetCurbsidesOk

`func (o *SymmetricalMatrixRequest) GetCurbsidesOk() (*[]string, bool)`

GetCurbsidesOk returns a tuple with the Curbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurbsides

`func (o *SymmetricalMatrixRequest) SetCurbsides(v []string)`

SetCurbsides sets Curbsides field to given value.

### HasCurbsides

`func (o *SymmetricalMatrixRequest) HasCurbsides() bool`

HasCurbsides returns a boolean if a field has been set.

### GetOutArrays

`func (o *SymmetricalMatrixRequest) GetOutArrays() []string`

GetOutArrays returns the OutArrays field if non-nil, zero value otherwise.

### GetOutArraysOk

`func (o *SymmetricalMatrixRequest) GetOutArraysOk() (*[]string, bool)`

GetOutArraysOk returns a tuple with the OutArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutArrays

`func (o *SymmetricalMatrixRequest) SetOutArrays(v []string)`

SetOutArrays sets OutArrays field to given value.

### HasOutArrays

`func (o *SymmetricalMatrixRequest) HasOutArrays() bool`

HasOutArrays returns a boolean if a field has been set.

### GetFailFast

`func (o *SymmetricalMatrixRequest) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *SymmetricalMatrixRequest) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *SymmetricalMatrixRequest) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *SymmetricalMatrixRequest) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


