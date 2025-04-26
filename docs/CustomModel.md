# CustomModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Speed** | Pointer to **[]map[string]interface{}** | See [speed customization](#tag/Custom-Model/Customizing-speed) | [optional] 
**Priority** | Pointer to **[]map[string]interface{}** | See [priority customization](#tag/Custom-Model/Customizing-priority) | [optional] 
**DistanceInfluence** | Pointer to **float32** | Use higher values to prefer shorter routes. See [here](#tag/Custom-Model/Customizing-distance_influence) for more details. | [optional] [default to 70]
**Areas** | Pointer to [**FeatureCollection**](FeatureCollection.md) | Areas are given in a GeoJson format as FeatureCollection. See more details and an example [here](#tag/Custom-Model/Define-areas).  | [optional] 

## Methods

### NewCustomModel

`func NewCustomModel() *CustomModel`

NewCustomModel instantiates a new CustomModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomModelWithDefaults

`func NewCustomModelWithDefaults() *CustomModel`

NewCustomModelWithDefaults instantiates a new CustomModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpeed

`func (o *CustomModel) GetSpeed() []map[string]interface{}`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CustomModel) GetSpeedOk() (*[]map[string]interface{}, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CustomModel) SetSpeed(v []map[string]interface{})`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CustomModel) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPriority

`func (o *CustomModel) GetPriority() []map[string]interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CustomModel) GetPriorityOk() (*[]map[string]interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CustomModel) SetPriority(v []map[string]interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CustomModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDistanceInfluence

`func (o *CustomModel) GetDistanceInfluence() float32`

GetDistanceInfluence returns the DistanceInfluence field if non-nil, zero value otherwise.

### GetDistanceInfluenceOk

`func (o *CustomModel) GetDistanceInfluenceOk() (*float32, bool)`

GetDistanceInfluenceOk returns a tuple with the DistanceInfluence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceInfluence

`func (o *CustomModel) SetDistanceInfluence(v float32)`

SetDistanceInfluence sets DistanceInfluence field to given value.

### HasDistanceInfluence

`func (o *CustomModel) HasDistanceInfluence() bool`

HasDistanceInfluence returns a boolean if a field has been set.

### GetAreas

`func (o *CustomModel) GetAreas() FeatureCollection`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *CustomModel) GetAreasOk() (*FeatureCollection, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *CustomModel) SetAreas(v FeatureCollection)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *CustomModel) HasAreas() bool`

HasAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


