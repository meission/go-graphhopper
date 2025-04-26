# RouteResponsePathInstructionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | A description what the user has to do in order to follow the route. The language depends on the locale parameter.  | [optional] 
**StreetName** | Pointer to **string** | The name of the street to turn onto in order to follow the route.  | [optional] 
**Distance** | Pointer to **float64** | The distance for this instruction, in meters.  | [optional] 
**Time** | Pointer to **int32** | The duration for this instruction, in milliseconds.  | [optional] 
**Interval** | Pointer to **[]int32** | Two indices into &#x60;points&#x60;, referring to the beginning and the end of the segment of the route this instruction refers to.  | [optional] 
**Sign** | Pointer to **int32** | A number which specifies the sign to show:  | sign | description  | |---|---| |-98| an U-turn without the knowledge if it is a right or left U-turn | | -8| a left U-turn | | -7| keep left | | -6| **not yet used**: leave roundabout | | -3| turn sharp left | | -2| turn left | | -1| turn slight left | |  0| continue on street | |  1| turn slight right | |  2| turn right | |  3| turn sharp right | |  4| the finish instruction before the last point | |  5| the instruction before a via point | |  6| the instruction before entering a roundabout | |  7| keep right | |  8| a right U-turn | |  *| **For future compatibility** it is important that all clients are able to handle also unknown instruction sign numbers  | [optional] 
**ExitNumber** | Pointer to **int32** | Only available for roundabout instructions (sign is 6). The count of exits at which the route leaves the roundabout.  | [optional] 
**TurnAngle** | Pointer to **float64** | Only available for roundabout instructions (sign is 6). The radian of the route within the roundabout &#x60;0 &lt; r &lt; 2*PI&#x60; for clockwise and &#x60;-2*PI &lt; r &lt; 0&#x60; for counterclockwise turns.  | [optional] 

## Methods

### NewRouteResponsePathInstructionsInner

`func NewRouteResponsePathInstructionsInner() *RouteResponsePathInstructionsInner`

NewRouteResponsePathInstructionsInner instantiates a new RouteResponsePathInstructionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteResponsePathInstructionsInnerWithDefaults

`func NewRouteResponsePathInstructionsInnerWithDefaults() *RouteResponsePathInstructionsInner`

NewRouteResponsePathInstructionsInnerWithDefaults instantiates a new RouteResponsePathInstructionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *RouteResponsePathInstructionsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RouteResponsePathInstructionsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RouteResponsePathInstructionsInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RouteResponsePathInstructionsInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetStreetName

`func (o *RouteResponsePathInstructionsInner) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *RouteResponsePathInstructionsInner) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *RouteResponsePathInstructionsInner) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *RouteResponsePathInstructionsInner) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetDistance

`func (o *RouteResponsePathInstructionsInner) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RouteResponsePathInstructionsInner) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RouteResponsePathInstructionsInner) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RouteResponsePathInstructionsInner) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetTime

`func (o *RouteResponsePathInstructionsInner) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RouteResponsePathInstructionsInner) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RouteResponsePathInstructionsInner) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *RouteResponsePathInstructionsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetInterval

`func (o *RouteResponsePathInstructionsInner) GetInterval() []int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RouteResponsePathInstructionsInner) GetIntervalOk() (*[]int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RouteResponsePathInstructionsInner) SetInterval(v []int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RouteResponsePathInstructionsInner) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetSign

`func (o *RouteResponsePathInstructionsInner) GetSign() int32`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *RouteResponsePathInstructionsInner) GetSignOk() (*int32, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *RouteResponsePathInstructionsInner) SetSign(v int32)`

SetSign sets Sign field to given value.

### HasSign

`func (o *RouteResponsePathInstructionsInner) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetExitNumber

`func (o *RouteResponsePathInstructionsInner) GetExitNumber() int32`

GetExitNumber returns the ExitNumber field if non-nil, zero value otherwise.

### GetExitNumberOk

`func (o *RouteResponsePathInstructionsInner) GetExitNumberOk() (*int32, bool)`

GetExitNumberOk returns a tuple with the ExitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitNumber

`func (o *RouteResponsePathInstructionsInner) SetExitNumber(v int32)`

SetExitNumber sets ExitNumber field to given value.

### HasExitNumber

`func (o *RouteResponsePathInstructionsInner) HasExitNumber() bool`

HasExitNumber returns a boolean if a field has been set.

### GetTurnAngle

`func (o *RouteResponsePathInstructionsInner) GetTurnAngle() float64`

GetTurnAngle returns the TurnAngle field if non-nil, zero value otherwise.

### GetTurnAngleOk

`func (o *RouteResponsePathInstructionsInner) GetTurnAngleOk() (*float64, bool)`

GetTurnAngleOk returns a tuple with the TurnAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnAngle

`func (o *RouteResponsePathInstructionsInner) SetTurnAngle(v float64)`

SetTurnAngle sets TurnAngle field to given value.

### HasTurnAngle

`func (o *RouteResponsePathInstructionsInner) HasTurnAngle() bool`

HasTurnAngle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


