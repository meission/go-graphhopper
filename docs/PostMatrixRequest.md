# PostMatrixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The routing profile. It determines the network, speed and other physical attributes used when computing the route. See the section about [routing profiles](#tag/Map-Data-and-Routing-Profiles) for more details and valid profile values. | [optional] [default to "car"]
**FromPoints** | Pointer to **[][]float64** | The origin points for the routes in an array of &#x60;[longitude,latitude]&#x60;. | [optional] 
**ToPoints** | Pointer to **[][]float64** | The destination points for the routes in an array of &#x60;[longitude,latitude]&#x60;. | [optional] 
**FromPointHints** | Pointer to **[]string** | See &#x60;point_hints&#x60;of symmetrical matrix | [optional] 
**ToPointHints** | Pointer to **[]string** | See &#x60;point_hints&#x60;of symmetrical matrix | [optional] 
**SnapPreventions** | Pointer to **[]string** | Optional parameter. &#39;Snapping&#39; is the process of finding the closest road location for GPS coordinates provided in the &#x60;points&#x60; array. The &#x60;snap_preventions&#x60; array allows you to prevent snapping to specific types of roads. For example, if the array includes &#x60;bridge&#x60;, then the routing engine will avoid snapping to a bridge, even if it is the closest road for the given point. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the &#x60;custom_model&#x60;. | [optional] 
**FromCurbsides** | Pointer to **[]string** | See &#x60;curbsides&#x60;of symmetrical matrix | [optional] 
**ToCurbsides** | Pointer to **[]string** | See &#x60;curbsides&#x60;of symmetrical matrix | [optional] 
**OutArrays** | Pointer to **[]string** | Specifies which matrices should be included in the response. Specify one or more of the following options &#x60;weights&#x60;, &#x60;times&#x60;, &#x60;distances&#x60;. The units of the entries of &#x60;distances&#x60; are meters, of &#x60;times&#x60; are seconds and of &#x60;weights&#x60; is arbitrary and it can differ for different vehicle profiles or versions of this API. | [optional] 
**FailFast** | Pointer to **bool** | Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to &#x60;false&#x60; the time/weight/distance matrix will be calculated for all valid points and contain the &#x60;null&#x60; value for all entries that could not be calculated. The &#x60;hint&#x60; field of the response will also contain additional information about what went wrong (see its documentation). | [optional] [default to true]
**Points** | Pointer to **[][]float64** | Specify multiple points for which the weight-, route-, time- or distance-matrix should be calculated as follows: &#x60;[longitude,latitude]&#x60;. In this case the origins are identical to the destinations. Thus, if there are N points, NxN entries are calculated. The order of the point parameter is important. Specify at least three points. Cannot be used together with &#x60;from_point&#x60; or &#x60;to_point.&#x60;. | [optional] 
**PointHints** | Pointer to **[]string** | Optional parameter. Specifies a hint for each point in the &#x60;points&#x60; array to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | [optional] 
**Curbsides** | Pointer to **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap. | [optional] 

## Methods

### NewPostMatrixRequest

`func NewPostMatrixRequest() *PostMatrixRequest`

NewPostMatrixRequest instantiates a new PostMatrixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMatrixRequestWithDefaults

`func NewPostMatrixRequestWithDefaults() *PostMatrixRequest`

NewPostMatrixRequestWithDefaults instantiates a new PostMatrixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *PostMatrixRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PostMatrixRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PostMatrixRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PostMatrixRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetFromPoints

`func (o *PostMatrixRequest) GetFromPoints() [][]float64`

GetFromPoints returns the FromPoints field if non-nil, zero value otherwise.

### GetFromPointsOk

`func (o *PostMatrixRequest) GetFromPointsOk() (*[][]float64, bool)`

GetFromPointsOk returns a tuple with the FromPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPoints

`func (o *PostMatrixRequest) SetFromPoints(v [][]float64)`

SetFromPoints sets FromPoints field to given value.

### HasFromPoints

`func (o *PostMatrixRequest) HasFromPoints() bool`

HasFromPoints returns a boolean if a field has been set.

### GetToPoints

`func (o *PostMatrixRequest) GetToPoints() [][]float64`

GetToPoints returns the ToPoints field if non-nil, zero value otherwise.

### GetToPointsOk

`func (o *PostMatrixRequest) GetToPointsOk() (*[][]float64, bool)`

GetToPointsOk returns a tuple with the ToPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPoints

`func (o *PostMatrixRequest) SetToPoints(v [][]float64)`

SetToPoints sets ToPoints field to given value.

### HasToPoints

`func (o *PostMatrixRequest) HasToPoints() bool`

HasToPoints returns a boolean if a field has been set.

### GetFromPointHints

`func (o *PostMatrixRequest) GetFromPointHints() []string`

GetFromPointHints returns the FromPointHints field if non-nil, zero value otherwise.

### GetFromPointHintsOk

`func (o *PostMatrixRequest) GetFromPointHintsOk() (*[]string, bool)`

GetFromPointHintsOk returns a tuple with the FromPointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPointHints

`func (o *PostMatrixRequest) SetFromPointHints(v []string)`

