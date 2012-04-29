
格式为`G|P URL`。其中`G`代表Get方式，`P`代表Post方式

## blah 操作

### G blah/all 

返回所有公开`blah`的基本信息列表。


### G blah/home

返回所有自己为参与者的`blah`的基本信息列表。

### G blah/:id

返回ID为:id的`blah`的详细信息

### G blah/:id/members

返回ID为:id的`blah`参与者信息列表。

### G blah/:id/timeline

返回ID为:id的`blah`时间轴。


### P blah/create

创建新blah。初始化参与者为创建者。接受一个参数text，为blah的内容

### P blah/destroy

删除一个blah。接收一个参数id，用于指定被删除的Blah。

参数:

 - visibility: public|privated，是否是公开的blah
 - text: 任意长度文本，内容

返回: 则返回该`blah`的详细信息

### P blah/destroy

参数:

 - 该blah的ID

返回: 则返回该`blah`的详细信息




