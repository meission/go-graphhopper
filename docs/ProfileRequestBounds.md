# ProfileRequestBounds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bbox** | Pointer to **[]float64** | A rectangular area given as an array [minLon, minLat, maxLon, maxLat]. The created profile will only work in this area. The maximum area size is 160 000 square kilometers. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]FeatureCollectionFeaturesInner**](FeatureCollectionFeaturesInner.md) | A FeatureCollection is an array of type \&quot;Feature\&quot; from GeoJSON | [optional] 

## Methods

### NewProfileRequestBounds

`func NewProfileRequestBounds() *ProfileRequestBounds`

NewProfileRequestBounds instantiates a new ProfileRequestBounds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileRequestBoundsWithDefaults

`func NewProfileRequestBoundsWithDefaults() *ProfileRequestBounds`

NewProfileRequestBoundsWithDefaults instantiates a new ProfileRequestBounds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBbox

`func (o *ProfileRequestBounds) GetBbox() []float64`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *ProfileRequestBounds) GetBboxOk() (*[]float64, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *ProfileRequestBounds) SetBbox(v []float64)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *ProfileRequestBounds) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### GetType

`func (o *ProfileRequestBounds) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileRequestBounds) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileRequestBounds) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProfileRequestBounds) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFeatures

`func (o *ProfileRequestBounds) GetFeatures() []FeatureCollectionFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProfileRequestBounds) GetFeaturesOk() (*[]FeatureCollectionFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProfileRequestBounds) SetFeatures(v []FeatureCollectionFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ProfileRequestBounds) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