SetFromPointHints sets FromPointHints field to given value.

### HasFromPointHints

`func (o *PostMatrixRequest) HasFromPointHints() bool`

HasFromPointHints returns a boolean if a field has been set.

### GetToPointHints

`func (o *PostMatrixRequest) GetToPointHints() []string`

GetToPointHints returns the ToPointHints field if non-nil, zero value otherwise.

### GetToPointHintsOk

`func (o *PostMatrixRequest) GetToPointHintsOk() (*[]string, bool)`

GetToPointHintsOk returns a tuple with the ToPointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPointHints

`func (o *PostMatrixRequest) SetToPointHints(v []string)`

SetToPointHints sets ToPointHints field to given value.

### HasToPointHints

`func (o *PostMatrixRequest) HasToPointHints() bool`

HasToPointHints returns a boolean if a field has been set.

### GetSnapPreventions

`func (o *PostMatrixRequest) GetSnapPreventions() []string`

GetSnapPreventions returns the SnapPreventions field if non-nil, zero value otherwise.

### GetSnapPreventionsOk

`func (o *PostMatrixRequest) GetSnapPreventionsOk() (*[]string, bool)`

GetSnapPreventionsOk returns a tuple with the SnapPreventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPreventions

`func (o *PostMatrixRequest) SetSnapPreventions(v []string)`

SetSnapPreventions sets SnapPreventions field to given value.

### HasSnapPreventions

`func (o *PostMatrixRequest) HasSnapPreventions() bool`

HasSnapPreventions returns a boolean if a field has been set.

### GetFromCurbsides

`func (o *PostMatrixRequest) GetFromCurbsides() []string`

GetFromCurbsides returns the FromCurbsides field if non-nil, zero value otherwise.

### GetFromCurbsidesOk

`func (o *PostMatrixRequest) GetFromCurbsidesOk() (*[]string, bool)`

GetFromCurbsidesOk returns a tuple with the FromCurbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurbsides

`func (o *PostMatrixRequest) SetFromCurbsides(v []string)`

SetFromCurbsides sets FromCurbsides field to given value.

### HasFromCurbsides

`func (o *PostMatrixRequest) HasFromCurbsides() bool`

HasFromCurbsides returns a boolean if a field has been set.

### GetToCurbsides

`func (o *PostMatrixRequest) GetToCurbsides() []string`

GetToCurbsides returns the ToCurbsides field if non-nil, zero value otherwise.

### GetToCurbsidesOk

`func (o *PostMatrixRequest) GetToCurbsidesOk() (*[]string, bool)`

GetToCurbsidesOk returns a tuple with the ToCurbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurbsides

`func (o *PostMatrixRequest) SetToCurbsides(v []string)`

SetToCurbsides sets ToCurbsides field to given value.

### HasToCurbsides

`func (o *PostMatrixRequest) HasToCurbsides() bool`

HasToCurbsides returns a boolean if a field has been set.

### GetOutArrays

`func (o *PostMatrixRequest) GetOutArrays() []string`

GetOutArrays returns the OutArrays field if non-nil, zero value otherwise.

### GetOutArraysOk

`func (o *PostMatrixRequest) GetOutArraysOk() (*[]string, bool)`

GetOutArraysOk returns a tuple with the OutArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutArrays

`func (o *PostMatrixRequest) SetOutArrays(v []string)`

SetOutArrays sets OutArrays field to given value.

### HasOutArrays

`func (o *PostMatrixRequest) HasOutArrays() bool`

HasOutArrays returns a boolean if a field has been set.

### GetFailFast

`func (o *PostMatrixRequest) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *PostMatrixRequest) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *PostMatrixRequest) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *PostMatrixRequest) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.

### GetPoints

`func (o *PostMatrixRequest) GetPoints() [][]float64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *PostMatrixRequest) GetPointsOk() (*[][]float64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *PostMatrixRequest) SetPoints(v [][]float64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *PostMatrixRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPointHints

`func (o *PostMatrixRequest) GetPointHints() []string`

GetPointHints returns the PointHints field if non-nil, zero value otherwise.

### GetPointHintsOk

`func (o *PostMatrixRequest) GetPointHintsOk() (*[]string, bool)`

GetPointHintsOk returns a tuple with the PointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointHints

`func (o *PostMatrixRequest) SetPointHints(v []string)`

SetPointHints sets PointHints field to given value.

### HasPointHints

`func (o *PostMatrixRequest) HasPointHints() bool`

HasPointHints returns a boolean if a field has been set.

### GetCurbsides

`func (o *PostMatrixRequest) GetCurbsides() []string`

GetCurbsides returns the Curbsides field if non-nil, zero value otherwise.

### GetCurbsidesOk

`func (o *PostMatrixRequest) GetCurbsidesOk() (*[]string, bool)`

GetCurbsidesOk returns a tuple with the Curbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurbsides

`func (o *PostMatrixRequest) SetCurbsides(v []string)`

SetCurbsides sets Curbsides field to given value.

### HasCurbsides

`func (o *PostMatrixRequest) HasCurbsides() bool`

HasCurbsides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


