# FeatureCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]FeatureCollectionFeaturesInner**](FeatureCollectionFeaturesInner.md) | A FeatureCollection is an array of type \&quot;Feature\&quot; from GeoJSON | [optional] 

## Methods

### NewFeatureCollection

`func NewFeatureCollection() *FeatureCollection`

NewFeatureCollection instantiates a new FeatureCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureCollectionWithDefaults

`func NewFeatureCollectionWithDefaults() *FeatureCollection`

NewFeatureCollectionWithDefaults instantiates a new FeatureCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FeatureCollection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeatureCollection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeatureCollection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeatureCollection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFeatures

`func (o *FeatureCollection) GetFeatures() []FeatureCollectionFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeatureCollection) GetFeaturesOk() (*[]FeatureCollectionFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeatureCollection) SetFeatures(v []FeatureCollectionFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FeatureCollection) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


