#Draft of Blahdy

## Summary

This is only a draft. Our goal is to design an alternative to IM/Email in companies and organizations: Blahdy.

If you are an author of this draft, your duty is keeping the Blah's design clean and easy to use. Don't panic.

## Technological Outline

**Blahdy** is a B/S Rich Internet Application. In browser, it use standard Web technology to provide a Rich User Interface. In server part, it use go language and key-value database to provide stable service. Websocket is the key to connect those two parts, which make it realtime and high performace.

## Plan

### Stage 1

** API **

 - Create/Destroy `Blah`
 - Create/Update `Action`
 - Add/Remove Members to/from `Blah`

** UI **

 - `Blah` Basic Operations
 - `Action` Basic Operations
 - Member Basic Operations
 - UI Framework
 - Update Service

### Stage 2

** API **

 - MultiMedia Action Support: Create, Update
   - Supported Formats: Image, Video, URL, Docs, etc.
	 - Blah only store the metadata of MultiMedia
	 - A Media Data Server is needed to store the entities of MultiMedia
 - Privilege of Blah
   - There is a privilege mask in each Blah.
	 - The mask is similar to Unix privilege mask. For more details of Blah privilege model, read the following sections.

** UI **
 
 - MultiMedia Action: Create, Update, Preview MultiMedia.
 - Privilege Basic Operations
 - Notification System 

### Stage 3

** API **

 - Friendship Management: Create/Destroy/GetInfo
 - ...

** UI **

 - Friendship Basic Operations
 - ...
	

## Database Design.

### Blah and Action

The basic entity of Blah is `Blah`. Each `Blah` contains zero or more `Actions`.

Each `Actions` can be identified by an unique 64-bit integer, which is the ID of the `Action`.

Each `Blah` can be identified by an unique 64-bit integer, which is the ID of the `Blah`.

For each `Blah`, system has to maintains a sequence which contains the own `Actions` in time order.

**Structure of Action**

 - Name: Create|Update:String
 - ScreenName: Author.ScreenName:String
 - DisplayName: Author.DisplayName:String
 - Text: :String
 - Attachments: :List of Metadata
	 
**Structure of Blah**

 - Title: :String
 - Owner: Author.ScreenName:String
 - Participators: :List of Author.ScreenName
 - Timeline: :List of Action
 - Group: ...
 - ACM: :AccessControlMatrix
 - Public: true|false:Boolean 

## Access Control

The access control model of Blah is similar to Unix privilege model, but simper. Blah uses a matrix to describe the authorization information of a specified Blah.

Here we go:

We describe all users in Blahdy as 4 roles: owner, group, participators and others.

 - Owner: the creator of this `Blah`.
 - Group: a specified group assigned with this Blah.
 - Participator: all participators of this Blah.
 - Others: other people.

Each role can apply some of following action:

 - Read this `Blah`.
 - Write this `Blah`.
 - Modify `Blah` information: subject, title, etc.
 - Member Management.
 - Destroy this `Blah`.

The matrix:

					Write		Modify	MemberManagment	
	Owner		?				?				?
	Group   ?				?				?
	Partic  ?			  ?				?
	Others	?				? 			?

 - 1 donates the role has this privilege and it's fixed
 - ? donates the role can be assigned this privilege.
 - 0 donates the role cannot be assigned this privilege and it' fixed.

In implement, the role `Owner` has full privilege so can omit it too.

The mask(Fuzzy):

For `Group` and `Participators`, each one use a quadruple to describe the privilege: The 1st digit denotes the visibility, public or privated. And the followed three digit denote the privilege of three roles.

Hence, we use a integer as the mask for each roles:
	
	W		Mo	MM	Digit		Comment
	0		0		0		0				Read only
	1		0		0		1				Read and Write
	0		1		0   2				Read and Modify
	1		1		0		3			  Read, Write and Modify
	0		0		1		4				Read and MemberManagment
	1		0		1		5				Read, Write and MemberManagment
	0		1		1		6				Read, Modify and MemberManagment
	1		1		1   7				Read, Write, Modify and MemberManagment

For example:

 - 0777 means all people can Write/Modify/MemberManagment this `Blah`, but it's privated.
 - 1711 means a specified group has full privilege to this `Blah` but participators and other people can only Read and Write

## API

信息获取

 - G blah/:id
   获取ID为:id的 blah 信息
 - G blah/all
   获取你的所有 blah

成员操作

 - P blah/members/destroy
   从某个 blah 中删除某些成员
 - P blah/members/create
   将某些成员从某个 blah 中删除

blah 操作

 - P blah/update
   更新某个 blah 的信息
 - P blah/create
   创建一个 blah
 - P blah/destroy
   删除一个指定的 blah

好友操作

 - P friendships/create
   添加一个好友
 - P friendships/destroy
   删除一个好友
 - G friendships/show/:id
   获取某个ID为:id的好友信息

信息发布

 - P blah/actions/create
   在某个 blah 上发布一个消息
 - P blah/actions/create
   在某个 blah 上更新一个消息
