# BBox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bbox** | Pointer to **[]float64** | A rectangular area given as an array [minLon, minLat, maxLon, maxLat]. The created profile will only work in this area. The maximum area size is 160 000 square kilometers. | [optional] 

## Methods

### NewBBox

`func NewBBox() *BBox`

NewBBox instantiates a new BBox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBBoxWithDefaults

`func NewBBoxWithDefaults() *BBox`

NewBBoxWithDefaults instantiates a new BBox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBbox

`func (o *BBox) GetBbox() []float64`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *BBox) GetBboxOk() (*[]float64, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *BBox) SetBbox(v []float64)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *BBox) HasBbox() bool`

HasBbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


