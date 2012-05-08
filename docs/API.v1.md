
# REST API

格式为`G|P URL`。其中`G`代表Get方式，`P`代表Post方式

## blah 操作

### G blah/all 

返回所有公开`blah`的基本信息列表。

**Resource URL**

http://host/api/blah/all.json

**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。
 - since_id: 返回的blah的ID都大于这个指定的值。
 - max_id: 返回的blah的ID都小于或等于这个指定的值。


### G blah/home

返回所有自己为参与者的`blah`的基本信息列表。

**Resource URL**

http://host/api/blah/home.json

**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。
 - since_id: 返回的blah的ID都大于这个指定的值。
 - max_id: 返回的blah的ID都小于或等于这个指定的值。

### G blah/:id

返回ID为:id的`blah`的详细信息

**Resource URL**

http://host/api/blah/:id.json


### P blah/create

创建新blah。初始化参与者为创建者。接受一个参数text，为blah的内容

**Resource URL**

http://host/api/blah/create.json

**Parameters**

 - visibility: public|privated，是否是公开的blah
 - text: 任意长度文本，内容


### P blah/destroy

删除一个blah。接收一个参数id，用于指定被删除的Blah。

**Resource URL**

http://host/api/blah/destroy.json

**Parameters**

 - id: 要删除的blah 的Id

### P blah/update

修改一个blah的文本或者可见性。

**Resource URL**

http://host/api/blah/update.json

**Parameters**

 - visibility: public|privated，是否是公开的blah
 - text: 任意长度文本，内容
 - id: 该blah的ID


## Blah Memeber 操作

### G blah/:id/members

返回ID为:id的`blah`参与者信息列表。

**Resource URL**

http://host/api/blah/:id/members.json
 
**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。

### P blah/:id/members/create

将一个用户加入指定的blah

**Resource URL**

http://host/api/blah/:id/members/create.json

**Parameters**

 - user_id: 被加入的用户ID


### P blah/:id/members/destroy

将一个用户从指定的blah删除

**Resource URL**

http://host/api/blah/:id/members/destroy.json

**Parameters**

 - user_id: 被加入的用户ID

## Message 操作

### G blah/:id/timeline

返回ID为:id的`blah`时间轴。

**Resource URL**

http://host/api/blah/:id/timeline.json

**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。
 - since_id: 返回的message的ID都大于这个指定的值。
 - max_id: 返回的message的ID都小于或等于这个指定的值。

### P blah/:id/create

在指定blah创建新message。

**Resource URL**

http://host/api/blah/:id/create.json

**Parameters**

 - text: 任意长度文本，内容
 - action: new|update，表示该动作是对之前 message 的修改还是仅仅添加新的 message
 - target_id: 如果action是update，则必须指定该参数

## User 操作

### G user/all

返回系统中所有用户。

**Resource URL**

http://host/api/user/all.json

**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。
 - cursor: 指定返回记录的游标。

### G user/star

返回系统中自己加星的用户。

**Resource URL**

http://host/api/user/star.json

**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。
 - cursor: 指定返回记录的游标。

### G user/:id

返回系统中的指定用户信息。

**Resource URL**

http://host/api/user/:id.json

**Parameters**

 - count: 指定接受多少个记录。默认为20，不得多于100。
 - cursor: 指定返回记录的游标。

# Streaming API
 ...
