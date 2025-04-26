# ProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The built-in profile your new custom profile shall be based on. Currently we only support &#x60;car&#x60; but you can add size restrictions and speed reduction to get similar profiles like &#x60;small_truck&#x60; and &#x60;truck&#x60;. Contact us to learn more about this. | [optional] 
**Bounds** | Pointer to [**ProfileRequestBounds**](ProfileRequestBounds.md) |  | [optional] 
**CustomModel** | Pointer to [**CustomModelForProfile**](CustomModelForProfile.md) |  | [optional] 

## Methods

### NewProfileRequest

`func NewProfileRequest() *ProfileRequest`

NewProfileRequest instantiates a new ProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileRequestWithDefaults

`func NewProfileRequestWithDefaults() *ProfileRequest`

NewProfileRequestWithDefaults instantiates a new ProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ProfileRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ProfileRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ProfileRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ProfileRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetBounds

`func (o *ProfileRequest) GetBounds() ProfileRequestBounds`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *ProfileRequest) GetBoundsOk() (*ProfileRequestBounds, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *ProfileRequest) SetBounds(v ProfileRequestBounds)`

SetBounds sets Bounds field to given value.

### HasBounds

`func (o *ProfileRequest) HasBounds() bool`

HasBounds returns a boolean if a field has been set.

### GetCustomModel

`func (o *ProfileRequest) GetCustomModel() CustomModelForProfile`

GetCustomModel returns the CustomModel field if non-nil, zero value otherwise.

### GetCustomModelOk

`func (o *ProfileRequest) GetCustomModelOk() (*CustomModelForProfile, bool)`

GetCustomModelOk returns a tuple with the CustomModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomModel

`func (o *ProfileRequest) SetCustomModel(v CustomModelForProfile)`

SetCustomModel sets CustomModel field to given value.

### HasCustomModel

`func (o *ProfileRequest) HasCustomModel() bool`

HasCustomModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


