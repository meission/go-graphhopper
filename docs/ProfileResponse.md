# ProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **string** | The built-in profile this custom profile is based on. | [optional] 
**Bounds** | Pointer to [**ProfileResponseBounds**](ProfileResponseBounds.md) |  | [optional] 
**CustomModel** | Pointer to [**CustomModelForProfile**](CustomModelForProfile.md) |  | [optional] 
**Id** | Pointer to **string** | The name of the created profile. Use this as the &#x60;profile&#x60; parameter for the /route API etc. For route optimization requests you need to define a vehicle_type where you can enter the custom profile. | [optional] 

## Methods

### NewProfileResponse

`func NewProfileResponse() *ProfileResponse`

NewProfileResponse instantiates a new ProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResponseWithDefaults

`func NewProfileResponseWithDefaults() *ProfileResponse`

NewProfileResponseWithDefaults instantiates a new ProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ProfileResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ProfileResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ProfileResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ProfileResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetBounds

`func (o *ProfileResponse) GetBounds() ProfileResponseBounds`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *ProfileResponse) GetBoundsOk() (*ProfileResponseBounds, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *ProfileResponse) SetBounds(v ProfileResponseBounds)`

SetBounds sets Bounds field to given value.

### HasBounds

`func (o *ProfileResponse) HasBounds() bool`

HasBounds returns a boolean if a field has been set.

### GetCustomModel

`func (o *ProfileResponse) GetCustomModel() CustomModelForProfile`

GetCustomModel returns the CustomModel field if non-nil, zero value otherwise.

### GetCustomModelOk

`func (o *ProfileResponse) GetCustomModelOk() (*CustomModelForProfile, bool)`

GetCustomModelOk returns a tuple with the CustomModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomModel

`func (o *ProfileResponse) SetCustomModel(v CustomModelForProfile)`

SetCustomModel sets CustomModel field to given value.

### HasCustomModel

`func (o *ProfileResponse) HasCustomModel() bool`

HasCustomModel returns a boolean if a field has been set.

### GetId

`func (o *ProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


