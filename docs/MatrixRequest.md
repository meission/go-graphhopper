# MatrixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The routing profile. It determines the network, speed and other physical attributes used when computing the route. See the section about [routing profiles](#tag/Map-Data-and-Routing-Profiles) for more details and valid profile values. | [optional] [default to "car"]
**FromPoints** | Pointer to **[][]float64** | The origin points for the routes in an array of &#x60;[longitude,latitude]&#x60;. | [optional] 
**ToPoints** | Pointer to **[][]float64** | The destination points for the routes in an array of &#x60;[longitude,latitude]&#x60;. | [optional] 
**FromPointHints** | Pointer to **[]string** | See &#x60;point_hints&#x60;of symmetrical matrix | [optional] 
**ToPointHints** | Pointer to **[]string** | See &#x60;point_hints&#x60;of symmetrical matrix | [optional] 
**SnapPreventions** | Pointer to **[]string** | See &#x60;snap_preventions&#x60; of symmetrical matrix | [optional] 
**FromCurbsides** | Pointer to **[]string** | See &#x60;curbsides&#x60;of symmetrical matrix | [optional] 
**ToCurbsides** | Pointer to **[]string** | See &#x60;curbsides&#x60;of symmetrical matrix | [optional] 
**OutArrays** | Pointer to **[]string** | Specifies which matrices should be included in the response. Specify one or more of the following options &#x60;weights&#x60;, &#x60;times&#x60;, &#x60;distances&#x60;. The units of the entries of &#x60;distances&#x60; are meters, of &#x60;times&#x60; are seconds and of &#x60;weights&#x60; is arbitrary and it can differ for different vehicles or versions of this API. | [optional] 
**FailFast** | Pointer to **bool** | Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to &#x60;false&#x60; the time/weight/distance matrix will be calculated for all valid points and contain the &#x60;null&#x60; value for all entries that could not be calculated. The &#x60;hint&#x60; field of the response will also contain additional information about what went wrong (see its documentation). | [optional] [default to true]

## Methods

### NewMatrixRequest

`func NewMatrixRequest() *MatrixRequest`

NewMatrixRequest instantiates a new MatrixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatrixRequestWithDefaults

`func NewMatrixRequestWithDefaults() *MatrixRequest`

NewMatrixRequestWithDefaults instantiates a new MatrixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *MatrixRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MatrixRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MatrixRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MatrixRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetFromPoints

`func (o *MatrixRequest) GetFromPoints() [][]float64`

GetFromPoints returns the FromPoints field if non-nil, zero value otherwise.

### GetFromPointsOk

`func (o *MatrixRequest) GetFromPointsOk() (*[][]float64, bool)`

GetFromPointsOk returns a tuple with the FromPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPoints

`func (o *MatrixRequest) SetFromPoints(v [][]float64)`

SetFromPoints sets FromPoints field to given value.

### HasFromPoints

`func (o *MatrixRequest) HasFromPoints() bool`

HasFromPoints returns a boolean if a field has been set.

### GetToPoints

`func (o *MatrixRequest) GetToPoints() [][]float64`

GetToPoints returns the ToPoints field if non-nil, zero value otherwise.

### GetToPointsOk

`func (o *MatrixRequest) GetToPointsOk() (*[][]float64, bool)`

GetToPointsOk returns a tuple with the ToPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPoints

`func (o *MatrixRequest) SetToPoints(v [][]float64)`

SetToPoints sets ToPoints field to given value.

### HasToPoints

`func (o *MatrixRequest) HasToPoints() bool`

HasToPoints returns a boolean if a field has been set.

### GetFromPointHints

`func (o *MatrixRequest) GetFromPointHints() []string`

GetFromPointHints returns the FromPointHints field if non-nil, zero value otherwise.

### GetFromPointHintsOk

`func (o *MatrixRequest) GetFromPointHintsOk() (*[]string, bool)`

GetFromPointHintsOk returns a tuple with the FromPointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPointHints

`func (o *MatrixRequest) SetFromPointHints(v []string)`

SetFromPointHints sets FromPointHints field to given value.

### HasFromPointHints

`func (o *MatrixRequest) HasFromPointHints() bool`

HasFromPointHints returns a boolean if a field has been set.

### GetToPointHints

`func (o *MatrixRequest) GetToPointHints() []string`

GetToPointHints returns the ToPointHints field if non-nil, zero value otherwise.

### GetToPointHintsOk

`func (o *MatrixRequest) GetToPointHintsOk() (*[]string, bool)`

GetToPointHintsOk returns a tuple with the ToPointHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPointHints

`func (o *MatrixRequest) SetToPointHints(v []string)`

SetToPointHints sets ToPointHints field to given value.

### HasToPointHints

`func (o *MatrixRequest) HasToPointHints() bool`

HasToPointHints returns a boolean if a field has been set.

### GetSnapPreventions

`func (o *MatrixRequest) GetSnapPreventions() []string`

GetSnapPreventions returns the SnapPreventions field if non-nil, zero value otherwise.

### GetSnapPreventionsOk

`func (o *MatrixRequest) GetSnapPreventionsOk() (*[]string, bool)`

GetSnapPreventionsOk returns a tuple with the SnapPreventions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPreventions

`func (o *MatrixRequest) SetSnapPreventions(v []string)`

SetSnapPreventions sets SnapPreventions field to given value.

### HasSnapPreventions

`func (o *MatrixRequest) HasSnapPreventions() bool`

HasSnapPreventions returns a boolean if a field has been set.

### GetFromCurbsides

`func (o *MatrixRequest) GetFromCurbsides() []string`

GetFromCurbsides returns the FromCurbsides field if non-nil, zero value otherwise.

### GetFromCurbsidesOk

`func (o *MatrixRequest) GetFromCurbsidesOk() (*[]string, bool)`

GetFromCurbsidesOk returns a tuple with the FromCurbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurbsides

`func (o *MatrixRequest) SetFromCurbsides(v []string)`

SetFromCurbsides sets FromCurbsides field to given value.

### HasFromCurbsides

`func (o *MatrixRequest) HasFromCurbsides() bool`

HasFromCurbsides returns a boolean if a field has been set.

### GetToCurbsides

`func (o *MatrixRequest) GetToCurbsides() []string`

GetToCurbsides returns the ToCurbsides field if non-nil, zero value otherwise.

### GetToCurbsidesOk

`func (o *MatrixRequest) GetToCurbsidesOk() (*[]string, bool)`

GetToCurbsidesOk returns a tuple with the ToCurbsides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurbsides

`func (o *MatrixRequest) SetToCurbsides(v []string)`

SetToCurbsides sets ToCurbsides field to given value.

### HasToCurbsides

`func (o *MatrixRequest) HasToCurbsides() bool`

HasToCurbsides returns a boolean if a field has been set.

### GetOutArrays

`func (o *MatrixRequest) GetOutArrays() []string`

GetOutArrays returns the OutArrays field if non-nil, zero value otherwise.

### GetOutArraysOk

`func (o *MatrixRequest) GetOutArraysOk() (*[]string, bool)`

GetOutArraysOk returns a tuple with the OutArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutArrays

`func (o *MatrixRequest) SetOutArrays(v []string)`

SetOutArrays sets OutArrays field to given value.

### HasOutArrays

`func (o *MatrixRequest) HasOutArrays() bool`

HasOutArrays returns a boolean if a field has been set.

### GetFailFast

`func (o *MatrixRequest) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *MatrixRequest) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *MatrixRequest) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *MatrixRequest) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


